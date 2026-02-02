"""Generate visualization charts for statistics."""

import platform
from pathlib import Path

import matplotlib.pyplot as plt
from matplotlib import font_manager
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


def setup_chinese_font() -> None:
    """Setup matplotlib to use Chinese font."""
    font_path = get_chinese_font_path()
    if font_path:
        font_manager.fontManager.addfont(font_path)
        prop = font_manager.FontProperties(fname=font_path)
        plt.rcParams["font.family"] = prop.get_name()


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

    setup_chinese_font()

    # Create figure with transparent background
    fig, ax = plt.subplots(figsize=(5, 5), facecolor="none")
    ax.set_facecolor("none")

    wedges, texts, autotexts = ax.pie(
        sizes,
        labels=labels,
        colors=colors,
        autopct="%1.1f%%",
        startangle=90,
        textprops={"fontsize": 11},
    )

    # Style the percentage text
    for autotext in autotexts:
        autotext.set_color("white")
        autotext.set_fontweight("bold")

    ax.set_title(f"Difficulty Distribution", fontsize=13, pad=15)

    # Save with tight layout
    output_path.parent.mkdir(parents=True, exist_ok=True)
    plt.savefig(output_path, dpi=150, bbox_inches="tight", transparent=True)
    plt.close()


def generate_bar_chart(stats: Statistics, output_path: Path, top_n: int = 10) -> None:
    """Generate a horizontal bar chart for top N tags.

    Args:
        stats: Statistics instance
        output_path: Path to save the chart image
        top_n: Number of top tags to show
    """
    if not stats.tag_counts:
        return

    setup_chinese_font()

    # Get top N tags sorted by count
    sorted_tags = sorted(stats.tag_counts.items(), key=lambda x: x[1], reverse=True)[:top_n]

    # Reverse for horizontal bar chart (top item at top)
    sorted_tags = sorted_tags[::-1]

    tags = [item[0] for item in sorted_tags]
    counts = [item[1] for item in sorted_tags]

    # Create figure
    fig, ax = plt.subplots(figsize=(6, 5), facecolor="none")
    ax.set_facecolor("none")

    # Create horizontal bar chart with gradient colors
    colors = plt.cm.viridis([i / len(tags) for i in range(len(tags))])
    bars = ax.barh(tags, counts, color=colors)

    # Add count labels on bars
    for bar, count in zip(bars, counts):
        ax.text(
            bar.get_width() + 0.1,
            bar.get_y() + bar.get_height() / 2,
            str(count),
            va="center",
            fontsize=10,
        )

    ax.set_xlabel("Count", fontsize=11)
    ax.set_title(f"Top {len(tags)} Tags", fontsize=13, pad=15)
    ax.spines["top"].set_visible(False)
    ax.spines["right"].set_visible(False)

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
        width=900,
        height=350,
        background_color=None,
        mode="RGBA",
        font_path=font_path,
        max_words=100,
        min_font_size=12,
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

    bar_path = assets_dir / "top_tags.png"
    generate_bar_chart(stats, bar_path)
    outputs["bar"] = bar_path

    cloud_path = assets_dir / "tag_cloud.png"
    generate_word_cloud(stats, cloud_path)
    outputs["cloud"] = cloud_path

    return outputs
