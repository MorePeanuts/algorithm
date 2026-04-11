---
link: https://leetcode.com/problems/remove-nth-node-from-end-of-list/
tags:
  - Medium
  - Linked_List
  - Two_Pointers
---
## Description
Given the `head` of a linked list, remove the `nth` node from the end of the list and return its head.

**Example 1:**

![](https://assets.leetcode.com/uploads/2020/10/03/remove_ex1.jpg)

```
Input: head = [1,2,3,4,5], n = 2
Output: [1,2,3,5]
```

**Example 2:**

```
Input: head = [1], n = 1
Output: []
```

**Example 3:**

```
Input: head = [1,2], n = 1
Output: [1]
```

**Constraints:**

- The number of nodes in the list is `sz`.
- `1 <= sz <= 30`
- `0 <= Node.val <= 100`
- `1 <= n <= sz`

**Follow up:** Could you do this in one pass?

## Solution
### Approach 1
**Method Name**: Two Pointers One Pass

**Principle:**
Use two pointers, fast and slow. The fast pointer moves n steps first, then both pointers move together until the fast pointer reaches the end of the list. At this point, the slow pointer exactly points to the predecessor of the nth node from the end.

**Steps:**
1. Initialize both slow pointer and fast pointer to head node
2. Move fast pointer n steps forward first
3. Initialize prev pointer to nil to record the predecessor of slow pointer
4. Move both fast and slow pointers together until fast reaches nil
5. If prev is not nil, the node to delete is not the head node, set prev.Next to slow.Next
6. If prev is nil, the node to delete is the head node, set head to slow.Next directly
7. Return head node

**Example:**
- Input: head = [1,2,3,4,5], n = 2
  - Fast pointer moves 2 steps first: fast points to node 3
  - Move both pointers until fast is nil: slow points to node 4, prev points to node 3
  - Delete node: prev.Next = slow.Next → node 3's Next points to node 5
- Output: [1,2,3,5]

```embed-go
PATH: "vault://leetcode/0001-0100/0019_remove_nth_node_from_end_of_list/solution.go"
TITLE: "leetcode 19. Remove Nth Node From End of List"
```
