---
link: https://leetcode.cn/problems/maximum-depth-of-binary-tree/
tags:
  - 简单
  - 树
  - 深度优先搜索
  - 广度优先搜索
  - 二叉树
---
## 题目描述
给定一个二叉树 `root` ，返回其最大深度。

二叉树的 **最大深度** 是指从根节点到最远叶子节点的最长路径上的节点数。

**示例 1：**

![](https://assets.leetcode.com/uploads/2020/11/26/tmp-tree.jpg)

```
输入：root = [3,9,20,null,null,15,7]
输出：3
```

**示例 2：**

```
输入：root = [1,null,2]
输出：2
```

**提示：**

- 树中节点的数量在 `[0, 104]` 区间内。
- `-100 <= Node.val <= 100`

## 题目解析
### 解法1
**方法名**：深度优先搜索（递归）

**原理：**
通过递归分别计算左右子树的最大深度，取较大值加1即为当前树的最大深度。

**步骤：**
1. 如果当前节点为空，返回深度0
2. 递归计算左子树的深度
3. 递归计算右子树的深度
4. 返回左右子树深度的最大值加1

```embed-go
PATH: "vault://leetcode/0101-0200/0104_maximum_depth_of_binary_tree/solution.go"
TITLE: "leetcode 104. Maximum Depth of Binary Tree"
```

### 解法2
**方法名**：广度优先搜索（迭代）

**原理：**
使用队列进行层序遍历，每遍历完一层，深度加1。

**步骤：**
1. 如果根节点为空，返回深度0
2. 将根节点加入队列
3. 当队列不为空时，记录当前层的节点数
4. 遍历当前层的所有节点，将它们的子节点加入队列
5. 每完成一层的遍历，深度加1
6. 遍历结束后返回深度

```embed-go
PATH: "vault://leetcode/0101-0200/0104_maximum_depth_of_binary_tree/solution2.go"
TITLE: "leetcode 104. Maximum Depth of Binary Tree"
```
