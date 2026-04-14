---
link: https://leetcode.com/problems/reverse-nodes-in-k-group/
tags:
  - Hard
  - Recursion
  - Linked_List
---
## Description
Given the `head` of a linked list, reverse the nodes of the list `k` at a time, and return *the modified list*.

`k` is a positive integer and is less than or equal to the length of the linked list. If the number of nodes is not a multiple of `k` then left-out nodes, in the end, should remain as it is.

You may not alter the values in the list's nodes, only nodes themselves may be changed.

**Example 1:**

![](https://assets.leetcode.com/uploads/2020/10/03/reverse_ex1.jpg)

```
Input: head = [1,2,3,4,5], k = 2
Output: [2,1,4,3,5]
```

**Example 2:**

![](https://assets.leetcode.com/uploads/2020/10/03/reverse_ex2.jpg)

```
Input: head = [1,2,3,4,5], k = 3
Output: [3,2,1,4,5]
```

**Constraints:**

- The number of nodes in the list is `n`.
- `1 <= k <= n <= 5000`
- `0 <= Node.val <= 1000`

**Follow-up:** Can you solve the problem in `O(1)` extra memory space?

## Solution
### Approach 1
**Method Name**: Iterative reversal with O(1) space

**Principle:**
Traverse the linked list and reverse each group of k nodes iteratively. Track the connections between reversed groups while using constant extra space.

**Steps:**
1. For each group, first check if there are at least k nodes remaining
2. Reverse the k nodes if sufficient nodes exist
3. Connect the reversed group with the previous group
4. Continue with the next group until the end of the list is reached

**Example:**
- Input: head = [1,2,3,4,5], k = 2
- Reverse first group [1,2] → [2,1]
- Reverse second group [3,4] → [4,3]
- Leave remaining [5] unchanged
- Output: [2,1,4,3,5]

> **Note**: When fewer than k nodes remain, they are left in their original order as required by the problem constraints.

```embed-go
PATH: "vault://leetcode/0001-0100/0025_reverse_nodes_in_k_group/solution.go"
TITLE: "leetcode 25. Reverse Nodes in k-Group"
```
