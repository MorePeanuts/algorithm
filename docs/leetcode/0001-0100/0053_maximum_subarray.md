---
link: https://leetcode.com/problems/maximum-subarray/
tags:
  - Medium
  - Array
  - Divide_and_Conquer
  - Dynamic_Programming
---
## Description
Given an integer array `nums`, find the subarray with the largest sum, and return *its sum*.

**Example 1:**

```
Input: nums = [-2,1,-3,4,-1,2,1,-5,4]
Output: 6
Explanation: The subarray [4,-1,2,1] has the largest sum 6.
```

**Example 2:**

```
Input: nums = [1]
Output: 1
Explanation: The subarray [1] has the largest sum 1.
```

**Example 3:**

```
Input: nums = [5,4,-1,7,8]
Output: 23
Explanation: The subarray [5,4,-1,7,8] has the largest sum 23.
```

**Constraints:**

- `1 <= nums.length <= 105`
- `-104 <= nums[i] <= 104`

**Follow up:** If you have figured out the `O(n)` solution, try coding another solution using the **divide and conquer** approach, which is more subtle.

## Solution

### Approach 1

**Dynamic Programming (Kadane's Algorithm)**: Maintain the current subarray sum, restart calculation when encountering a negative prefix.

**Principle:**
If the cumulative sum at the previous position is negative, it provides no benefit to the current position and should be discarded, restarting accumulation from the current position.

**Steps:**
1. Initialize `thisSum = 0`, `maxSum = nums[0]`
2. Iterate through the array, accumulate current element to `thisSum`
3. If the previous cumulative sum `lastSum < 0`, subtract it from `thisSum` (equivalent to restarting)
4. Update `maxSum` to be the larger of `thisSum` and `maxSum`

**Example:**
- Input `[-2,1,-3,4,-1,2,1,-5,4]`
- Traversal process: `-2 → 1 (reset) → -2 → 4 (reset) → 3 → 5 → 6 → 1 → 5`
- Maximum value is `6`

```embed-go
PATH: "vault://leetcode/0001-0100/0053_maximum_subarray/solution.go"
TITLE: "leetcode 53. Maximum Subarray"
```

### Approach 2

**Divide and Conquer**: Split the array into left and right halves; the maximum subarray is either in the left half, right half, or crosses the midpoint.

**Principle:**
Recursively decompose the problem into smaller subproblems, then merge results. The case crossing the midpoint requires extending from the midpoint both left and right to find the maximum sum.

**Steps:**
1. Base case: when array length is 1, return that element
2. Split the array at the midpoint into left and right halves
3. Recursively solve for the maximum subarray sum in the left half
4. Recursively solve for the maximum subarray sum in the right half
5. Extend from the midpoint leftward to find the maximum left boundary sum
6. Extend from the midpoint rightward to find the maximum right boundary sum
7. Return the maximum of the three values

```embed-go
PATH: "vault://leetcode/0001-0100/0053_maximum_subarray/solution2.go"
TITLE: "leetcode 53. Maximum Subarray"
```

### Approach 3

**Brute Force (Optimized)**: Enumerate all subarray start and end points, using prefix sum concept for cumulative summation.

**Principle:**
Fix starting point i, extend ending point j from i, each time only adding the new element `nums[j]` to get the sum of subarray `[i,j]`.

**Steps:**
1. Outer loop enumerates starting point i
2. Inner loop enumerates ending point j, accumulating `nums[j]` to `thisSum`
3. Update `maxSum` after each accumulation

```embed-go
PATH: "vault://leetcode/0001-0100/0053_maximum_subarray/solution3.go"
TITLE: "leetcode 53. Maximum Subarray"
```

### Approach 4

**Brute Force (Naive)**: Enumerate all subarrays, independently calculate the sum of each subarray.

**Principle:**
Triple loop: outer two loops enumerate subarray start i and end j, inner loop calculates the element sum from i to j.

**Steps:**
1. Outer loop enumerates starting point i
2. Middle loop enumerates ending point j
3. Inner loop accumulates sum from i to j
4. Update `maxSum` after each sum calculation

```embed-go
PATH: "vault://leetcode/0001-0100/0053_maximum_subarray/solution4.go"
TITLE: "leetcode 53. Maximum Subarray"
```
