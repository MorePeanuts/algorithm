---
link: https://leetcode.com/problems/find-median-from-data-stream/
tags:
  - Hard
  - Design
  - Two_Pointers
  - Data_Stream
  - Sorting
  - Heap_(Priority_Queue)
---
## Description
The **median** is the middle value in an ordered integer list. If the size of the list is even, there is no middle value, and the median is the mean of the two middle values.

- For example, for `arr = [2,3,4]`, the median is `3`.
- For example, for `arr = [2,3]`, the median is `(2 + 3) / 2 = 2.5`.

Implement the MedianFinder class:

- `MedianFinder()` initializes the `MedianFinder` object.
- `void addNum(int num)` adds the integer `num` from the data stream to the data structure.
- `double findMedian()` returns the median of all elements so far. Answers within `10-5` of the actual answer will be accepted.

**Example 1:**

```
Input
["MedianFinder", "addNum", "addNum", "findMedian", "addNum", "findMedian"]
[[], [1], [2], [], [3], []]
Output
[null, null, null, 1.5, null, 2.0]

Explanation
MedianFinder medianFinder = new MedianFinder();
medianFinder.addNum(1);    // arr = [1]
medianFinder.addNum(2);    // arr = [1, 2]
medianFinder.findMedian(); // return 1.5 (i.e., (1 + 2) / 2)
medianFinder.addNum(3);    // arr[1, 2, 3]
medianFinder.findMedian(); // return 2.0
```

**Constraints:**

- `-105 <= num <= 105`
- There will be at least one element in the data structure before calling `findMedian`.
- At most `5 * 104` calls will be made to `addNum` and `findMedian`.

**Follow up:**

- If all integer numbers from the stream are in the range `[0, 100]`, how would you optimize your solution?
- If `99%` of all integer numbers from the stream are in the range `[0, 100]`, how would you optimize your solution?

## Solution
### Approach 1
**AddNum / FindMedian**: Dual-heap approach — max-heap for the smaller half, min-heap for the larger half

**Principle:**
Maintain two heaps that split the data stream into a smaller half and a larger half. The max-heap stores the smaller half, and the min-heap stores the larger half. Keep the min-heap's size equal to or one greater than the max-heap's size. The median is the min-heap's top (odd count) or the average of both tops (even count).

**Steps:**
1. Initialize an empty max-heap and min-heap
2. On addNum, decide insertion strategy based on heap sizes:
   - Equal sizes: if num < max-heap top, move max-heap top to min-heap and push num into max-heap; otherwise push num directly into min-heap
   - Min-heap has one more: if num ≤ min-heap top, push num into max-heap; otherwise move min-heap top to max-heap and push num into min-heap
3. On findMedian: if equal sizes, return average of both tops; otherwise return min-heap top

**Example:**
- addNum(1): both heaps empty, push into min-heap → minHeap=[1], maxHeap=[]
- addNum(2): min-heap has one more, 2 > minHeap[0]=1, move 1 to max-heap, push 2 into min-heap → minHeap=[2], maxHeap=[1]
- findMedian(): equal sizes, (1+2)/2 = 1.5
- addNum(3): equal sizes, 3 ≥ maxHeap[0]=1, push into min-heap → minHeap=[2,3], maxHeap=[1]
- findMedian(): min-heap has one more, return minHeap[0]=2.0

**Complexity Analysis:**
- addNum: O(log n) time, at most two heap operations per insertion
- findMedian: O(1) time, direct access to heap tops
- Space: O(n) to store all elements

> **Note**: The insertion logic maintains balance and ordering by "pop-then-push" across heaps, ensuring the max of the smaller half never exceeds the min of the larger half.

```embed-go
PATH: "vault://leetcode/0201-0300/0295_find_median_from_data_stream/solution.go"
TITLE: "leetcode 0295.Find Median from Data Stream"
```
