---
link: https://leetcode.com/problems/binary-tree-right-side-view/
tags:
  - Medium
  - Tree
  - Depth-First_Search
  - Breadth-First_Search
  - Binary_Tree
---
## Description
Given the `root` of a binary tree, imagine yourself standing on the **right side** of it, return *the values of the nodes you can see ordered from top to bottom*.

**Example 1:**

**Input:** root = [1,2,3,null,5,null,4]

**Output:** [1,3,4]

**Explanation:**

![](https://assets.leetcode.com/uploads/2024/11/24/tmpd5jn43fs-1.png)

**Example 2:**

**Input:** root = [1,2,3,4,null,null,null,5]

**Output:** [1,3,4,5]

**Explanation:**

![](https://assets.leetcode.com/uploads/2024/11/24/tmpkpe40xeh-1.png)

**Example 3:**

**Input:** root = [1,null,3]

**Output:** [1,3]

**Example 4:**

**Input:** root = []

**Output:** []

**Constraints:**

- The number of nodes in the tree is in the range `[0, 100]`.
- `-100 <= Node.val <= 100`

## Solution
### Approach 1
**Method Name**: Breadth-First Search (Level Order Traversal)

**Principle:**
Traverse the tree level by level, and for each level, only keep the rightmost node's value.

**Steps:**
1. Initialize a queue with the root node if it exists.
2. While the queue is not empty, process each level:
   a. Record the number of nodes in the current level.
   b. For each node in the level, if it's the last node (rightmost), add its value to the result.
   c. Add the left and right children of each node to the queue for the next level.

**Example:**
- Input: [1,2,3,null,5,null,4]
- Level 0: [1] → add 1
- Level 1: [2,3] → add 3
- Level 2: [5,4] → add 4
- Output: [1,3,4]

```embed-go
PATH: "vault://leetcode/0101-0200/0199_binary_tree_right_side_view/solution.go"
TITLE: "leetcode 199. Binary Tree Right Side View"
```

### Approach 2
**Method Name**: Depth-First Search (Right-First Preorder)

**Principle:**
Traverse the tree using DFS, visiting the right child first. The first node encountered at each depth is the rightmost node for that level.

**Steps:**
1. Perform a preorder traversal (root → right → left).
2. At each node, if the current depth equals the length of the result list, this is the first node at this depth (rightmost), so add its value to the result.
3. Recursively visit the right child first, then the left child, incrementing the depth each time.

**Example:**
- Input: [1,2,3,null,5,null,4]
- Depth 0: node 1 → add 1
- Depth 1: node 3 → add 3
- Depth 2: node 4 → add 4
- Output: [1,3,4]

```embed-go
PATH: "vault://leetcode/0101-0200/0199_binary_tree_right_side_view/solution2.go"
TITLE: "leetcode 199. Binary Tree Right Side View"
```
