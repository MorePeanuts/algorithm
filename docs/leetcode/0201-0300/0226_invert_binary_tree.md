---
link: https://leetcode.com/problems/invert-binary-tree/
tags:
  - Easy
  - Tree
  - Depth-First_Search
  - Breadth-First_Search
  - Binary_Tree
---
## Description
Given the `root` of a binary tree, invert the tree, and return *its root*.

**Example 1:**

![](https://assets.leetcode.com/uploads/2021/03/14/invert1-tree.jpg)

```
Input: root = [4,2,7,1,3,6,9]
Output: [4,7,2,9,6,3,1]
```

**Example 2:**

![](https://assets.leetcode.com/uploads/2021/03/14/invert2-tree.jpg)

```
Input: root = [2,1,3]
Output: [2,3,1]
```

**Example 3:**

```
Input: root = []
Output: []
```

**Constraints:**

- The number of nodes in the tree is in the range `[0, 100]`.
- `-100 <= Node.val <= 100`

## Solution
### Approach 1
**Method Name**: Recursive Inversion (Post-order Traversal)

**Principle:**
Recursively invert the left and right subtrees first, then swap the left and right child nodes of the current node.

**Steps:**
1. If the current node is nil, return nil
2. Recursively invert the left subtree
3. Recursively invert the right subtree
4. Swap the left and right child nodes of the current node
5. Return the current node

```embed-go
PATH: "vault://leetcode/0201-0300/0226_invert_binary_tree/solution.go"
TITLE: "leetcode 226. Invert Binary Tree"
```

### Approach 2
**Method Name**: Recursive Inversion (Pre-order Traversal)

**Principle:**
Recursively swap the left and right child nodes of the current node first, then invert the left and right subtrees respectively.

**Steps:**
1. If the current node is nil, return nil
2. Swap the left and right child nodes of the current node
3. Recursively invert the left subtree
4. Recursively invert the right subtree
5. Return the current node

```embed-go
PATH: "vault://leetcode/0201-0300/0226_invert_binary_tree/solution2.go"
TITLE: "leetcode 226. Invert Binary Tree"
```
