---
link: https://leetcode.com/problems/valid-parentheses/
tags:
  - Easy
  - Stack
  - String
---
## Description
Given a string `s` containing just the characters `'('`, `')'`, `'{'`, `'}'`, `'['` and `']'`, determine if the input string is valid.

An input string is valid if:

1. Open brackets must be closed by the same type of brackets.
2. Open brackets must be closed in the correct order.
3. Every close bracket has a corresponding open bracket of the same type.

**Example 1:**

**Input:** `s = "()"`

**Output:** `true`

**Example 2:**

**Input:** `s = "()[]{}"`

**Output:** `true`

**Example 3:**

**Input:** `s = "(]"`

**Output:** `false`

**Example 4:**

**Input:** `s = "([])"`

**Output:** `true`

**Example 5:**

**Input:** `s = "([)]"`

**Output:** `false`

**Constraints:**

- `1 <= s.length <= 10^4`
- `s` consists of parentheses only `'()[]{}'`.

## Solution

### Approach 1

**Stack Matching**: Use stack's LIFO property to match bracket pairs.

**Principle:**
Push left brackets onto the stack, and when encountering a right bracket, check if the stack top is the corresponding left bracket. If matched, pop the stack; otherwise, the string is invalid. Use ASCII code difference (`)`-`(`=1, `]`-`[`=2, `}`-`{`=2) to quickly determine if brackets match.

**Steps:**
1. If the string length is odd, return false immediately
2. Iterate through the string, push left brackets `(`, `{`, `[` onto the stack
3. When encountering a right bracket, check if the stack is non-empty and the ASCII difference between the current right bracket and stack top is 1 or 2. If matched, pop; otherwise return false
4. After iteration, the string is valid if the stack is empty

```embed-go
PATH: "vault://leetcode/0001-0100/0020_valid_parentheses/solution.go"
TITLE: "leetcode 20. Valid Parentheses"
```
