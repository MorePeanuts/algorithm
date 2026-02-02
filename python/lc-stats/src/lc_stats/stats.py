"""Calculate statistics from parsed problem data."""

from collections import Counter
from dataclasses import dataclass

from .parser import Problem


@dataclass
class Statistics:
    """Holds computed statistics."""

    total: int
    difficulty_counts: dict[str, int]
    tag_counts: dict[str, int]


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
