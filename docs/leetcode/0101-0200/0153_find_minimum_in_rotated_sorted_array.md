---
link: https://leetcode.com/problems/find-minimum-in-rotated-sorted-array/
tags:
  - Medium
  - Array
  - Binary_Search
---
## Description
Suppose an array of length `n` sorted in ascending order is **rotated** between `1` and `n` times. For example, the array `nums = [0,1,2,4,5,6,7]` might become:

- `[4,5,6,7,0,1,2]` if it was rotated `4` times.
- `[0,1,2,4,5,6,7]` if it was rotated `7` times.

Notice that **rotating** an array `[a[0], a[1], a[2], ..., a[n-1]]` 1 time results in the array `[a[n-1], a[0], a[1], a[2], ..., a[n-2]]`.

Given the sorted rotated array `nums` of **unique** elements, return *the minimum element of this array*.

You must write an algorithm that runs in `O(log n) time`.

**Example 1:**

```
Input: nums = [3,4,5,1,2]
Output: 1
Explanation: The original array was [1,2,3,4,5] rotated 3 times.
```

**Example 2:**

```
Input: nums = [4,5,6,7,0,1,2]
Output: 0
Explanation: The original array was [0,1,2,4,5,6,7] and it was rotated 4 times.
```

**Example 3:**

```
Input: nums = [11,13,15,17]
Output: 11
Explanation: The original array was [11,13,15,17] and it was rotated 4 times.
```

**Constraints:**

- `n == nums.length`
- `1 <= n <= 5000`
- `-5000 <= nums[i] <= 5000`
- All the integers of `nums` are **unique**.
- `nums` is sorted and rotated between `1` and `n` times.

## Solution
### Approach 1
**Method Name**: Binary Search

**Principle:**
Leverage the properties of rotated sorted arrays to quickly locate the minimum value using binary search. The rotated array can be divided into two ascending subarrays, with the minimum value being the boundary point between them. By comparing the middle element with the right boundary element, we can determine the interval where the minimum value resides.

**Steps:**
1. First check if the array is not rotated (first element < last element), if so return the first element directly
2. Initialize left and right pointers `l=0`, `r=len(nums)`
3. While `l < r`, compute middle position `m = l + (r-l)/2`
4. If `nums[m] > rightNum` (right boundary element), the minimum is in the right half, adjust `l = m + 1`
5. Otherwise the minimum is in the left half (including m), adjust `r = m`
6. When `l == r`, `nums[l]` is the minimum value

**Example:**
- Input `nums = [4,5,6,7,0,1,2]`, rightNum = 2
- nums[0] = 4 > 2, enter binary search
- m=3, nums[3]=7 > 2 → l=4
- m=5, nums[5]=1 ≤ 2 → r=5
- m=4, nums[4]=0 ≤ 2 → r=4
- l==r=4, return nums[4]=0

```embed-go
PATH: "vault://leetcode/0101-0200/0153_find_minimum_in_rotated_sorted_array/solution.go"
TITLE: "leetcode 153. Find Minimum in Rotated Sorted Array"
```
