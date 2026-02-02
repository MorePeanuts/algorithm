---
link: https://leetcode.com/problems/top-k-frequent-elements/
tags:
  - Medium
  - Array
  - Hash_Table
  - Divide_and_Conquer
  - Bucket_Sort
  - Counting
  - Quickselect
  - Sorting
  - Heap_(Priority_Queue)
---
## Description
Given an integer array `nums` and an integer `k`, return *the* `k` *most frequent elements*. You may return the answer in **any order**.

**Example 1:**

**Input:** nums = [1,1,1,2,2,3], k = 2

**Output:** [1,2]

**Example 2:**

**Input:** nums = [1], k = 1

**Output:** [1]

**Example 3:**

**Input:** nums = [1,2,1,2,1,2,3,1,3,2], k = 2

**Output:** [1,2]

**Constraints:**

- `1 <= nums.length <= 105`
- `-104 <= nums[i] <= 104`
- `k` is in the range `[1, the number of unique elements in the array]`.
- It is **guaranteed** that the answer is **unique**.

**Follow up:** Your algorithm's time complexity must be better than `O(n log n)`, where n is the array's size.

## Solution

### Approach 1

**Hash Table + Max Heap**: First count element frequencies, then use a heap to extract the k elements with highest frequencies.

**Principle:**
Use a hash table to count each element's occurrences in O(n), then build a max heap sorted by frequency. Pop the heap top k times to get the k most frequent elements.

**Steps:**
1. Traverse the array, use hash table `freq` to count each element's occurrence frequency
2. Convert `(element, frequency)` pairs to a slice
3. Build a max heap based on this slice (sorted by frequency in descending order)
4. Pop k elements from the heap, collect their values as the result

**Example:**
- Input: `nums = [1,1,1,2,2,3], k = 2`
- Frequency count: `{1: 3, 2: 2, 3: 1}`
- Heap pop order: `(1,3) → (2,2)`
- Output: `[1, 2]`

```embed-go
PATH: "vault://leetcode/0301-0400/0347_top_k_frequent_elements/solution.go"
TITLE: "leetcode 347. Top K Frequent Elements"
```
