---
link: https://leetcode.com/problems/sliding-window-maximum/
tags:
  - Hard
  - Queue
  - Array
  - Sliding_Window
  - Monotonic_Queue
  - Heap_(Priority_Queue)
---
## Description
You are given an array of integers `nums`, there is a sliding window of size `k` which is moving from the very left of the array to the very right. You can only see the `k` numbers in the window. Each time the sliding window moves right by one position.

Return *the max sliding window*.

**Example 1:**

```
Input: nums = [1,3,-1,-3,5,3,6,7], k = 3
Output: [3,3,5,5,6,7]
Explanation: 
Window position                Max
---------------               -----
[1  3  -1] -3  5  3  6  7       3
 1 [3  -1  -3] 5  3  6  7       3
 1  3 [-1  -3  5] 3  6  7       5
 1  3  -1 [-3  5  3] 6  7       5
 1  3  -1  -3 [5  3  6] 7       6
 1  3  -1  -3  5 [3  6  7]      7
```

**Example 2:**

```
Input: nums = [1], k = 1
Output: [1]
```

**Constraints:**

- `1 <= nums.length <= 10^5`
- `-10^4 <= nums[i] <= 10^4`
- `1 <= k <= nums.length`

## Solution
### Approach 1
**Method Name**: Monotonic Queue

**Principle:**
Use a deque (double-ended queue) to maintain indices of elements that could be the maximum in the sliding window. The queue keeps elements in monotonically decreasing order, so the front of the queue is always the maximum value of the current window.

**Steps:**
1. Initialize an empty deque
2. Process the first k elements to build the initial monotonic queue
3. Iterate through the remaining elements:
   - Remove elements from the front that are out of the window bounds
   - Remove all elements smaller than the current element from the back (they can't be maximum in any future window)
   - Add the current element to the back of the queue
   - The front of the queue is the maximum for the current window, add to result

**Example:**
Input nums = [1,3,-1,-3,5,3,6,7], k = 3
- Window [1,3,-1]: Queue [3,-1] → Max 3
- Window [3,-1,-3]: Queue [3,-1,-3] → Max 3
- Window [-1,-3,5]: Queue [5] → Max 5
- Window [-3,5,3]: Queue [5,3] → Max 5
- Window [5,3,6]: Queue [6] → Max 6
- Window [3,6,7]: Queue [7] → Max 7
Output: [3,3,5,5,6,7]

**Complexity Analysis:**
- Time Complexity: O(n), each element is enqueued and dequeued at most once
- Space Complexity: O(k), the queue stores at most k elements

> **Note:**
> - The queue stores element values, but indices can also be stored to more precisely determine if an element is within the window
> - When encountering equal elements, there's no need to remove elements from the queue because they occupy different positions in the window

```embed-go
PATH: "vault://leetcode/0201-0300/0239_sliding_window_maximum/solution.go"
TITLE: "leetcode 239. Sliding Window Maximum"
```
