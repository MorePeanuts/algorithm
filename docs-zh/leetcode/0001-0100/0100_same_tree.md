---
link: https://leetcode.cn/problems/same-tree/
tags:
  - 简单
  - 树
  - 深度优先搜索
  - 广度优先搜索
  - 二叉树
---
## 题目描述
给你两棵二叉树的根节点 `p` 和 `q` ，编写一个函数来检验这两棵树是否相同。

如果两个树在结构上相同，并且节点具有相同的值，则认为它们是相同的。

**示例 1：**

![](https://assets.leetcode.com/uploads/2020/12/20/ex1.jpg)

```
输入：p = [1,2,3], q = [1,2,3]
输出：true
```

**示例 2：**

![](https://assets.leetcode.com/uploads/2020/12/20/ex2.jpg)

```
输入：p = [1,2], q = [1,null,2]
输出：false
```

**示例 3：**

![](https://assets.leetcode.com/uploads/2020/12/20/ex3.jpg)

```
输入：p = [1,2,1], q = [1,1,2]
输出：false
```

**提示：**

- 两棵树上的节点数目都在范围 `[0, 100]` 内
- `-104 <= Node.val <= 104`

## 题目解析
### 解法1
**方法名**：深度优先搜索

**原理：**
通过递归方式，同时遍历两棵树的每个节点，比较节点值是否相等，并递归比较左右子树。

**步骤：**
1. 如果两个节点都为 nil，说明两棵树在该位置结构相同，返回 true
2. 如果其中一个节点为 nil 而另一个不为 nil，说明结构不同，返回 false
3. 比较当前节点的值是否相等
4. 递归比较左子树和右子树

```embed-go
PATH: "vault://leetcode/0001-0100/0100_same_tree/solution.go"
TITLE: "leetcode 100. Same Tree"
```
