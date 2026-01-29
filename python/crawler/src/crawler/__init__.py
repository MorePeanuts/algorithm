"""Problem crawler for generating markdown docs and Go templates."""

import argparse
import sys
from pathlib import Path

import httpx

from . import leetcode
from .utils import generate_go_template, generate_markdown, get_directory_range, merge_existing_content


def get_project_root() -> Path:
    """Get the project root directory."""
    return Path(__file__).parent.parent.parent.parent.parent


def crawl_leetcode(url: str) -> None:
    """Crawl a LeetCode problem and generate files.

    Args:
        url: LeetCode problem URL
    """
    slug, site = leetcode.parse_url(url)
    print(f"Fetching problem: {slug} from {site}")

    raw_problem = leetcode.fetch_problem(slug, site)
    use_chinese = site == "leetcode.cn"
    problem = leetcode.normalize_problem(raw_problem, use_chinese=use_chinese)

    print(f"Found: {problem['id']}. {problem['title']}")

    project_root = get_project_root()
    dir_range = get_directory_range(problem["id"])
    file_name = f"{problem['id']:04d}_{problem['slug'].replace('-', '_')}"

    # Generate markdown doc
    doc_dir = project_root / "docs" / "leetcode" / dir_range
    doc_dir.mkdir(parents=True, exist_ok=True)
    doc_path = doc_dir / f"{file_name}.md"

    markdown_content = generate_markdown(problem)
    markdown_content = merge_existing_content(markdown_content, doc_path)
    doc_path.write_text(markdown_content, encoding="utf-8")
    print(f"Generated doc: {doc_path.relative_to(project_root)}")

    # Generate Go template
    go_dir = project_root / "leetcode" / dir_range / file_name
    go_path = go_dir / "solution.go"

    if not go_path.exists():
        go_dir.mkdir(parents=True, exist_ok=True)
        go_content = generate_go_template(problem)
        go_path.write_text(go_content, encoding="utf-8")
        print(f"Generated Go template: {go_path.relative_to(project_root)}")
    else:
        print(f"Go file already exists: {go_path.relative_to(project_root)}")


def main() -> None:
    """Main entry point for the crawler."""
    parser = argparse.ArgumentParser(
        description="Crawl coding problems and generate docs"
    )
    parser.add_argument("url", help="Problem URL (LeetCode)")
    args = parser.parse_args()

    url = args.url

    try:
        if "leetcode" in url:
            crawl_leetcode(url)
        else:
            print(f"Error: Unsupported URL: {url}", file=sys.stderr)
            sys.exit(1)

        print("Done!")

    except ValueError as e:
        print(f"Error: {e}", file=sys.stderr)
        sys.exit(1)
    except httpx.HTTPError as e:
        print(f"HTTP Error: {e}", file=sys.stderr)
        sys.exit(1)
    except RuntimeError as e:
        print(f"Error: {e}", file=sys.stderr)
        sys.exit(1)
