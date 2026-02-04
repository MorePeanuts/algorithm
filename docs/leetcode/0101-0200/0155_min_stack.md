---
link: https://leetcode.com/problems/min-stack/
tags:
  - Medium
  - Stack
  - Design
---
## Description
Design a stack that supports push, pop, top, and retrieving the minimum element in constant time.

Implement the `MinStack` class:

- `MinStack()` initializes the stack object.
- `void push(int val)` pushes the element `val` onto the stack.
- `void pop()` removes the element on the top of the stack.
- `int top()` gets the top element of the stack.
- `int getMin()` retrieves the minimum element in the stack.

You must implement a solution with `O(1)` time complexity for each function.

**Example 1:**

```
Input
["MinStack","push","push","push","getMin","pop","top","getMin"]
[[],[-2],[0],[-3],[],[],[],[]]

Output
[null,null,null,null,-3,null,0,-2]

Explanation
MinStack minStack = new MinStack();
minStack.push(-2);
minStack.push(0);
minStack.push(-3);
minStack.getMin(); // return -3
minStack.pop();
minStack.top();    // return 0
minStack.getMin(); // return -2
```

**Constraints:**

- `-2^31 <= val <= 2^31 - 1`
- Methods `pop`, `top` and `getMin` operations will always be called on **non-empty** stacks.
- At most `3 * 10^4` calls will be made to `push`, `pop`, `top`, and `getMin`.

## Solution

### Approach 1

**Auxiliary Stack**: Use an auxiliary stack to track the minimum value synchronously.

**Principle:**
Maintain two stacks: one stores all elements, and another only pushes when the current element is less than or equal to its top, ensuring the auxiliary stack's top is always the current minimum.

**Steps:**
1. Push: Add element to main stack; if auxiliary stack is empty or element ≤ auxiliary top, also push to auxiliary stack
2. Pop: Pop from main stack; if popped element equals auxiliary top, also pop from auxiliary stack
3. Top: Return main stack's top element
4. GetMin: Return auxiliary stack's top element

**Example:**
- Push(-2, 0, -3): main stack [-2, 0, -3], auxiliary stack [-2, -3]
- GetMin() → -3 (auxiliary stack top)
- Pop(): main stack [-2, 0], auxiliary stack [-2] (-3 equals auxiliary top, so also popped)
- GetMin() → -2

```embed-go
PATH: "vault://leetcode/0101-0200/0155_min_stack/solution.go"
TITLE: "leetcode 155. Min Stack"
```

### Approach 2

**Difference Method**: Use a single stack to store differences, with a variable tracking the current minimum.

**Principle:**
Store "current value - current minimum" as the difference. Negative differences mark where min changed and record the gap between old and new min for restoration during Pop.

**Steps:**
1. Push: Calculate and push difference; if difference < 0, update min
2. Pop: If top < 0, restore old min = min - top
3. Top: If top < 0 return min, otherwise return min + top
4. GetMin: Return min

**Example:**
- Push(-2): min = -2, diff = 0, stack [0]
- Push(0): diff = 0 - (-2) = 2, stack [0, 2]
- Push(-3): diff = -3 - (-2) = -1 < 0, update min = -3, stack [0, 2, -1]
- GetMin() → -3
- Pop(): top -1 < 0, restore min = -3 - (-1) = -2, stack [0, 2]
- GetMin() → -2

```embed-go
PATH: "vault://leetcode/0101-0200/0155_min_stack/solution2.go"
TITLE: "leetcode 155. Min Stack"
```
