---
link: https://leetcode.com/problems/median-of-two-sorted-arrays/
tags:
  - Hard
  - Array
  - Binary_Search
  - Divide_and_Conquer
---
## Description
Given two sorted arrays `nums1` and `nums2` of size `m` and `n` respectively, return **the median** of the two sorted arrays.

The overall run time complexity should be `O(log (m+n))`.

**Example 1:**

```
Input: nums1 = [1,3], nums2 = [2]
Output: 2.00000
Explanation: merged array = [1,2,3] and median is 2.
```

**Example 2:**

```
Input: nums1 = [1,2], nums2 = [3,4]
Output: 2.50000
Explanation: merged array = [1,2,3,4] and median is (2 + 3) / 2 = 2.5.
```

**Constraints:**

- `nums1.length == m`
- `nums2.length == n`
- `0 <= m <= 1000`
- `0 <= n <= 1000`
- `1 <= m + n <= 2000`
- `-106 <= nums1[i], nums2[i] <= 106`

## Solution
### Approach 1
**Method Name**: Binary Search on Partition Point

**Principle:**
By performing binary search on the shorter array, find a partition point such that all elements on the left are less than or equal to all elements on the right, and the total number of elements on both sides are equal (or differ by 1). The median can then be calculated from the elements near the partition points.

**Steps:**
1. Ensure nums1 is the shorter array, swap if not. This reduces the number of binary search iterations.
2. Initialize binary search boundaries l=-1, r=len(nums1)
3. Perform binary search on nums1:
   - Calculate partition point p1 = (l + r) / 2 for nums1
   - Calculate partition point p2 = (len1+len2+1)/2 - (p1+1) - 1 for nums2 based on equal element count principle
   - If nums1[p1] > nums2[p2+1], p1 is too far right, search left by adjusting r = p1
   - Otherwise, p1 is valid or can be searched right by adjusting l = p1
4. After finding valid partition points, calculate leftMax (maximum of left elements) and rightMin (minimum of right elements)
5. Return median based on total length parity:
   - Odd length: return leftMax
   - Even length: return (leftMax + rightMin) / 2

**Example:**
Take nums1 = [1,3], nums2 = [2]:
- Total length 3 (odd), need 2 elements on left, 1 element on right
- Binary search on nums1, final p1=0 (element 1), p2=0 (element 2)
- leftMax = max(1, 2) = 2, rightMin = min(3, +∞) = 3
- Total length is odd, return leftMax = 2.0

Take nums1 = [1,2], nums2 = [3,4]:
- Total length 4 (even), need 2 elements on left, 2 elements on right
- Final partition p1=1, p2=-1: leftMax = max(2, -∞) = 2, rightMin = min(+∞, 3) = 3
- Return (2+3)/2 = 2.5

> **Note**: Need to handle edge cases such as one array being empty, partition points at array ends, etc. The code uses math.MinInt and math.MaxInt to handle these boundary conditions.

```embed-go
PATH: "vault://leetcode/0001-0100/0004_median_of_two_sorted_arrays/solution.go"
TITLE: "leetcode 4. Median of Two Sorted Arrays"
```
