---
link: https://leetcode.cn/problems/min-stack/
tags:
  - 中等
  - 栈
  - 设计
---
## 题目描述
设计一个支持 `push` ，`pop` ，`top` 操作，并能在常数时间内检索到最小元素的栈。

实现 `MinStack` 类:

- `MinStack()` 初始化堆栈对象。
- `void push(int val)` 将元素val推入堆栈。
- `void pop()` 删除堆栈顶部的元素。
- `int top()` 获取堆栈顶部的元素。
- `int getMin()` 获取堆栈中的最小元素。

**示例 1:**

```
输入：
["MinStack","push","push","push","getMin","pop","top","getMin"]
[[],[-2],[0],[-3],[],[],[],[]]

输出：
[null,null,null,null,-3,null,0,-2]

解释：
MinStack minStack = new MinStack();
minStack.push(-2);
minStack.push(0);
minStack.push(-3);
minStack.getMin();   --> 返回 -3.
minStack.pop();
minStack.top();      --> 返回 0.
minStack.getMin();   --> 返回 -2.
```

**提示：**

- `-2^31 <= val <= 2^31 - 1`
- `pop`、`top` 和 `getMin` 操作总是在 **非空栈** 上调用
- `push`, `pop`, `top`, and `getMin`最多被调用 `3 * 10^4` 次

## 题目解析

### 解法1

**辅助栈法**：使用一个辅助栈同步记录当前栈的最小值。

**原理：**
维护两个栈，一个存储所有元素，另一个仅在当前元素小于等于栈顶时入栈，从而保持辅助栈栈顶始终是当前栈的最小值。

**步骤：**
1. Push 操作：元素入主栈；若辅助栈为空或元素 ≤ 辅助栈栈顶，则同时入辅助栈
2. Pop 操作：主栈出栈；若出栈元素等于辅助栈栈顶，辅助栈也出栈
3. Top 操作：返回主栈栈顶
4. GetMin 操作：返回辅助栈栈顶

**示例：**
- 依次 Push(-2, 0, -3)：主栈 [-2, 0, -3]，辅助栈 [-2, -3]
- GetMin() → -3（辅助栈栈顶）
- Pop()：主栈 [-2, 0]，辅助栈 [-2]（-3 等于辅助栈顶，同时出栈）
- GetMin() → -2

```embed-go
PATH: "vault://leetcode/0101-0200/0155_min_stack/solution.go"
TITLE: "leetcode 155.最小栈"
```

### 解法2

**差值法**：只用一个栈存储差值，用一个变量记录当前最小值。

**原理：**
栈中存储「当前值 - 当前最小值」的差值。负差值标记最小值变化的位置，同时记录新旧最小值的差距，用于 Pop 时还原。

**步骤：**
1. Push 操作：计算差值入栈；若差值 < 0，更新 min
2. Pop 操作：若栈顶 < 0，还原旧 min = min - 栈顶
3. Top 操作：若栈顶 < 0 返回 min，否则返回 min + 栈顶
4. GetMin 操作：返回 min

**示例：**
- Push(-2)：min = -2，差值 = 0，栈 [0]
- Push(0)：差值 = 0 - (-2) = 2，栈 [0, 2]
- Push(-3)：差值 = -3 - (-2) = -1 < 0，更新 min = -3，栈 [0, 2, -1]
- GetMin() → -3
- Pop()：栈顶 -1 < 0，还原 min = -3 - (-1) = -2，栈 [0, 2]
- GetMin() → -2

```embed-go
PATH: "vault://leetcode/0101-0200/0155_min_stack/solution2.go"
TITLE: "leetcode 155.最小栈"
```
