"""Common utilities for problem crawlers."""

import re
from pathlib import Path

from markdownify import markdownify as md


def html_to_markdown(html: str) -> str:
    """Convert HTML content to Markdown.

    Args:
        html: HTML string

    Returns:
        Markdown string
    """
    markdown = md(html, heading_style="ATX", bullets="-")
    markdown = re.sub(r"\n{3,}", "\n\n", markdown)
    return markdown.strip()


def get_directory_range(problem_id: int) -> str:
    """Get directory range string for problem ID.

    Args:
        problem_id: Problem ID

    Returns:
        Directory range string (e.g., '0001-0100')
    """
    start = ((problem_id - 1) // 100) * 100 + 1
    end = start + 99
    return f"{start:04d}-{end:04d}"


def generate_markdown(problem: dict) -> str:
    """Generate markdown content for problem.

    Args:
        problem: Normalized problem dictionary with keys:
            - id, title, slug, difficulty, tags, content, lang

    Returns:
        Markdown content string
    """
    lang = problem.get("lang", "zh")
    slug = problem["slug"]

    # Generate link based on language
    if lang == "zh":
        link = f"https://leetcode.cn/problems/{slug}/"
    else:
        link = f"https://leetcode.com/problems/{slug}/"

    tags = [problem["difficulty"]] + problem["tags"]
    tags_yaml = "\n".join(f"  - {tag}" for tag in tags)
    frontmatter = f"---\nlink: {link}\ntags:\n{tags_yaml}\n---"

    content_md = html_to_markdown(problem["content"])

    if lang == "zh":
        desc_header = "## 题目描述"
        analysis_header = "## 题目解析"
    else:
        desc_header = "## Description"
        analysis_header = "## Solution"

    return f"""{frontmatter}
{desc_header}
{content_md}

{analysis_header}
"""


def get_analysis_marker(lang: str) -> str:
    """Get the analysis section marker for a language.

    Args:
        lang: Language code ("zh" or "en")

    Returns:
        Analysis section header string
    """
    return "## 题目解析" if lang == "zh" else "## Solution"


def generate_go_template(problem: dict) -> str:
    """Generate Go solution template for problem.

    Args:
        problem: Normalized problem dictionary

    Returns:
        Go code string
    """
    problem_id = problem["id"]
    title = problem.get("title_en") or problem["title"]

    return f"""// Package leetcode{problem_id:04d} solves LeetCode {problem_id}. {title}
package leetcode{problem_id:04d}

// TODO: implement solution
"""


def merge_existing_content(new_content: str, existing_path: Path, lang: str = "zh") -> str:
    """Merge new content with existing file, preserving user analysis section.

    Args:
        new_content: New generated content
        existing_path: Path to existing file
        lang: Language code ("zh" or "en")

    Returns:
        Merged content string
    """
    if not existing_path.exists():
        return new_content

    existing = existing_path.read_text(encoding="utf-8")

    analysis_marker = get_analysis_marker(lang)
    if analysis_marker in existing:
        analysis_start = existing.index(analysis_marker)
        existing_analysis = existing[analysis_start:]

        new_analysis_start = new_content.index(analysis_marker)
        return new_content[:new_analysis_start] + existing_analysis

    return new_content
