---
tags:
  - 中等
  - 数组
  - 哈希表
  - 字符串
  - 排序
---
## 题目描述
给你一个字符串数组，请你将 **字母异位词** 组合在一起。可以按任意顺序返回结果列表。

---
示例 1:
- 输入: `strs = ["eat", "tea", "tan", "ate", "nat", "bat"]`
- 输出: `[["bat"],["nat","tan"],["ate","eat","tea"]]`
- 解释：
	- 在 strs 中没有字符串可以通过重新排列来形成 "bat"。
	- 字符串 "nat" 和 "tan" 是字母异位词，因为它们可以重新排列以形成彼此。
	- 字符串 "ate" ，"eat" 和 "tea" 是字母异位词，因为它们可以重新排列以形成彼此。

示例 2:
- 输入: `strs = [""]`
- 输出: `[[""]]`

示例 3:
- 输入: `strs = ["a"]`
- 输出: `[["a"]]`
 
---
提示：
- `1 <= strs.length <= 10^4`
- `0 <= strs[i].length <= 100`
- `strs[i]` 仅包含小写字母
## 题解
### 解法1
```embed-go
PATH: "vault://leetcode/0001-0100/0049_group_anagrams/solution1.go"
TITLE: "leetcode 49.字母异位词分组"
```
### 解法2
**算术基本定理**（Fundamental Theorem of Arithmetic）：**任何一个大于1的自然数都可以唯一分解成有限个质数的乘积**。

**原理：**
乘法满足交换律（`a * b = b * a`），这天然解决了“顺序无关”的问题。利用质数的特性，可以保证不同的字母组合乘积一定不同。

**步骤：**
1. 建立一个映射表，将 26 个字母映射到 26 个不同的**质数**上。
    - `a` -> 2, `b` -> 3, `c` -> 5, `d` -> 7, ...
2. 遍历字符串，将每个字符对应的质数相乘。
3. 乘积结果就是唯一的 Key。

**示例：**
- `a` = 2, `e` = 11, `t` = 71 (假设)
- `"eat"` 的 Key = $11 \times 2 \times 71 = 1562$
- `"ate"` 的 Key = $2 \times 71 \times 11 = 1562$ 
- `"aba"` 的 Key = $2 \times 3 \times 2 = 12$

> **整型溢出风险**：题目提示 `strs[i].length <= 100`。如果是 100 个 'z'，那么乘积就是 `(第26个质数)^100`，会导致整数溢出，因此需要使用大数类型。

```embed-go
PATH: "vault://leetcode/0001-0100/0049_group_anagrams/solution2.go"
TITLE: "leetcode 49.字母异位词分组"
```
### 解法3
```embed-go
PATH: "vault://leetcode/0001-0100/0049_group_anagrams/solution3.go"
TITLE: "leetcode 49.字母异位词分组"
```
