---
link: https://leetcode.cn/problems/reverse-linked-list/
tags:
  - 简单
  - 递归
  - 链表
---
## 题目描述
给你单链表的头节点 `head` ，请你反转链表，并返回反转后的链表。

**示例 1：**

![](https://assets.leetcode.com/uploads/2021/02/19/rev1ex1.jpg)

```
输入：head = [1,2,3,4,5]
输出：[5,4,3,2,1]
```

**示例 2：**

![](https://assets.leetcode.com/uploads/2021/02/19/rev1ex2.jpg)

```
输入：head = [1,2]
输出：[2,1]
```

**示例 3：**

```
输入：head = []
输出：[]
```

**提示：**

- 链表中节点的数目范围是 `[0, 5000]`
- `-5000 <= Node.val <= 5000`

**进阶：**链表可以选用迭代或递归方式完成反转。你能否用两种方法解决这道题？

## 题目解析
### 解法1
**方法名**：迭代法（三指针遍历）

**原理：**
使用三个指针遍历链表，逐个改变节点的指针方向，将当前节点的 Next 指向前一个节点，从而实现链表反转。

**步骤：**
1. 初始化三个指针：p1（前一个节点）为 nil，p2（当前节点）为 head
2. 当 p2 不为空时，继续遍历：
   - 保存 p2 的下一个节点到 p3
   - 将 p2 的 Next 指向 p1（反转指针方向）
   - p1 移动到 p2 的位置
   - p2 移动到 p3 的位置
3. 当 p2 为空时，p1 即为新的头节点，返回 p1

```embed-go
PATH: "vault://leetcode/0201-0300/0206_reverse_linked_list/solution.go"
TITLE: "leetcode 206.反转链表"
```
