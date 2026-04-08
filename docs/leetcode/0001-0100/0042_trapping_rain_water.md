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
**Dynamic Programming + Prefix/Suffix Maxima**: Precompute the maximum height to the left and right of each position to directly determine the amount of water each position can trap.

**Principle:**
For each position, the amount of water it can trap is determined by the shorter of the two tallest pillars on either side. By precomputing a prefix maximum array (from left to right) and a suffix maximum array (from right to left), we can get the left and right maximum heights for any position in O(1) time.

**Steps:**
1. Create prefix maximum array `prefix`, where `prefix[i]` is the maximum height to the left of position `i`
2. Create suffix maximum array `suffix`, where `suffix[i]` is the maximum height to the right of position `i`
3. Iterate through each position, calculate the water that can be trapped: `min(prefix[i], suffix[i]) - height[i]`
4. Sum all positive water amounts to get the result

**Example:**
- Input `[0,1,0,2,1,0,1,3,2,1,2,1]`
- `prefix` array: `[0,0,1,1,2,2,2,2,3,3,3,3]`
- `suffix` array: `[3,3,3,3,3,3,3,2,2,2,1,0]`
- Position 2: `min(1,3) - 0 = 1`, water +1
- Position 5: `min(2,3) - 0 = 2`, water +2
- Total accumulated water is 6

> **Note**: Time complexity O(n), space complexity O(n). This method does not modify the original array.

```embed-go
PATH: "vault://leetcode/0001-0100/0042_trapping_rain_water/solution.go"
TITLE: "leetcode 42. Trapping Rain Water"
```

### Approach 2
**Two Pointers + Space Optimization**: Based on dynamic programming, use two pointers traversing from both ends towards the center, using two variables instead of arrays to record left and right maximum heights.

**Principle:**
When the maximum height at the left pointer is less than the maximum height at the right pointer, the water trapped at the left pointer position is determined by the left maximum height (because there must be a taller pillar on the right). Conversely, the water trapped at the right pointer position is determined by the right maximum height. This way, we can complete the calculation in one pass with only constant space.

**Steps:**
1. Initialize left pointer `left=0`, right pointer `right=n-1`
2. Initialize left maximum height `prefix=height[0]`, right maximum height `suffix=height[n-1]`
3. Loop while `right-left > 1`:
   - If `height[left] <= height[right]`, process the position to the right of the left pointer, update left maximum height, move left pointer right
   - Otherwise, process the position to the left of the right pointer, update right maximum height, move right pointer left
4. Sum the water amounts to get the result

**Example:**
- Input `[0,1,0,2,1,0,1,3,2,1,2,1]`
- `left=0, right=11`, `prefix=0, suffix=1`, process right side
- `left=0, right=10`, `prefix=0, suffix=2`, process left side
- Gradually move towards the center, total accumulated water is 6

> **Note**: Time complexity O(n), space complexity O(1). This method is the optimal solution and does not modify the original array.

```embed-go
PATH: "vault://leetcode/0001-0100/0042_trapping_rain_water/solution2.go"
TITLE: "leetcode 42. Trapping Rain Water"
```

### Approach 3
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
PATH: "vault://leetcode/0001-0100/0042_trapping_rain_water/solution3.go"
TITLE: "leetcode 42. Trapping Rain Water"
```