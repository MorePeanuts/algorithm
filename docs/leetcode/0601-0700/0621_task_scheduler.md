---
link: https://leetcode.com/problems/task-scheduler/
tags:
  - Medium
  - Greedy
  - Array
  - Hash_Table
  - Counting
  - Sorting
  - Heap_(Priority_Queue)
---
## Description
You are given an array of CPU `tasks`, each labeled with a letter from A to Z, and a number `n`. Each CPU interval can be idle or allow the completion of one task. Tasks can be completed in any order, but there's a constraint: there has to be a gap of **at least** `n` intervals between two tasks with the same label.

Return the **minimum** number of CPU intervals required to complete all tasks.

**Example 1:**

**Input:** tasks = ["A","A","A","B","B","B"], n = 2

**Output:** 8

**Explanation:** A possible sequence is: A -> B -> idle -> A -> B -> idle -> A -> B.

After completing task A, you must wait two intervals before doing A again. The same applies to task B. In the 3rd interval, neither A nor B can be done, so you idle. By the 4th interval, you can do A again as 2 intervals have passed.

**Example 2:**

**Input:** tasks = ["A","C","A","B","D","B"], n = 1

**Output:** 6

**Explanation:** A possible sequence is: A -> B -> C -> D -> A -> B.

With a cooling interval of 1, you can repeat a task after just one other task.

**Example 3:**

**Input:** tasks = ["A","A","A", "B","B","B"], n = 3

**Output:** 10

**Explanation:** A possible sequence is: A -> B -> idle -> idle -> A -> B -> idle -> idle -> A -> B.

There are only two types of tasks, A and B, which need to be separated by 3 intervals. This leads to idling twice between repetitions of these tasks.

**Constraints:**

- `1 <= tasks.length <= 104`
- `tasks[i]` is an uppercase English letter.
- `0 <= n <= 100`

## Solution
### Approach 1
**Method Name**: Max Heap + Queue Simulation

**Principle:**
Simulates the CPU scheduling process using a max-heap to prioritize tasks with the most remaining occurrences, and a queue to manage tasks in cooldown, ensuring the gap between identical tasks meets the cooling time requirement.

**Steps:**
1. Count the frequency of each task
2. Build a max-heap with all tasks sorted by frequency
3. Use a queue of length `n+1` to manage tasks in cooldown
4. Each time interval:
   - Take a task from the front of the queue (push back to heap if it has remaining count)
   - Extract the most frequent task from the heap and execute it
   - If the task still has remaining count, add it to the end of the queue to wait for cooldown
   - If no task is available, enter idle state
5. Accumulate time until all tasks are completed

**Example:**
Input: tasks = ["A","A","A","B","B","B"], n = 2
- Initial: heap [A(3), B(3)], queue [nil, nil, nil]
- Time 1: execute A(2) → queue [nil, nil, A(2)]
- Time 2: execute B(2) → queue [nil, A(2), B(2)]
- Time 3: idle → queue [A(2), B(2), nil]
- Time 4: A cooled down, execute A(1) → queue [B(2), nil, A(1)]
- Time 5: B cooled down, execute B(1) → queue [nil, A(1), B(1)]
- Time 6: idle → queue [A(1), B(1), nil]
- Time 7: A cooled down, execute A(0) → queue [B(1), nil, nil]
- Time 8: B cooled down, execute B(0) → done
- Total time: 8

```embed-go
PATH: "vault://leetcode/0601-0700/0621_task_scheduler/solution.go"
TITLE: "leetcode 621. Task Scheduler"
```

### Approach 2
**Method Name**: Greedy Algorithm

**Principle:**
Uses a greedy strategy to first arrange the most frequent tasks to form the basic scheduling framework, then fills other tasks into the gaps. When there are enough task types, gaps can be completely filled without idle time.

**Steps:**
1. Count the frequency of each task
2. Find the maximum frequency `maxFreq` and how many tasks have this maximum frequency `maxCount`
3. Calculate the base framework time: `(maxFreq - 1) * (n + 1) + maxCount`
4. The final result is the maximum of the above value and the total number of tasks (when there are many task types, no idle time may be needed)

**Example:**
Input: tasks = ["A","A","A","B","B","B"], n = 2
- Max frequency maxFreq = 3 (A and B each appear 3 times)
- maxCount = 2 (two tasks have the maximum frequency)
- Base framework time = (3-1) * (2+1) + 2 = 2*3+2 = 8
- Total tasks = 6
- Final result = max(8, 6) = 8

```embed-go
PATH: "vault://leetcode/0601-0700/0621_task_scheduler/solution2.go"
TITLE: "leetcode 621. Task Scheduler"
```
