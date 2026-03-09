---
link: https://leetcode.com/problems/permutation-in-string/
tags:
  - Medium
  - Hash_Table
  - Two_Pointers
  - String
  - Sliding_Window
---
## Description
Given two strings `s1` and `s2`, return `true` if `s2` contains a permutation of `s1`, or `false` otherwise.

In other words, return `true` if one of `s1`'s permutations is the substring of `s2`.

**Example 1:**

```
Input: s1 = "ab", s2 = "eidbaooo"
Output: true
Explanation: s2 contains one permutation of s1 ("ba").
```

**Example 2:**

```
Input: s1 = "ab", s2 = "eidboaoo"
Output: false
```

**Constraints:**

- `1 <= s1.length, s2.length <= 10^4`
- `s1` and `s2` consist of lowercase English letters.

## Solution
### Approach 1
**Sliding Window + Character Frequency Count**: Maintain a fixed-size window and compare character frequencies within the window against the target string.

**Principle:**
A permutation means having the same characters with the same counts. We only need to find a substring in s2 of length len(s1) whose character frequency matches s1 exactly.

**Steps:**
1. If len(s1) > len(s2), return false immediately
2. Count frequency of each character in s1, store in count1 array
3. Initialize window: count frequency of first len(s1) characters in s2, store in count2
4. Compare count1 and count2, return true if equal
5. Slide window: move one position right, decrement left character count, increment right character count
6. After each slide, compare frequency arrays, return true if equal
7. Return false if no match found after traversal

**Example:**
- Input: s1 = "ab", s2 = "eidbaooo"
- Initial window "ei": count2 = [0,0,0,0,1,0,0,0,1,...], doesn't match count1 = [1,1,0,...]
- Slide to "id", "db", "ba": when window is "ba", count2 = [1,1,0,...] matches count1
- Output: true

```embed-go
PATH: "vault://leetcode/0501-0600/0567_permutation_in_string/solution.go"
TITLE: "leetcode 567. Permutation in String"
```
