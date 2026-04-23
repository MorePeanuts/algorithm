---
link: https://leetcode.com/problems/kth-largest-element-in-a-stream/
tags:
  - Easy
  - Tree
  - Design
  - Binary_Search_Tree
  - Binary_Tree
  - Data_Stream
  - Heap_(Priority_Queue)
---
## Description
You are part of a university admissions office and need to keep track of the `kth` highest test score from applicants in real-time. This helps to determine cut-off marks for interviews and admissions dynamically as new applicants submit their scores.

You are tasked to implement a class which, for a given integer `k`, maintains a stream of test scores and continuously returns the `k`th highest test score **after** a new score has been submitted. More specifically, we are looking for the `k`th highest score in the sorted list of all scores.

Implement the `KthLargest` class:

- `KthLargest(int k, int[] nums)` Initializes the object with the integer `k` and the stream of test scores `nums`.
- `int add(int val)` Adds a new test score `val` to the stream and returns the element representing the `kth` largest element in the pool of test scores so far.

**Example 1:**

**Input:**  
["KthLargest", "add", "add", "add", "add", "add"]  
[[3, [4, 5, 8, 2]], [3], [5], [10], [9], [4]]

**Output:** [null, 4, 5, 5, 8, 8]

**Explanation:**

KthLargest kthLargest = new KthLargest(3, [4, 5, 8, 2]);  
kthLargest.add(3); // return 4  
kthLargest.add(5); // return 5  
kthLargest.add(10); // return 5  
kthLargest.add(9); // return 8  
kthLargest.add(4); // return 8

**Example 2:**

**Input:**  
["KthLargest", "add", "add", "add", "add"]  
[[4, [7, 7, 7, 7, 8, 3]], [2], [10], [9], [9]]

**Output:** [null, 7, 7, 7, 8]

**Explanation:**

KthLargest kthLargest = new KthLargest(4, [7, 7, 7, 7, 8, 3]);  
kthLargest.add(2); // return 7  
kthLargest.add(10); // return 7  
kthLargest.add(9); // return 7  
kthLargest.add(9); // return 8

**Constraints:**

- `0 <= nums.length <= 104`
- `1 <= k <= nums.length + 1`
- `-104 <= nums[i] <= 104`
- `-104 <= val <= 104`
- At most `104` calls will be made to `add`.

## Solution
### Approach 1
**Method Name**: Min Heap

**Principle:**
Use a min-heap of size k to maintain the k largest elements in the stream, with the top of the heap being the kth largest element.

**Steps:**
1. During initialization, add all elements to the heap and keep popping until the heap size is k
2. When adding a new element:
   - If heap size is less than k, add it directly
   - Otherwise, if the new element is larger than the heap top, add it and pop the heap top
3. The heap top is always the kth largest element in the current stream

```embed-go
PATH: "vault://leetcode/0701-0800/0703_kth_largest_element_in_a_stream/solution.go"
TITLE: "leetcode 703. Kth Largest Element in a Stream"
```
