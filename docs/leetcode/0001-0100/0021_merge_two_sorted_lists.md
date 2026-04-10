---
link: https://leetcode.com/problems/merge-two-sorted-lists/
tags:
  - Easy
  - Recursion
  - Linked_List
---
## Description
You are given the heads of two sorted linked lists `list1` and `list2`.

Merge the two lists into one **sorted** list. The list should be made by splicing together the nodes of the first two lists.

Return *the head of the merged linked list*.

**Example 1:**

![](https://assets.leetcode.com/uploads/2020/10/03/merge_ex1.jpg)

```
Input: list1 = [1,2,4], list2 = [1,3,4]
Output: [1,1,2,3,4,4]
```

**Example 2:**

```
Input: list1 = [], list2 = []
Output: []
```

**Example 3:**

```
Input: list1 = [], list2 = [0]
Output: [0]
```

**Constraints:**

- The number of nodes in both lists is in the range `[0, 50]`.
- `-100 <= Node.val <= 100`
- Both `list1` and `list2` are sorted in **non-decreasing** order.

## Solution
### Approach 1
**Method Name**: Iterative Merge

**Principle:**
Use two pointers to traverse both lists, selecting the node with the smaller value at each step to add to the result list until both lists are fully traversed.

**Steps:**
1. Initialize result list head pointer `res` and current pointer `p`
2. Use `p1` and `p2` to traverse the two lists respectively
3. Compare current node values, select the smaller node as "winner"
4. Add the "winner" node to the result list
5. Move the corresponding pointer to continue traversal
6. Return the merged list head pointer

```embed-go
PATH: "vault://leetcode/0001-0100/0021_merge_two_sorted_lists/solution.go"
TITLE: "leetcode 21. Merge Two Sorted Lists"
```
