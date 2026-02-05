---
link: https://leetcode.com/problems/trapping-rain-water/
tags:
  - Hard
  - Stack
  - Array
  - Two_Pointers
  - Dynamic_Programming
  - Monotonic_Stack
---
## Description
Given `n` non-negative integers representing an elevation map where the width of each bar is `1`, compute how much water it can trap after raining.

**Example 1:**

![](https://assets.leetcode.com/uploads/2018/10/22/rainwatertrap.png)

```
Input: height = [0,1,0,2,1,0,1,3,2,1,2,1]
Output: 6
Explanation: The above elevation map (black section) is represented by array [0,1,0,2,1,0,1,3,2,1,2,1]. In this case, 6 units of rain water (blue section) are being trapped.
```

**Example 2:**

```
Input: height = [4,2,0,3,2,5]
Output: 9
```

**Constraints:**

- `n == height.length`
- `1 <= n <= 2 * 10^4`
- `0 <= height[i] <= 10^5`

## Solution
### Approach 1
**Monotonic Stack + Horizontal Filling**: Use a monotonic decreasing stack to find left and right boundaries, filling water horizontally layer by layer.

**Principle:**
Maintain a monotonic decreasing stack storing column indices. When encountering a column taller than the stack top, a water-trapping valley is formed. Fill the valley bottom up to the height of the lower boundary and accumulate the water volume.

**Steps:**
1. Initialize a monotonic decreasing stack and iterate through each column
2. While the stack is non-empty and the current column height exceeds the stack top's height, pop the stack
3. Determine the fill height: use the recorded top's height if the stack is empty, otherwise use the current column's height
4. Fill all columns between the stack top and the current position to the fill height, accumulating water volume
5. Push the current index onto the stack

**Example:**
- Input `[0,1,0,2,1,0,1,3,2,1,2,1]`
- At `i=3`, `height[3]=2 > height[2]=0`, fill position 2 to height 1 (left boundary height), water +1
- At `i=7`, `height[7]=3` triggers multiple fills, progressively leveling the valley
- Total accumulated water is 6

> **Note**: This approach modifies the original array; copy it first if preservation is needed. Time complexity O(n), space complexity O(n).

```embed-go
PATH: "vault://leetcode/0001-0100/0042_trapping_rain_water/solution.go"
TITLE: "leetcode 42. Trapping Rain Water"
```
