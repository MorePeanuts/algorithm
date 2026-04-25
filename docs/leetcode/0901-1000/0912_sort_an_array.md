---
link: https://leetcode.com/problems/sort-an-array/
tags:
  - Medium
  - Array
  - Divide_and_Conquer
  - Bucket_Sort
  - Counting_Sort
  - Radix_Sort
  - Sorting
  - Heap_(Priority_Queue)
  - Merge_Sort
---
## Description
Given an array of integers `nums`, sort the array in ascending order and return it.

You must solve the problem **without using any built-in** functions in `O(nlog(n))` time complexity and with the smallest space complexity possible.

**Example 1:**

```
Input: nums = [5,2,3,1]
Output: [1,2,3,5]
Explanation: After sorting the array, the positions of some numbers are not changed (for example, 2 and 3), while the positions of other numbers are changed (for example, 1 and 5).
```

**Example 2:**

```
Input: nums = [5,1,1,2,0,0]
Output: [0,0,1,1,2,5]
Explanation: Note that the values of nums are not necessarily unique.
```

**Constraints:**

- `1 <= nums.length <= 5 * 104`
- `-5 * 104 <= nums[i] <= 5 * 104`

## Solution
### Approach 1
**Method Name**: QuickSort (2-way partition + insertion sort optimization)

**Principle:**
QuickSort selects a pivot to partition the array, sorts recursively, and uses insertion sort to optimize small arrays.

**Steps:**
1. If array length ≤ 10, use insertion sort
2. Otherwise, select pivot using median-of-three and place it properly
3. 2-way partition: move elements smaller than pivot to left, larger to right
4. Recursively sort left and right partitions

**Example:**
- Input [5,2,3,1] → pivot 3 → partition [2,1,3,5] → recursive sort → [1,2,3,5]

```embed-go
PATH: "vault://leetcode/0901-1000/0912_sort_an_array/solution.go"
TITLE: "leetcode 912. Sort an Array"
```

### Approach 2
**Method Name**: QuickSort (3-way partition + insertion sort optimization)

**Principle:**
3-way partition divides array into less than, equal to, and greater than pivot, efficiently handling duplicates.

**Steps:**
1. If array length ≤ 10, use insertion sort
2. Select pivot using median-of-three
3. 3-way partition: traverse array and distribute elements to three regions
4. Recursively sort less than and greater than partitions

**Example:**
- Input [5,1,1,2,0,0] → pivot 1 → partition [0,0,1,1,2,5] → complete sorting recursively

```embed-go
PATH: "vault://leetcode/0901-1000/0912_sort_an_array/solution2.go"
TITLE: "leetcode 912. Sort an Array"
```

### Approach 3
**Method Name**: MergeSort

**Principle:**
MergeSort uses divide-and-conquer, recursively splitting then merging, ensuring stable O(n log n) time complexity.

**Steps:**
1. If array length ≤ 1, return directly
2. Otherwise, split array into two parts from middle
3. Recursively sort left and right parts
4. Merge two sorted arrays into one sorted array

**Example:**
- Input [5,2,3,1] → split into [5,2] and [3,1] → sorted [2,5] and [1,3] → merge to [1,2,3,5]

```embed-go
PATH: "vault://leetcode/0901-1000/0912_sort_an_array/solution3.go"
TITLE: "leetcode 912. Sort an Array"
```
