---
link: https://leetcode.com/problems/car-fleet/
tags:
  - Medium
  - Stack
  - Array
  - Sorting
  - Monotonic_Stack
---
## Description
There are `n` cars at given miles away from the starting mile 0, traveling to reach the mile `target`.

You are given two integer arrays `position` and `speed`, both of length `n`, where `position[i]` is the starting mile of the `ith` car and `speed[i]` is the speed of the `ith` car in miles per hour.

A car cannot pass another car, but it can catch up and then travel next to it at the speed of the slower car.

A **car fleet** is a single car or a group of cars driving next to each other. The speed of the car fleet is the **minimum** speed of any car in the fleet.

If a car catches up to a car fleet at the mile `target`, it will still be considered as part of the car fleet.

Return the number of car fleets that will arrive at the destination.

**Example 1:**

**Input:** `target = 12, position = [10,8,0,5,3], speed = [2,4,1,1,3]`

**Output:** `3`

**Explanation:**

- The cars starting at 10 (speed 2) and 8 (speed 4) become a fleet, meeting each other at 12. The fleet forms at `target`.
- The car starting at 0 (speed 1) does not catch up to any other car, so it is a fleet by itself.
- The cars starting at 5 (speed 1) and 3 (speed 3) become a fleet, meeting each other at 6. The fleet moves at speed 1 until it reaches `target`.

**Example 2:**

**Input:** `target = 10, position = [3], speed = [3]`

**Output:** `1`

**Explanation:**

There is only one car, hence there is only one fleet.

**Example 3:**

**Input:** `target = 100, position = [0,2,4], speed = [4,2,1]`

**Output:** `1`

**Explanation:**

- The cars starting at 0 (speed 4) and 2 (speed 2) become a fleet, meeting each other at 4. The car starting at 4 (speed 1) travels to 5.
- Then, the fleet at 4 (speed 2) and the car at position 5 (speed 1) become one fleet, meeting each other at 6. The fleet moves at speed 1 until it reaches `target`.

**Constraints:**

- `n == position.length == speed.length`
- `1 <= n <= 10^5`
- `0 < target <= 10^6`
- `0 <= position[i] < target`
- All the values of `position` are **unique**.
- `0 < speed[i] <= 10^6`

## Solution
### Approach 1
**Sorting + Arrival Time Comparison**: Sort cars by position in descending order and compare arrival times to determine fleet formation.

**Principle:**
If a car behind takes less or equal time to reach the target compared to the car ahead, it will catch up and form a fleet; otherwise, it forms a new fleet.

**Steps:**
1. Create a position-to-speed mapping and calculate arrival time for each car
2. Sort positions in descending order (process from the car closest to target)
3. Iterate through each car, comparing its arrival time with the current fleet leader's arrival time
4. If the current car's arrival time is longer, it cannot catch up and forms a new fleet; update the leader's arrival time

**Example:**
- Input: `target = 12, position = [10,8,0,5,3], speed = [2,4,1,1,3]`
- Sort by position descending: `[10, 8, 5, 3, 0]`
- Calculate arrival times:
  - Position 10, speed 2: time = (12-10)/2 = 1.0 (fleet 1)
  - Position 8, speed 4: time = (12-8)/4 = 1.0 ≤ 1.0 (joins fleet 1)
  - Position 5, speed 1: time = (12-5)/1 = 7.0 > 1.0 (fleet 2)
  - Position 3, speed 3: time = (12-3)/3 = 3.0 ≤ 7.0 (joins fleet 2)
  - Position 0, speed 1: time = (12-0)/1 = 12.0 > 7.0 (fleet 3)
- Output: 3 fleets

```embed-go
PATH: "vault://leetcode/0801-0900/0853_car_fleet/solution.go"
TITLE: "leetcode 853. Car Fleet"
```

### Approach 2
**Bucket Sort Optimization**: Use bucket sort to avoid the standard library sort, achieving linear time complexity when the position range is small.

**Principle:**
Since position values are in the range `[0, target)`, we can use a boolean array as buckets to mark positions with cars, then iterate backwards to obtain a sorted position sequence without explicit sorting.

**Steps:**
1. Create a boolean array of length target as buckets, marking positions with cars
2. Build a position-to-speed mapping
3. Iterate the bucket array from target-1 to 0, collecting positions with cars (automatically in descending order)
4. Apply the same arrival time comparison logic as Approach 1 to count fleets

**Example:**
- Input: `target = 12, position = [10,8,0,5,3], speed = [2,4,1,1,3]`
- Bucket array: indices 0, 3, 5, 8, 10 are true
- Reverse iteration yields: `[10, 8, 5, 3, 0]` (same as Approach 1's sorted result)
- Subsequent calculation is identical to Approach 1, output: 3 fleets

> **Note**: When target is very large, this method consumes significant memory, making Approach 1 preferable; when n approaches target, this method achieves O(target) time complexity.

```embed-go
PATH: "vault://leetcode/0801-0900/0853_car_fleet/solution2.go"
TITLE: "leetcode 853. Car Fleet"
```
