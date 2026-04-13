---
link: https://leetcode.com/problems/basic-calculator/
tags:
  - Hard
  - Stack
  - Recursion
  - Math
  - String
---
## Description
Given a string `s` representing a valid expression, implement a basic calculator to evaluate it, and return *the result of the evaluation*.

**Note:** You are **not** allowed to use any built-in function which evaluates strings as mathematical expressions, such as `eval()`.

**Example 1:**

```
Input: s = "1 + 1"
Output: 2
```

**Example 2:**

```
Input: s = " 2-1 + 2 "
Output: 3
```

**Example 3:**

```
Input: s = "(1+(4+5+2)-3)+(6+8)"
Output: 23
```

**Constraints:**

- `1 <= s.length <= 3 * 105`
- `s` consists of digits, `'+'`, `'-'`, `'('`, `')'`, and `' '`.
- `s` represents a valid expression.
- `'+'` is **not** used as a unary operation (i.e., `"+1"` and `"+(2 + 3)"` is invalid).
- `'-'` could be used as a unary operation (i.e., `"-1"` and `"-(2 + 3)"` is valid).
- There will be no two consecutive operators in the input.
- Every number and running calculation will fit in a signed 32-bit integer.

## Solution
### Approach 1
**Method Name**: Infix to Postfix Conversion + Postfix Evaluation

**Principle:**
Convert the infix expression to postfix notation (Reverse Polish Notation), then use a stack to evaluate the postfix expression. This approach transforms operator precedence issues into linear processing, simplifying the calculation logic.

**Steps:**
1. **Tokenization**: Convert the input string into a list of tokens (numbers, operators, parentheses), handling unary minus by prepending 0 before `-` that follows `(`
2. **Infix to Postfix**: Use a stack to handle operator precedence. Numbers are output directly, `(` is pushed to stack, `)` pops from stack until `(` is found, operators pop higher or equal precedence operators before being pushed
3. **Postfix Evaluation**: Traverse the postfix expression, push numbers to stack, pop two numbers for each operator and push result back. The final stack top is the answer

**Example:**
Input `"(1+(4+5+2)-3)+(6+8)"`:
- Tokenization → `["(", "1", "+", "(", "4", "+", "5", "+", "2", ")", "-", "3", ")", "+", "(", "6", "+", "8", ")"]`
- To Postfix → `["1", "4", "5", "+", "2", "+", "+", "3", "-", "6", "8", "+", "+"]`
- Evaluation → 23

**Complexity Analysis:**
- Time Complexity: O(n), three linear passes over the expression
- Space Complexity: O(n), stack storage for intermediate results

> **Note**: Unary minus handling is crucial—check the previous non-space character to determine if 0 needs to be prepended

```embed-go
PATH: "vault://leetcode/0201-0300/0224_basic_calculator/solution.go"
TITLE: "leetcode 224. Basic Calculator"
```

### Approach 2
**Method Name**: Dual Stack Direct Evaluation (Simplified)

**Principle:**
Directly evaluate the tokenized infix expression using two stacks (number stack + operator stack), avoiding explicit conversion to postfix. Calculate in real-time during traversal based on operator precedence and parentheses rules.

**Steps:**
1. **Tokenization**: Same as Approach 1, handle unary minus
2. **Dual Stack Traversal**:
   - Number: push to number stack
   - `(`: push to operator stack
   - `)`: pop and calculate until `(` is found
   - `+/-`: pop all calculable operators (until `(`) before pushing
3. **Clear Stacks**: After traversal, pop remaining operators and calculate

**Example:**
Input `"1 + 1"`:
- Tokenization → `["1", "+", "1"]`
- Traversal → push 1, push `+`, push 1
- Clear Stacks → pop `+` and two 1s, calculate to get 2

**Complexity Analysis:**
- Time Complexity: O(n), each token is pushed and popped once
- Space Complexity: O(n), dual stack storage

> **Note**: Since only `+` and `-` with same precedence exist, operator precedence logic can be simplified

```embed-go
PATH: "vault://leetcode/0201-0300/0224_basic_calculator/solution2.go"
TITLE: "leetcode 224. Basic Calculator"
```

### Approach 3
**Method Name**: Single Pass + Dual Stack Evaluation (Optimized)

**Principle:**
Combine tokenization and evaluation into a single pass. Complete token recognition and dual stack evaluation while traversing the string, further optimizing space and time efficiency.

**Steps:**
1. **Single Pass**: Use two pointers to recognize number tokens, while handling operators and parentheses
2. **Real-time Processing**: Push complete numbers to stack immediately, process operators/parentheses with stack rules as they are recognized
3. **Clear Stacks**: Same as Approach 2

**Example:**
Input `" 2-1 + 2 "`:
- Scan `2`: push to number stack [2]
- Scan `-`: push to operator stack [`-`]
- Scan `1`: push to number stack [2, 1]
- Scan `+`: pop `-` to calculate 1, push 1 and then `+`. Stack state: numbers [1], operators [`+`]
- Scan `2`: push to number stack [1, 2]
- Clear Stacks: pop `+` to calculate 3

**Complexity Analysis:**
- Time Complexity: O(n), single pass + stack operations
- Space Complexity: O(n), dual stack storage

> **Note**: Single pass requires careful handling of spaces and number boundaries to ensure correct token recognition

```embed-go
PATH: "vault://leetcode/0201-0300/0224_basic_calculator/solution3.go"
TITLE: "leetcode 224. Basic Calculator"
```

