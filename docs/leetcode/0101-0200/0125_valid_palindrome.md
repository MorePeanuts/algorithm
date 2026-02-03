---
link: https://leetcode.com/problems/valid-palindrome/
tags:
  - Easy
  - Two_Pointers
  - String
---
## Description
A phrase is a **palindrome** if, after converting all uppercase letters into lowercase letters and removing all non-alphanumeric characters, it reads the same forward and backward. Alphanumeric characters include letters and numbers.

Given a string `s`, return `true` *if it is a **palindrome**, or* `false` *otherwise*.

**Example 1:**

```
Input: s = "A man, a plan, a canal: Panama"
Output: true
Explanation: "amanaplanacanalpanama" is a palindrome.
```

**Example 2:**

```
Input: s = "race a car"
Output: false
Explanation: "raceacar" is not a palindrome.
```

**Example 3:**

```
Input: s = " "
Output: true
Explanation: s is an empty string "" after removing non-alphanumeric characters.
Since an empty string reads the same forward and backward, it is a palindrome.
```

**Constraints:**

- `1 <= s.length <= 2 * 10^5`
- `s` consists only of printable ASCII characters.

## Solution

### Approach 1

**Two Pointers**: Use left and right pointers moving from both ends toward the center, skipping non-alphanumeric characters, and comparing characters one by one.

**Principle:**
A palindrome reads the same forward and backward. By using two pointers starting from both ends of the string and moving toward the center, we can compare valid characters (letters or digits) at each step. If all comparisons match, the string is a palindrome.

**Steps:**
1. Initialize left pointer `l` at the start of the string and right pointer `r` at the end
2. While `l < r`:
   - Move left pointer to skip all non-alphanumeric characters
   - Move right pointer to skip all non-alphanumeric characters
   - If `l >= r`, comparison is complete, break the loop
   - Convert both characters to lowercase and compare; if not equal, return `false`
   - Move left pointer right and right pointer left
3. Return `true` after the loop completes

```embed-go
PATH: "vault://leetcode/0101-0200/0125_valid_palindrome/solution.go"
TITLE: "leetcode 125. Valid Palindrome"
```
