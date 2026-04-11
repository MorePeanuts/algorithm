---
link: https://leetcode.com/problems/find-the-duplicate-number/
tags:
  - Medium
  - Bit_Manipulation
  - Array
  - Two_Pointers
  - Binary_Search
---
## Description
Given an array of integers `nums` containing `n + 1` integers where each integer is in the range `[1, n]` inclusive.

There is only **one repeated number** in `nums`, return *this repeated number*.

You must solve the problem **without** modifying the array `nums` and using only constant extra space.

**Example 1:**

```
Input: nums = [1,3,4,2,2]
Output: 2
```

**Example 2:**

```
Input: nums = [3,1,3,4,2]
Output: 3
```

**Example 3:**

```
Input: nums = [3,3,3,3,3]
Output: 3
```

**Constraints:**

- `1 <= n <= 105`
- `nums.length == n + 1`
- `1 <= nums[i] <= n`
- All the integers in `nums` appear only **once** except for **precisely one integer** which appears **two or more** times.

**Follow up:**

- How can we prove that at least one duplicate number must exist in `nums`?
- Can you solve the problem in linear runtime complexity?

## Solution
### Approach 1
**Method Name**: Fast and Slow Pointers (Floyd's Tortoise and Hare Algorithm)

**Principle:**
Treat the array as a linked list where indices are nodes and values are next pointers.

**Why a duplicate number guarantees a cycle:**
- The array has `n+1` elements, each value in range `[1, n]`
- Suppose there's a duplicate number `d`, then there exist at least two different indices `i` and `j` such that `nums[i] = d` and `nums[j] = d`
- This means node `i` and node `j` both point to node `d`, forming a "convergence point"
- Since every node has out-degree 1 (each index has exactly one `nums[i]`), once two paths converge, the subsequent path will be identical, creating a cycle

**Why starting from 0 guarantees entering the cycle:**
- 0 is not in the value range `[1, n]`, so no node points to 0, making 0 the unique starting point of the linked list
- Traversing from 0, with `n+1` nodes but only `n` possible values, by the pigeonhole principle, we must eventually visit a node that has been visited before
- Once we visit a duplicate node, we have found the entry point of the cycle

Find the cycle entry point using fast and slow pointers, which is the duplicate number.

**Steps:**
1. Initialize both slow and fast pointers to index 0
2. Slow pointer moves one step (slow = nums[slow]), fast pointer moves two steps (fast = nums[nums[fast]])
3. When slow and fast meet, reset fast pointer to index 0
4. Move both pointers one step at a time until they meet again
5. The value at the meeting position is the duplicate number

**Example:**
- Input `[1,3,4,2,2]` -> Pointers first meet at index 4 -> Reset fast to 0 -> Move synchronously and meet at index 2 (value 2) -> Output 2

```embed-go
PATH: "vault://leetcode/0201-0300/0287_find_the_duplicate_number/solution.go"
TITLE: "leetcode 287. Find the Duplicate Number"
```
