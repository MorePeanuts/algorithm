---
link: https://leetcode.cn/problems/subtree-of-another-tree/
tags:
  - 简单
  - 树
  - 深度优先搜索
  - 二叉树
  - 字符串匹配
  - 哈希函数
---
## 题目描述
给你两棵二叉树 `root` 和 `subRoot` 。检验 `root` 中是否包含和 `subRoot` 具有相同结构和节点值的子树。如果存在，返回 `true` ；否则，返回 `false` 。

二叉树 `tree` 的一棵子树包括 `tree` 的某个节点和这个节点的所有后代节点。`tree` 也可以看做它自身的一棵子树。

**示例 1：**

![](https://pic.leetcode.cn/1724998676-cATjhe-image.png)

```
输入：root = [3,4,5,1,2], subRoot = [4,1,2]
输出：true
```

**示例 2：**

![](https://pic.leetcode.cn/1724998698-sEJWnq-image.png)

```
输入：root = [3,4,5,1,2,null,null,null,null,0], subRoot = [4,1,2]
输出：false
```

**提示：**

- `root` 树上的节点数量范围是 `[1, 2000]`
- `subRoot` 树上的节点数量范围是 `[1, 1000]`
- `-104 <= root.val <= 104`
- `-104 <= subRoot.val <= 104`

## 题目解析
### 解法1
**方法名**：深度优先搜索 + 相同树判断

**原理：**
通过深度优先搜索遍历主树的每个节点，对于每个节点，判断以该节点为根的子树是否与目标子树完全相同。

**步骤：**
1. 如果主树当前节点为空，则只有当子树也为空时才匹配
2. 判断当前节点开始的树是否与子树相同（使用 `isSameTree` 函数）
3. 如果不匹配，递归检查左子树
4. 如果左子树不匹配，递归检查右子树
5. 都不匹配则返回 false

```embed-go
PATH: "vault://leetcode/0501-0600/0572_subtree_of_another_tree/solution.go"
TITLE: "leetcode 572. Subtree of Another Tree"
```

### 解法2
**方法名**：高度优化的深度优先搜索

**原理：**
先计算子树的高度，然后在遍历主树时只对高度匹配的节点进行相同树判断，减少不必要的比较。

**步骤：**
1. 先计算目标子树的高度
2. 后序遍历主树，同时计算每个节点的高度
3. 当某节点的高度与子树高度相同时，才进行相同树判断
4. 如果找到匹配，提前返回结果

```embed-go
PATH: "vault://leetcode/0501-0600/0572_subtree_of_another_tree/solution2.go"
TITLE: "leetcode 572. Subtree of Another Tree"
```
