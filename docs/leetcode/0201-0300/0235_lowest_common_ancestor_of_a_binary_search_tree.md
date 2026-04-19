---
link: https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-search-tree/
tags:
  - Medium
  - Tree
  - Depth-First_Search
  - Binary_Search_Tree
  - Binary_Tree
---
## Description
Given a binary search tree (BST), find the lowest common ancestor (LCA) node of two given nodes in the BST.

According to the [definition of LCA on Wikipedia](https://en.wikipedia.org/wiki/Lowest_common_ancestor): “The lowest common ancestor is defined between two nodes `p` and `q` as the lowest node in `T` that has both `p` and `q` as descendants (where we allow **a node to be a descendant of itself**).”

**Example 1:**

![](https://assets.leetcode.com/uploads/2018/12/14/binarysearchtree_improved.png)

```
Input: root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 8
Output: 6
Explanation: The LCA of nodes 2 and 8 is 6.
```

**Example 2:**

![](https://assets.leetcode.com/uploads/2018/12/14/binarysearchtree_improved.png)

```
Input: root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 4
Output: 2
Explanation: The LCA of nodes 2 and 4 is 2, since a node can be a descendant of itself according to the LCA definition.
```

**Example 3:**

```
Input: root = [2,1], p = 2, q = 1
Output: 2
```

**Constraints:**

- The number of nodes in the tree is in the range `[2, 105]`.
- `-109 <= Node.val <= 109`
- All `Node.val` are **unique**.
- `p != q`
- `p` and `q` will exist in the BST.

## Solution
### Approach 1
**Method Name**: Path Finding

**Principle:**
Obtain the paths from root to p and root to q respectively, then compare the two paths to find the last common node, which is the lowest common ancestor.

**Steps:**
1. Starting from root, traverse to target node p using BST property (left subtree nodes < root, right subtree nodes > root), record the path
2. Obtain the path to q using the same method
3. Traverse both paths simultaneously to find the last common node

**Example:**
- Input: root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 4
- Path to p: [6, 2]
- Path to q: [6, 2, 4]
- Compare paths, last common node is 2, output 2

```embed-go
PATH: "vault://leetcode/0201-0300/0235_lowest_common_ancestor_of_a_binary_search_tree/solution.go"
TITLE: "leetcode 235. Lowest Common Ancestor of a Binary Search Tree"
```

### Approach 2
**Method Name**: Recursive Search

**Principle:**
Utilize BST property to search directly in the tree: if both p and q are in the right subtree, recursively go right; if both are in the left subtree, recursively go left; otherwise current node is the lowest common ancestor.

**Steps:**
1. If both p.Val and q.Val > root.Val, both nodes are in right subtree, recursively process right subtree
2. If both p.Val and q.Val < root.Val, both nodes are in left subtree, recursively process left subtree
3. Otherwise, current node is the divergence point, which is the lowest common ancestor

**Example:**
- Input: root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 8
- p.Val=2 < 6, q.Val=8 > 6, divergence condition met, return 6 directly

```embed-go
PATH: "vault://leetcode/0201-0300/0235_lowest_common_ancestor_of_a_binary_search_tree/solution2.go"
TITLE: "leetcode 235. Lowest Common Ancestor of a Binary Search Tree"
```

### Approach 3
**Method Name**: Iterative Search

**Principle:**
Same idea as recursive solution, uses iterative approach to traverse the tree and find the divergence point of p and q.

**Steps:**
1. Ensure p.Val <= q.Val to simplify subsequent judgments
2. Start traversing from root
3. If p.Val > current node value, both nodes are in right subtree, move right
4. If q.Val < current node value, both nodes are in left subtree, move left
5. Otherwise, current node is the lowest common ancestor

**Example:**
- Input: root = [6,2,8,0,4,7,9,null,null,3,5], p = 4, q = 5
- Ensure p=4, q=5
- 4 < 6 and 5 < 6, move left to 2
- 4 > 2 and 5 > 2, move right to 4
- Now p=4 is not greater than current node value, q=5 is not less than current node value, return 4

```embed-go
PATH: "vault://leetcode/0201-0300/0235_lowest_common_ancestor_of_a_binary_search_tree/solution3.go"
TITLE: "leetcode 235. Lowest Common Ancestor of a Binary Search Tree"
```
