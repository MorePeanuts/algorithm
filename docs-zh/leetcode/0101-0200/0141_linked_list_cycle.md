---
link: https://leetcode.cn/problems/linked-list-cycle/
tags:
  - 简单
  - 哈希表
  - 链表
  - 双指针
---
## 题目描述
给你一个链表的头节点 `head` ，判断链表中是否有环。

如果链表中有某个节点，可以通过连续跟踪 `next` 指针再次到达，则链表中存在环。 为了表示给定链表中的环，评测系统内部使用整数 `pos` 来表示链表尾连接到链表中的位置（索引从 0 开始）。**注意：`pos` 不作为参数进行传递**。仅仅是为了标识链表的实际情况。

*如果链表中存在环* ，则返回 `true` 。 否则，返回 `false` 。

**示例 1：**

![](https://assets.leetcode.cn/aliyun-lc-upload/uploads/2018/12/07/circularlinkedlist.png)

```
输入：head = [3,2,0,-4], pos = 1
输出：true
解释：链表中有一个环，其尾部连接到第二个节点。
```

**示例 2：**

![](https://assets.leetcode.cn/aliyun-lc-upload/uploads/2018/12/07/circularlinkedlist_test2.png)

```
输入：head = [1,2], pos = 0
输出：true
解释：链表中有一个环，其尾部连接到第一个节点。
```

**示例 3：**

![](https://assets.leetcode.cn/aliyun-lc-upload/uploads/2018/12/07/circularlinkedlist_test3.png)

```
输入：head = [1], pos = -1
输出：false
解释：链表中没有环。
```

**提示：**

- 链表中节点的数目范围是 `[0, 104]`
- `-105 <= Node.val <= 105`
- `pos` 为 `-1` 或者链表中的一个 **有效索引** 。

**进阶：**你能用 `O(1)`（即，常量）内存解决此问题吗？

## 题目解析
### 解法1
**方法名**：快慢指针（Floyd判圈算法）

**原理：**
使用两个指针以不同速度遍历链表，若存在环，则快指针最终会追上慢指针；若不存在环，则快指针会先到达链表末尾。

**步骤：**
1. 初始化快指针（fast）和慢指针（slow），均指向链表头节点
2. 快指针每次移动两步，慢指针每次移动一步
3. 遍历过程中，若快指针或快指针的下一个节点为 nil，说明链表无环，返回 false
4. 若快指针与慢指针相遇，说明链表存在环，返回 true

```embed-go
PATH: "vault://leetcode/0101-0200/0141_linked_list_cycle/solution.go"
TITLE: "leetcode 141. Linked List Cycle"
```
