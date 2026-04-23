---
link: https://leetcode.com/problems/k-closest-points-to-origin/
tags:
  - Medium
  - Geometry
  - Array
  - Math
  - Divide_and_Conquer
  - Quickselect
  - Sorting
  - Heap_(Priority_Queue)
---
## Description
Given an array of `points` where `points[i] = [xi, yi]` represents a point on the **X-Y** plane and an integer `k`, return the `k` closest points to the origin `(0, 0)`.

The distance between two points on the **X-Y** plane is the Euclidean distance (i.e., `√(x1 - x2)2 + (y1 - y2)2`).

You may return the answer in **any order**. The answer is **guaranteed** to be **unique** (except for the order that it is in).

**Example 1:**

![](https://assets.leetcode.com/uploads/2021/03/03/closestplane1.jpg)

```
Input: points = [[1,3],[-2,2]], k = 1
Output: [[-2,2]]
Explanation:
The distance between (1, 3) and the origin is sqrt(10).
The distance between (-2, 2) and the origin is sqrt(8).
Since sqrt(8) < sqrt(10), (-2, 2) is closer to the origin.
We only want the closest k = 1 points from the origin, so the answer is just [[-2,2]].
```

**Example 2:**

```
Input: points = [[3,3],[5,-1],[-2,4]], k = 2
Output: [[3,3],[-2,4]]
Explanation: The answer [[-2,4],[3,3]] would also be accepted.
```

**Constraints:**

- `1 <= k <= points.length <= 104`
- `-104 <= xi, yi <= 104`

## Solution
### Approach 1
**Method Name**: Min Heap

**Principle:**
Use a min heap (priority queue) to quickly find the k closest points to the origin. The heap comparison is based on the squared distance from the point to the origin (avoids square root operation for efficiency).

**Steps:**
1. Build a min heap with all points, the top of the heap is the closest point to origin
2. Pop k elements from the heap sequentially, these k elements are the answer

**Example:**
Input: points = [[1,3],[-2,2]], k = 1
- Calculate squared distance: (1,3) → 1+9=10, (-2,2) → 4+4=8
- Build min heap, top is [-2,2]
- Pop 1 element → [[-2,2]]

```embed-go
PATH: "vault://leetcode/0901-1000/0973_k_closest_points_to_origin/solution.go"
TITLE: "leetcode 973. K Closest Points to Origin"
```
