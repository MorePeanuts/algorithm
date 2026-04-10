---
link: https://leetcode.com/problems/linked-list-cycle/
tags:
  - Easy
  - Hash_Table
  - Linked_List
  - Two_Pointers
---
## Description
Given `head`, the head of a linked list, determine if the linked list has a cycle in it.

There is a cycle in a linked list if there is some node in the list that can be reached again by continuously following the `next` pointer. Internally, `pos` is used to denote the index of the node that tail's `next` pointer is connected to. **Note that `pos` is not passed as a parameter**.

Return `true` *if there is a cycle in the linked list*. Otherwise, return `false`.

**Example 1:**

![](https://assets.leetcode.com/uploads/2018/12/07/circularlinkedlist.png)

```
Input: head = [3,2,0,-4], pos = 1
Output: true
Explanation: There is a cycle in the linked list, where the tail connects to the 1st node (0-indexed).
```

**Example 2:**

![](https://assets.leetcode.com/uploads/2018/12/07/circularlinkedlist_test2.png)

```
Input: head = [1,2], pos = 0
Output: true
Explanation: There is a cycle in the linked list, where the tail connects to the 0th node.
```

**Example 3:**

![](https://assets.leetcode.com/uploads/2018/12/07/circularlinkedlist_test3.png)

```
Input: head = [1], pos = -1
Output: false
Explanation: There is no cycle in the linked list.
```

**Constraints:**

- The number of the nodes in the list is in the range `[0, 104]`.
- `-105 <= Node.val <= 105`
- `pos` is `-1` or a **valid index** in the linked-list.

**Follow up:** Can you solve it using `O(1)` (i.e. constant) memory?

## Solution
### Approach 1
**Method Name**: Fast and Slow Pointers (Floyd's Tortoise and Hare Algorithm)

**Principle:**
Use two pointers traversing the linked list at different speeds. If there's a cycle, the fast pointer will eventually catch up to the slow pointer; if not, the fast pointer will reach the end of the list first.

**Steps:**
1. Initialize both fast pointer and slow pointer to the head of the linked list
2. Fast pointer moves two steps at a time, slow pointer moves one step at a time
3. During traversal, if fast pointer or fast pointer's next node is nil, there's no cycle, return false
4. If fast pointer meets slow pointer, there's a cycle, return true

```embed-go
PATH: "vault://leetcode/0101-0200/0141_linked_list_cycle/solution.go"
TITLE: "leetcode 141. Linked List Cycle"
```
