---
link: https://leetcode.com/problems/contains-duplicate/
tags:
  - Easy
  - Array
  - Hash_Table
  - Sorting
---
## Description
Given an integer array `nums`, return `true` if any value appears **at least twice** in the array, and return `false` if every element is distinct.

**Example 1:**

**Input:** nums = [1,2,3,1]

**Output:** true

**Explanation:**

The element 1 occurs at the indices 0 and 3.

**Example 2:**

**Input:** nums = [1,2,3,4]

**Output:** false

**Explanation:**

All elements are distinct.

**Example 3:**

**Input:** nums = [1,1,1,3,3,4,3,2,4,2]

**Output:** true

**Constraints:**

- `1 <= nums.length <= 105`
- `-109 <= nums[i] <= 109`

## Solution

### Approach 1

**Hash Set Method**: Use a hash set to record traversed elements, return immediately when a duplicate is encountered.

**Principle:**
While traversing the array, store each element in a hash set. Before storing, check if the element already exists in the set. If it does, there's a duplicate element.

**Steps:**
1. Create an empty hash set
2. Iterate through each element in the array
3. Check if the current element already exists in the set
4. If exists, return `true` directly
5. If not, add the current element to the set
6. Return `false` after traversal completes

```embed-go
PATH: "vault://leetcode/0201-0300/0217_contains_duplicate/solution.go"
TITLE: "leetcode 217. Contains Duplicate"
```
