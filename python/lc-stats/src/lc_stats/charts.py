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

# Theme configurations
THEMES = {
    "light": {
        "text_color": "#1f2328",
        "axis_color": "#1f2328",
        "grid_color": "#d0d7de",
        "spine_color": "#d0d7de",
    },
    "dark": {
        "text_color": "#e6edf3",
        "axis_color": "#e6edf3",
        "grid_color": "#3d444d",
        "spine_color": "#3d444d",
    },
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


def generate_pie_chart(stats: Statistics, output_path: Path, theme: str = "light") -> None:
    """Generate a pie chart for difficulty distribution.

    Args:
        stats: Statistics instance
        output_path: Path to save the chart image
        theme: Color theme ("light" or "dark")
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
    theme_config = THEMES[theme]

    # Create figure with transparent background
    fig, ax = plt.subplots(figsize=(5, 5), facecolor="none")
    ax.set_facecolor("none")

    wedges, texts, autotexts = ax.pie(
        sizes,
        labels=labels,
        colors=colors,
        autopct="%1.1f%%",
        startangle=90,
        textprops={"fontsize": 11, "color": theme_config["text_color"]},
    )

    # Style the percentage text
    for autotext in autotexts:
        autotext.set_color("white")
        autotext.set_fontweight("bold")

    ax.set_title(
        "Difficulty Distribution",
        fontsize=13,
        pad=15,
        color=theme_config["text_color"],
    )

    # Save with tight layout
    output_path.parent.mkdir(parents=True, exist_ok=True)
    plt.savefig(output_path, dpi=150, bbox_inches="tight", transparent=True)
    plt.close()


def generate_bar_chart(
    stats: Statistics, output_path: Path, top_n: int = 10, theme: str = "light"
) -> None:
    """Generate a horizontal bar chart for top N tags.

    Args:
        stats: Statistics instance
        output_path: Path to save the chart image
        top_n: Number of top tags to show
        theme: Color theme ("light" or "dark")
    """
    if not stats.tag_counts:
        return

    setup_chinese_font()
    theme_config = THEMES[theme]

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
            color=theme_config["text_color"],
        )

    # Style axis
    ax.set_xlabel("Count", fontsize=11, color=theme_config["text_color"])
    ax.set_title(f"Top {len(tags)} Tags", fontsize=13, pad=15, color=theme_config["text_color"])
    ax.tick_params(axis="both", colors=theme_config["axis_color"])

    # Style spines
    for spine in ["top", "right"]:
        ax.spines[spine].set_visible(False)
    for spine in ["bottom", "left"]:
        ax.spines[spine].set_color(theme_config["spine_color"])

    # Save with tight layout
    output_path.parent.mkdir(parents=True, exist_ok=True)
    plt.savefig(output_path, dpi=150, bbox_inches="tight", transparent=True)
    plt.close()


def generate_word_cloud(stats: Statistics, output_path: Path, theme: str = "light") -> None:
    """Generate a word cloud for tag distribution.

    Args:
        stats: Statistics instance
        output_path: Path to save the word cloud image
        theme: Color theme ("light" or "dark")
    """
    if not stats.tag_counts:
        return

    font_path = get_chinese_font_path()

    # Use different colormap for different themes
    colormap = "viridis" if theme == "light" else "plasma"

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
        colormap=colormap,
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

    # Generate charts for both themes in stats subdirectory
    stats_dir = assets_dir / "stats"
    stats_dir.mkdir(parents=True, exist_ok=True)

    for theme in ["light", "dark"]:
        suffix = f"_{theme}"

        pie_path = stats_dir / f"difficulty_distribution{suffix}.png"
        generate_pie_chart(stats, pie_path, theme=theme)
        outputs[f"pie_{theme}"] = pie_path

        bar_path = stats_dir / f"top_tags{suffix}.png"
        generate_bar_chart(stats, bar_path, theme=theme)
        outputs[f"bar_{theme}"] = bar_path

        cloud_path = stats_dir / f"tag_cloud{suffix}.png"
        generate_word_cloud(stats, cloud_path, theme=theme)
        outputs[f"cloud_{theme}"] = cloud_path

    return outputs
