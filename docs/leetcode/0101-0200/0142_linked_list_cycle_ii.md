---
link: https://leetcode.com/problems/linked-list-cycle-ii/
tags:
  - Medium
  - Hash_Table
  - Linked_List
  - Two_Pointers
---
## Description
Given the `head` of a linked list, return *the node where the cycle begins. If there is no cycle, return* `null`.

There is a cycle in a linked list if there is some node in the list that can be reached again by continuously following the `next` pointer. Internally, `pos` is used to denote the index of the node that tail's `next` pointer is connected to (**0-indexed**). It is `-1` if there is no cycle. **Note that** `pos` **is not passed as a parameter**.

**Do not modify** the linked list.

**Example 1:**

![](https://assets.leetcode.com/uploads/2018/12/07/circularlinkedlist.png)

```
Input: head = [3,2,0,-4], pos = 1
Output: tail connects to node index 1
Explanation: There is a cycle in the linked list, where tail connects to the second node.
```

**Example 2:**

![](https://assets.leetcode.com/uploads/2018/12/07/circularlinkedlist_test2.png)

```
Input: head = [1,2], pos = 0
Output: tail connects to node index 0
Explanation: There is a cycle in the linked list, where tail connects to the first node.
```

**Example 3:**

![](https://assets.leetcode.com/uploads/2018/12/07/circularlinkedlist_test3.png)

```
Input: head = [1], pos = -1
Output: no cycle
Explanation: There is no cycle in the linked list.
```

**Constraints:**

- The number of the nodes in the list is in the range `[0, 104]`.
- `-105 <= Node.val <= 105`
- `pos` is `-1` or a **valid index** in the linked-list.

**Follow up:** Can you solve it using `O(1)` (i.e. constant) memory?

## Solution
### Approach 1
**Method Name**: Calculate Cycle Length + Two Pointers

**Principle:**
First use fast and slow pointers to detect if the linked list has a cycle. If a cycle exists, calculate its length. Then use two pointers: one pointer starts from the head and moves forward by the cycle length steps, while another pointer starts from the head. Both pointers move at the same speed, and their meeting point is the cycle entrance.

**Steps:**
1. Use fast and slow pointers (slow moves 1 step, fast moves 2 steps each iteration) to traverse the linked list. If the two pointers meet, a cycle exists.
2. From the meeting point, continue moving one pointer until it returns to the meeting point, counting the cycle length.
3. Let the fast pointer start from the head and move forward by the cycle length steps.
4. Let the slow pointer start from the head, and both pointers move 1 step at a time simultaneously.
5. When the two pointers meet, that node is the cycle entrance.

**Example:**
- Input: `[3,2,0,-4]`, cycle starts at index 1
- After fast and slow pointers meet, the cycle length is calculated as 3
- fast starts from head and moves 3 steps to reach node -4
- slow starts from head, both move simultaneously:
  - slow: 3 → 2 → 0
  - fast: -4 → 2 → 0
  - Meet at node 2, which is the cycle entrance

```embed-go
PATH: "vault://leetcode/0101-0200/0142_linked_list_cycle_ii/solution.go"
TITLE: "leetcode 142. Linked List Cycle II"
```

### Approach 2
**Method Name**: Two Pointers - Two Meetings

**Principle:**
Use fast and slow pointers to meet twice to find the cycle entrance. After the first meeting, reset the fast pointer to the head of the linked list, then both pointers move forward at the same speed. The second meeting point is the cycle entrance. The core mathematical principle is: the distance from the head to the cycle entrance equals the distance from the first meeting point to the cycle entrance (possibly plus several cycle lengths).

**Mathematical Analysis of First Meeting:**
Let the linked list have \(a + b\) nodes, where:
- \(a\): number of nodes from head to cycle entrance (excluding the entrance node)
- \(b\): number of nodes in the cycle

Let the two pointers have taken \(f\) (fast) and \(s\) (slow) steps, then:
1. \(f = 2s\) (fast speed is twice slow speed)
2. \(f = s + nb\) (fast has taken \(n\) more cycle lengths than slow)

Subtracting the two equations: \(f = 2nb\), \(s = nb\)

**Principle of Second Meeting:**
If a pointer moves from the head continuously, the number of steps to reach the cycle entrance is \(k = a + nb\) (first \(a\) steps to reach entrance, then every \(b\) steps returns to entrance).

Currently, slow has already taken \(nb\) steps, so letting slow take \(a\) more steps will reach the entrance.

Although we don't know the exact value of \(a\), we can use the fact that it takes exactly \(a\) steps from the head to the entrance. Therefore, reset fast to the head, let both pointers move forward at the same speed. When fast has taken \(a\) steps to reach the entrance, slow will have also taken \(a + nb\) steps to reach the entrance, and the two pointers meet for the second time at the entrance.

**Steps:**
1. Use fast and slow pointers (slow moves 1 step, fast moves 2 steps each iteration) to traverse the linked list.
2. If fast reaches the end of the list, there is no cycle, return null.
3. If slow and fast meet for the first time, a cycle exists, proceed to next step.
4. Reset fast to the head of the linked list.
5. Both slow and fast move forward 1 step at a time simultaneously.
6. When the two pointers meet for the second time, that node is the cycle entrance.

**Example:**
- Input: `[3,2,0,-4]`, cycle starts at index 1 (\(a = 1\), \(b = 3\))
- First meeting: slow and fast meet at node -4 (slow took 3 steps = \(nb = 1 \times 3\), fast took 6 steps = \(2nb = 2 \times 3\))
- fast reset to head node 3
- Second meeting:
  - slow: -4 → 2
  - fast: 3 → 2
  - Meet at node 2, which is the cycle entrance

```embed-go
PATH: "vault://leetcode/0101-0200/0142_linked_list_cycle_ii/solution2.go"
TITLE: "leetcode 142. Linked List Cycle II"
```
