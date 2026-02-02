#!/bin/bash
# Sync English docs from Chinese docs
# Crawls English version for docs that exist in docs-zh but not in docs

set -e

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PROJECT_ROOT="$(cd "$SCRIPT_DIR/../.." && pwd)"

DOCS_ZH="$PROJECT_ROOT/docs-zh/leetcode"
DOCS_EN="$PROJECT_ROOT/docs/leetcode"

# Find all Chinese docs
for zh_file in $(find "$DOCS_ZH" -name "*.md" | sort); do
    # Get relative path
    rel_path="${zh_file#$DOCS_ZH/}"
    en_file="$DOCS_EN/$rel_path"

    # Skip if English doc already exists
    if [[ -f "$en_file" ]]; then
        echo "Skip: $rel_path (already exists)"
        continue
    fi

    # Extract link from frontmatter
    link=$(grep -m1 "^link:" "$zh_file" | sed 's/^link: *//')

    if [[ -z "$link" ]]; then
        echo "Skip: $rel_path (no link found)"
        continue
    fi

    echo "Crawling: $rel_path"
    echo "  Link: $link"

    # Run crawler for English only
    (cd "$PROJECT_ROOT" && uv run --package crawler crawler "$link" --lang en)

    echo ""
done

echo "Done!"
