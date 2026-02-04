---
link: https://leetcode.com/problems/daily-temperatures/
tags:
  - Medium
  - Stack
  - Array
  - Monotonic_Stack
---
## Description
Given an array of integers `temperatures` represents the daily temperatures, return *an array* `answer` *such that* `answer[i]` *is the number of days you have to wait after the* `ith` *day to get a warmer temperature*. If there is no future day for which this is possible, keep `answer[i] == 0` instead.

**Example 1:**

```
Input: temperatures = [73,74,75,71,69,72,76,73]
Output: [1,1,4,2,1,1,0,0]
```

**Example 2:**

```
Input: temperatures = [30,40,50,60]
Output: [1,1,1,0]
```

**Example 3:**

```
Input: temperatures = [30,60,90]
Output: [1,1,0]
```

**Constraints:**

- `1 <= temperatures.length <= 10^5`
- `30 <= temperatures[i] <= 100`

## Solution

### Approach 1

**Monotonic Decreasing Stack**: Use a monotonic stack to efficiently find the next greater element for each position.

**Principle:**
Maintain a stack storing indices in decreasing order of temperatures. When encountering a temperature higher than the one at the stack's top, we've found the "next warmer day" for that index, and the difference in indices gives us the answer.

**Steps:**
1. Initialize result array `answer` and empty `stack` (storing indices)
2. Iterate through the temperatures array, for each temperature `t`:
   - While stack is non-empty and current temperature is higher than temperature at stack top, pop the top and calculate the day difference as the answer
   - Repeat until stack is empty or top temperature is not less than current
   - Push current index onto the stack
3. After iteration, remaining indices in stack keep their answer as 0

**Example:**
- Input `[73,74,75,71,69,72,76,73]`
- Processing:
  - `i=0`: Stack empty, push `[0]`
  - `i=1`: 74>73, pop 0, `answer[0]=1`, push `[1]`
  - `i=2`: 75>74, pop 1, `answer[1]=1`, push `[2]`
  - `i=3`: 71<75, push `[2,3]`
  - `i=4`: 69<71, push `[2,3,4]`
  - `i=5`: 72>69, pop 4, `answer[4]=1`; 72>71, pop 3, `answer[3]=2`; 72<75, push `[2,5]`
  - `i=6`: 76>72, pop 5, `answer[5]=1`; 76>75, pop 2, `answer[2]=4`, push `[6]`
  - `i=7`: 73<76, push `[6,7]`
- Output `[1,1,4,2,1,1,0,0]`

```embed-go
PATH: "vault://leetcode/0701-0800/0739_daily_temperatures/solution.go"
TITLE: "leetcode 739. Daily Temperatures"
```

### Approach 2

**Reverse Traversal with Jump Optimization**: Traverse from back to front, leveraging pre-computed answers to jump efficiently.

**Principle:**
Iterate from the end of the array backwards. For each position, use the already computed "waiting days" of subsequent positions to jump directly, quickly locating the next warmer temperature.

**Steps:**
1. Initialize result array `answer`, iterate from second-to-last element backwards
2. For each position `i`, start searching from `j = i+1`:
   - If `temperatures[j] > temperatures[i]`, found answer `answer[i] = j - i`
   - If `temperatures[j] <= temperatures[i]` and `answer[j] == 0`, no warmer day exists after `j`, so none for `i` either
   - Otherwise jump to `j += answer[j]`, using pre-computed answers to accelerate search
3. Return `answer` after iteration

**Example:**
- Input `[73,74,75,71,69,72,76,73]`
- Reverse processing:
  - `i=7`: Last element, `answer[7]=0`
  - `i=6`: j=7, 73<76, `answer[6]=0`
  - `i=5`: j=6, 76>72, `answer[5]=1`
  - `i=4`: j=5, 72>69, `answer[4]=1`
  - `i=3`: j=4, 69<71, jump j=4+1=5; 72>71, `answer[3]=2`
  - `i=2`: j=3, 71<75, jump j=3+2=5; 72<75, jump j=5+1=6; 76>75, `answer[2]=4`
  - `i=1`: j=2, 75>74, `answer[1]=1`
  - `i=0`: j=1, 74>73, `answer[0]=1`
- Output `[1,1,4,2,1,1,0,0]`

```embed-go
PATH: "vault://leetcode/0701-0800/0739_daily_temperatures/solution2.go"
TITLE: "leetcode 739. Daily Temperatures"
```
