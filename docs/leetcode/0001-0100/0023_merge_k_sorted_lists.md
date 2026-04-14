---
link: https://leetcode.com/problems/merge-k-sorted-lists/
tags:
  - Hard
  - Linked_List
  - Divide_and_Conquer
  - Heap_(Priority_Queue)
  - Merge_Sort
---
## Description
You are given an array of `k` linked-lists `lists`, each linked-list is sorted in ascending order.

*Merge all the linked-lists into one sorted linked-list and return it.*

**Example 1:**

```
Input: lists = [[1,4,5],[1,3,4],[2,6]]
Output: [1,1,2,3,4,4,5,6]
Explanation: The linked-lists are:
[
  1->4->5,
  1->3->4,
  2->6
]
merging them into one sorted linked list:
1->1->2->3->4->4->5->6
```

**Example 2:**

```
Input: lists = []
Output: []
```

**Example 3:**

```
Input: lists = [[]]
Output: []
```

**Constraints:**

- `k == lists.length`
- `0 <= k <= 104`
- `0 <= lists[i].length <= 500`
- `-104 <= lists[i][j] <= 104`
- `lists[i]` is sorted in **ascending order**.
- The sum of `lists[i].length` will not exceed `104`.

## Solution
### Approach 1
**Method Name**: Divide and Conquer with Pairwise Merging

**Principle:**
Using divide and conquer strategy, decompose the problem of merging k sorted lists into multiple subproblems of merging two lists. By merging level by level, we finally obtain a single sorted linked list. This approach borrows the idea from merge sort and has better time complexity.

**Steps:**
1. Check if the input list array is empty, return nil if true
2. While the list array length is greater than 1, perform merging iteratively:
   - Traverse the array, merge every two adjacent lists into one
   - Place the merged lists at the front of the array
   - Truncate the array to keep only the merged portion
3. The remaining list in the array is the final merged result

**Example:**
- Input: `[[1,4,5],[1,3,4],[2,6]]`
- First round merge: merge `[1,4,5]` and `[1,3,4]` to get `[1,1,3,4,4,5]`, `[2,6]` remains as is
- Second round merge: merge `[1,1,3,4,4,5]` and `[2,6]` to get `[1,1,2,3,4,4,5,6]`
- Output: `[1,1,2,3,4,4,5,6]`

**Complexity Analysis:**
- Time Complexity: O(N log k), where N is the total number of nodes in all lists, k is the number of linked lists. Each round takes O(N), and we need O(log k) rounds
- Space Complexity: O(1), only constant extra space is used

> **Note**: During the merging process, when the number of lists is odd, the last list doesn't need pairing and directly proceeds to the next round.

```embed-go
PATH: "vault://leetcode/0001-0100/0023_merge_k_sorted_lists/solution.go"
TITLE: "leetcode 23. Merge k Sorted Lists"
```
