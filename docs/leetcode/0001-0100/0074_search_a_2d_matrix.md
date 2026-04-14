---
link: https://leetcode.com/problems/search-a-2d-matrix/
tags:
  - Medium
  - Array
  - Binary_Search
  - Matrix
---
## Description
You are given an `m x n` integer matrix `matrix` with the following two properties:

- Each row is sorted in non-decreasing order.
- The first integer of each row is greater than the last integer of the previous row.

Given an integer `target`, return `true` *if* `target` *is in* `matrix` *or* `false` *otherwise*.

You must write a solution in `O(log(m * n))` time complexity.

**Example 1:**

![](https://assets.leetcode.com/uploads/2020/10/05/mat.jpg)

```
Input: matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 3
Output: true
```

**Example 2:**

![](https://assets.leetcode.com/uploads/2020/10/05/mat2.jpg)

```
Input: matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 13
Output: false
```

**Constraints:**

- `m == matrix.length`
- `n == matrix[i].length`
- `1 <= m, n <= 100`
- `-104 <= matrix[i][j], target <= 104`

## Solution
### Approach 1
**Method Name**: Double Binary Search

**Principle:**
Leverage the sorted property of the matrix: first use binary search to determine which row might contain the target, then use binary search again on that row to check if the target exists.

**Steps:**
1. Extract the first element of each row to form a sorted array `firstNums`
2. Perform binary search on `firstNums` to determine which row might contain the target
3. If target is found, return true; otherwise continue binary search on the determined row
4. Return the result of the second search

**Example:**
Input: matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 13
- firstNums = [1, 10, 23]
- First binary search: 13 not in firstNums, determined to be in row 1 ([10,11,16,20])
- Second binary search: 13 not in that row, return false

> **Note**: The key is the design of the `binarySearchPos` function, which not only finds whether the target exists but also returns the index before where the target would be inserted when not found.

```embed-go
PATH: "vault://leetcode/0001-0100/0074_search_a_2d_matrix/solution.go"
TITLE: "leetcode 74. Search a 2D Matrix"
```

---

### Detailed Explanation of `binarySearchPos` Function

The `binarySearchPos` function is a carefully designed binary search variant. Its core feature is: **whether the target is found or not, it returns meaningful position information**.

#### Why does the returned pos satisfy `nums[pos] < target < nums[pos+1]` when target is not found?

Let's understand this conclusion by analyzing loop invariants:

**1. Initialization and Loop Condition**
```go
l, r := 0, len(nums)  // Note: r is initialized to len(nums), not len(nums)-1
for l < r {
```
- The search interval is half-open `[l, r)`
- `l` is the left boundary, `r` is the right boundary (exclusive)

**2. Loop Invariants**
During each iteration, the following properties are maintained:
- For all `i < l`, we have `nums[i] < target`
- For all `i >= r`, we have `nums[i] > target`

**3. State at Loop Termination**
When `l == r`, the loop terminates. At this point:
- `l` (equal to `r`) is the **insertion position** where target should be inserted to keep the array sorted
- All elements to the left of `l` are `< target`
- All elements to the right of `l` are `> target`

**4. Return Value Calculation**
```go
return l - 1, false
```
- Returns `l - 1`, which is the index of the largest element smaller than target
- At this point: `nums[l-1] < target < nums[l]` (assuming both `l-1` and `l` are within array bounds)

#### Boundary Case Handling

| Case | Value of l | Return Value | Meaning |
|------|------------|--------------|---------|
| target < all elements | 0 | -1 | target is smaller than all elements |
| target > all elements | len(nums) | len(nums)-1 | target is larger than all elements |
| target in the middle | insertion position | insertion position-1 | nums[pos] < target < nums[pos+1] |

#### Example Verification

Take `nums = [1, 10, 23]`, `target = 13` as an example:

| Iteration | l | r | m | nums[m] | Comparison | Action |
|-----------|---|---|---|---------|------------|--------|
| Initial | 0 | 3 | - | - | - | - |
| 1 | 0 | 3 | 1 | 10 | 13 > 10 | l = m + 1 = 2 |
| 2 | 2 | 3 | 2 | 23 | 13 < 23 | r = m = 2 |
| End | 2 | 2 | - | - | - | Return l-1 = 1 |

Result: `pos = 1`, satisfying `nums[1] = 10 < 13 < 23 = nums[2]`
