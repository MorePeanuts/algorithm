---
link: https://leetcode.com/problems/minimum-size-subarray-sum/
tags:
  - Medium
  - Array
  - Binary_Search
  - Prefix_Sum
  - Sliding_Window
---
## Description
Given an array of positive integers `nums` and a positive integer `target`, return *the **minimal length** of a* *subarray* *whose sum is greater than or equal to* `target`. If there is no such subarray, return `0` instead.

**Example 1:**

```
Input: target = 7, nums = [2,3,1,2,4,3]
Output: 2
Explanation: The subarray [4,3] has the minimal length under the problem constraint.
```

**Example 2:**

```
Input: target = 4, nums = [1,4,4]
Output: 1
```

**Example 3:**

```
Input: target = 11, nums = [1,1,1,1,1,1,1,1]
Output: 0
```

**Constraints:**

- `1 <= target <= 109`
- `1 <= nums.length <= 105`
- `1 <= nums[i] <= 104`

**Follow up:** If you have figured out the `O(n)` solution, try coding another solution of which the time complexity is `O(n log(n))`.

## Solution
### Approach 1
**Method Name**: Sliding Window (Two Pointers)

**Principle:**
Use two pointers (left and right) to maintain a window. Expand the window by moving the right pointer until the sum reaches the target, then shrink the window by moving the left pointer to find the minimum length.

**Steps:**
1. Initialize left and right pointers at the start of the array, current window sum as the first element, and minimum length as array length + 1 (indicating no valid subarray initially)
2. Loop while right pointer is within bounds and left pointer doesn't exceed right pointer:
   - If current window sum is less than target, move right pointer right and add the new element to window sum
   - If current window sum is greater than or equal to target, update minimum length, then move left pointer right and subtract the removed element from window sum
3. Finally, check if minimum length is still the initial value to determine if a valid subarray exists

**Example:**
- Input: target = 7, nums = [2,3,1,2,4,3]
- Processing:
  - Window [2] → sum=2 < 7 → expand
  - Window [2,3] → sum=5 < 7 → expand
  - Window [2,3,1] → sum=6 < 7 → expand
  - Window [2,3,1,2] → sum=8 ≥ 7 → record length 4, shrink → [3,1,2]
  - Window [3,1,2] → sum=6 < 7 → expand
  - Window [3,1,2,4] → sum=10 ≥ 7 → record length 4, shrink → [1,2,4]
  - Window [1,2,4] → sum=7 ≥ 7 → record length 3, shrink → [2,4]
  - Window [2,4] → sum=6 < 7 → expand
  - Window [2,4,3] → sum=9 ≥ 7 → record length 3, shrink → [4,3]
  - Window [4,3] → sum=7 ≥ 7 → record length 2, shrink → [3]
- Output: 2

```embed-go
PATH: "vault://leetcode/0201-0300/0209_minimum_size_subarray_sum/solution.go"
TITLE: "leetcode 209. Minimum Size Subarray Sum"
```

### Approach 2
**Method Name**: Sliding Window (Optimized)

**Principle:**
Outer loop uses right pointer to traverse the array, inner loop continuously tries to shrink the left pointer, completing window expansion and shrinkage in a single pass.

**Steps:**
1. Initialize left pointer at the start of the array, current window sum as 0, and minimum length as array length + 1
2. Right pointer traverses the array, adding each element to window sum
3. When window sum is greater than or equal to target, loop:
   - Update minimum length
   - Subtract the element at left pointer from window sum
   - Move left pointer right
4. Finally, check if minimum length is still the initial value to determine if a valid subarray exists

**Example:**
- Input: target = 7, nums = [2,3,1,2,4,3]
- Processing:
  - Right at 0, sum=2 → no shrink
  - Right at 1, sum=5 → no shrink
  - Right at 2, sum=6 → no shrink
  - Right at 3, sum=8 ≥ 7 → record length 4, shrink to sum=6 (left=1)
  - Right at 4, sum=10 ≥ 7 → record length 4, shrink to sum=7 (left=2), record length 3, shrink to sum=6 (left=3)
  - Right at 5, sum=9 ≥ 7 → record length 3, shrink to sum=7 (left=4), record length 2, shrink to sum=3 (left=5)
- Output: 2

```embed-go
PATH: "vault://leetcode/0201-0300/0209_minimum_size_subarray_sum/solution2.go"
TITLE: "leetcode 209. Minimum Size Subarray Sum"
```
