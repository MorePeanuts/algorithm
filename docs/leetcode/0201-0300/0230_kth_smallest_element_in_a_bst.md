---
link: https://leetcode.com/problems/kth-smallest-element-in-a-bst/
tags:
  - Medium
  - Tree
  - Depth-First_Search
  - Binary_Search_Tree
  - Binary_Tree
---
## Description
Given the `root` of a binary search tree, and an integer `k`, return *the* `kth` *smallest value (**1-indexed**) of all the values of the nodes in the tree*.

**Example 1:**

![](https://assets.leetcode.com/uploads/2021/01/28/kthtree1.jpg)

```
Input: root = [3,1,4,null,2], k = 1
Output: 1
```

**Example 2:**

![](https://assets.leetcode.com/uploads/2021/01/28/kthtree2.jpg)

```
Input: root = [5,3,6,2,4,null,null,1], k = 3
Output: 3
```

**Constraints:**

- The number of nodes in the tree is `n`.
- `1 <= k <= n <= 104`
- `0 <= Node.val <= 104`

**Follow up:** If the BST is modified often (i.e., we can do insert and delete operations) and you need to find the kth smallest frequently, how would you optimize?

## Solution
### Approach 1
**Method Name**: Inorder Traversal with Early Termination

**Principle:**
Leverage the property of Binary Search Tree (BST) where inorder traversal yields a strictly increasing sequence. By visiting nodes in inorder and counting, we can return the answer as soon as we reach the kth node, allowing early termination.

**Steps:**
1. Define a counter `count` to record the number of visited nodes
2. Perform inorder traversal (left-root-right):
   - First recursively traverse the left subtree
   - Increment counter when visiting current node
   - If counter equals k, return current node value and terminate traversal
   - If not found, continue recursively traversing the right subtree
3. Use a helper function returning two values (node value, found) to achieve early termination

**Example:**
Take Example 1: `root = [3,1,4,null,2], k = 1`
- Traversal order: 1 → 2 → 3 → 4
- The 1st visited node is 1, return 1 directly

```embed-go
PATH: "vault://leetcode/0201-0300/0230_kth_smallest_element_in_a_bst/solution.go"
TITLE: "leetcode 230. Kth Smallest Element in a BST"
```
