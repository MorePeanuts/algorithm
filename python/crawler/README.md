# Crawler

Fetch coding problems and generate markdown docs with Go solution templates.

## Usage

```bash
uv run --package crawler crawler <url> [--lang zh|en]

# Generate both Chinese and English docs
uv run --package crawler crawler https://leetcode.cn/problems/two-sum/

# Generate Chinese doc only
uv run --package crawler crawler https://leetcode.cn/problems/two-sum/ --lang zh

# Generate English doc only
uv run --package crawler crawler https://leetcode.com/problems/add-two-numbers/ --lang en
```

## Options

- `--lang zh` - Generate Chinese doc only (to `docs-zh/`)
- `--lang en` - Generate English doc only (to `docs/`)
- No `--lang` - Generate both Chinese and English docs

## Generated Files

- `docs-zh/leetcode/XXXX-XXXX/XXXX_slug.md` - Chinese problem description
- `docs/leetcode/XXXX-XXXX/XXXX_slug.md` - English problem description
- `leetcode/XXXX-XXXX/XXXX_slug/solution.go` - Go solution template
