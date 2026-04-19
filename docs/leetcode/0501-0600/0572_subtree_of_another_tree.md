---
link: https://leetcode.com/problems/subtree-of-another-tree/
tags:
  - Easy
  - Tree
  - Depth-First_Search
  - Binary_Tree
  - String_Matching
  - Hash_Function
---
## Description
Given the roots of two binary trees `root` and `subRoot`, return `true` if there is a subtree of `root` with the same structure and node values of `subRoot` and `false` otherwise.

A subtree of a binary tree `tree` is a tree that consists of a node in `tree` and all of this node's descendants. The tree `tree` could also be considered as a subtree of itself.

**Example 1:**

![](https://assets.leetcode.com/uploads/2021/04/28/subtree1-tree.jpg)

```
Input: root = [3,4,5,1,2], subRoot = [4,1,2]
Output: true
```

**Example 2:**

![](https://assets.leetcode.com/uploads/2021/04/28/subtree2-tree.jpg)

```
Input: root = [3,4,5,1,2,null,null,null,null,0], subRoot = [4,1,2]
Output: false
```

**Constraints:**

- The number of nodes in the `root` tree is in the range `[1, 2000]`.
- The number of nodes in the `subRoot` tree is in the range `[1, 1000]`.
- `-104 <= root.val <= 104`
- `-104 <= subRoot.val <= 104`

## Solution
### Approach 1
**Method Name**: DFS + Same Tree Check

**Principle:**
Traverse each node of the main tree via DFS, and for each node, check if the subtree rooted at that node is identical to the target subtree.

**Steps:**
1. If current node of main tree is nil, only match if subtree is also nil
2. Check if tree starting at current node is identical to subtree (using `isSameTree` function)
3. If no match, recursively check left subtree
4. If left subtree doesn't match, recursively check right subtree
5. Return false if none match

```embed-go
PATH: "vault://leetcode/0501-0600/0572_subtree_of_another_tree/solution.go"
TITLE: "leetcode 572. Subtree of Another Tree"
```

### Approach 2
**Method Name**: Height-Optimized DFS

**Principle:**
First compute the height of the subtree, then only perform same-tree checks on nodes with matching height during main tree traversal, reducing unnecessary comparisons.

**Steps:**
1. First compute the height of the target subtree
2. Post-order traverse the main tree while computing each node's height
3. Only perform same-tree check when a node's height matches the subtree height
4. Early return if match is found

```embed-go
PATH: "vault://leetcode/0501-0600/0572_subtree_of_another_tree/solution2.go"
TITLE: "leetcode 572. Subtree of Another Tree"
```
