"""Parse frontmatter from LeetCode markdown files."""

import re
from dataclasses import dataclass
from pathlib import Path

import yaml

# Difficulty level mapping for English docs
DIFFICULTY_LEVELS_EN = {"Easy", "Medium", "Hard"}

# Difficulty level mapping for Chinese docs (Chinese to English)
DIFFICULTY_MAP_ZH = {
    "简单": "Easy",
    "中等": "Medium",
    "困难": "Hard",
}

# Embed code pattern (e.g., ```embed-go, ```embed-python, etc.)
EMBED_CODE_PATTERN = re.compile(r"```embed-\w+")


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


def has_embed_code(content: str) -> bool:
    """Check if the document contains embedded code references.

    Args:
        content: Markdown file content

    Returns:
        True if embed code block exists (e.g., ```embed-go)
    """
    return bool(EMBED_CODE_PATTERN.search(content))


def normalize_tag(tag: str, is_chinese: bool) -> str:
    """Normalize a tag for display.

    Args:
        tag: Raw tag from frontmatter
        is_chinese: Whether this is a Chinese document

    Returns:
        Normalized tag (underscores replaced with spaces for English)
    """
    if is_chinese:
        return tag
    # Replace underscores with spaces for English tags
    return tag.replace("_", " ")


def parse_problem(filepath: Path, is_chinese: bool = False) -> Problem | None:
    """Parse a single markdown file and extract problem metadata.

    Args:
        filepath: Path to the markdown file
        is_chinese: Whether this is a Chinese document

    Returns:
        Problem instance, or None if parsing fails or no solution
    """
    try:
        content = filepath.read_text(encoding="utf-8")
    except OSError:
        return None

    # Check if the problem has embedded code (meaning it's solved)
    if not has_embed_code(content):
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

    if is_chinese:
        # Chinese docs: difficulty is in Chinese
        for tag in tags:
            if tag in DIFFICULTY_MAP_ZH:
                difficulty = DIFFICULTY_MAP_ZH[tag]
            else:
                other_tags.append(normalize_tag(tag, is_chinese))
    else:
        # English docs: difficulty is in English
        for tag in tags:
            if tag in DIFFICULTY_LEVELS_EN:
                difficulty = tag
            else:
                other_tags.append(normalize_tag(tag, is_chinese))

    if not difficulty:
        return None

    return Problem(
        filename=filepath.name,
        difficulty=difficulty,
        tags=other_tags,
    )


def scan_problems(docs_dir: Path, is_chinese: bool = False) -> list[Problem]:
    """Scan all LeetCode problem markdown files.

    Args:
        docs_dir: Path to docs/leetcode or docs-zh/leetcode directory
        is_chinese: Whether this is the Chinese docs directory

    Returns:
        List of parsed Problem instances (only those with embedded code)
    """
    problems = []

    if not docs_dir.exists():
        return problems

    for md_file in docs_dir.rglob("*.md"):
        problem = parse_problem(md_file, is_chinese)
        if problem:
            problems.append(problem)

    return problems
