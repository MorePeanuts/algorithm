# lc-stats

LeetCode progress statistics generator with visualizations.

## Features

- Scans `docs/leetcode/` (English) and `docs-zh/leetcode/` (Chinese) separately
- Generates difficulty distribution pie chart
- Generates top 10 tags horizontal bar chart
- Generates tag word cloud
- Updates README.md and README_zh.md automatically
- Supports light/dark theme for GitHub

## Usage

```bash
uv run --package lc-stats lc-stats
```

## Technical Highlights

### 1. Safe README Modification with Marker Comments

The tool uses HTML comment markers to safely update README files without causing conflicts during multiple runs:

```markdown
<!-- STATS_START -->
## 📊 LeetCode Progress
... (auto-generated content)
<!-- STATS_END -->
```

**How it works:**

```python
STATS_START_MARKER = "<!-- STATS_START -->"
STATS_END_MARKER = "<!-- STATS_END -->"

def update_readme_content(content: str, stats_section: str) -> str:
    # Use regex to find and replace content between markers
    pattern = f"{re.escape(STATS_START_MARKER)}.*?{re.escape(STATS_END_MARKER)}"

    if re.search(pattern, content, re.DOTALL):
        # Replace existing section
        return re.sub(pattern, stats_section, content, flags=re.DOTALL)
    else:
        # Insert new section at appropriate location
        ...
```

This approach ensures:
- **Idempotent updates**: Running the tool multiple times produces the same result
- **No content corruption**: Only the marked section is modified
- **Flexible insertion**: If markers don't exist, the tool inserts them at a logical location

### 2. Automatic Trigger via Git Post-Commit Hook

The tool can be automatically triggered when committing changes that add new LeetCode solutions.

**Setup:**

```bash
# Configure git to use custom hooks directory
git config core.hooksPath .githooks
```

**How `.githooks/post-commit` works:**

```bash
#!/bin/bash
# Get the commit message
commit_msg=$(git log -1 --pretty=%B)

# Check if message contains "Add leetcode" (case insensitive)
if echo "$commit_msg" | grep -iq "Add leetcode"; then
    echo "Detected LeetCode addition, updating statistics..."

    # Run the stats generator
    uv run --package lc-stats lc-stats

    # Amend the commit with updated files
    git add assets/stats/*.png README.md README_zh.md
    git commit --amend --no-edit
fi
```

**Key points:**
- Uses `grep -iq` for case-insensitive matching
- Checks if `uv` is available before running
- Uses `git commit --amend --no-edit` to include generated files in the same commit
- The hook runs **after** the commit, then amends it with updated statistics

### 3. Embed Code Detection for Solved Problems

Instead of checking for specific section headers, the tool detects if a problem is solved by looking for embedded code references:

```python
# Pattern matches ```embed-go, ```embed-python, etc.
EMBED_CODE_PATTERN = re.compile(r"```embed-\w+")

def has_embed_code(content: str) -> bool:
    return bool(EMBED_CODE_PATTERN.search(content))
```

This is more reliable because:
- A problem document might exist but not have a solution yet
- The presence of embedded code definitively indicates a solution exists

### 4. Theme-Aware Chart Generation

Charts are generated in both light and dark variants for GitHub's theme switching:

```python
def generate_all_charts(stats, assets_dir, lang="en"):
    for theme in ["light", "dark"]:
        suffix = f"_{lang}_{theme}"
        # Generate: difficulty_distribution_en_light.png, etc.
```

README uses GitHub's `<picture>` element for automatic theme detection:

```html
<picture>
  <source media="(prefers-color-scheme: dark)" srcset="./assets/stats/chart_en_dark.png">
  <source media="(prefers-color-scheme: light)" srcset="./assets/stats/chart_en_light.png">
  <img src="./assets/stats/chart_en_light.png" alt="Chart">
</picture>
```

## Project Structure

```
python/lc-stats/
├── pyproject.toml
├── README.md
└── src/lc_stats/
    ├── __init__.py      # Main entry point
    ├── parser.py        # Frontmatter & embed code parser
    ├── stats.py         # Statistics calculator
    ├── charts.py        # Chart generators (pie, bar, wordcloud)
    └── readme.py        # README updater with marker system
```
