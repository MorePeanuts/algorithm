---
link: https://leetcode.cn/problems/construct-binary-tree-from-preorder-and-inorder-traversal/
tags:
  - 中等
  - 树
  - 数组
  - 哈希表
  - 分治
  - 二叉树
---
## 题目描述
给定两个整数数组 `preorder` 和 `inorder` ，其中 `preorder` 是二叉树的**先序遍历**， `inorder` 是同一棵树的**中序遍历**，请构造二叉树并返回其根节点。

**示例 1:**

![](https://assets.leetcode.com/uploads/2021/02/19/tree.jpg)

```
输入: preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
输出: [3,9,20,null,null,15,7]
```

**示例 2:**

```
输入: preorder = [-1], inorder = [-1]
输出: [-1]
```

**提示:**

- `1 <= preorder.length <= 3000`
- `inorder.length == preorder.length`
- `-3000 <= preorder[i], inorder[i] <= 3000`
- `preorder` 和 `inorder` 均 **无重复** 元素
- `inorder` 均出现在 `preorder`
- `preorder` **保证** 为二叉树的前序遍历序列
- `inorder` **保证** 为二叉树的中序遍历序列

## 题目解析
### 解法1
**方法名**：递归分治（数组切片）

**原理：**
利用前序遍历第一个元素为根节点的特性，在中序遍历中找到根节点位置，从而划分左右子树的范围，递归构建二叉树。

**步骤：**
1. 若前序遍历数组为空，返回 nil
2. 取前序遍历第一个元素作为根节点
3. 在中序遍历中找到根节点的位置 rootIdx
4. 左子树的前序遍历为 preorder[1:rootIdx+1]，中序遍历为 inorder[:rootIdx]
5. 右子树的前序遍历为 preorder[rootIdx+1:]，中序遍历为 inorder[rootIdx+1:]
6. 递归构建左右子树并连接到根节点

**示例：**
- 输入：preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
- 根节点为 3，在 inorder 中索引为 1
- 左子树：preorder = [9], inorder = [9]
- 右子树：preorder = [20,15,7], inorder = [15,20,7]
- 递归处理直至构建完整棵树

```embed-go
PATH: "vault://leetcode/0101-0200/0105_construct_binary_tree_from_preorder_and_inorder_traversal/solution.go"
TITLE: "leetcode 105. Construct Binary Tree from Preorder and Inorder Traversal"
```

### 解法2
**方法名**：递归分治（哈希表优化）

**原理：**
在解法1的基础上，使用哈希表预先存储中序遍历的值到索引的映射，避免每次递归都线性查找根节点位置；同时使用索引边界代替数组切片，减少内存开销。

**步骤：**
1. 构建哈希表 idxMap，存储 inorder 中值到索引的映射
2. 定义递归函数 dfs，参数为根节点在 preorder 中的位置、当前子树在 inorder 中的左右边界
3. 若左边界大于右边界，返回 nil
4. 创建根节点，通过哈希表快速找到根节点在 inorder 中的位置
5. 左子树的根节点在 preorder 中的位置为 rootAtPreorder + 1
6. 右子树的根节点在 preorder 中的位置为 rootAtPreorder + 1 + (rootAtInorder - leftAtInorder)
7. 递归构建左右子树并连接到根节点

**示例：**
- 输入：preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
- 构建 idxMap：{9:0, 3:1, 15:2, 20:3, 7:4}
- 初始调用 dfs(0, 0, 4)，根节点为 3，在 inorder 中索引为 1
- 左子树调用 dfs(1, 0, 0)
- 右子树调用 dfs(2, 2, 4)
- 递归处理直至构建完整棵树

```embed-go
PATH: "vault://leetcode/0101-0200/0105_construct_binary_tree_from_preorder_and_inorder_traversal/solution2.go"
TITLE: "leetcode 105. Construct Binary Tree from Preorder and Inorder Traversal"
```
