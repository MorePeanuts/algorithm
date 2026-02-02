---
link: https://leetcode.com/problems/add-two-numbers/
tags:
  - Medium
  - Recursion
  - Linked_List
  - Math
---
## Description
You are given two **non-empty** linked lists representing two non-negative integers. The digits are stored in **reverse order**, and each of their nodes contains a single digit. Add the two numbers and return the sum as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.

**Example 1:**

![](https://assets.leetcode.com/uploads/2020/10/02/addtwonumber1.jpg)

```
Input: l1 = [2,4,3], l2 = [5,6,4]
Output: [7,0,8]
Explanation: 342 + 465 = 807.
```

**Example 2:**

```
Input: l1 = [0], l2 = [0]
Output: [0]
```

**Example 3:**

```
Input: l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
Output: [8,9,9,9,0,0,0,1]
```

**Constraints:**

- The number of nodes in each linked list is in the range `[1, 100]`.
- `0 <= Node.val <= 9`
- It is guaranteed that the list represents a number that does not have leading zeros.

## Solution

### Approach 1

**Simulate Addition**: Add digit by digit while handling carry.

**Principle:**
Simulate the vertical addition process, starting from the least significant digit. Use a variable to track the carry until both linked lists are fully traversed and there's no remaining carry.

**Steps:**
1. Initialize a result linked list and a carry variable `carry = 0`
2. Traverse both linked lists simultaneously, computing the sum: `sum = carry + l1.Val + l2.Val`
3. Create a new node storing `sum % 10`, update carry as `carry = sum / 10`
4. Link the new node to the result list
5. After traversal, if `carry > 0`, append a node with value `carry`

**Example:**
- Input: `l1 = [2,4,3]` (represents 342), `l2 = [5,6,4]` (represents 465)
- Process:
  - Digit 1: `2 + 5 = 7`, carry 0, result node value 7
  - Digit 2: `4 + 6 = 10`, carry 1, result node value 0
  - Digit 3: `3 + 4 + 1 = 8`, carry 0, result node value 8
- Output: `[7,0,8]` (represents 807)

```embed-go
PATH: "vault://leetcode/0001-0100/0002_add_two_numbers/solution.go"
TITLE: "leetcode 2. Add Two Numbers"
```
