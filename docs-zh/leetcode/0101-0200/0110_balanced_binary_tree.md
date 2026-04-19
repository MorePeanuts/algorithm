---
link: https://leetcode.cn/problems/balanced-binary-tree/
tags:
  - 简单
  - 树
  - 深度优先搜索
  - 二叉树
---
## 题目描述
给定一个二叉树，判断它是否是 平衡二叉树

**示例 1：**

![](https://assets.leetcode.com/uploads/2020/10/06/balance_1.jpg)

```
输入：root = [3,9,20,null,null,15,7]
输出：true
```

**示例 2：**

![](https://assets.leetcode.com/uploads/2020/10/06/balance_2.jpg)

```
输入：root = [1,2,2,3,3,null,null,4,4]
输出：false
```

**示例 3：**

```
输入：root = []
输出：true
```

**提示：**

- 树中的节点数在范围 `[0, 5000]` 内
- `-104 <= Node.val <= 104`

## 题目解析
### 解法1
**方法名**：后序遍历 + 提前终止

**原理：**
通过后序遍历计算每个子树的高度，同时检查左右子树的高度差是否超过 1。使用指针跟踪平衡状态，一旦发现不平衡立即提前终止递归，避免不必要的计算。

**步骤：**
1. 初始化布尔变量 `balance` 为 `true`
2. 调用辅助函数 `height`，传入根节点和 `balance` 的地址
3. 在辅助函数中：
   - 若节点为空，返回高度 0（基准情况）
   - 递归计算左子树高度
   - 若已发现不平衡，返回 -1 提前终止
   - 递归计算右子树高度
   - 若已发现不平衡，返回 -1 提前终止
   - 检查左右子树高度差，若超过 1 则设置 `balance` 为 `false`
   - 返回当前节点的高度（左右子树最大高度 + 1）
4. 返回 `balance` 变量

```embed-go
PATH: "vault://leetcode/0101-0200/0110_balanced_binary_tree/solution.go"
TITLE: "leetcode 110.平衡二叉树"
```
