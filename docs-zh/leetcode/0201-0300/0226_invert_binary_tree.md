---
link: https://leetcode.cn/problems/invert-binary-tree/
tags:
  - 简单
  - 树
  - 深度优先搜索
  - 广度优先搜索
  - 二叉树
---
## 题目描述
给你一棵二叉树的根节点 `root` ，翻转这棵二叉树，并返回其根节点。

**示例 1：**

![](https://assets.leetcode.com/uploads/2021/03/14/invert1-tree.jpg)

```
输入：root = [4,2,7,1,3,6,9]
输出：[4,7,2,9,6,3,1]
```

**示例 2：**

![](https://assets.leetcode.com/uploads/2021/03/14/invert2-tree.jpg)

```
输入：root = [2,1,3]
输出：[2,3,1]
```

**示例 3：**

```
输入：root = []
输出：[]
```

**提示：**

- 树中节点数目范围在 `[0, 100]` 内
- `-100 <= Node.val <= 100`

## 题目解析
### 解法1
**方法名**：递归翻转（后序遍历）

**原理：**
通过递归方式，先分别翻转左子树和右子树，再交换当前节点的左右子节点。

**步骤：**
1. 如果当前节点为空，直接返回 nil
2. 递归翻转左子树
3. 递归翻转右子树
4. 交换当前节点的左右子节点
5. 返回当前节点

```embed-go
PATH: "vault://leetcode/0201-0300/0226_invert_binary_tree/solution.go"
TITLE: "leetcode 226.翻转二叉树"
```

### 解法2
**方法名**：递归翻转（前序遍历）

**原理：**
通过递归方式，先交换当前节点的左右子节点，再分别翻转左子树和右子树。

**步骤：**
1. 如果当前节点为空，直接返回 nil
2. 交换当前节点的左右子节点
3. 递归翻转左子树
4. 递归翻转右子树
5. 返回当前节点

```embed-go
PATH: "vault://leetcode/0201-0300/0226_invert_binary_tree/solution2.go"
TITLE: "leetcode 226.翻转二叉树"
```
