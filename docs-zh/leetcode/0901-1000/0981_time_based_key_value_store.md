---
link: https://leetcode.cn/problems/time-based-key-value-store/
tags:
  - 中等
  - 设计
  - 哈希表
  - 字符串
  - 二分查找
---
## 题目描述
设计一个基于时间的键值数据结构，该结构可以在不同时间戳存储对应同一个键的多个值，并针对特定时间戳检索键对应的值。

实现 `TimeMap` 类：

- `TimeMap()` 初始化数据结构对象
- `void set(String key, String value, int timestamp)` 存储给定时间戳 `timestamp` 时的键 `key` 和值 `value`。
- `String get(String key, int timestamp)` 返回一个值，该值在之前调用了 `set`，其中 `timestamp_prev <= timestamp` 。如果有多个这样的值，它将返回与最大  `timestamp_prev` 关联的值。如果没有值，则返回空字符串（`""`）。

**示例 1：**

```
输入：
["TimeMap", "set", "get", "get", "set", "get", "get"]
[[], ["foo", "bar", 1], ["foo", 1], ["foo", 3], ["foo", "bar2", 4], ["foo", 4], ["foo", 5]]
输出：
[null, null, "bar", "bar", null, "bar2", "bar2"]

解释：
TimeMap timeMap = new TimeMap();
timeMap.set("foo", "bar", 1);  // 存储键 "foo" 和值 "bar" ，时间戳 timestamp = 1   
timeMap.get("foo", 1);         // 返回 "bar"
timeMap.get("foo", 3);         // 返回 "bar", 因为在时间戳 3 和时间戳 2 处没有对应 "foo" 的值，所以唯一的值位于时间戳 1 处（即 "bar"） 。
timeMap.set("foo", "bar2", 4); // 存储键 "foo" 和值 "bar2" ，时间戳 timestamp = 4  
timeMap.get("foo", 4);         // 返回 "bar2"
timeMap.get("foo", 5);         // 返回 "bar2"
```

**提示：**

- `1 <= key.length, value.length <= 100`
- `key` 和 `value` 由小写英文字母和数字组成
- `1 <= timestamp <= 107`
- `set` 操作中的时间戳 `timestamp` 都是严格递增的
- 最多调用 `set` 和 `get` 操作 `2 * 105` 次

## 题目解析
### 解法1
**方法名**：哈希表 + 二分查找

**原理：**
使用两个哈希表分别存储每个键对应的时间戳数组和值数组，利用 set 操作时间戳严格递增的特性，通过二分查找快速定位到最接近且不大于目标时间戳的位置。

**步骤：**
1. 初始化两个哈希表，`timestamps` 存储键到时间戳数组的映射，`values` 存储键到值数组的映射
2. `set` 操作：直接将时间戳和值追加到对应键的数组末尾（利用时间戳严格递增的特性）
3. `get` 操作：使用二分查找在时间戳数组中找到最接近且不大于目标时间戳的位置，返回对应位置的值

**示例：**
- 输入：set("foo", "bar", 1) → timestamps["foo"] = [1], values["foo"] = ["bar"]
- 输入：get("foo", 3) → 二分查找找到位置 0（时间戳 1 ≤ 3）→ 输出 "bar"
- 输入：set("foo", "bar2", 4) → timestamps["foo"] = [1, 4], values["foo"] = ["bar", "bar2"]
- 输入：get("foo", 4) → 二分查找找到位置 1（时间戳 4 ≤ 4）→ 输出 "bar2"

```embed-go
PATH: "vault://leetcode/0901-1000/0981_time_based_key_value_store/solution.go"
TITLE: "leetcode 981. Time Based Key-Value Store"
```
