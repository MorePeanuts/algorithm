---
link: https://leetcode.cn/problems/search-in-rotated-sorted-array/
tags:
  - 中等
  - 数组
  - 二分查找
---
## 题目描述
整数数组 `nums` 按升序排列，数组中的值 **互不相同** 。

在传递给函数之前，`nums` 在预先未知的某个下标 `k`（`0 <= k < nums.length`）上进行了 **向左旋转**，使数组变为 `[nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]]`（下标 **从 0 开始** 计数）。例如， `[0,1,2,4,5,6,7]` 下标 `3` 上向左旋转后可能变为 `[4,5,6,7,0,1,2]` 。

给你 **旋转后** 的数组 `nums` 和一个整数 `target` ，如果 `nums` 中存在这个目标值 `target` ，则返回它的下标，否则返回 `-1` 。

你必须设计一个时间复杂度为 `O(log n)` 的算法解决此问题。

**示例 1：**

```
输入：nums = [4,5,6,7,0,1,2], target = 0
输出：4
```

**示例 2：**

```
输入：nums = [4,5,6,7,0,1,2], target = 3
输出：-1
```

**示例 3：**

```
输入：nums = [1], target = 0
输出：-1
```

**提示：**

- `1 <= nums.length <= 5000`
- `-104 <= nums[i] <= 104`
- `nums` 中的每个值都 **独一无二**
- 题目数据保证 `nums` 在预先未知的某个下标上进行了旋转
- `-104 <= target <= 104`

## 题目解析
### 解法1
**方法名**：先找旋转点再二分查找

**原理：**
旋转后的数组可以看作两个有序子数组，先通过二分查找找到旋转点（最小值的位置），然后根据目标值与第一个元素的比较，确定在哪个子数组中进行二分查找。

**步骤：**
1. 检查数组是否未旋转（nums[0] < nums[len(nums)-1]），若是则直接对整个数组二分查找
2. 二分查找找到旋转点 l（最小值的索引）
3. 比较 target 与 nums[0]：
   - 若 target < nums[0]，在右半部分 nums[l:] 中查找
   - 若 target == nums[0]，直接返回 0
   - 若 target > nums[0]，在左半部分 nums[:l] 中查找

**示例：**
- 输入 nums = [4,5,6,7,0,1,2], target = 0
- 找到旋转点 l = 4（值为 0）
- target < nums[0] (0 < 4)，在 nums[4:] = [0,1,2] 中查找，找到索引 0
- 返回 0 + 4 = 4

​```embed-go
PATH: "vault://leetcode/0001-0100/0033_search_in_rotated_sorted_array/solution.go"
TITLE: "leetcode 33. Search in Rotated Sorted Array"
​```

### 解法2
**方法名**：一次二分查找

**原理：**
在旋转排序数组中，任意选取一个中间点，至少有一侧是有序的。利用这一性质，通过判断中间点在哪一侧有序，以及目标值是否在该有序区间内，来调整搜索范围。

**步骤：**
1. 初始化左右指针 l = 0, r = len(nums)
2. 当 l < r 时：
   - 计算中间点 m = l + (r-l)/2
   - 若 nums[m] == target，直接返回 m
   - 若 nums[m] >= nums[l]，说明左半部分 [l, m) 有序：
     - 若 target 在 [nums[l], nums[m]) 范围内，在左半部分查找
     - 否则在右半部分查找
   - 否则，说明右半部分 [m, r) 有序：
     - 若 target < nums[m] 或 target > nums[r-1]，在左半部分查找
     - 否则在右半部分查找
3. 未找到返回 -1

**示例：**
- 输入 nums = [4,5,6,7,0,1,2], target = 0
- l=0, r=7, m=3, nums[m]=7 >= nums[0]=4，左半有序
- target=0 不在 [4,7) 范围内，l=4
- l=4, r=7, m=5, nums[m]=1 < nums[4]=0？不，nums[5]=1 >= nums[4]=0，左半有序
- target=0 < nums[4]=0？不，target=0 == nums[4]，但 nums[m]=1 != 0
- target=0 < nums[m]=1，r=5
- l=4, r=5, m=4, nums[m]=0 == target，返回 4

​```embed-go
PATH: "vault://leetcode/0001-0100/0033_search_in_rotated_sorted_array/solution2.go"
TITLE: "leetcode 33. Search in Rotated Sorted Array"
​```

