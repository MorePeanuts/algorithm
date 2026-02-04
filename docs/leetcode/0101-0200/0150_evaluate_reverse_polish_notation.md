---
link: https://leetcode.com/problems/evaluate-reverse-polish-notation/
tags:
  - Medium
  - Stack
  - Array
  - Math
---
## Description
You are given an array of strings `tokens` that represents an arithmetic expression in a [Reverse Polish Notation](http://en.wikipedia.org/wiki/Reverse_Polish_notation).

Evaluate the expression. Return *an integer that represents the value of the expression*.

**Note** that:

- The valid operators are `'+'`, `'-'`, `'*'`, and `'/'`.
- Each operand may be an integer or another expression.
- The division between two integers always **truncates toward zero**.
- There will not be any division by zero.
- The input represents a valid arithmetic expression in a reverse polish notation.
- The answer and all the intermediate calculations can be represented in a **32-bit** integer.

**Example 1:**

```
Input: tokens = ["2","1","+","3","*"]
Output: 9
Explanation: ((2 + 1) * 3) = 9
```

**Example 2:**

```
Input: tokens = ["4","13","5","/","+"]
Output: 6
Explanation: (4 + (13 / 5)) = 6
```

**Example 3:**

```
Input: tokens = ["10","6","9","3","+","-11","*","/","*","17","+","5","+"]
Output: 22
Explanation: ((10 * (6 / ((9 + 3) * -11))) + 17) + 5
= ((10 * (6 / (12 * -11))) + 17) + 5
= ((10 * (6 / -132)) + 17) + 5
= ((10 * 0) + 17) + 5
= (0 + 17) + 5
= 17 + 5
= 22
```

**Constraints:**

- `1 <= tokens.length <= 10^4`
- `tokens[i]` is either an operator: `"+"`, `"-"`, `"*"`, or `"/"`, or an integer in the range `[-200, 200]`.

## Solution

### Approach 1

**Stack Simulation**: Utilize the LIFO property of stack to process the Reverse Polish Notation sequentially.

**Principle:**
Reverse Polish Notation is naturally suited for stack operations—push numbers onto the stack, and when encountering an operator, pop two numbers, perform the calculation, and push the result back. Using an integer stack instead of a string stack avoids repeated type conversions during operations.

**Steps:**
1. Initialize an integer stack
2. Iterate through the tokens array:
   - If the current element is a number, convert it to an integer and push onto the stack
   - If the current element is an operator, pop two numbers from the stack (the first popped is the second operand), perform the operation, and push the result back
3. After iteration, the top element of the stack is the result

**Example:**
- Input `["4","13","5","/","+"]`
- Process: `push 4` → `push 13` → `push 5` → `/ pops 13 and 5, computes 13/5=2, push 2` → `+ pops 4 and 2, computes 4+2=6, push 6`
- Output `6`

```embed-go
PATH: "vault://leetcode/0101-0200/0150_evaluate_reverse_polish_notation/solution.go"
TITLE: "leetcode 150. Evaluate Reverse Polish Notation"
```
