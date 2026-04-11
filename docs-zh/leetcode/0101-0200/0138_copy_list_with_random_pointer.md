---
link: https://leetcode.cn/problems/copy-list-with-random-pointer/
tags:
  - 中等
  - 哈希表
  - 链表
---
## 题目描述
给你一个长度为 `n` 的链表，每个节点包含一个额外增加的随机指针 `random` ，该指针可以指向链表中的任何节点或空节点。

构造这个链表的 **[深拷贝](https://baike.baidu.com/item/深拷贝/22785317?fr=aladdin)**。 深拷贝应该正好由 `n` 个 **全新** 节点组成，其中每个新节点的值都设为其对应的原节点的值。新节点的 `next` 指针和 `random` 指针也都应指向复制链表中的新节点，并使原链表和复制链表中的这些指针能够表示相同的链表状态。**复制链表中的指针都不应指向原链表中的节点** 。

例如，如果原链表中有 `X` 和 `Y` 两个节点，其中 `X.random --> Y` 。那么在复制链表中对应的两个节点 `x` 和 `y` ，同样有 `x.random --> y` 。

返回复制链表的头节点。

用一个由 `n` 个节点组成的链表来表示输入/输出中的链表。每个节点用一个 `[val, random_index]` 表示：

- `val`：一个表示 `Node.val` 的整数。
- `random_index`：随机指针指向的节点索引（范围从 `0` 到 `n-1`）；如果不指向任何节点，则为  `null` 。

你的代码 **只** 接受原链表的头节点 `head` 作为传入参数。

**示例 1：**

![](https://assets.leetcode.cn/aliyun-lc-upload/uploads/2020/01/09/e1.png)

```
输入：head = [[7,null],[13,0],[11,4],[10,2],[1,0]]
输出：[[7,null],[13,0],[11,4],[10,2],[1,0]]
```

**示例 2：**

![](https://assets.leetcode.cn/aliyun-lc-upload/uploads/2020/01/09/e2.png)

```
输入：head = [[1,1],[2,1]]
输出：[[1,1],[2,1]]
```

**示例 3：**

**![](https://assets.leetcode.cn/aliyun-lc-upload/uploads/2020/01/09/e3.png)**

```
输入：head = [[3,null],[3,0],[3,null]]
输出：[[3,null],[3,0],[3,null]]
```

**提示：**

- `0 <= n <= 1000`
- `-104 <= Node.val <= 104`
- `Node.random` 为 `null` 或指向链表中的节点。

## 题目解析
### 解法1
**方法名**：哈希表映射节点索引

**原理：**
使用哈希表记录原节点与其索引的对应关系，同时记录每个节点的 random 指针指向的索引，然后根据这些索引信息构建新链表。

**步骤：**
1. 第一次遍历原链表，用哈希表记录每个节点的索引
2. 第二次遍历哈希表，为每个节点记录其 random 指针指向的索引
3. 根据哈希表中的索引信息创建新节点数组
4. 设置新节点的 next 和 random 指针

**示例：**
输入链表: `[[7,null],[13,0],[11,4],[10,2],[1,0]]`
- 哈希表记录: 节点0→[0], 节点1→[1], 节点2→[2], 节点3→[3], 节点4→[4]
- 添加 random 索引: 节点1→[1,0], 节点2→[2,4], 节点3→[3,2], 节点4→[4,0]
- 创建新节点数组并设置指针

```embed-go
PATH: "vault://leetcode/0101-0200/0138_copy_list_with_random_pointer/solution.go"
TITLE: "leetcode 138. Copy List with Random Pointer"
```

### 解法2
**方法名**：递归+哈希表

**原理：**
使用递归深度优先遍历链表，同时用哈希表记录已创建的新节点，避免重复创建。递归地复制 next 和 random 指针。

**步骤：**
1. 定义哈希表 old2new 记录原节点到新节点的映射
2. 定义递归函数 deepCopy：
   - 如果节点为空，返回 nil
   - 如果节点已在哈希表中，返回已创建的新节点
   - 否则创建新节点并存入哈希表
   - 递归复制 next 和 random 指针
3. 调用递归函数返回新链表头

**示例：**
输入链表: `[[1,1],[2,1]]`
- 复制节点1，创建新节点1，递归复制 next 和 random
- 复制节点2，创建新节点2，递归复制 next 和 random
- 递归返回，设置指针关系

```embed-go
PATH: "vault://leetcode/0101-0200/0138_copy_list_with_random_pointer/solution2.go"
TITLE: "leetcode 138. Copy List with Random Pointer"
```

### 解法3
**方法名**：原地修改+拆分

**原理：**
将每个新节点直接插入到原节点后面，形成旧→新→旧→新...的结构，然后利用这个结构设置 random 指针，最后将新旧链表拆分。不需要额外的哈希表空间。

**步骤：**
1. 第一次遍历：在每个原节点后面插入一个新节点
2. 第二次遍历：设置新节点的 random 指针（原节点 random 的 next）
3. 第三次遍历：拆分新旧链表，恢复原链表，得到新链表

**示例：**
输入链表: `[[3,null],[3,0],[3,null]]`
- 插入后: 3→3'→3→3'→3→3'
- 设置 random: 第二个3'的 random 指向第一个3'
- 拆分后得到新链表: 3'→3'→3'

> **注意**：拆分时需要同时恢复原链表的结构，避免原链表被破坏。

```embed-go
PATH: "vault://leetcode/0101-0200/0138_copy_list_with_random_pointer/solution3.go"
TITLE: "leetcode 138. Copy List with Random Pointer"
```
