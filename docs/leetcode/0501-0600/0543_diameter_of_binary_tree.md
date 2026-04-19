---
link: https://leetcode.com/problems/diameter-of-binary-tree/
tags:
  - Easy
  - Tree
  - Depth-First_Search
  - Binary_Tree
---
## Description
Given the `root` of a binary tree, return *the length of the **diameter** of the tree*.

The **diameter** of a binary tree is the **length** of the longest path between any two nodes in a tree. This path may or may not pass through the `root`.

The **length** of a path between two nodes is represented by the number of edges between them.

**Example 1:**

![](https://assets.leetcode.com/uploads/2021/03/06/diamtree.jpg)

```
Input: root = [1,2,3,4,5]
Output: 3
Explanation: 3 is the length of the path [4,2,1,3] or [5,2,1,3].
```

**Example 2:**

```
Input: root = [1,2]
Output: 1
```

**Constraints:**

- The number of nodes in the tree is in the range `[1, 104]`.
- `-100 <= Node.val <= 100`

## Solution
### Approach 1
**Method Name**: Depth-First Search (DFS)

**Principle:**
While calculating the maximum depth of each node in the binary tree, we track the maximum value of the sum of the depths of the left and right subtrees centered at that node, which gives the diameter of the tree.

**Steps:**
1. Initialize the diameter variable to 0
2. Recursively calculate the maximum depth of each node:
   - If the node is nil, return depth 0
   - Recursively calculate the left subtree depth
   - Recursively calculate the right subtree depth
   - Update the diameter to be the maximum of the current diameter and (left depth + right depth)
   - Return the maximum depth of the current node (max(left depth, right depth) + 1)
3. After recursion completes, the diameter variable holds the result

```embed-go
PATH: "vault://leetcode/0501-0600/0543_diameter_of_binary_tree/solution.go"
TITLE: "leetcode 543. Diameter of Binary Tree"
```
