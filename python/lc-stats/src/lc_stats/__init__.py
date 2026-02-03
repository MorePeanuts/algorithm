"""LeetCode progress statistics generator."""

from pathlib import Path

from .charts import generate_all_charts
from .parser import scan_problems
from .readme import update_readme
from .stats import calculate_stats, load_cached_stats, save_stats_cache


def get_project_root() -> Path:
    """Get the project root directory.

    Returns:
        Path to project root (where README.md is located)
    """
    return Path(__file__).parent.parent.parent.parent.parent


def main() -> None:
    """Main entry point for lc-stats."""
    project_root = get_project_root()
    assets_dir = project_root / "assets"

    stats_dir = assets_dir / "stats"

    # Process English docs (docs/leetcode) -> README.md
    docs_en_dir = project_root / "docs" / "leetcode"
    print("Scanning English LeetCode docs...")
    problems_en = scan_problems(docs_en_dir, is_chinese=False)
    print(f"Found {len(problems_en)} solved problems (with embed code)")

    if problems_en:
        stats_en = calculate_stats(problems_en)
        print(f"  Total: {stats_en.total}")
        print(f"  Easy: {stats_en.difficulty_counts['Easy']}")
        print(f"  Medium: {stats_en.difficulty_counts['Medium']}")
        print(f"  Hard: {stats_en.difficulty_counts['Hard']}")
        print(f"  Unique tags: {len(stats_en.tag_counts)}")

        # Check if stats changed before regenerating charts
        cache_path_en = stats_dir / ".stats_cache_en.json"
        cached_stats_en = load_cached_stats(cache_path_en)

        if cached_stats_en == stats_en:
            print("Statistics unchanged, skipping chart generation (English).")
        else:
            print("Generating charts (English)...")
            charts = generate_all_charts(stats_en, assets_dir, lang="en")
            for chart_type, path in charts.items():
                print(f"  Generated: {path.relative_to(project_root)}")
            save_stats_cache(stats_en, cache_path_en)

        print("Updating README.md...")
        readme_en = project_root / "README.md"
        update_readme(readme_en, stats_en, is_chinese=False)
        print("  Updated: README.md")

    # Process Chinese docs (docs-zh/leetcode) -> README_zh.md
    docs_zh_dir = project_root / "docs-zh" / "leetcode"
    print("\nScanning Chinese LeetCode docs...")
    problems_zh = scan_problems(docs_zh_dir, is_chinese=True)
    print(f"Found {len(problems_zh)} solved problems (with embed code)")

    if problems_zh:
        stats_zh = calculate_stats(problems_zh)
        print(f"  Total: {stats_zh.total}")
        print(f"  Easy: {stats_zh.difficulty_counts['Easy']}")
        print(f"  Medium: {stats_zh.difficulty_counts['Medium']}")
        print(f"  Hard: {stats_zh.difficulty_counts['Hard']}")
        print(f"  Unique tags: {len(stats_zh.tag_counts)}")

        # Check if stats changed before regenerating charts
        cache_path_zh = stats_dir / ".stats_cache_zh.json"
        cached_stats_zh = load_cached_stats(cache_path_zh)

        if cached_stats_zh == stats_zh:
            print("Statistics unchanged, skipping chart generation (Chinese).")
        else:
            print("Generating charts (Chinese)...")
            charts = generate_all_charts(stats_zh, assets_dir, lang="zh")
            for chart_type, path in charts.items():
                print(f"  Generated: {path.relative_to(project_root)}")
            save_stats_cache(stats_zh, cache_path_zh)

        print("Updating README_zh.md...")
        readme_zh = project_root / "README_zh.md"
        update_readme(readme_zh, stats_zh, is_chinese=True)
        print("  Updated: README_zh.md")

    if not problems_en and not problems_zh:
        print("No solved problems found. Exiting.")
        return

    print("\nDone!")
