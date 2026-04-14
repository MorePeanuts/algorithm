---
link: https://leetcode.com/problems/binary-search/
tags:
  - Easy
  - Array
  - Binary_Search
---
## Description
Given an array of integers `nums` which is sorted in ascending order, and an integer `target`, write a function to search `target` in `nums`. If `target` exists, then return its index. Otherwise, return `-1`.

You must write an algorithm with `O(log n)` runtime complexity.

**Example 1:**

```
Input: nums = [-1,0,3,5,9,12], target = 9
Output: 4
Explanation: 9 exists in nums and its index is 4
```

**Example 2:**

```
Input: nums = [-1,0,3,5,9,12], target = 2
Output: -1
Explanation: 2 does not exist in nums so return -1
```

**Constraints:**

- `1 <= nums.length <= 104`
- `-104 < nums[i], target < 104`
- All the integers in `nums` are **unique**.
- `nums` is sorted in ascending order.

## Solution

### Approach 1

**Recursive Binary Search**: Recursively narrow down the search interval, halving the search range each time.

**Principle:**
Utilize the sorted property of the array. Each time compare the middle element with the target value, then recursively search the left or right half based on the comparison result.

**Steps:**
1. If array is empty, return -1
2. Get middle position `center = n / 2`
3. If middle element equals target, return `center`
4. If target is less than middle element, recursively search left half `nums[:center]`
5. If target is greater than middle element, recursively search right half `nums[center+1:]`, adding offset `center + 1` to the returned index

```embed-go
PATH: "vault://leetcode/0701-0800/0704_binary_search/solution.go"
TITLE: "leetcode 704. Binary Search"
```

### Approach 2

**Iterative Binary Search**: Use loops and two pointers to implement binary search, avoiding recursion overhead.

**Principle:**
Maintain left and right boundaries `[left, right)` as a left-closed, right-open interval. Each time adjust boundaries based on comparison of middle element with target, until the target is found or the interval becomes empty.

**Steps:**
1. Initialize `left = 0`, `right = len(nums)`
2. While `left < right`:
   - Calculate middle position `mid = left + (right - left) / 2` (avoids integer overflow)
   - If `target < nums[mid]`, shrink right boundary `right = mid`
   - If `target > nums[mid]`, shrink left boundary `left = mid + 1`
   - If equal, return `mid`
3. If loop ends without finding, return -1

```embed-go
PATH: "vault://leetcode/0701-0800/0704_binary_search/solution2.go"
TITLE: "leetcode 704. Binary Search"
```

### Approach 3

**Two-Way Comparison Binary Search**: Only do two-way comparison in the loop, check for target match outside the loop.

**Principle:**
Using a left-closed, right-open interval `[left, right)`, each iteration only checks if the target is less than the middle element. After the loop, `left` points to the first position where the element is greater than or equal to the target. We need to check if the element before this position is the target.

**Steps:**
1. Initialize `left = 0`, `right = len(nums)`
2. While `left < right`:
   - Calculate middle position `mid = left + (right - left) / 2`
   - If `target < nums[mid]`, shrink right boundary `right = mid` (target is in the left half)
   - Otherwise, shrink left boundary `left = mid + 1` (target might be to the right of mid)
3. After loop ends, check `left > 0 && nums[left-1] == target`, return `left-1` if equal, otherwise return -1

```embed-go
PATH: "vault://leetcode/0701-0800/0704_binary_search/solution3.go"
TITLE: "leetcode 704. Binary Search"
```
