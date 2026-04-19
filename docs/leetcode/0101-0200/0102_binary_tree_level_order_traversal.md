---
link: https://leetcode.com/problems/binary-tree-level-order-traversal/
tags:
  - Medium
  - Tree
  - Breadth-First_Search
  - Binary_Tree
---
## Description
Given the `root` of a binary tree, return *the level order traversal of its nodes' values*. (i.e., from left to right, level by level).

**Example 1:**

![](https://assets.leetcode.com/uploads/2021/02/19/tree1.jpg)

```
Input: root = [3,9,20,null,null,15,7]
Output: [[3],[9,20],[15,7]]
```

**Example 2:**

```
Input: root = [1]
Output: [[1]]
```

**Example 3:**

```
Input: root = []
Output: []
```

**Constraints:**

- The number of nodes in the tree is in the range `[0, 2000]`.
- `-1000 <= Node.val <= 1000`

## Solution
### Approach 1
**Method Name**: Breadth-First Search (BFS) using Queue

**Principle:**
Use a queue for level order traversal, processing nodes level by level by tracking the number of nodes at each level.

**Steps:**
1. Initialize result list, return empty if root is null
2. Create queue and enqueue root node
3. While queue is not empty, record current queue length (number of nodes at current level)
4. Traverse all nodes at current level, add node values to current level list, and enqueue left and right children
5. Add current level list to result
6. Repeat steps 3-5 until queue is empty

**Example:**
- Input: `[3,9,20,null,null,15,7]`
- Processing:
  - Level 1: Queue `[3]` → Process `3` → Enqueue `9, 20` → Result `[[3]]`
  - Level 2: Queue `[9, 20]` → Process `9, 20` → Enqueue `15, 7` → Result `[[3], [9, 20]]`
  - Level 3: Queue `[15, 7]` → Process `15, 7` → No children → Result `[[3], [9, 20], [15, 7]]`
- Output: `[[3],[9,20],[15,7]]`

```embed-go
PATH: "vault://leetcode/0101-0200/0102_binary_tree_level_order_traversal/solution.go"
TITLE: "leetcode 102. Binary Tree Level Order Traversal"
```
