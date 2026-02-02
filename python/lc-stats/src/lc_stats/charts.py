"""Generate visualization charts for statistics."""

import platform
from pathlib import Path

import matplotlib.pyplot as plt
from wordcloud import WordCloud

from .stats import Statistics

# Color scheme for difficulty levels
DIFFICULTY_COLORS = {
    "Easy": "#00b8a3",  # Green (LeetCode style)
    "Medium": "#ffc01e",  # Yellow/Orange
    "Hard": "#ff375f",  # Red
}


def get_chinese_font_path() -> str | None:
    """Get a system font path that supports Chinese characters.

    Returns:
        Path to a Chinese font, or None if not found
    """
    system = platform.system()

    if system == "Darwin":  # macOS
        candidates = [
            "/System/Library/Fonts/PingFang.ttc",
            "/System/Library/Fonts/STHeiti Light.ttc",
            "/System/Library/Fonts/Hiragino Sans GB.ttc",
        ]
    elif system == "Windows":
        candidates = [
            "C:/Windows/Fonts/msyh.ttc",  # Microsoft YaHei
            "C:/Windows/Fonts/simsun.ttc",  # SimSun
        ]
    else:  # Linux
        candidates = [
            "/usr/share/fonts/truetype/wqy/wqy-microhei.ttc",
            "/usr/share/fonts/opentype/noto/NotoSansCJK-Regular.ttc",
        ]

    for font_path in candidates:
        if Path(font_path).exists():
            return font_path

    return None


def generate_pie_chart(stats: Statistics, output_path: Path) -> None:
    """Generate a pie chart for difficulty distribution.

    Args:
        stats: Statistics instance
        output_path: Path to save the chart image
    """
    # Filter out zero counts
    labels = []
    sizes = []
    colors = []

    for difficulty in ["Easy", "Medium", "Hard"]:
        count = stats.difficulty_counts.get(difficulty, 0)
        if count > 0:
            labels.append(f"{difficulty}\n({count})")
            sizes.append(count)
            colors.append(DIFFICULTY_COLORS[difficulty])

    if not sizes:
        return

    # Create figure with transparent background
    fig, ax = plt.subplots(figsize=(6, 6), facecolor="none")
    ax.set_facecolor("none")

    wedges, texts, autotexts = ax.pie(
        sizes,
        labels=labels,
        colors=colors,
        autopct="%1.1f%%",
        startangle=90,
        textprops={"fontsize": 12},
    )

    # Style the percentage text
    for autotext in autotexts:
        autotext.set_color("white")
        autotext.set_fontweight("bold")

    ax.set_title(f"Difficulty Distribution (Total: {stats.total})", fontsize=14, pad=20)

    # Save with tight layout
    output_path.parent.mkdir(parents=True, exist_ok=True)
    plt.savefig(output_path, dpi=150, bbox_inches="tight", transparent=True)
    plt.close()


def generate_word_cloud(stats: Statistics, output_path: Path) -> None:
    """Generate a word cloud for tag distribution.

    Args:
        stats: Statistics instance
        output_path: Path to save the word cloud image
    """
    if not stats.tag_counts:
        return

    font_path = get_chinese_font_path()

    # Create word cloud
    wc = WordCloud(
        width=800,
        height=400,
        background_color=None,
        mode="RGBA",
        font_path=font_path,
        max_words=100,
        min_font_size=10,
        max_font_size=80,
        colormap="viridis",
        prefer_horizontal=0.7,
    )

    wc.generate_from_frequencies(stats.tag_counts)

    # Save the word cloud
    output_path.parent.mkdir(parents=True, exist_ok=True)
    wc.to_file(str(output_path))


def generate_all_charts(stats: Statistics, assets_dir: Path) -> dict[str, Path]:
    """Generate all charts and save to assets directory.

    Args:
        stats: Statistics instance
        assets_dir: Path to assets directory

    Returns:
        Dict mapping chart type to output path
    """
    outputs = {}

    pie_path = assets_dir / "difficulty_distribution.png"
    generate_pie_chart(stats, pie_path)
    outputs["pie"] = pie_path

    cloud_path = assets_dir / "tag_cloud.png"
    generate_word_cloud(stats, cloud_path)
    outputs["cloud"] = cloud_path

    return outputs
