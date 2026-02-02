---
link: https://leetcode.com/problems/two-sum/
tags:
  - Easy
  - Array
  - Hash_Table
---
## Description
Given an array of integers `nums` and an integer `target`, return *indices of the two numbers such that they add up to `target`*.

You may assume that each input would have ***exactly* one solution**, and you may not use the *same* element twice.

You can return the answer in any order.

**Example 1:**

```
Input: nums = [2,7,11,15], target = 9
Output: [0,1]
Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].
```

**Example 2:**

```
Input: nums = [3,2,4], target = 6
Output: [1,2]
```

**Example 3:**

```
Input: nums = [3,3], target = 6
Output: [0,1]
```

**Constraints:**

- `2 <= nums.length <= 104`
- `-109 <= nums[i] <= 109`
- `-109 <= target <= 109`
- **Only one valid answer exists.**

**Follow-up:**Can you come up with an algorithm that is less than `O(n2)` time complexity?

## Solution

### Approach 1

**Hash Table Method**: Use a hash table to store traversed elements' values and indices, achieving O(1) time lookup for paired elements.

**Principle:**
For each element `num`, calculate its complement `target - num`. Use the hash table to quickly check if the complement already exists. If it does, return both indices directly.

**Steps:**
1. Create a hash table with array element values as keys and their indices as values
2. Iterate through the array, calculating `diff = target - num` for each element
3. Check if `diff` exists in the hash table
4. If exists, return `[hashTable[diff], i]`
5. If not, store the current element and its index in the hash table

```embed-go
PATH: "vault://leetcode/0001-0100/0001_two_sum/solution.go"
TITLE: "leetcode 1. Two Sum"
```
