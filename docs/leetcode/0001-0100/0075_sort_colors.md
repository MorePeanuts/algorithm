---
link: https://leetcode.com/problems/sort-colors/
tags:
  - Medium
  - Array
  - Two_Pointers
  - Sorting
---
## Description
Given an array `nums` with `n` objects colored red, white, or blue, sort them **[in-place](https://en.wikipedia.org/wiki/In-place_algorithm)** so that objects of the same color are adjacent, with the colors in the order red, white, and blue.

We will use the integers `0`, `1`, and `2` to represent the color red, white, and blue, respectively.

You must solve this problem without using the library's sort function.

**Example 1:**

```
Input: nums = [2,0,2,1,1,0]
Output: [0,0,1,1,2,2]
```

**Example 2:**

```
Input: nums = [2,0,1]
Output: [0,1,2]
```

**Constraints:**

- `n == nums.length`
- `1 <= n <= 300`
- `nums[i]` is either `0`, `1`, or `2`.

**Follow up:** Could you come up with a one-pass algorithm using only constant extra space?

## Solution
### Approach 1
**Method Name**: Three Pointers One-Pass

**Principle:**
Use three pointers to divide the array into three regions: 0's region [0, left), 1's region [left, i), unprocessed region [i, right], and 2's region (right, end). Swap elements to their corresponding regions in one pass.

**Steps:**
1. Initialize left=0 (right boundary of 0s), right=len(nums)-1 (left boundary of 2s), pivot=1 (middle value)
2. Traverse the array with i starting from left:
   - If nums[i] < pivot (i.e., 0): swap with left position, increment both left and i
   - If nums[i] == pivot (i.e., 1): just increment i
   - If nums[i] > pivot (i.e., 2): swap with right position, decrement right (don't move i because the swapped element hasn't been checked)
3. End when i > right

**Example:**
- Input: [2,0,2,1,1,0]
- Processing:
  - i=0, nums[i]=2 → swap with right=5 → [0,0,2,1,1,2], right=4
  - i=0, nums[i]=0 → swap with left=0 → [0,0,2,1,1,2], left=1, i=1
  - i=1, nums[i]=0 → swap with left=1 → [0,0,2,1,1,2], left=2, i=2
  - i=2, nums[i]=2 → swap with right=4 → [0,0,1,1,2,2], right=3
  - i=2, nums[i]=1 → i=3
  - i=3, nums[i]=1 → i=4 > right=3, done
- Output: [0,0,1,1,2,2]

```embed-go
PATH: "vault://leetcode/0001-0100/0075_sort_colors/solution.go"
TITLE: "leetcode 75. Sort Colors"
```
