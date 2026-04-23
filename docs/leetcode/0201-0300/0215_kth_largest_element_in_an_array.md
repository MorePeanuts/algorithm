---
link: https://leetcode.com/problems/kth-largest-element-in-an-array/
tags:
  - Medium
  - Array
  - Divide_and_Conquer
  - Quickselect
  - Sorting
  - Heap_(Priority_Queue)
---
## Description
Given an integer array `nums` and an integer `k`, return *the* `kth` *largest element in the array*.

Note that it is the `kth` largest element in the sorted order, not the `kth` distinct element.

Can you solve it without sorting?

**Example 1:**

```
Input: nums = [3,2,1,5,6,4], k = 2
Output: 5
```

**Example 2:**

```
Input: nums = [3,2,3,1,2,4,5,5,6], k = 4
Output: 4
```

**Constraints:**

- `1 <= k <= nums.length <= 105`
- `-104 <= nums[i] <= 104`

## Solution
### Approach 1
**Method Name**: Max Heap

**Principle:**
Using the property of a max heap where the top element is always the maximum value. By building a max heap and popping k times, we get the kth largest element.

**Steps:**
1. Build a max heap from the entire array
2. Perform k pop operations
3. The kth popped element is the kth largest element

**Example:**
- Input: [3,2,1,5,6,4], k = 2 -> Build max heap [6,5,4,2,3,1] -> First pop 6, second pop 5 -> Output: 5

```embed-go
PATH: "vault://leetcode/0201-0300/0215_kth_largest_element_in_an_array/solution.go"
TITLE: "leetcode 215. Kth Largest Element in an Array"
```

### Approach 2
**Method Name**: Min Heap (size k)

**Principle:**
Maintain a min heap of size k where the top element is the kth largest among the elements processed so far. Iterate through remaining elements, replacing the top if a larger element is found. The final top is the answer.

**Steps:**
1. Build a min heap with the first k elements
2. Iterate through remaining elements in the array
3. If current element is larger than heap top, pop the top and insert current element
4. After iteration, the heap top is the kth largest element

**Example:**
- Input: [3,2,1,5,6,4], k = 2 -> Initial heap [2,3] -> Process 5 (replace 2) -> Heap [3,5] -> Process 6 (replace 3) -> Heap [5,6] -> Process 4 (skip) -> Heap top 5 -> Output: 5

```embed-go
PATH: "vault://leetcode/0201-0300/0215_kth_largest_element_in_an_array/solution2.go"
TITLE: "leetcode 215. Kth Largest Element in an Array"
```
