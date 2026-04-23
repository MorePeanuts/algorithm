---
link: https://leetcode.cn/problems/kth-largest-element-in-a-stream/
tags:
  - 简单
  - 树
  - 设计
  - 二叉搜索树
  - 二叉树
  - 数据流
  - 堆（优先队列）
---
## 题目描述
设计一个找到数据流中第 `k` 大元素的类（class）。注意是排序后的第 `k` 大元素，不是第 `k` 个不同的元素。

请实现 `KthLargest` 类：

- `KthLargest(int k, int[] nums)` 使用整数 `k` 和整数流 `nums` 初始化对象。
- `int add(int val)` 将 `val` 插入数据流 `nums` 后，返回当前数据流中第 `k` 大的元素。

**示例 1：**

**输入：**  
["KthLargest", "add", "add", "add", "add", "add"]  
[[3, [4, 5, 8, 2]], [3], [5], [10], [9], [4]]

**输出：**[null, 4, 5, 5, 8, 8]

**解释：**

KthLargest kthLargest = new KthLargest(3, [4, 5, 8, 2]);  
kthLargest.add(3); // 返回 4  
kthLargest.add(5); // 返回 5  
kthLargest.add(10); // 返回 5  
kthLargest.add(9); // 返回 8  
kthLargest.add(4); // 返回 8

**示例 2：**

**输入：**  
["KthLargest", "add", "add", "add", "add"]  
[[4, [7, 7, 7, 7, 8, 3]], [2], [10], [9], [9]]

**输出：**[null, 7, 7, 7, 8]

**解释：**

KthLargest kthLargest = new KthLargest(4, [7, 7, 7, 7, 8, 3]);  
kthLargest.add(2); // 返回 7  
kthLargest.add(10); // 返回 7  
kthLargest.add(9); // 返回 7  
kthLargest.add(9); // 返回 8

**提示：**

- `0 <= nums.length <= 104`
- `1 <= k <= nums.length + 1`
- `-104 <= nums[i] <= 104`
- `-104 <= val <= 104`
- 最多调用 `add` 方法 `104` 次

## 题目解析
### 解法1
**方法名**：最小堆（Min Heap）

**原理：**
使用一个大小为 k 的最小堆来维护数据流中前 k 大的元素，堆顶元素即为第 k 大的元素。

**步骤：**
1. 初始化时，将所有元素加入堆，然后不断弹出堆顶元素直到堆的大小为 k
2. 添加新元素时：
   - 如果堆的大小小于 k，直接加入堆
   - 否则，如果新元素大于堆顶元素，则加入堆并弹出堆顶元素
3. 堆顶元素始终为当前数据流中的第 k 大元素

```embed-go
PATH: "vault://leetcode/0701-0800/0703_kth_largest_element_in_a_stream/solution.go"
TITLE: "leetcode 703. 数据流中的第 K 大元素"
```
