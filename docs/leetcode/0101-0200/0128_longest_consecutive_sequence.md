---
link: https://leetcode.com/problems/longest-consecutive-sequence/
tags:
  - Medium
  - Union_Find
  - Array
  - Hash_Table
---
## Description
Given an unsorted array of integers `nums`, return *the length of the longest consecutive elements sequence.*

You must write an algorithm that runs in `O(n)` time.

**Example 1:**

```
Input: nums = [100,4,200,1,3,2]
Output: 4
Explanation: The longest consecutive elements sequence is [1, 2, 3, 4]. Therefore its length is 4.
```

**Example 2:**

```
Input: nums = [0,3,7,2,5,8,4,6,0,1]
Output: 9
```

**Example 3:**

```
Input: nums = [1,0,1,2]
Output: 3
```

**Constraints:**

- `0 <= nums.length <= 105`
- `-109 <= nums[i] <= 109`

## Solution

### Approach 1

**Bidirectional Boundary Hash Tables**: Use two hash tables to record mappings from left boundary to right boundary and vice versa. Dynamically extend or merge intervals while traversing the array.

**Principle:**
Maintain two hash tables `left2right` and `right2left`, recording "left boundary outer → right boundary outer" and "right boundary outer → left boundary outer" mappings respectively. When a new number is added, extend or merge intervals based on its adjacency to existing intervals. Finally, traverse all intervals to calculate the maximum length.

**Steps:**
1. Initialize `left2right`, `right2left` hash tables and a `set` for deduplication
2. Traverse each number `num` in the array (skip duplicates):
   - Check if `num` is the left boundary outer (`left2right[num]`) or right boundary outer (`right2left[num]`) of some interval
   - **Only left interval exists**: Extend that interval one position to the right
   - **Only right interval exists**: Extend that interval one position to the left
   - **Both sides have intervals**: Merge the two intervals
   - **Neither side has interval**: Create a new single-element interval
3. Traverse `left2right`, calculate each interval's length (`right boundary outer - left boundary outer - 1`), take the maximum

> Note: Both hash tables need to be updated when modifying an interval.

**Example:**
Using `nums = [100, 4, 200, 1, 3, 2]`:
- `100`: Create interval, `left2right[99]=101`, `right2left[101]=99`
- `4`: Create interval, `left2right[3]=5`, `right2left[5]=3`
- `200`: Create interval, `left2right[199]=201`, `right2left[201]=199`
- `1`: Create interval, `left2right[0]=2`, `right2left[2]=0`
- `3`: 3 exists in `left2right` (interval [4,4]), extend left to [3,4]
- `2`: 2 exists in both `left2right` (interval [3,4]) and `right2left` (interval [1,1]), merge to [1,4]
- Final `left2right[0]=5`, interval length = 5 - 0 - 1 = 4

```embed-go
PATH: "vault://leetcode/0101-0200/0128_longest_consecutive_sequence/solution.go"
TITLE: "leetcode 128. Longest Consecutive Sequence"
```

### Approach 2

**Hash Set + Starting Point Enumeration**: Store all numbers in a hash set, only extend rightward from each consecutive sequence's starting point.

**Principle:**
A number is the starting point of a consecutive sequence if and only if `num-1` is not in the set. By skipping non-starting numbers, each number is visited at most twice (once for building the set, once for counting), ensuring O(n) time complexity.

**Steps:**
1. Store all numbers in hash set `set` (automatic deduplication)
2. Traverse each number `num` in the set:
   - If `num-1` exists in the set, skip (not a starting point)
   - Otherwise, starting from `num`, continuously check if `num+1`, `num+2`, ... exist, count sequence length
3. Return maximum length

**Example:**
Using `nums = [100, 4, 200, 1, 3, 2]`:
- Set: `{100, 4, 200, 1, 3, 2}`
- `100`: 99 doesn't exist, is starting point, extend right to get length 1
- `4`: 3 exists, skip
- `200`: 199 doesn't exist, is starting point, extend right to get length 1
- `1`: 0 doesn't exist, is starting point, extend 1→2→3→4, length 4
- `3`: 2 exists, skip
- `2`: 1 exists, skip
- Maximum length = 4

```embed-go
PATH: "vault://leetcode/0101-0200/0128_longest_consecutive_sequence/solution2.go"
TITLE: "leetcode 128. Longest Consecutive Sequence"
```
