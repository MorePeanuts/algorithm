---
link: https://leetcode.com/problems/reverse-linked-list/
tags:
  - Easy
  - Recursion
  - Linked_List
---
## Description
Given the `head` of a singly linked list, reverse the list, and return *the reversed list*.

**Example 1:**

![](https://assets.leetcode.com/uploads/2021/02/19/rev1ex1.jpg)

```
Input: head = [1,2,3,4,5]
Output: [5,4,3,2,1]
```

**Example 2:**

![](https://assets.leetcode.com/uploads/2021/02/19/rev1ex2.jpg)

```
Input: head = [1,2]
Output: [2,1]
```

**Example 3:**

```
Input: head = []
Output: []
```

**Constraints:**

- The number of nodes in the list is the range `[0, 5000]`.
- `-5000 <= Node.val <= 5000`

**Follow up:** A linked list can be reversed either iteratively or recursively. Could you implement both?

## Solution
### Approach 1
**Method Name**: Iterative method (Three-pointer traversal)

**Principle:**
Use three pointers to traverse the linked list, changing the direction of each node's pointer one by one, pointing the current node's Next to the previous node, thereby achieving linked list reversal.

**Steps:**
1. Initialize three pointers: p1 (previous node) as nil, p2 (current node) as head
2. While p2 is not nil, continue traversing:
   - Save p2's next node to p3
   - Point p2's Next to p1 (reverse pointer direction)
   - Move p1 to p2's position
   - Move p2 to p3's position
3. When p2 is nil, p1 is the new head node, return p1

```embed-go
PATH: "vault://leetcode/0201-0300/0206_reverse_linked_list/solution.go"
TITLE: "leetcode 206. Reverse Linked List"
```
