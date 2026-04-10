---
link: https://leetcode.com/problems/reorder-list/
tags:
  - Medium
  - Stack
  - Recursion
  - Linked_List
  - Two_Pointers
---
## Description
You are given the head of a singly linked-list. The list can be represented as:

```
L0 → L1 → … → Ln - 1 → Ln
```

*Reorder the list to be on the following form:*

```
L0 → Ln → L1 → Ln - 1 → L2 → Ln - 2 → …
```

You may not modify the values in the list's nodes. Only nodes themselves may be changed.

**Example 1:**

![](https://assets.leetcode.com/uploads/2021/03/04/reorder1linked-list.jpg)

```
Input: head = [1,2,3,4]
Output: [1,4,2,3]
```

**Example 2:**

![](https://assets.leetcode.com/uploads/2021/03/09/reorder2-linked-list.jpg)

```
Input: head = [1,2,3,4,5]
Output: [1,5,2,4,3]
```

**Constraints:**

- The number of nodes in the list is in the range `[1, 5 * 104]`.
- `1 <= Node.val <= 1000`

## Solution
### Approach 1
**Method Name**: Slow and Fast Pointers + Reverse Linked List + Cross Merge

**Principle:**
Find the middle of the linked list using slow and fast pointers, split the list into two halves, reverse the second half, then cross merge the two halves to get the final result.

**Steps:**
1. Use slow and fast pointers to find the middle - slow moves one step, fast moves two steps
2. Split the list at the middle into first half and second half
3. Reverse the second half of the list
4. Cross merge the first half with the reversed second half

**Example:**
- Input `[1,2,3,4]` -> find middle at `2` -> reverse second half to `[4,3]` -> cross merge -> `[1,4,2,3]`

```embed-go
PATH: "vault://leetcode/0101-0200/0143_reorder_list/solution.go"
TITLE: "leetcode 143. Reorder List"
```

### Approach 2
**Method Name**: Split List + Reverse + Alternate Connection

**Principle:**
Split the linked list into two halves, reverse the second half, then alternately connect nodes from both lists.

**Steps:**
1. Traverse to find the middle node and split the list into two halves
2. Reverse the second half of the list
3. Use three pointers to alternately connect nodes from the first half and reversed second half

**Example:**
- Input `[1,2,3,4,5]` -> split into `[1,2]` and `[3,4,5]` -> reverse second half to `[5,4,3]` -> alternate connection -> `[1,5,2,4,3]`

```embed-go
PATH: "vault://leetcode/0101-0200/0143_reorder_list/solution2.go"
TITLE: "leetcode 143. Reorder List"
```

### Approach 3
**Method Name**: Array Storage + Two Pointers Rearrangement

**Principle:**
Store linked list nodes in an array, leverage random access capability of arrays, use two pointers traversing from both ends towards the center to reconnect nodes.

**Steps:**
1. Traverse the linked list and store all nodes in an array
2. Use left pointer starting from head, right pointer starting from tail
3. Connect left pointer node to right pointer node, right pointer node to left's next node
4. Move pointers until they meet, finally handle the tail node's Next pointer

**Example:**
- Input `[1,2,3,4]` -> store in array `[1,2,3,4]` -> two pointers `l=0,r=3` -> `1→4→2→3` -> done

```embed-go
PATH: "vault://leetcode/0101-0200/0143_reorder_list/solution3.go"
TITLE: "leetcode 143. Reorder List"
```
