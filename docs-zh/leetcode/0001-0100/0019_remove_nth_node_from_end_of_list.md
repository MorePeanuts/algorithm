---
link: https://leetcode.cn/problems/remove-nth-node-from-end-of-list/
tags:
  - 中等
  - 链表
  - 双指针
---
## 题目描述
给你一个链表，删除链表的倒数第 `n`个结点，并且返回链表的头结点。

**示例 1：**

![](https://assets.leetcode.com/uploads/2020/10/03/remove_ex1.jpg)

```
输入：head = [1,2,3,4,5], n = 2
输出：[1,2,3,5]
```

**示例 2：**

```
输入：head = [1], n = 1
输出：[]
```

**示例 3：**

```
输入：head = [1,2], n = 1
输出：[1]
```

**提示：**

- 链表中结点的数目为 `sz`
- `1 <= sz <= 30`
- `0 <= Node.val <= 100`
- `1 <= n <= sz`

**进阶：**你能尝试使用一趟扫描实现吗？

## 题目解析
### 解法1
**方法名**：双指针一趟扫描

**原理：**
使用快慢两个指针，快指针先向前移动 n 步，然后快慢指针同时移动，直到快指针到达链表末尾。此时慢指针恰好指向倒数第 n 个节点的前驱节点。

**步骤：**
1. 初始化慢指针 slow 和快指针 fast 都指向头节点 head
2. 快指针 fast 先向前移动 n 步
3. 初始化 prev 指针为 nil，用于记录慢指针的前驱节点
4. 同时移动快指针和慢指针，直到快指针 fast 到达 nil
5. 如果 prev 不为 nil，说明要删除的不是头节点，将 prev.Next 指向 slow.Next
6. 如果 prev 为 nil，说明要删除的是头节点，直接将 head 指向 slow.Next
7. 返回头节点 head

**示例：**
- 输入：head = [1,2,3,4,5], n = 2
  - 快指针先移动 2 步：fast 指向节点 3
  - 同时移动快慢指针直到 fast 为 nil：slow 指向节点 4，prev 指向节点 3
  - 删除节点：prev.Next = slow.Next → 节点 3 的 Next 指向节点 5
- 输出：[1,2,3,5]

```embed-go
PATH: "vault://leetcode/0001-0100/0019_remove_nth_node_from_end_of_list/solution.go"
TITLE: "leetcode 19. Remove Nth Node From End of List"
```
