---
link: https://leetcode.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/
tags:
  - Medium
  - Tree
  - Array
  - Hash_Table
  - Divide_and_Conquer
  - Binary_Tree
---
## Description
Given two integer arrays `preorder` and `inorder` where `preorder` is the preorder traversal of a binary tree and `inorder` is the inorder traversal of the same tree, construct and return *the binary tree*.

**Example 1:**

![](https://assets.leetcode.com/uploads/2021/02/19/tree.jpg)

```
Input: preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
Output: [3,9,20,null,null,15,7]
```

**Example 2:**

```
Input: preorder = [-1], inorder = [-1]
Output: [-1]
```

**Constraints:**

- `1 <= preorder.length <= 3000`
- `inorder.length == preorder.length`
- `-3000 <= preorder[i], inorder[i] <= 3000`
- `preorder` and `inorder` consist of **unique** values.
- Each value of `inorder` also appears in `preorder`.
- `preorder` is **guaranteed** to be the preorder traversal of the tree.
- `inorder` is **guaranteed** to be the inorder traversal of the tree.

## Solution
### Approach 1
**Method Name**: Recursive Divide and Conquer (Array Slicing)

**Principle:**
Using the property that the first element in preorder traversal is the root node, find the root position in inorder traversal to divide the left and right subtree ranges, then recursively construct the binary tree.

**Steps:**
1. If preorder array is empty, return nil
2. Take the first element of preorder as the root node
3. Find the root position rootIdx in inorder traversal
4. Left subtree preorder is preorder[1:rootIdx+1], inorder is inorder[:rootIdx]
5. Right subtree preorder is preorder[rootIdx+1:], inorder is inorder[rootIdx+1:]
6. Recursively construct left and right subtrees and connect to root

**Example:**
- Input: preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
- Root is 3, index 1 in inorder
- Left subtree: preorder = [9], inorder = [9]
- Right subtree: preorder = [20,15,7], inorder = [15,20,7]
- Recursively process until the entire tree is constructed

```embed-go
PATH: "vault://leetcode/0101-0200/0105_construct_binary_tree_from_preorder_and_inorder_traversal/solution.go"
TITLE: "leetcode 105. Construct Binary Tree from Preorder and Inorder Traversal"
```

### Approach 2
**Method Name**: Recursive Divide and Conquer (Hash Map Optimization)

**Principle:**
Based on Approach 1, use a hash map to pre-store the mapping from values to indices in inorder traversal to avoid linear search for root position in each recursion; meanwhile use index boundaries instead of array slicing to reduce memory overhead.

**Steps:**
1. Build hash map idxMap, storing the mapping from values to indices in inorder
2. Define recursive function dfs with parameters: root position in preorder, left and right boundaries of current subtree in inorder
3. If left boundary > right boundary, return nil
4. Create root node, quickly find root position in inorder through hash map
5. Left subtree root position in preorder is rootAtPreorder + 1
6. Right subtree root position in preorder is rootAtPreorder + 1 + (rootAtInorder - leftAtInorder)
7. Recursively construct left and right subtrees and connect to root

**Example:**
- Input: preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
- Build idxMap: {9:0, 3:1, 15:2, 20:3, 7:4}
- Initial call dfs(0, 0, 4), root is 3, index 1 in inorder
- Left subtree calls dfs(1, 0, 0)
- Right subtree calls dfs(2, 2, 4)
- Recursively process until the entire tree is constructed

```embed-go
PATH: "vault://leetcode/0101-0200/0105_construct_binary_tree_from_preorder_and_inorder_traversal/solution2.go"
TITLE: "leetcode 105. Construct Binary Tree from Preorder and Inorder Traversal"
```
