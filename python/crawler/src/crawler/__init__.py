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


def crawl_leetcode(url: str, lang: str | None = None) -> None:
    """Crawl a LeetCode problem and generate files.

    Args:
        url: LeetCode problem URL
        lang: Language code ("zh", "en", or None for both)
    """
    slug, site = leetcode.parse_url(url)
    print(f"Fetching problem: {slug} from {site}")

    # Always fetch from CN site first (has both Chinese and English content)
    # Fall back to COM site if CN fails
    try:
        raw_problem = leetcode.fetch_problem(slug, "leetcode.cn")
    except RuntimeError:
        print("CN site failed, trying COM site...")
        raw_problem = leetcode.fetch_problem(slug, "leetcode.com")

    # Determine which languages to generate
    langs = [lang] if lang else ["zh", "en"]

    project_root = get_project_root()
    problem_id = int(raw_problem["questionFrontendId"])
    slug_name = raw_problem["titleSlug"]
    dir_range = get_directory_range(problem_id)
    file_name = f"{problem_id:04d}_{slug_name.replace('-', '_')}"

    # Print problem info
    title_zh = raw_problem.get("translatedTitle") or raw_problem["title"]
    title_en = raw_problem["title"]
    if lang == "zh":
        print(f"Found: {problem_id}. {title_zh}")
    elif lang == "en":
        print(f"Found: {problem_id}. {title_en}")
    else:
        print(f"Found: {problem_id}. {title_zh} / {title_en}")

    # Generate markdown docs for each language
    for l in langs:
        problem = leetcode.normalize_problem(raw_problem, lang=l)
        if l == "zh":
            doc_dir = project_root / "docs-zh" / "leetcode" / dir_range
        else:
            doc_dir = project_root / "docs" / "leetcode" / dir_range

        doc_dir.mkdir(parents=True, exist_ok=True)
        doc_path = doc_dir / f"{file_name}.md"

        markdown = generate_markdown(problem)
        markdown = merge_existing_content(markdown, doc_path, lang=l)
        doc_path.write_text(markdown, encoding="utf-8")
        print(f"Generated doc ({l}): {doc_path.relative_to(project_root)}")

    # Generate Go template (use English title)
    go_dir = project_root / "leetcode" / dir_range / file_name
    go_path = go_dir / "solution.go"

    if not go_path.exists():
        go_dir.mkdir(parents=True, exist_ok=True)
        problem_en = leetcode.normalize_problem(raw_problem, lang="en")
        go_content = generate_go_template(problem_en)
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
    parser.add_argument(
        "--lang",
        choices=["zh", "en"],
        default=None,
        help="Language for docs (zh=Chinese, en=English). If not specified, generates both.",
    )
    args = parser.parse_args()

    url = args.url

    try:
        if "leetcode" in url:
            crawl_leetcode(url, lang=args.lang)
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
