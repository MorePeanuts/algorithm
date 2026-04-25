---
link: https://leetcode.cn/problems/sort-colors/
tags:
  - 中等
  - 数组
  - 双指针
  - 排序
---
## 题目描述
给定一个包含红色、白色和蓝色、共 `n`个元素的数组 `nums` ，**[原地](https://baike.baidu.com/item/%E5%8E%9F%E5%9C%B0%E7%AE%97%E6%B3%95)**对它们进行排序，使得相同颜色的元素相邻，并按照红色、白色、蓝色顺序排列。

我们使用整数 `0`、 `1` 和 `2` 分别表示红色、白色和蓝色。

必须在不使用库内置的 sort 函数的情况下解决这个问题。

**示例 1：**

```
输入：nums = [2,0,2,1,1,0]
输出：[0,0,1,1,2,2]
```

**示例 2：**

```
输入：nums = [2,0,1]
输出：[0,1,2]
```

**提示：**

- `n == nums.length`
- `1 <= n <= 300`
- `nums[i]` 为 `0`、`1` 或 `2`

**进阶：**

- 你能想出一个仅使用常数空间的一趟扫描算法吗？

## 题目解析
### 解法1
**方法名**：三指针一趟扫描

**原理：**
使用三个指针将数组分为三个区域：0的区域[0, left)、1的区域[left, i)、待处理区域[i, right]、2的区域(right, end)。通过一次遍历将元素交换到对应区域。

**步骤：**
1. 初始化 left=0（0的右边界），right=len(nums)-1（2的左边界），pivot=1（中间值）
2. 遍历数组，i 从 left 开始：
   - 若 nums[i] < pivot（即0）：交换到 left 位置，left 和 i 都右移
   - 若 nums[i] == pivot（即1）：直接 i 右移
   - 若 nums[i] > pivot（即2）：交换到 right 位置，right 左移（i 不移动，因为交换来的元素还未检查）
3. 当 i > right 时结束

**示例：**
- 输入：[2,0,2,1,1,0]
- 处理过程：
  - i=0, nums[i]=2 → 与 right=5 交换 → [0,0,2,1,1,2], right=4
  - i=0, nums[i]=0 → 与 left=0 交换 → [0,0,2,1,1,2], left=1, i=1
  - i=1, nums[i]=0 → 与 left=1 交换 → [0,0,2,1,1,2], left=2, i=2
  - i=2, nums[i]=2 → 与 right=4 交换 → [0,0,1,1,2,2], right=3
  - i=2, nums[i]=1 → i=3
  - i=3, nums[i]=1 → i=4 > right=3，结束
- 输出：[0,0,1,1,2,2]

```embed-go
PATH: "vault://leetcode/0001-0100/0075_sort_colors/solution.go"
TITLE: "leetcode 75. Sort Colors"
```
