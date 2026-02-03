"""Calculate statistics from parsed problem data."""

import json
from collections import Counter
from dataclasses import dataclass
from pathlib import Path

from .parser import Problem


@dataclass
class Statistics:
    """Holds computed statistics."""

    total: int
    difficulty_counts: dict[str, int]
    tag_counts: dict[str, int]

    def to_dict(self) -> dict:
        """Convert to a JSON-serializable dictionary."""
        return {
            "total": self.total,
            "difficulty_counts": self.difficulty_counts,
            "tag_counts": self.tag_counts,
        }

    @classmethod
    def from_dict(cls, data: dict) -> "Statistics":
        """Create Statistics from a dictionary."""
        return cls(
            total=data["total"],
            difficulty_counts=data["difficulty_counts"],
            tag_counts=data["tag_counts"],
        )

    def __eq__(self, other: object) -> bool:
        """Check equality with another Statistics instance."""
        if not isinstance(other, Statistics):
            return False
        return (
            self.total == other.total
            and self.difficulty_counts == other.difficulty_counts
            and self.tag_counts == other.tag_counts
        )


def load_cached_stats(cache_path: Path) -> Statistics | None:
    """Load cached statistics from file.

    Args:
        cache_path: Path to the cache file

    Returns:
        Statistics instance if cache exists and is valid, None otherwise
    """
    if not cache_path.exists():
        return None
    try:
        with open(cache_path) as f:
            data = json.load(f)
        return Statistics.from_dict(data)
    except (json.JSONDecodeError, KeyError):
        return None


def save_stats_cache(stats: Statistics, cache_path: Path) -> None:
    """Save statistics to cache file.

    Args:
        stats: Statistics instance to cache
        cache_path: Path to save the cache
    """
    cache_path.parent.mkdir(parents=True, exist_ok=True)
    with open(cache_path, "w") as f:
        json.dump(stats.to_dict(), f, indent=2)


def calculate_stats(problems: list[Problem]) -> Statistics:
    """Calculate statistics from a list of problems.

    Args:
        problems: List of parsed Problem instances

    Returns:
        Statistics instance with computed values
    """
    # Count difficulties
    difficulty_counter: Counter[str] = Counter()
    for problem in problems:
        difficulty_counter[problem.difficulty] += 1

    # Ensure all difficulty levels are present (even if 0)
    difficulty_counts = {
        "Easy": difficulty_counter.get("Easy", 0),
        "Medium": difficulty_counter.get("Medium", 0),
        "Hard": difficulty_counter.get("Hard", 0),
    }

    # Count tags
    tag_counter: Counter[str] = Counter()
    for problem in problems:
        for tag in problem.tags:
            tag_counter[tag] += 1

    return Statistics(
        total=len(problems),
        difficulty_counts=difficulty_counts,
        tag_counts=dict(tag_counter),
    )
