---
link: https://leetcode.com/problems/binary-tree-maximum-path-sum/
tags:
  - Hard
  - Tree
  - Depth-First_Search
  - Dynamic_Programming
  - Binary_Tree
---
## Description
A **path** in a binary tree is a sequence of nodes where each pair of adjacent nodes in the sequence has an edge connecting them. A node can only appear in the sequence **at most once**. Note that the path does not need to pass through the root.

The **path sum** of a path is the sum of the node's values in the path.

Given the `root` of a binary tree, return *the maximum **path sum** of any **non-empty** path*.

**Example 1:**

![](https://assets.leetcode.com/uploads/2020/10/13/exx1.jpg)

```
Input: root = [1,2,3]
Output: 6
Explanation: The optimal path is 2 -> 1 -> 3 with a path sum of 2 + 1 + 3 = 6.
```

**Example 2:**

![](https://assets.leetcode.com/uploads/2020/10/13/exx2.jpg)

```
Input: root = [-10,9,20,null,null,15,7]
Output: 42
Explanation: The optimal path is 15 -> 20 -> 7 with a path sum of 15 + 20 + 7 = 42.
```

**Constraints:**

- The number of nodes in the tree is in the range `[1, 3 * 104]`.
- `-1000 <= Node.val <= 1000`

## Solution
### Approach 1
**Method Name**: Depth-First Search + Dynamic Programming

**Principle:**
Use post-order traversal to recursively compute the maximum path sum with each node as the highest point of the path, while maintaining a global maximum. For each node, the path can be in three forms: only the current node, current node plus left subtree path, or current node plus right subtree path.

**Steps:**
1. Initialize global maximum path sum with the root node's value
2. Perform depth-first search (post-order traversal):
   - Recursively compute the maximum contribution from the left subtree (discard if negative, take current node value directly)
   - Recursively compute the maximum contribution from the right subtree (discard if negative, take current node value directly)
   - Update global maximum path sum: current node value + left contribution + right contribution
   - Return the maximum contribution this node can provide to its parent (the larger of left and right contributions)
3. Return the global maximum path sum

**Example:**
- Input: `[-10,9,20,null,null,15,7]`
- Processing:
  - Node 9: returns 9, global max is 9
  - Node 15: returns 15, global max is 15
  - Node 7: returns 7, global max is 15
  - Node 20: left contribution 15, right contribution 7, path sum 15+20+7=42, global max is 42, returns max(15+20,7+20)=35
  - Node -10: left contribution 9, right contribution 35, path sum 9+(-10)+35=34, global max remains 42, returns max(9-10,35-10)=25
- Output: 42

**Complexity Analysis:**
- Time Complexity: O(n), where n is the number of nodes in the binary tree, each node is visited exactly once
- Space Complexity: O(h), where h is the height of the binary tree, recursion stack depth is O(n) in worst case (linked list shape)

> **Note**:
> 1. Node values can be negative, need to consider discarding subtrees with negative contributions
> 2. The path doesn't necessarily pass through the root, need to update global maximum at each node
> 3. Only a single path (left or right) can be returned to the parent node, cannot include both left and right subtrees simultaneously

```embed-go
PATH: "vault://leetcode/0101-0200/0124_binary_tree_maximum_path_sum/solution.go"
TITLE: "leetcode 124. Binary Tree Maximum Path Sum"
```
