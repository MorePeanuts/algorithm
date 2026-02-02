"""Parse frontmatter from LeetCode markdown files."""

import re
from dataclasses import dataclass
from pathlib import Path

import yaml

# Difficulty level mapping (Chinese to English)
DIFFICULTY_MAP = {
    "简单": "Easy",
    "中等": "Medium",
    "困难": "Hard",
}

DIFFICULTY_LEVELS = {"简单", "中等", "困难", "Easy", "Medium", "Hard"}


@dataclass
class Problem:
    """Represents a parsed LeetCode problem."""

    filename: str
    difficulty: str
    tags: list[str]


def extract_frontmatter(content: str) -> dict | None:
    """Extract YAML frontmatter from markdown content.

    Args:
        content: Markdown file content

    Returns:
        Parsed frontmatter as dict, or None if not found
    """
    pattern = r"^---\s*\n(.*?)\n---"
    match = re.match(pattern, content, re.DOTALL)
    if match:
        try:
            return yaml.safe_load(match.group(1))
        except yaml.YAMLError:
            return None
    return None


def parse_problem(filepath: Path) -> Problem | None:
    """Parse a single markdown file and extract problem metadata.

    Args:
        filepath: Path to the markdown file

    Returns:
        Problem instance, or None if parsing fails
    """
    try:
        content = filepath.read_text(encoding="utf-8")
    except OSError:
        return None

    frontmatter = extract_frontmatter(content)
    if not frontmatter or "tags" not in frontmatter:
        return None

    tags = frontmatter["tags"]
    if not isinstance(tags, list):
        return None

    # Separate difficulty from other tags
    difficulty = None
    other_tags = []

    for tag in tags:
        if tag in DIFFICULTY_LEVELS:
            difficulty = DIFFICULTY_MAP.get(tag, tag)
        else:
            other_tags.append(tag)

    if not difficulty:
        return None

    return Problem(
        filename=filepath.name,
        difficulty=difficulty,
        tags=other_tags,
    )


def scan_problems(docs_dir: Path) -> list[Problem]:
    """Scan all LeetCode problem markdown files.

    Args:
        docs_dir: Path to docs/leetcode directory

    Returns:
        List of parsed Problem instances
    """
    problems = []

    if not docs_dir.exists():
        return problems

    for md_file in docs_dir.rglob("*.md"):
        problem = parse_problem(md_file)
        if problem:
            problems.append(problem)

    return problems
