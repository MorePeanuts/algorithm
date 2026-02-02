"""LeetCode progress statistics generator."""

from pathlib import Path

from .charts import generate_all_charts
from .parser import scan_problems
from .readme import update_all_readmes
from .stats import calculate_stats


def get_project_root() -> Path:
    """Get the project root directory.

    Returns:
        Path to project root (where README.md is located)
    """
    return Path(__file__).parent.parent.parent.parent.parent


def main() -> None:
    """Main entry point for lc-stats."""
    project_root = get_project_root()
    docs_dir = project_root / "docs" / "leetcode"
    assets_dir = project_root / "assets"

    print("Scanning LeetCode problems...")
    problems = scan_problems(docs_dir)
    print(f"Found {len(problems)} problems")

    if not problems:
        print("No problems found. Exiting.")
        return

    print("Calculating statistics...")
    stats = calculate_stats(problems)

    print(f"  Total: {stats.total}")
    print(f"  Easy: {stats.difficulty_counts['Easy']}")
    print(f"  Medium: {stats.difficulty_counts['Medium']}")
    print(f"  Hard: {stats.difficulty_counts['Hard']}")
    print(f"  Unique tags: {len(stats.tag_counts)}")

    print("Generating charts...")
    charts = generate_all_charts(stats, assets_dir)
    for chart_type, path in charts.items():
        print(f"  Generated: {path.relative_to(project_root)}")

    print("Updating README files...")
    updated = update_all_readmes(project_root, stats)
    for path in updated:
        print(f"  Updated: {path.relative_to(project_root)}")

    print("Done!")
