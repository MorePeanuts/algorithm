---
link: https://leetcode.cn/problems/find-median-from-data-stream/
tags:
  - 困难
  - 设计
  - 双指针
  - 数据流
  - 排序
  - 堆（优先队列）
---
## 题目描述
**中位数**是有序整数列表中的中间值。如果列表的大小是偶数，则没有中间值，中位数是两个中间值的平均值。

- 例如 `arr = [2,3,4]` 的中位数是 `3` 。
- 例如 `arr = [2,3]` 的中位数是 `(2 + 3) / 2 = 2.5` 。

实现 MedianFinder 类:

- `MedianFinder()` 初始化 `MedianFinder` 对象。
- `void addNum(int num)` 将数据流中的整数 `num` 添加到数据结构中。
- `double findMedian()` 返回到目前为止所有元素的中位数。与实际答案相差 `10-5` 以内的答案将被接受。

**示例 1：**

```
输入
["MedianFinder", "addNum", "addNum", "findMedian", "addNum", "findMedian"]
[[], [1], [2], [], [3], []]
输出
[null, null, null, 1.5, null, 2.0]

解释
MedianFinder medianFinder = new MedianFinder();
medianFinder.addNum(1);    // arr = [1]
medianFinder.addNum(2);    // arr = [1, 2]
medianFinder.findMedian(); // 返回 1.5 ((1 + 2) / 2)
medianFinder.addNum(3);    // arr[1, 2, 3]
medianFinder.findMedian(); // return 2.0
```

**提示:**

- `-105 <= num <= 105`
- 在调用 `findMedian` 之前，数据结构中至少有一个元素
- 最多 `5 * 104` 次调用 `addNum` 和 `findMedian`

## 题目解析
### 解法1
**AddNum / FindMedian**：双堆法——大顶堆存较小半部分，小顶堆存较大半部分

**原理：**
维护两个堆，将数据流分成较小的一半和较大的一半。大顶堆（maxHeap）存储较小的一半，小顶堆（minHeap）存储较大的一半。始终保持 minHeap 的元素数等于 maxHeap 或多一个，中位数即为 minHeap 堆顶（奇数个元素）或两堆顶的平均值（偶数个元素）。

**步骤：**
1. 初始化时创建空的大顶堆和小顶堆
2. addNum 时根据两堆大小关系决定插入策略：
   - 两堆元素数相等：若 num 小于 maxHeap 堆顶，则将 maxHeap 堆顶移至 minHeap，再将 num 放入 maxHeap；否则直接放入 minHeap
   - minHeap 多一个元素：若 num 不大于 minHeap 堆顶，直接放入 maxHeap；否则将 minHeap 堆顶移至 maxHeap，再将 num 放入 minHeap
3. findMedian 时：若两堆元素数相等取两堆顶的平均值，否则取 minHeap 堆顶

**示例：**
- addNum(1)：两堆为空，直接放入 minHeap → minHeap=[1], maxHeap=[]
- addNum(2)：minHeap 多一个元素，2 > minHeap[0]=1，将 1 移至 maxHeap，2 放入 minHeap → minHeap=[2], maxHeap=[1]
- findMedian()：两堆元素数相等，(1+2)/2 = 1.5
- addNum(3)：两堆元素数相等，3 ≥ maxHeap[0]=1，直接放入 minHeap → minHeap=[2,3], maxHeap=[1]
- findMedian()：minHeap 多一个元素，返回 minHeap[0]=2.0

**复杂度分析：**
- addNum：时间 O(log n)，每次最多执行两次堆操作
- findMedian：时间 O(1)，直接访问堆顶
- 空间 O(n)，存储所有元素

> **注意**：插入时通过"先弹出再压入"的方式维持两堆的平衡和有序性，确保较小半部分的最大值始终不超过较大半部分的最小值。

```embed-go
PATH: "vault://leetcode/0201-0300/0295_find_median_from_data_stream/solution.go"
TITLE: "leetcode 0295.Find Median from Data Stream"
```
