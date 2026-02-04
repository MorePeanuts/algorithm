---
link: https://leetcode.com/problems/largest-rectangle-in-histogram/
tags:
  - Hard
  - Stack
  - Array
  - Monotonic_Stack
---
## Description
Given an array of integers `heights` representing the histogram's bar height where the width of each bar is `1`, return *the area of the largest rectangle in the histogram*.

**Example 1:**

![](https://assets.leetcode.com/uploads/2021/01/04/histogram.jpg)

```
Input: heights = [2,1,5,6,2,3]
Output: 10
Explanation: The above is a histogram where width of each bar is 1.
The largest rectangle is shown in the red area, which has an area = 10 units.
```

**Example 2:**

![](https://assets.leetcode.com/uploads/2021/01/04/histogram-1.jpg)

```
Input: heights = [2,4]
Output: 4
```

**Constraints:**

- `1 <= heights.length <= 10^5`
- `0 <= heights[i] <= 10^4`

## Solution

### Approach 1

**Monotonic Stack for Left/Right Boundaries**: Use monotonic stacks to precompute the left and right boundaries for each bar (the first shorter bar on each side), then calculate the maximum rectangle area using each bar's height.

**Principle:**
For each bar, the maximum rectangle with that bar's height is bounded by the first shorter bar on the left and right. A monotonically increasing stack efficiently finds these boundaries in linear time by processing elements in order and popping when a shorter bar is encountered.

**Steps:**
1. Initialize `left` and `right` arrays to store the index of the first shorter bar on each side
2. Traverse left-to-right maintaining a monotonically increasing stack (storing indices). When a shorter bar is encountered, the top element's right boundary is found
3. Simultaneously traverse right-to-left (using symmetry) with another stack to determine left boundaries
4. After traversal, elements remaining in the stack have no shorter bars on that side; set their boundary to the array edge
5. For each bar, calculate `height[i] * (right[i] - left[i] - 1)` and track the maximum

**Example:**
For `heights = [2,1,5,6,2,3]`:
- Bar `5` (index 2): left boundary at index 1 (height 1), right boundary at index 4 (height 2)
- Width = 4 - 1 - 1 = 2, Area = 5 × 2 = 10
- Bar `6` (index 3): left boundary at index 2 (height 5), right boundary at index 4 (height 2)
- Width = 4 - 2 - 1 = 1, Area = 6 × 1 = 6
- Maximum area is 10

**Complexity Analysis:**
- Time: O(n) - each element is pushed and popped at most once
- Space: O(n) - for left, right arrays and stacks

> **Note**:
> 1. Store indices (not heights) in the stack to facilitate width calculation
> 2. Use `<=` comparison to correctly handle adjacent bars with equal heights
> 3. Elements remaining in the stack after traversal extend to the array boundaries

```embed-go
PATH: "vault://leetcode/0001-0100/0084_largest_rectangle_in_histogram/solution.go"
TITLE: "leetcode 84. Largest Rectangle in Histogram"
```
