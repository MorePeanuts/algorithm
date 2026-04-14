---
link: https://leetcode.cn/problems/search-a-2d-matrix/
tags:
  - 中等
  - 数组
  - 二分查找
  - 矩阵
---
## 题目描述
给你一个满足下述两条属性的 `m x n` 整数矩阵：

- 每行中的整数从左到右按非严格递增顺序排列。
- 每行的第一个整数大于前一行的最后一个整数。

给你一个整数 `target` ，如果 `target` 在矩阵中，返回 `true` ；否则，返回 `false` 。

**示例 1：**

![](https://assets.leetcode.com/uploads/2020/10/05/mat.jpg)

```
输入：matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 3
输出：true
```

**示例 2：**

![](https://assets.leetcode.cn/aliyun-lc-upload/uploads/2020/11/25/mat2.jpg)

```
输入：matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 13
输出：false
```

**提示：**

- `m == matrix.length`
- `n == matrix[i].length`
- `1 <= m, n <= 100`
- `-104 <= matrix[i][j], target <= 104`

## 题目解析
### 解法1
**方法名**：两次二分查找

**原理：**
利用矩阵的有序性，先通过二分查找确定 target 可能所在的行，再在该行中进行二分查找确定 target 是否存在。

**步骤：**
1. 提取矩阵每行的第一个元素，构成一个有序数组 `firstNums`
2. 对 `firstNums` 进行二分查找，确定 target 可能在哪一行
3. 若找到目标值，直接返回 true；否则在确定的行中继续二分查找
4. 返回第二次查找的结果

**示例：**
输入：matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 13
- firstNums = [1, 10, 23]
- 第一次二分查找：13 不在 firstNums 中，确定在第 1 行（值为 [10,11,16,20]）
- 第二次二分查找：13 不在该行中，返回 false

> **注意**：关键在于 `binarySearchPos` 函数的设计，它不仅能查找目标值是否存在，还能在未找到时返回目标值应该插入的位置的前一个索引。

```embed-go
PATH: "vault://leetcode/0001-0100/0074_search_a_2d_matrix/solution.go"
TITLE: "leetcode 74. Search a 2D Matrix"
```

---

### 关于 `binarySearchPos` 函数的详细说明

`binarySearchPos` 函数是一个精心设计的二分查找变体，其核心特点是：**无论是否找到 target，都能返回有意义的位置信息**。

#### 为什么未找到时返回的 pos 满足 `nums[pos] < target < nums[pos+1]`？

让我们通过分析循环不变量来理解这个结论：

**1. 初始化与循环条件**
```go
l, r := 0, len(nums)  // 注意：r 初始化为 len(nums)，不是 len(nums)-1
for l < r {
```
- 搜索区间是左闭右开区间 `[l, r)`
- `l` 是左边界，`r` 是右边界（不包含）

**2. 循环中的不变量**
在每一次循环迭代中，始终保持以下性质：
- 对于所有 `i < l`，有 `nums[i] < target`
- 对于所有 `i >= r`，有 `nums[i] > target`

**3. 循环结束时的状态**
当 `l == r` 时，循环终止，此时：
- `l`（等于 `r`）就是 **target 应该插入的位置**，使得插入后数组仍然有序
- 所有在 `l` 左边的元素都 `< target`
- 所有在 `l` 右边的元素都 `> target`

**4. 返回值的计算**
```go
return l - 1, false
```
- 返回 `l - 1` 即小于 target 的最大元素的索引
- 此时有：`nums[l-1] < target < nums[l]`（假设 `l-1` 和 `l` 都在数组范围内）

#### 边界情况处理

| 情况 | l 的值 | 返回值 | 含义 |
|------|--------|--------|------|
| target < 所有元素 | 0 | -1 | target 比所有元素都小 |
| target > 所有元素 | len(nums) | len(nums)-1 | target 比所有元素都大 |
| target 在中间 | 插入位置 | 插入位置-1 | nums[pos] < target < nums[pos+1] |

#### 实例验证

以 `nums = [1, 10, 23]`，`target = 13` 为例：

| 迭代 | l | r | m | nums[m] | 比较 | 操作 |
|------|---|---|---|---------|------|------|
| 初始 | 0 | 3 | - | - | - | - |
| 1 | 0 | 3 | 1 | 10 | 13 > 10 | l = m + 1 = 2 |
| 2 | 2 | 3 | 2 | 23 | 13 < 23 | r = m = 2 |
| 结束 | 2 | 2 | - | - | - | 返回 l-1 = 1 |

结果：`pos = 1`，满足 `nums[1] = 10 < 13 < 23 = nums[2]`
