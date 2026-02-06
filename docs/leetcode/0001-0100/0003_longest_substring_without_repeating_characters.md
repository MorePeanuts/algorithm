---
link: https://leetcode.com/problems/longest-substring-without-repeating-characters/
tags:
  - Medium
  - Hash_Table
  - String
  - Sliding_Window
---
## Description
Given a string `s`, find the length of the **longest** **substring** without duplicate characters.

**Example 1:**

```
Input: s = "abcabcbb"
Output: 3
Explanation: The answer is "abc", with the length of 3. Note that "bca" and "cab" are also correct answers.
```

**Example 2:**

```
Input: s = "bbbbb"
Output: 1
Explanation: The answer is "b", with the length of 1.
```

**Example 3:**

```
Input: s = "pwwkew"
Output: 3
Explanation: The answer is "wke", with the length of 3.
Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.
```

**Constraints:**

- `0 <= s.length <= 5 * 10^4`
- `s` consists of English letters, digits, symbols and spaces.

## Solution

### Approach 1

**Sliding Window + Hash Table**: Use two pointers to maintain a window without duplicate characters, with a hash table recording character positions.

**Principle:**
Maintain a sliding window `[l, r]` ensuring no duplicate characters within. When the right pointer encounters a duplicate, move the left pointer to the position after the duplicate and clean up removed characters from the hash table.

**Steps:**
1. Initialize left pointer `l=0`, right pointer `r=1`, hash table stores characters and their positions
2. Expand the window by moving the right pointer
3. If current character exists in the window, delete all characters from left pointer to duplicate position, update left pointer
4. Add current character to hash table, update maximum length

**Example:**
- Input `s = "abcabcbb"`
- Initial: `l=0, r=1, pos={a:0}`
- `r=1`: 'b' not duplicate, `pos={a:0, b:1}`, length 2
- `r=2`: 'c' not duplicate, `pos={a:0, b:1, c:2}`, length 3
- `r=3`: 'a' duplicate, delete `pos[a]`, `l=1`, `pos={b:1, c:2, a:3}`, length 3
- Continue until end, maximum length is 3

```embed-go
PATH: "vault://leetcode/0001-0100/0003_longest_substring_without_repeating_characters/solution.go"
TITLE: "leetcode 3. Longest Substring Without Repeating Characters"
```
