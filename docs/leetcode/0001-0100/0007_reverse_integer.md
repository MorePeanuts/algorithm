---
link: https://leetcode.com/problems/reverse-integer/
tags:
  - Medium
  - Math
---
## Description
Given a signed 32-bit integer `x`, return `x` *with its digits reversed*. If reversing `x` causes the value to go outside the signed 32-bit integer range `[-231, 231 - 1]`, then return `0`.

**Assume the environment does not allow you to store 64-bit integers (signed or unsigned).**

**Example 1:**

```
Input: x = 123
Output: 321
```

**Example 2:**

```
Input: x = -123
Output: -321
```

**Example 3:**

```
Input: x = 120
Output: 21
```

**Constraints:**

- `-231 <= x <= 231 - 1`

## Solution

### Approach 1

**Mathematical Modulo Method**: Extract digits one by one and rebuild the reversed result, checking for overflow at each step.

**Principle:**
Extract the least significant digit using modulo 10, accumulate it to the result, and remove the digit by dividing by 10. Since the problem requires not using 64-bit integers, overflow checks must be performed before each operation.

**Steps:**
1. Initialize result `res = 0`
2. Loop until `x` becomes 0:
   - Get the last digit `last = x % 10`
   - Multiplication overflow check: if `res > MaxInt32/10` or `res < MinInt32/10`, multiplying by 10 will overflow, return 0
   - Addition overflow check: if `res*10 > MaxInt32-last` or `res*10 < MinInt32-last`, return 0
   - Calculate new result `res = res*10 + last`
   - Divide `x` by 10 to remove the processed digit
3. Return `res`

**Example:**
- Input `x = 123`:
  - Round 1: `last = 3`, `res = 3`, `x = 12`
  - Round 2: `last = 2`, `res = 32`, `x = 1`
  - Round 3: `last = 1`, `res = 321`, `x = 0`
  - Output `321`

- Input `x = -123`:
  - Round 1: `last = -3`, `res = -3`, `x = -12`
  - Round 2: `last = -2`, `res = -32`, `x = -1`
  - Round 3: `last = -1`, `res = -321`, `x = 0`
  - Output `-321`

```embed-go
PATH: "vault://leetcode/0001-0100/0007_reverse_integer/solution.go"
TITLE: "leetcode 7. Reverse Integer"
```
