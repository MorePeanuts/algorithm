---
link: https://leetcode.cn/problems/reverse-linked-list-ii/
tags:
  - 中等
  - 链表
---
## 题目描述
给你单链表的头指针 `head` 和两个整数 `left` 和 `right` ，其中 `left <= right` 。请你反转从位置 `left` 到位置 `right` 的链表节点，返回 **反转后的链表** 。

**示例 1：**

![](https://assets.leetcode.com/uploads/2021/02/19/rev2ex2.jpg)

```
输入：head = [1,2,3,4,5], left = 2, right = 4
输出：[1,4,3,2,5]
```

**示例 2：**

```
输入：head = [5], left = 1, right = 1
输出：[5]
```

**提示：**

- 链表中节点数目为 `n`
- `1 <= n <= 500`
- `-500 <= Node.val <= 500`
- `1 <= left <= right <= n`

**进阶：** 你可以使用一趟扫描完成反转吗？

## 题目解析
### 解法1
**方法名**：一趟扫描反转链表

**原理：**
先找到需要反转区间的起始位置，然后在该区间内进行指针反转，最后调整前后连接关系，整个过程只需要一趟扫描。

**步骤：**
1. 找到反转区间的前一个节点 `p1` 和起始节点 `p2`（即第 `left` 个节点）
2. 从 `left` 到 `right` 区间内逐个反转指针方向
3. 调整反转后区间的前后连接：将原起始节点的 `Next` 指向 `right` 之后的节点，将 `p1`（或 `head` 若 `left=1`）指向反转后的区间头节点

**示例：**
- 输入 `head = [1,2,3,4,5], left = 2, right = 4`
  - 第一步：找到 `p1=1`, `p2=2`
  - 第二步：反转区间 `2-4`，得到 `4->3->2`
  - 第三步：连接 `1->4` 和 `2->5`
- 输出 `[1,4,3,2,5]`

```embed-go
PATH: "vault://leetcode/0001-0100/0092_reverse_linked_list_ii/solution.go"
TITLE: "leetcode 92. Reverse Linked List II"
```
