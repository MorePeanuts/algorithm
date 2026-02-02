"""Update README files with statistics."""

import re
from pathlib import Path

from .stats import Statistics

STATS_START_MARKER = "<!-- STATS_START -->"
STATS_END_MARKER = "<!-- STATS_END -->"


def generate_stats_section_en(stats: Statistics) -> str:
    """Generate English statistics section content.

    Args:
        stats: Statistics instance

    Returns:
        Markdown content for statistics section
    """
    easy = stats.difficulty_counts.get("Easy", 0)
    medium = stats.difficulty_counts.get("Medium", 0)
    hard = stats.difficulty_counts.get("Hard", 0)

    return f"""{STATS_START_MARKER}
## 📊 Progress

| Total | Easy | Medium | Hard |
|:-----:|:----:|:------:|:----:|
| {stats.total} | {easy} | {medium} | {hard} |

<p align="center">
  <img src="./assets/difficulty_distribution.png" alt="Difficulty Distribution" width="400">
  <img src="./assets/tag_cloud.png" alt="Tag Cloud" width="400">
</p>

{STATS_END_MARKER}"""


def generate_stats_section_zh(stats: Statistics) -> str:
    """Generate Chinese statistics section content.

    Args:
        stats: Statistics instance

    Returns:
        Markdown content for statistics section
    """
    easy = stats.difficulty_counts.get("Easy", 0)
    medium = stats.difficulty_counts.get("Medium", 0)
    hard = stats.difficulty_counts.get("Hard", 0)

    return f"""{STATS_START_MARKER}
## 📊 刷题进度

| 总计 | 简单 | 中等 | 困难 |
|:----:|:----:|:----:|:----:|
| {stats.total} | {easy} | {medium} | {hard} |

<p align="center">
  <img src="./assets/difficulty_distribution.png" alt="难度分布" width="400">
  <img src="./assets/tag_cloud.png" alt="标签云" width="400">
</p>

{STATS_END_MARKER}"""


def update_readme_content(content: str, stats_section: str) -> str:
    """Update README content with new statistics section.

    Args:
        content: Original README content
        stats_section: New statistics section content

    Returns:
        Updated README content
    """
    # Pattern to match existing stats section
    pattern = f"{re.escape(STATS_START_MARKER)}.*?{re.escape(STATS_END_MARKER)}"

    if re.search(pattern, content, re.DOTALL):
        # Replace existing section
        return re.sub(pattern, stats_section, content, flags=re.DOTALL)
    else:
        # Insert after Project Introduction section
        # Look for the next ## heading after Project Introduction
        insert_pattern = r"(## 📖 Project Introduction.*?)(\n## |\n## )"
        insert_pattern_zh = r"(## 📖 项目简介.*?)(\n## |\n## )"

        for pat in [insert_pattern, insert_pattern_zh]:
            match = re.search(pat, content, re.DOTALL)
            if match:
                insert_pos = match.end(1)
                return content[:insert_pos] + "\n\n" + stats_section + "\n" + content[insert_pos:]

        # Fallback: append at the end
        return content + "\n\n" + stats_section


def update_readme(readme_path: Path, stats: Statistics, is_chinese: bool = False) -> None:
    """Update a README file with statistics.

    Args:
        readme_path: Path to the README file
        stats: Statistics instance
        is_chinese: Whether this is the Chinese version
    """
    if not readme_path.exists():
        return

    content = readme_path.read_text(encoding="utf-8")

    if is_chinese:
        stats_section = generate_stats_section_zh(stats)
    else:
        stats_section = generate_stats_section_en(stats)

    updated_content = update_readme_content(content, stats_section)
    readme_path.write_text(updated_content, encoding="utf-8")


def update_all_readmes(project_root: Path, stats: Statistics) -> list[Path]:
    """Update all README files with statistics.

    Args:
        project_root: Path to project root directory
        stats: Statistics instance

    Returns:
        List of updated file paths
    """
    updated = []

    readme_en = project_root / "README.md"
    if readme_en.exists():
        update_readme(readme_en, stats, is_chinese=False)
        updated.append(readme_en)

    readme_zh = project_root / "README_zh.md"
    if readme_zh.exists():
        update_readme(readme_zh, stats, is_chinese=True)
        updated.append(readme_zh)

    return updated
