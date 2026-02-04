---
link: https://leetcode.cn/problems/two-sum-ii-input-array-is-sorted/
tags:
  - 中等
  - 数组
  - 双指针
  - 二分查找
---
## 题目描述
给你一个下标从 **1** 开始的整数数组 `numbers` ，该数组已按**非递减顺序排列** ，请你从数组中找出满足相加之和等于目标数 `target` 的两个数。如果设这两个数分别是 `numbers[index1]` 和 `numbers[index2]` ，则 `1 <= index1 < index2 <= numbers.length` 。

以长度为 2 的整数数组 `[index1, index2]` 的形式返回这两个整数的下标 `index1`和`index2`。

你可以假设每个输入 **只对应唯一的答案** ，而且你 **不可以** 重复使用相同的元素。

你所设计的解决方案必须只使用常量级的额外空间。

**示例 1：**

```
输入：numbers = [2,7,11,15], target = 9
输出：[1,2]
解释：2 与 7 之和等于目标数 9 。因此 index1 = 1, index2 = 2 。返回 [1, 2] 。
```

**示例 2：**

```
输入：numbers = [2,3,4], target = 6
输出：[1,3]
解释：2 与 4 之和等于目标数 6 。因此 index1 = 1, index2 = 3 。返回 [1, 3] 。
```

**示例 3：**

```
输入：numbers = [-1,0], target = -1
输出：[1,2]
解释：-1 与 0 之和等于目标数 -1 。因此 index1 = 1, index2 = 2 。返回 [1, 2] 。
```

**提示：**

- `2 <= numbers.length <= 3 * 10^4`
- `-1000 <= numbers[i] <= 1000`
- `numbers` 按 **非递减顺序** 排列
- `-1000 <= target <= 1000`
- **仅存在一个有效答案**

## 题目解析

### 解法1

**双指针法**：利用数组有序的特性，通过左右指针相向移动来找到目标和。

**原理：**
由于数组已按非递减顺序排列，使用双指针从两端向中间逼近。当两数之和小于目标值时，左指针右移以增大和；当两数之和大于目标值时，右指针左移以减小和；相等时即找到答案。

**步骤：**
1. 初始化左指针 `left` 指向数组开头，右指针 `right` 指向数组末尾
2. 计算 `numbers[left] + numbers[right]` 的和
3. 若和小于 `target`，左指针右移 `left++`
4. 若和大于 `target`，右指针左移 `right--`
5. 若和等于 `target`，返回 `[left+1, right+1]`（题目要求下标从 1 开始）

**示例：**
- 输入：`numbers = [2,7,11,15], target = 9`
- 初始：`left=0, right=3`，和 `2+15=17 > 9`，右指针左移
- `left=0, right=2`，和 `2+11=13 > 9`，右指针左移
- `left=0, right=1`，和 `2+7=9 == 9`，找到答案
- 输出：`[1, 2]`

```embed-go
PATH: "vault://leetcode/0101-0200/0167_two_sum_ii_input_array_is_sorted/solution.go"
TITLE: "leetcode 167.两数之和 II - 输入有序数组"
```
