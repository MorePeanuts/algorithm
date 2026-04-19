---
link: https://leetcode.com/problems/validate-binary-search-tree/
tags:
  - Medium
  - Tree
  - Depth-First_Search
  - Binary_Search_Tree
  - Binary_Tree
---
## Description
Given the `root` of a binary tree, *determine if it is a valid binary search tree (BST)*.

A **valid BST** is defined as follows:

- The left subtree of a node contains only nodes with keys **strictly less than** the node's key.
- The right subtree of a node contains only nodes with keys **strictly greater than** the node's key.
- Both the left and right subtrees must also be binary search trees.

**Example 1:**

![](https://assets.leetcode.com/uploads/2020/12/01/tree1.jpg)

```
Input: root = [2,1,3]
Output: true
```

**Example 2:**

![](https://assets.leetcode.com/uploads/2020/12/01/tree2.jpg)

```
Input: root = [5,1,4,null,null,3,6]
Output: false
Explanation: The root node's value is 5 but its right child's value is 4.
```

**Constraints:**

- The number of nodes in the tree is in the range `[1, 104]`.
- `-231 <= Node.val <= 231 - 1`

## Solution
### Approach 1
**Method Name**: Recursive validation with predecessor/successor check

**Principle:**
For each node, verify that the largest node in its left subtree (predecessor) is smaller than it, and the smallest node in its right subtree (successor) is larger than it, then recursively check both subtrees.

**Steps:**
1. If the root is nil, return true (empty tree is valid BST).
2. Find the leftmost node in the right subtree (smallest node in right subtree).
3. If that node exists and its value is <= root's value, return false.
4. Find the rightmost node in the left subtree (largest node in left subtree).
5. If that node exists and its value is >= root's value, return false.
6. Recursively validate the left and right subtrees.

**Example:**
- Input: [2,1,3] -> Check predecessor of 2 (1 < 2) and successor (3 > 2), then validate left (1) and right (3) which are leaves -> Output: true
- Input: [5,1,4,null,null,3,6] -> Check successor of 5 is 3 (which is < 5) -> Output: false

```embed-go
PATH: "vault://leetcode/0001-0100/0098_validate_binary_search_tree/solution.go"
TITLE: "leetcode 98. Validate Binary Search Tree"
```
