---
link: https://leetcode.com/problems/search-in-rotated-sorted-array/
tags:
  - Medium
  - Array
  - Binary_Search
---
## Description
There is an integer array `nums` sorted in ascending order (with **distinct** values).

Prior to being passed to your function, `nums` is **possibly left rotated** at an unknown index `k` (`1 <= k < nums.length`) such that the resulting array is `[nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]]` (**0-indexed**). For example, `[0,1,2,4,5,6,7]` might be left rotated by `3` indices and become `[4,5,6,7,0,1,2]`.

Given the array `nums` **after** the possible rotation and an integer `target`, return *the index of* `target` *if it is in* `nums`*, or* `-1` *if it is not in* `nums`.

You must write an algorithm with `O(log n)` runtime complexity.

**Example 1:**

```
Input: nums = [4,5,6,7,0,1,2], target = 0
Output: 4
```

**Example 2:**

```
Input: nums = [4,5,6,7,0,1,2], target = 3
Output: -1
```

**Example 3:**

```
Input: nums = [1], target = 0
Output: -1
```

**Constraints:**

- `1 <= nums.length <= 5000`
- `-104 <= nums[i] <= 104`
- All values of `nums` are **unique**.
- `nums` is an ascending array that is possibly rotated.
- `-104 <= target <= 104`

## Solution
### Approach 1
**Method Name**: Find Rotation Point then Binary Search

**Principle:**
The rotated array can be viewed as two sorted subarrays. First find the rotation point (position of minimum value) using binary search, then determine which subarray to search in based on comparison between target and the first element.

**Steps:**
1. Check if array is not rotated (nums[0] < nums[len(nums)-1]), if so perform binary search on entire array
2. Find rotation point l (index of minimum value) using binary search
3. Compare target with nums[0]:
   - If target < nums[0], search in right half nums[l:]
   - If target == nums[0], return 0 directly
   - If target > nums[0], search in left half nums[:l]

**Example:**
- Input nums = [4,5,6,7,0,1,2], target = 0
- Find rotation point l = 4 (value 0)
- target < nums[0] (0 < 4), search in nums[4:] = [0,1,2], found at index 0
- Return 0 + 4 = 4

​```embed-go
PATH: "vault://leetcode/0001-0100/0033_search_in_rotated_sorted_array/solution.go"
TITLE: "leetcode 33. Search in Rotated Sorted Array"
​```

### Approach 2
**Method Name**: One-pass Binary Search

**Principle:**
In a rotated sorted array, at least one side of any midpoint is sorted. Using this property, we can adjust the search range by determining which side is sorted and whether the target lies within that sorted interval.

**Steps:**
1. Initialize left and right pointers l = 0, r = len(nums)
2. While l < r:
   - Calculate midpoint m = l + (r-l)/2
   - If nums[m] == target, return m directly
   - If nums[m] >= nums[l], left half [l, m) is sorted:
     - If target is in range [nums[l], nums[m]), search in left half
     - Otherwise search in right half
   - Otherwise, right half [m, r) is sorted:
     - If target < nums[m] or target > nums[r-1], search in left half
     - Otherwise search in right half
3. Return -1 if not found

**Example:**
- Input nums = [4,5,6,7,0,1,2], target = 0
- l=0, r=7, m=3, nums[m]=7 >= nums[0]=4, left half sorted
- target=0 not in range [4,7), l=4
- l=4, r=7, m=5, nums[m]=1 >= nums[4]=0, left half sorted
- target=0 == nums[4], but nums[m]=1 != 0
- target=0 < nums[m]=1, r=5
- l=4, r=5, m=4, nums[m]=0 == target, return 4

​```embed-go
PATH: "vault://leetcode/0001-0100/0033_search_in_rotated_sorted_array/solution2.go"
TITLE: "leetcode 33. Search in Rotated Sorted Array"
​```

