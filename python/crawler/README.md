# Crawler

Fetch coding problems and generate markdown docs with Go solution templates.

## Usage

```bash
uv run crawler <url>

# Examples
uv run crawler https://leetcode.cn/problems/two-sum/
uv run crawler https://leetcode.com/problems/add-two-numbers/
```

## Supported Sites

- LeetCode CN (`leetcode.cn`) - generates Chinese content
- LeetCode (`leetcode.com`) - generates English content

## Generated Files

- `docs/leetcode/XXXX-XXXX/XXXX_slug.md` - Problem description with tags
- `leetcode/XXXX-XXXX/XXXX_slug/solution.go` - Go solution template
