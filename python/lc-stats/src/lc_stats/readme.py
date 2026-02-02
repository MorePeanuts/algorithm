"""Update README files with statistics."""

import re
from pathlib import Path

from .stats import Statistics

STATS_START_MARKER = "<!-- STATS_START -->"
STATS_END_MARKER = "<!-- STATS_END -->"

# Chinese to English tag mapping for common tags
TAG_TRANSLATION = {
    "数组": "Array",
    "哈希表": "Hash Table",
    "排序": "Sorting",
    "字符串": "String",
    "链表": "Linked List",
    "树": "Tree",
    "二叉树": "Binary Tree",
    "动态规划": "Dynamic Programming",
    "贪心": "Greedy",
    "双指针": "Two Pointers",
    "栈": "Stack",
    "队列": "Queue",
    "堆（优先队列）": "Heap",
    "图": "Graph",
    "深度优先搜索": "DFS",
    "广度优先搜索": "BFS",
    "二分查找": "Binary Search",
    "回溯": "Backtracking",
    "分治": "Divide and Conquer",
    "位运算": "Bit Manipulation",
    "数学": "Math",
    "计数": "Counting",
    "桶排序": "Bucket Sort",
    "快速选择": "Quick Select",
    "矩阵": "Matrix",
}

# Difficulty name mapping
DIFFICULTY_NAME_EN = {"Easy": "Easy", "Medium": "Medium", "Hard": "Hard"}
DIFFICULTY_NAME_ZH = {"Easy": "简单", "Medium": "中等", "Hard": "困难"}


def get_dominant_difficulty(stats: Statistics) -> str:
    """Get the difficulty level with the most problems.

    Args:
        stats: Statistics instance

    Returns:
        The dominant difficulty level name
    """
    max_count = 0
    dominant = "Medium"  # Default fallback

    for difficulty, count in stats.difficulty_counts.items():
        if count > max_count:
            max_count = count
            dominant = difficulty

    return dominant


def translate_tag(tag: str) -> str:
    """Translate a Chinese tag to English.

    Args:
        tag: Tag name (possibly in Chinese)

    Returns:
        English tag name
    """
    return TAG_TRANSLATION.get(tag, tag)


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

    # Get dominant difficulty
    dominant = get_dominant_difficulty(stats)
    dominant_name = DIFFICULTY_NAME_EN[dominant]

    # Get top tags for description (translated to English)
    sorted_tags = sorted(stats.tag_counts.items(), key=lambda x: x[1], reverse=True)
    top_3_tags = [translate_tag(tag) for tag, _ in sorted_tags[:3]]
    top_tags_str = ", ".join(top_3_tags) if top_3_tags else "various topics"

    return f"""{STATS_START_MARKER}
## 📊 LeetCode Progress

Currently, this repository contains solutions to **{stats.total}** LeetCode problems, covering a range of difficulty levels and algorithmic topics.

| Total | Easy | Medium | Hard |
|:-----:|:----:|:------:|:----:|
| {stats.total} | {easy} | {medium} | {hard} |

The pie chart below shows the distribution of problems by difficulty level, while the bar chart highlights the top 10 most frequently encountered tags. The majority of problems fall into the **{dominant_name}** category, with a focus on fundamental topics like {top_tags_str}.

<p align="center">
  <img src="./assets/difficulty_distribution.png" alt="Difficulty Distribution" width="380">
  <img src="./assets/top_tags.png" alt="Top 10 Tags" width="420">
</p>

The word cloud below provides a visual overview of all tags covered in this repository, with larger words indicating more frequently practiced topics.

<p align="center">
  <img src="./assets/tag_cloud.png" alt="Tag Cloud" width="700">
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

    # Get dominant difficulty
    dominant = get_dominant_difficulty(stats)
    dominant_name = DIFFICULTY_NAME_ZH[dominant]

    # Get top tags for description
    sorted_tags = sorted(stats.tag_counts.items(), key=lambda x: x[1], reverse=True)
    top_3_tags = [tag for tag, _ in sorted_tags[:3]]
    top_tags_str = "、".join(top_3_tags) if top_3_tags else "多种主题"

    return f"""{STATS_START_MARKER}
## 📊 LeetCode 解题进度

目前，本仓库已收录 **{stats.total}** 道 LeetCode 题目的解答，涵盖不同难度级别和多种算法主题。

| 总计 | 简单 | 中等 | 困难 |
|:----:|:----:|:----:|:----:|
| {stats.total} | {easy} | {medium} | {hard} |

下方饼图展示了题目的难度分布，条形图则列出了出现频率最高的 10 个标签。可以看到，大部分题目集中在**{dominant_name}**难度，重点练习了{top_tags_str}等基础主题。

<p align="center">
  <img src="./assets/difficulty_distribution.png" alt="难度分布" width="380">
  <img src="./assets/top_tags.png" alt="Top 10 标签" width="420">
</p>

下方词云图直观呈现了本仓库涉及的所有标签，字体越大表示该主题的练习频率越高。

<p align="center">
  <img src="./assets/tag_cloud.png" alt="标签云" width="700">
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
