---
link: https://leetcode.com/problems/maximum-depth-of-binary-tree/
tags:
  - Easy
  - Tree
  - Depth-First_Search
  - Breadth-First_Search
  - Binary_Tree
---
## Description
Given the `root` of a binary tree, return *its maximum depth*.

A binary tree's **maximum depth** is the number of nodes along the longest path from the root node down to the farthest leaf node.

**Example 1:**

![](https://assets.leetcode.com/uploads/2020/11/26/tmp-tree.jpg)

```
Input: root = [3,9,20,null,null,15,7]
Output: 3
```

**Example 2:**

```
Input: root = [1,null,2]
Output: 2
```

**Constraints:**

- The number of nodes in the tree is in the range `[0, 104]`.
- `-100 <= Node.val <= 100`

## Solution
### Approach 1
**Method Name**: Depth-First Search (Recursive)

**Principle:**
Recursively calculate the maximum depth of the left and right subtrees, return the larger value plus one.

**Steps:**
1. If current node is null, return depth 0
2. Recursively calculate left subtree depth
3. Recursively calculate right subtree depth
4. Return the maximum of left and right depths plus 1

```embed-go
PATH: "vault://leetcode/0101-0200/0104_maximum_depth_of_binary_tree/solution.go"
TITLE: "leetcode 104. Maximum Depth of Binary Tree"
```

### Approach 2
**Method Name**: Breadth-First Search (Iterative)

**Principle:**
Use a queue for level-order traversal, increment depth after completing each level.

**Steps:**
1. If root is null, return depth 0
2. Add root to queue
3. While queue is not empty, record current level node count
4. Traverse all nodes in current level, add their children to queue
5. Increment depth after completing each level
6. Return depth after traversal completes

```embed-go
PATH: "vault://leetcode/0101-0200/0104_maximum_depth_of_binary_tree/solution2.go"
TITLE: "leetcode 104. Maximum Depth of Binary Tree"
```
