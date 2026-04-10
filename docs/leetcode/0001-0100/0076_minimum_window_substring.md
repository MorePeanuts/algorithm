---
link: https://leetcode.com/problems/minimum-window-substring/
tags:
  - Hard
  - Hash_Table
  - String
  - Sliding_Window
---
## Description
Given two strings `s` and `t` of lengths `m` and `n` respectively, return *the **minimum window*** ***substring*** *of* `s` *such that every character in* `t` *(**including duplicates**) is included in the window*. If there is no such substring, return *the empty string* `""`.

The testcases will be generated such that the answer is **unique**.

**Example 1:**

```
Input: s = "ADOBECODEBANC", t = "ABC"
Output: "BANC"
Explanation: The minimum window substring "BANC" includes 'A', 'B', and 'C' from string t.
```

**Example 2:**

```
Input: s = "a", t = "a"
Output: "a"
Explanation: The entire string s is the minimum window.
```

**Example 3:**

```
Input: s = "a", t = "aa"
Output: ""
Explanation: Both 'a's from t must be included in the window.
Since the largest window of s only has one 'a', return empty string.
```

**Constraints:**

- `m == s.length`
- `n == t.length`
- `1 <= m, n <= 10^5`
- `s` and `t` consist of uppercase and lowercase English letters.

**Follow up:** Could you find an algorithm that runs in `O(m + n)` time?

## Solution
### Approach 1
**Method Name**: Sliding Window + Hash Table

**Principle:**
Use a sliding window to find the shortest substring in s that contains all characters from t. Two hash tables are used to track the required count of characters from t and the characters missing in the current window, dynamically adjusting the window boundaries to find the minimum window.

**Steps:**
1. Initialize two hash tables: `count` records the required count of each character from t, `miss` records characters missing in the current window
2. Iterate through string t to populate the `count` and `miss` hash tables
3. Use left and right pointers l and r to represent the sliding window boundaries, both starting at 0
4. Move right pointer r to expand the window, update `count` when encountering characters from t, remove from `miss` when a character's count ≤ 0
5. When the window contains all required characters (`miss` is empty), try moving left pointer l to shrink the window and find a smaller valid window
6. During window shrinking, record the substring with the minimum window length
7. When the window no longer contains all required characters, continue moving the right pointer

**Example:**
Using s = "ADOBECODEBANC", t = "ABC" as an example:
- Initial: count={A:1, B:1, C:1}, miss={A:true, B:true, C:true}
- r moves to index 0 (A): count={A:0, B:1, C:1}, miss={B:true, C:true}
- r moves to index 3 (B): count={A:0, B:0, C:1}, miss={C:true}
- r moves to index 5 (C): count={A:0, B:0, C:0}, miss={}, window is "ADOBEC"
- Start shrinking window, l moves to 1, window becomes "DOBEC", now miss={A:true}
- Continue moving r, eventually find the minimum window "BANC" (indices 9-12)

**Complexity Analysis:**
- Time Complexity: O(m + n), where m is the length of s and n is the length of t. Each character is accessed at most once by both the left and right pointers.
- Space Complexity: O(1), since the character set is limited to English letters with at most 52 distinct characters, the hash table size is constant.

> **Note:**
> 1. Need to handle duplicate characters in t (e.g., t = "aa")
> 2. The window boundary movements need to correctly handle hash table updates
> 3. Setting the initial minimum length to len(s) + 1 ensures that it will be updated when a valid window is found

```embed-go
PATH: "vault://leetcode/0001-0100/0076_minimum_window_substring/solution.go"
TITLE: "leetcode 76. Minimum Window Substring"
```
