---
link: https://leetcode.com/problems/longest-repeating-character-replacement/
tags:
  - Medium
  - Hash_Table
  - String
  - Sliding_Window
---
## Description
You are given a string `s` and an integer `k`. You can choose any character of the string and change it to any other uppercase English character. You can perform this operation at most `k` times.

Return *the length of the longest substring containing the same letter you can get after performing the above operations*.

**Example 1:**

```
Input: s = "ABAB", k = 2
Output: 4
Explanation: Replace the two 'A's with two 'B's or vice versa.
```

**Example 2:**

```
Input: s = "AABABBA", k = 1
Output: 4
Explanation: Replace the one 'A' in the middle with 'B' and form "AABBBBA".
The substring "BBBB" has the longest repeating letters, which is 4.
There may exists other ways to achieve this answer too.
```

**Constraints:**

- `1 <= s.length <= 10^5`
- `s` consists of only uppercase English letters.
- `0 <= k <= s.length`

## Solution

### Approach 1

**Method**: Sliding Window

**Principle**: A window is valid when `window length - max frequency in window <= k`. The window only expands or slides, never shrinks. The final window size is the answer.

**Explanation**:

1. **Valid Window Condition**
   - Window has `len` characters, most frequent character appears `maxCount` times
   - To make all characters the same, we need to replace `len - maxCount` characters
   - If `len - maxCount <= k`, the window can become uniform with at most k replacements

2. **Window Only Expands or Slides, Never Shrinks**
   - Valid: right pointer advances, window expands
   - Invalid: both pointers advance together, window slides (size unchanged)
   - Why not shrink? We want the longest valid window. If current window size is n and invalid, shrinking to n-1 is pointless—we've already found valid windows of size n-1. Just slide and wait for maxCount to increase

3. **maxCount Only Increases**
   - When sliding, actual max frequency in window may decrease, but code doesn't update maxCount
   - This is correct: maxCount tracks "historical max frequency". Only when maxCount increases can we find longer valid windows
   - Old maxCount value just keeps window sliding, doesn't affect the answer

**Steps**:
1. Use `count[26]` to track frequency of each character in window
2. Right pointer `r` expands window, update `count[s[r]]` and `maxCount`
3. If `r - l + 1 - maxCount > k`, move left pointer `l` one step (window slides)
4. Return `len(s) - l` (final window size)

**Example**:

Using `s = "AABABBA", k = 1`:

| r | Window | maxCount | Length | Replace | Action |
|---|--------|----------|--------|---------|--------|
| 0 | A | 1 | 1 | 0 | Expand |
| 1 | AB | 1 | 2 | 1 | Expand |
| 2 | ABA | 2 | 3 | 1 | Expand |
| 3 | ABAB | 2 | 4 | 2 | Slide, l=1 |
| 4 | BABA | 2 | 4 | 2 | Slide, l=2 |
| 5 | ABBB | 3 | 4 | 1 | Expand (maxCount updated) |
| 6 | BBBA | 3 | 4 | 1 | Slide, l=3 |

Final `l = 3`, answer = `7 - 3 = 4`

**Complexity**:
- Time: O(n), each character visited once
- Space: O(1), fixed-size array

```embed-go
PATH: "vault://leetcode/0401-0500/0424_longest_repeating_character_replacement/solution.go"
TITLE: "leetcode 424. Longest Repeating Character Replacement"
```
