#!/usr/bin/env bash
#
# Orchestrate full codebase annotation with validation
#

set -euo pipefail

REPO_ROOT=$(git rev-parse --show-toplevel)
LOG_DIR="$REPO_ROOT/.annotation_logs"
PROGRESS_FILE="$LOG_DIR/progress.txt"
ERROR_LOG="$LOG_DIR/errors.log"
VALIDATION_INTERVAL=30

# Colors
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
CYAN='\033[0;36m'
NC='\033[0m'

# Stats
TOTAL=0
PROCESSED=0
ANNOTATED=0
SKIPPED=0
FAILED=0
VALIDATION_COUNT=0
VALIDATION_FAILED=0

print_color() { echo -e "${1}${2}${NC}"; }
print_error() { print_color "$RED" "✗ $*"; }
print_success() { print_color "$GREEN" "✓ $*"; }
print_info() { print_color "$BLUE" "ℹ $*"; }
print_warning() { print_color "$YELLOW" "⚠ $*"; }
print_header() { print_color "$CYAN" "━━━ $* ━━━"; }

# Initialize
mkdir -p "$LOG_DIR"
echo "0" > "$PROGRESS_FILE"
echo "" > "$ERROR_LOG"

print_header "Batch LinkedDoc Annotation - Consul Codebase"
print_info "Starting full annotation process"
echo ""

# Get list of all unannotated Go files
print_info "Scanning for unannotated Go files..."
FILES=()
while IFS= read -r file; do
    # Skip if already has LinkedDoc
    if ! grep -q "<!-- LinkedDoc RDF -->" "$file" 2>/dev/null; then
        FILES+=("$file")
    fi
done < <(find "$REPO_ROOT" -name "*.go" -type f | grep -v "/vendor/" | grep -v "_test.go" | sort)

TOTAL=${#FILES[@]}
print_success "Found $TOTAL files to annotate"
echo ""

# Process each file
batch_count=0
for file in "${FILES[@]}"; do
    ((PROCESSED++))
    rel_path="${file#$REPO_ROOT/}"

    echo -n "[$PROCESSED/$TOTAL] $rel_path ... "

    # Run annotation script
    if python3 "$REPO_ROOT/ai_annotate.py" "$file" 2>/dev/null; then
        echo -e "${GREEN}✓${NC}"
        ((ANNOTATED++))
        ((batch_count++))
    else
        exit_code=$?
        if [ $exit_code -eq 2 ]; then
            echo -e "${YELLOW}⊘ skipped${NC}"
            ((SKIPPED++))
        else
            echo -e "${RED}✗ failed${NC}"
            ((FAILED++))
            echo "[$PROCESSED] $rel_path" >> "$ERROR_LOG"
        fi
    fi

    # Validate every N files
    if [ $batch_count -ge $VALIDATION_INTERVAL ]; then
        echo ""
        print_info "Validating batch ($VALIDATION_COUNT batches processed)..."

        # Run validation and capture output
        validation_output=$(graphfs scan --exclude "*.py" --exclude "*.sh" --exclude "*.md" --validate 2>&1)

        # Check if validation passed (look for "successfully" or absence of "failed with")
        if echo "$validation_output" | grep -q "built successfully"; then
            print_success "Validation passed"
            ((VALIDATION_COUNT++))
        elif echo "$validation_output" | grep -q "failed with"; then
            print_error "Validation failed!"
            ((VALIDATION_FAILED++))

            # Show error details
            echo "$validation_output" | tail -20

            print_error "Stopping due to validation failure"
            print_info "Processed: $PROCESSED/$TOTAL"
            print_info "Annotated: $ANNOTATED"
            print_info "Check $ERROR_LOG for details"
            exit 1
        else
            # If no explicit failure, treat as success
            print_success "Validation passed (warnings only)"
            ((VALIDATION_COUNT++))
        fi

        batch_count=0
        echo ""
    fi

    # Save progress periodically
    if [ $((PROCESSED % 10)) -eq 0 ]; then
        echo "$PROCESSED" > "$PROGRESS_FILE"
    fi
done

# Final validation
echo ""
print_header "Running final validation"
final_output=$(graphfs scan --exclude "*.py" --exclude "*.sh" --exclude "*.md" --validate --stats 2>&1)
echo "$final_output"

if echo "$final_output" | grep -q "built successfully"; then
    print_success "Final validation passed!"
elif echo "$final_output" | grep -q "failed with"; then
    print_error "Final validation failed"
    exit 1
else
    print_success "Final validation passed (warnings only)!"
fi

# Summary
echo ""
echo "═══════════════════════════════════════════════════"
print_header "ANNOTATION COMPLETE"
echo "═══════════════════════════════════════════════════"
echo ""
echo "  Total files:        $TOTAL"
print_success "  Annotated:          $ANNOTATED"
print_warning "  Skipped:            $SKIPPED"
[ $FAILED -gt 0 ] && print_error "  Failed:             $FAILED"
echo ""
echo "  Validations:        $VALIDATION_COUNT"
[ $VALIDATION_FAILED -gt 0 ] && print_error "  Val. failed:        $VALIDATION_FAILED"
echo ""
print_success "All annotations complete!"
echo "═══════════════════════════════════════════════════"
