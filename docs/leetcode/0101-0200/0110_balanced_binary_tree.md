---
link: https://leetcode.com/problems/balanced-binary-tree/
tags:
  - Easy
  - Tree
  - Depth-First_Search
  - Binary_Tree
---
## Description
Given a binary tree, determine if it is **height-balanced**.

**Example 1:**

![](https://assets.leetcode.com/uploads/2020/10/06/balance_1.jpg)

```
Input: root = [3,9,20,null,null,15,7]
Output: true
```

**Example 2:**

![](https://assets.leetcode.com/uploads/2020/10/06/balance_2.jpg)

```
Input: root = [1,2,2,3,3,null,null,4,4]
Output: false
```

**Example 3:**

```
Input: root = []
Output: true
```

**Constraints:**

- The number of nodes in the tree is in the range `[0, 5000]`.
- `-104 <= Node.val <= 104`

## Solution
### Approach 1
**Method Name**: Post-order Traversal with Early Termination

**Principle:**
Use post-order traversal to compute the height of each subtree while checking if the height difference between left and right subtrees exceeds 1. Use a pointer to track the balance state and terminate recursion early once imbalance is detected, avoiding unnecessary computations.

**Steps:**
1. Initialize boolean variable `balance` to `true`
2. Call helper function `height` with root node and address of `balance`
3. In the helper function:
   - If node is nil, return height 0 (base case)
   - Recursively compute left subtree height
   - If imbalance already detected, return -1 for early termination
   - Recursively compute right subtree height
   - If imbalance already detected, return -1 for early termination
   - Check height difference between left and right, set `balance` to `false` if difference > 1
   - Return current node height (max of left/right heights + 1)
4. Return the `balance` variable

```embed-go
PATH: "vault://leetcode/0101-0200/0110_balanced_binary_tree/solution.go"
TITLE: "leetcode 110. Balanced Binary Tree"
```
