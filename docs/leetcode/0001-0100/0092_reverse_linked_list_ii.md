---
link: https://leetcode.com/problems/reverse-linked-list-ii/
tags:
  - Medium
  - Linked_List
---
## Description
Given the `head` of a singly linked list and two integers `left` and `right` where `left <= right`, reverse the nodes of the list from position `left` to position `right`, and return *the reversed list*.

**Example 1:**

![](https://assets.leetcode.com/uploads/2021/02/19/rev2ex2.jpg)

```
Input: head = [1,2,3,4,5], left = 2, right = 4
Output: [1,4,3,2,5]
```

**Example 2:**

```
Input: head = [5], left = 1, right = 1
Output: [5]
```

**Constraints:**

- The number of nodes in the list is `n`.
- `1 <= n <= 500`
- `-500 <= Node.val <= 500`
- `1 <= left <= right <= n`

**Follow up:** Could you do it in one pass?

## Solution
### Approach 1
**Method Name**: One-pass Linked List Reversal

**Principle:**
First locate the start of the reversal interval, then reverse the pointers within that interval, and finally adjust the connections before and after. The entire process only requires one pass.

**Steps:**
1. Find the node before the reversal interval `p1` and the start node `p2` (the `left`-th node)
2. Reverse the pointer direction one by one from `left` to `right`
3. Adjust the connections of the reversed interval: point the original start node's `Next` to the node after `right`, and point `p1` (or `head` if `left=1`) to the new head of the reversed interval

**Example:**
- Input `head = [1,2,3,4,5], left = 2, right = 4`
  - Step 1: Find `p1=1`, `p2=2`
  - Step 2: Reverse interval `2-4`, get `4->3->2`
  - Step 3: Connect `1->4` and `2->5`
- Output `[1,4,3,2,5]`

```embed-go
PATH: "vault://leetcode/0001-0100/0092_reverse_linked_list_ii/solution.go"
TITLE: "leetcode 92. Reverse Linked List II"
```
