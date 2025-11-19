#!/usr/bin/env bash
#
# Batch LinkedDoc annotation generator with GraphFS validation
# Uses Haiku for cost-effective bulk generation
#

set -euo pipefail

# Configuration
BATCH_SIZE=30
VALIDATION_INTERVAL=30
REPO_ROOT=$(git rev-parse --show-toplevel)
LOG_DIR="$REPO_ROOT/.annotation_logs"
PROGRESS_FILE="$LOG_DIR/progress.txt"
ERROR_LOG="$LOG_DIR/errors.log"
STATS_FILE="$LOG_DIR/stats.json"

# Colors
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
CYAN='\033[0;36m'
NC='\033[0m'

# Statistics
TOTAL_FILES=0
PROCESSED=0
ANNOTATED=0
SKIPPED=0
FAILED=0
VALIDATION_RUNS=0
VALIDATION_ERRORS=0

# Print utilities
print_color() { echo -e "${1}${2}${NC}"; }
print_error() { print_color "$RED" "✗ $*"; }
print_success() { print_color "$GREEN" "✓ $*"; }
print_info() { print_color "$BLUE" "ℹ $*"; }
print_warning() { print_color "$YELLOW" "⚠ $*"; }
print_header() { print_color "$CYAN" "━━━ $* ━━━"; }

# Initialize
init() {
    mkdir -p "$LOG_DIR"
    echo "0" > "$PROGRESS_FILE"
    echo "" > "$ERROR_LOG"

    print_header "Batch LinkedDoc Annotation - Consul Codebase"
    print_info "Repository: $REPO_ROOT"
    print_info "Batch size: $BATCH_SIZE"
    print_info "Validation interval: $VALIDATION_INTERVAL files"
    echo ""
}

# Check if file already has LinkedDoc
has_linkeddoc() {
    grep -q "<!-- LinkedDoc RDF -->" "$1" 2>/dev/null
}

# Get files to process
get_files_to_process() {
    find "$REPO_ROOT" -name "*.go" -type f \
        | grep -v "/vendor/" \
        | grep -v "_test.go" \
        | sort
}

# Save progress
save_progress() {
    cat > "$STATS_FILE" <<EOF
{
  "total_files": $TOTAL_FILES,
  "processed": $PROCESSED,
  "annotated": $ANNOTATED,
  "skipped": $SKIPPED,
  "failed": $FAILED,
  "validation_runs": $VALIDATION_RUNS,
  "validation_errors": $VALIDATION_ERRORS,
  "timestamp": "$(date -u +%Y-%m-%dT%H:%M:%SZ)"
}
EOF
}

# Print progress
print_progress() {
    local pct=0
    if [ $TOTAL_FILES -gt 0 ]; then
        pct=$((PROCESSED * 100 / TOTAL_FILES))
    fi

    echo ""
    print_header "Progress: $PROCESSED/$TOTAL_FILES ($pct%)"
    echo "  Annotated: $ANNOTATED"
    echo "  Skipped:   $SKIPPED"
    echo "  Failed:    $FAILED"
    echo "  Validations: $VALIDATION_RUNS (${VALIDATION_ERRORS} errors)"
    echo ""
}

# Validate with GraphFS
validate_graph() {
    print_info "Running GraphFS validation..."

    local output
    if output=$(graphfs scan --validate --stats 2>&1); then
        print_success "GraphFS validation passed"
        ((VALIDATION_RUNS++))
        echo "$output"
        return 0
    else
        print_error "GraphFS validation failed"
        ((VALIDATION_RUNS++))
        ((VALIDATION_ERRORS++))
        echo "$output" | tee -a "$ERROR_LOG"
        return 1
    fi
}

# Process single file using dox script
process_file_with_dox() {
    local file=$1

    # Skip if already has LinkedDoc
    if has_linkeddoc "$file"; then
        return 2  # Return code 2 = already annotated
    fi

    # Run dox script on this file
    if /Users/coolbeans/Development/dev/graphfs/scripts/dox --full-dox "$file" >/dev/null 2>&1; then
        return 0  # Success
    else
        return 1  # Failed
    fi
}

# Main processing loop
process_all_files() {
    local -a files
    mapfile -t files < <(get_files_to_process)

    TOTAL_FILES=${#files[@]}
    print_info "Found $TOTAL_FILES Go files to process"
    echo ""

    local batch_count=0

    for file in "${files[@]}"; do
        ((PROCESSED++))

        local rel_path="${file#$REPO_ROOT/}"
        echo -n "[$PROCESSED/$TOTAL_FILES] Processing $rel_path ... "

        local result
        set +e
        process_file_with_dox "$file"
        result=$?
        set -e

        case $result in
            0)
                echo -e "${GREEN}✓ annotated${NC}"
                ((ANNOTATED++))
                ((batch_count++))
                ;;
            2)
                echo -e "${YELLOW}○ skipped (already annotated)${NC}"
                ((SKIPPED++))
                ;;
            *)
                echo -e "${RED}✗ failed${NC}"
                ((FAILED++))
                echo "[$(date)] Failed: $rel_path" >> "$ERROR_LOG"
                ;;
        esac

        # Validate every N files
        if [ $batch_count -ge $VALIDATION_INTERVAL ] && [ $batch_count -gt 0 ]; then
            echo ""
            print_progress

            if ! validate_graph; then
                print_error "Validation failed after processing batch"
                print_warning "Stopping to allow manual review"
                print_info "Check $ERROR_LOG for details"
                save_progress
                return 1
            fi

            batch_count=0
            save_progress
        fi

        # Save progress periodically
        if [ $((PROCESSED % 10)) -eq 0 ]; then
            echo "$PROCESSED" > "$PROGRESS_FILE"
        fi
    done

    # Final validation
    echo ""
    print_header "Running final validation"
    if validate_graph; then
        print_success "All validations passed!"
    else
        print_error "Final validation failed"
        return 1
    fi
}

# Print summary
print_summary() {
    echo ""
    echo "═══════════════════════════════════════════════════"
    print_header "BATCH ANNOTATION COMPLETE"
    echo "═══════════════════════════════════════════════════"
    echo ""
    echo "  Total files:        $TOTAL_FILES"
    print_success "  Annotated:          $ANNOTATED"
    print_warning "  Skipped:            $SKIPPED"
    [ $FAILED -gt 0 ] && print_error "  Failed:             $FAILED"
    echo ""
    echo "  Validation runs:    $VALIDATION_RUNS"
    [ $VALIDATION_ERRORS -gt 0 ] && print_error "  Validation errors:  $VALIDATION_ERRORS"
    echo ""
    print_info "Logs saved to: $LOG_DIR"
    [ -s "$ERROR_LOG" ] && print_warning "Errors logged to: $ERROR_LOG"
    echo "═══════════════════════════════════════════════════"
    echo ""
}

# Cleanup on exit
cleanup() {
    save_progress
    print_summary
}

trap cleanup EXIT

# Main
main() {
    init

    # Check dependencies
    if ! command -v graphfs &> /dev/null; then
        print_error "graphfs command not found"
        print_info "Please ensure GraphFS is installed and in PATH"
        exit 1
    fi

    if ! [ -x /Users/coolbeans/Development/dev/graphfs/scripts/dox ]; then
        print_error "dox script not found or not executable"
        exit 1
    fi

    # Process all files
    process_all_files

    if [ $FAILED -eq 0 ] && [ $VALIDATION_ERRORS -eq 0 ]; then
        print_success "Batch annotation completed successfully!"
        exit 0
    else
        print_warning "Batch annotation completed with errors"
        exit 1
    fi
}

main "$@"
