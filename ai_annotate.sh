#!/usr/bin/env bash
#
# AI-enhanced LinkedDoc annotation generator
# Uses Claude Haiku for semantic understanding and relationship detection
#

set -euo pipefail

# This is a placeholder script that will be called by the main orchestrator
# The actual AI enhancement will be done by Claude Code using the Task tool

FILE="$1"
REPO_ROOT=$(git rev-parse --show-toplevel 2>/dev/null || pwd)

# For now, just use the dox script
/Users/coolbeans/Development/dev/graphfs/scripts/dox --full-dox "$FILE"
