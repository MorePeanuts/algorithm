---
link: https://leetcode.com/problems/valid-anagram/
tags:
  - Easy
  - Hash_Table
  - String
  - Sorting
---
## Description
Given two strings `s` and `t`, return `true` if `t` is an anagram of `s`, and `false` otherwise.

**Example 1:**

**Input:** s = "anagram", t = "nagaram"

**Output:** true

**Example 2:**

**Input:** s = "rat", t = "car"

**Output:** false

**Constraints:**

- `1 <= s.length, t.length <= 5 * 104`
- `s` and `t` consist of lowercase English letters.

**Follow up:** What if the inputs contain Unicode characters? How would you adapt your solution to such a case?

## Solution

### Approach 1

**Hash Table Counting Method**: Use a hash table to count character frequencies, determine if two strings are anagrams by incrementing and decrementing counts.

**Principle:**
The essence of anagrams is that two strings contain the same characters with the same occurrence count for each. Use a hash table to record the occurrence count of each character in the first string, then iterate through the second string to decrement counts. If any character count becomes negative, they are not anagrams.

**Steps:**
1. First compare the lengths of both strings; if unequal, return `false` directly
2. Create hash table `count`, traverse string `s`, count each character's occurrences
3. Traverse string `t`, decrement the count for each character in the hash table
4. If any character count becomes negative, `t` has more of that character than `s`, return `false`
5. Return `true` after traversal completes

```embed-go
PATH: "vault://leetcode/0201-0300/0242_valid_anagram/solution.go"
TITLE: "leetcode 242. Valid Anagram"
```
