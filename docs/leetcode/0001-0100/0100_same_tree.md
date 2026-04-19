---
link: https://leetcode.com/problems/same-tree/
tags:
  - Easy
  - Tree
  - Depth-First_Search
  - Breadth-First_Search
  - Binary_Tree
---
## Description
Given the roots of two binary trees `p` and `q`, write a function to check if they are the same or not.

Two binary trees are considered the same if they are structurally identical, and the nodes have the same value.

**Example 1:**

![](https://assets.leetcode.com/uploads/2020/12/20/ex1.jpg)

```
Input: p = [1,2,3], q = [1,2,3]
Output: true
```

**Example 2:**

![](https://assets.leetcode.com/uploads/2020/12/20/ex2.jpg)

```
Input: p = [1,2], q = [1,null,2]
Output: false
```

**Example 3:**

![](https://assets.leetcode.com/uploads/2020/12/20/ex3.jpg)

```
Input: p = [1,2,1], q = [1,1,2]
Output: false
```

**Constraints:**

- The number of nodes in both trees is in the range `[0, 100]`.
- `-104 <= Node.val <= 104`

## Solution
### Approach 1
**Method Name**: Depth-First Search

**Principle:**
Traverse both trees simultaneously using recursion, comparing node values and recursively checking left and right subtrees.

**Steps:**
1. If both nodes are nil, the trees are structurally identical at this position, return true
2. If one node is nil and the other is not, the structure differs, return false
3. Compare if the current node values are equal
4. Recursively compare the left subtree and right subtree

```embed-go
PATH: "vault://leetcode/0001-0100/0100_same_tree/solution.go"
TITLE: "leetcode 100. Same Tree"
```
