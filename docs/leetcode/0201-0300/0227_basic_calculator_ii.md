---
link: https://leetcode.com/problems/basic-calculator-ii/
tags:
  - Medium
  - Stack
  - Math
  - String
---
## Description
Given a string `s` which represents an expression, *evaluate this expression and return its value*.

The integer division should truncate toward zero.

You may assume that the given expression is always valid. All intermediate results will be in the range of `[-231, 231 - 1]`.

**Note:** You are not allowed to use any built-in function which evaluates strings as mathematical expressions, such as `eval()`.

**Example 1:**

```
Input: s = "3+2*2"
Output: 7
```

**Example 2:**

```
Input: s = " 3/2 "
Output: 1
```

**Example 3:**

```
Input: s = " 3+5 / 2 "
Output: 5
```

**Constraints:**

- `1 <= s.length <= 3 * 105`
- `s` consists of integers and operators `('+', '-', '*', '/')` separated by some number of spaces.
- `s` represents **a valid expression**.
- All the integers in the expression are non-negative integers in the range `[0, 231 - 1]`.
- The answer is **guaranteed** to fit in a **32-bit integer**.

## Solution
### Approach 1
**Method Name**: Infix to Postfix (Reverse Polish Notation) Evaluation

**Principle:**
Convert infix expressions (standard mathematical notation) to postfix expressions (operators follow operands) to eliminate operator precedence ambiguity, then use a stack to directly evaluate the postfix expression.

**Steps:**
1. **Expression Parsing**: Parse the input string into tokens (numbers and operators), skipping spaces
2. **Infix to Postfix**: Use the Shunting-yard algorithm:
   - Directly output numbers to the postfix expression
   - When encountering an operator, pop all operators from the stack with precedence greater than or equal to the current operator, then push the current operator
   - `*` and `/` have higher precedence than `+` and `-`
3. **Postfix Expression Evaluation**:
   - Push numbers onto the stack
   - When encountering an operator, pop two numbers, calculate, and push the result
   - The remaining number in the stack is the final result

**Example:**
- Input: `"3+2*2"`
- Parse tokens: `["3", "+", "2", "*", "2"]`
- Convert to postfix: `["3", "2", "2", "*", "+"]`
- Evaluation: push 3 → push 2 → push 2 → pop 2,2 calculate 2*2=4, push 4 → pop 3,4 calculate 3+4=7, push 7
- Output: `7`

```embed-go
PATH: "vault://leetcode/0201-0300/0227_basic_calculator_ii/solution.go"
TITLE: "leetcode 227. Basic Calculator II"
```

### Approach 2
**Method Name**: Two-pass Calculation (Multiplication/Division First, Then Addition/Subtraction)

**Principle:**
Leverage operator precedence with two rounds of calculation: first compute all multiplications and divisions, then compute additions and subtractions. Each round only processes simple expressions with same-precedence operators.

**Steps:**
1. **Expression Encoding**: Convert the string to an integer array, numbers remain as-is, operators encoded as negative numbers (`+`=-1, `-`=-2, `*`=-3, `/`=-4)
2. **First Round (Multiplication/Division)**: Traverse the array, immediately compute left and right operands when encountering `*` or `/` and replace them
3. **Second Round (Addition/Subtraction)**: Traverse the array again, compute left and right operands when encountering `+` or `-` and replace them
4. **Result**: The last remaining element in the array is the expression value

**Example:**
- Input: `"3+5/2"`
- Encoded: `[3, -1, 5, -4, 2]`
- First round (process -4 which is `/`): `[3, -1, 2]`
- Second round (process -1 which is `+`): `[5]`
- Output: `5`

```embed-go
PATH: "vault://leetcode/0201-0300/0227_basic_calculator_ii/solution2.go"
TITLE: "leetcode 227. Basic Calculator II"
```

### Approach 3
**Method Name**: Single Stack Real-time Calculation

**Principle:**
Use a single stack to store intermediate results to be accumulated. When encountering `+` and `-`, push signed numbers onto the stack. When encountering `*` and `/`, immediately compute with the stack top element. Finally, sum all numbers in the stack to get the result.

**Steps:**
1. **Initialization**: Empty stack, record the previous operator (defaults to `+`)
2. **Traverse the Expression**:
   - Parse complete numbers
   - Determine how to handle the current number based on the previous operator:
     - `+`: Push number directly onto stack
     - `-`: Push negated number onto stack
     - `*`: Pop stack top, multiply by current number, push result
     - `/`: Pop stack top, divide by current number, push result
   - Update the previous operator
3. **Sum**: Add all numbers in the stack for the final result

**Example:**
- Input: `"3+2*2"`
- Process number 3, previous operator `+`: stack `[3]`
- Process number 2, previous operator `+`: stack `[3, 2]`
- Process number 2, previous operator `*`: pop 2, calculate 2*2=4, stack `[3, 4]`
- Sum: 3+4=7
- Output: `7`

```embed-go
PATH: "vault://leetcode/0201-0300/0227_basic_calculator_ii/solution3.go"
TITLE: "leetcode 227. Basic Calculator II"
```
