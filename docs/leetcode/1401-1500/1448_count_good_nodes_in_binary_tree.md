---
link: https://leetcode.com/problems/count-good-nodes-in-binary-tree/
tags:
  - Medium
  - Tree
  - Depth-First_Search
  - Breadth-First_Search
  - Binary_Tree
---
## Description
Given a binary tree `root`, a node *X* in the tree is named **good** if in the path from root to *X* there are no nodes with a value *greater than* X.

Return the number of **good** nodes in the binary tree.

**Example 1:**

**![](https://assets.leetcode.com/uploads/2020/04/02/test_sample_1.png)**

```

Input: root = [3,1,4,3,null,1,5]
Output: 4
Explanation: Nodes in blue are good.
Root Node (3) is always a good node.
Node 4 -> (3,4) is the maximum value in the path starting from the root.
Node 5 -> (3,4,5) is the maximum value in the path
Node 3 -> (3,1,3) is the maximum value in the path.
```

**Example 2:**

**![](https://assets.leetcode.com/uploads/2020/04/02/test_sample_2.png)**

```

Input: root = [3,3,null,4,2]
Output: 3
Explanation: Node 2 -> (3, 3, 2) is not good, because "3" is higher than it.
```

**Example 3:**

```

Input: root = [1]
Output: 1
Explanation: Root is considered as good.
```

**Constraints:**

- The number of nodes in the binary tree is in the range `[1, 10^5]`.
- Each node's value is between `[-10^4, 10^4]`.

## Solution
### Approach 1
**Method Name**: Depth-First Search (DFS) + Path Maximum Tracking

**Principle:**
Traverse the binary tree using DFS while continuously tracking the maximum value along the path from the root to the current node. If the current node's value is greater than or equal to this maximum, increment the count and update the maximum.

**Steps:**
1. Start DFS traversal from the root, initializing the path maximum with the root's value
2. For each node:
   - If the node is null, return 0
   - If node value ≥ current path maximum, it's a good node, count +1 and update the maximum
   - Recursively traverse left and right subtrees, passing the updated maximum
3. Return the sum of left and right subtree counts plus the current node's result

**Example:**
- Input example: `[3,1,4,3,null,1,5]`
  - Root node 3: path max 3 → good node (count=1)
  - Left child 1: path max 3 → 1 < 3 → not good
  - Right child 4: path max 3 → 4 ≥ 3 → good node (count=2), update max to 4
  - Continue recursive traversal, accumulating 4 good nodes total

```embed-go
PATH: "vault://leetcode/1401-1500/1448_count_good_nodes_in_binary_tree/solution.go"
TITLE: "leetcode 1448. Count Good Nodes in Binary Tree"
```
