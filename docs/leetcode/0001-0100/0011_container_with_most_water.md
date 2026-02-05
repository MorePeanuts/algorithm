---
link: https://leetcode.com/problems/container-with-most-water/
tags:
  - Medium
  - Greedy
  - Array
  - Two_Pointers
---
## Description
You are given an integer array `height` of length `n`. There are `n` vertical lines drawn such that the two endpoints of the `ith` line are `(i, 0)` and `(i, height[i])`.

Find two lines that together with the x-axis form a container, such that the container contains the most water.

Return *the maximum amount of water a container can store*.

**Notice** that you may not slant the container.

**Example 1:**

![](https://s3-lc-upload.s3.amazonaws.com/uploads/2018/07/17/question_11.jpg)

```
Input: height = [1,8,6,2,5,4,8,3,7]
Output: 49
Explanation: The above vertical lines are represented by array [1,8,6,2,5,4,8,3,7]. In this case, the max area of water (blue section) the container can contain is 49.
```

**Example 2:**

```
Input: height = [1,1]
Output: 1
```

**Constraints:**

- `n == height.length`
- `2 <= n <= 10^5`
- `0 <= height[i] <= 10^4`

## Solution

### Approach 1

**Optimized Two Pointers Greedy**: Shrink from both ends, moving the shorter side and skipping shorter heights.

**Principle:**
The container area is determined by the shorter side and the width. Use two pointers starting from both ends, always moving the pointer at the shorter side, since moving the longer side only decreases width without increasing height, thus reducing area.

**Steps:**
1. Initialize left and right pointers at both ends of the array
2. Calculate current area and update the maximum
3. Move the pointer at the shorter side, skipping all positions with heights not greater than current
4. Repeat until pointers meet

**Example:**
- Input `[1,8,6,2,5,4,8,3,7]`
- Initial `l=0, r=8`, heights `1,7`, area `8*1=8`
- Move left pointer to `l=1` (skip height 1), heights `8,7`, area `7*7=49`
- Move right pointer to `r=6` (skip height 3), heights `8,8`, area `5*8=40`
- Continue shrinking... final max area is `49`

```embed-go
PATH: "vault://leetcode/0001-0100/0011_container_with_most_water/solution.go"
TITLE: "leetcode 11. Container With Most Water"
```
