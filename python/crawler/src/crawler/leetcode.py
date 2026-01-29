"""LeetCode crawler implementation."""

import re

import httpx

# GraphQL query to fetch problem details (includes translated fields for CN site)
GRAPHQL_QUERY = """
query questionData($titleSlug: String!) {
    question(titleSlug: $titleSlug) {
        questionFrontendId
        title
        translatedTitle
        titleSlug
        difficulty
        content
        translatedContent
        topicTags {
            name
            translatedName
        }
    }
}
"""

# API endpoints
API_ENDPOINTS = {
    "leetcode.cn": "https://leetcode.cn/graphql",
    "leetcode.com": "https://leetcode.com/graphql",
}

# Difficulty translation
DIFFICULTY_MAP = {
    "Easy": "简单",
    "Medium": "中等",
    "Hard": "困难",
}


def parse_url(url: str) -> tuple[str, str]:
    """Extract problem slug and site from LeetCode URL.

    Args:
        url: LeetCode problem URL

    Returns:
        Tuple of (slug, site)

    Raises:
        ValueError: If URL format is invalid
    """
    pattern = r"https?://(leetcode\.(?:cn|com))/problems/([a-z0-9-]+)"
    match = re.match(pattern, url)
    if not match:
        raise ValueError(f"Invalid LeetCode URL: {url}")
    return match.group(2), match.group(1)


def fetch_problem(slug: str, site: str) -> dict:
    """Fetch problem data from LeetCode GraphQL API.

    Args:
        slug: Problem slug (e.g., 'two-sum')
        site: Site domain (e.g., 'leetcode.cn')

    Returns:
        Problem data dictionary

    Raises:
        RuntimeError: If API request fails
    """
    endpoint = API_ENDPOINTS[site]
    headers = {
        "Content-Type": "application/json",
        "User-Agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7)",
        "Referer": f"https://{site}/problems/{slug}/",
    }

    with httpx.Client(timeout=30.0) as client:
        response = client.post(
            endpoint,
            json={"query": GRAPHQL_QUERY, "variables": {"titleSlug": slug}},
            headers=headers,
        )
        response.raise_for_status()
        data = response.json()

    question = data.get("data", {}).get("question")
    if not question:
        raise RuntimeError(f"Problem not found: {slug}")

    return question


def normalize_problem(problem: dict, use_chinese: bool = False) -> dict:
    """Normalize problem data to a common format.

    Args:
        problem: Raw problem data from API
        use_chinese: Whether to prefer Chinese content

    Returns:
        Normalized problem dictionary with keys:
        - id: Problem ID (int)
        - title: Problem title
        - slug: URL slug
        - difficulty: Difficulty in Chinese
        - tags: List of tag names
        - content: Problem content (HTML)
    """
    difficulty = problem["difficulty"]
    difficulty_cn = DIFFICULTY_MAP.get(difficulty, difficulty)

    tags = []
    for tag in problem["topicTags"]:
        tag_name = tag.get("translatedName") or tag.get("name")
        if tag_name:
            tags.append(tag_name)

    content = problem.get("translatedContent") if use_chinese else None
    if not content:
        content = problem["content"]

    return {
        "id": int(problem["questionFrontendId"]),
        "title": problem.get("translatedTitle") or problem["title"],
        "title_en": problem["title"],
        "slug": problem["titleSlug"],
        "difficulty": difficulty_cn,
        "tags": tags,
        "content": content,
    }
