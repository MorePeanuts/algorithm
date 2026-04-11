---
link: https://leetcode.cn/problems/lru-cache/
tags:
  - 中等
  - 设计
  - 哈希表
  - 链表
  - 双向链表
---
## 题目描述
请你设计并实现一个满足  [LRU (最近最少使用) 缓存](https://baike.baidu.com/item/LRU) 约束的数据结构。

实现 `LRUCache` 类：

- `LRUCache(int capacity)` 以 **正整数** 作为容量 `capacity` 初始化 LRU 缓存
- `int get(int key)` 如果关键字 `key` 存在于缓存中，则返回关键字的值，否则返回 `-1` 。
- `void put(int key, int value)` 如果关键字 `key` 已经存在，则变更其数据值 `value` ；如果不存在，则向缓存中插入该组 `key-value` 。如果插入操作导致关键字数量超过 `capacity` ，则应该 **逐出** 最久未使用的关键字。

函数 `get` 和 `put` 必须以 `O(1)` 的平均时间复杂度运行。

**示例：**

```
输入
["LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"]
[[2], [1, 1], [2, 2], [1], [3, 3], [2], [4, 4], [1], [3], [4]]
输出
[null, null, null, 1, null, -1, null, -1, 3, 4]

解释
LRUCache lRUCache = new LRUCache(2);
lRUCache.put(1, 1); // 缓存是 {1=1}
lRUCache.put(2, 2); // 缓存是 {1=1, 2=2}
lRUCache.get(1);    // 返回 1
lRUCache.put(3, 3); // 该操作会使得关键字 2 作废，缓存是 {1=1, 3=3}
lRUCache.get(2);    // 返回 -1 (未找到)
lRUCache.put(4, 4); // 该操作会使得关键字 1 作废，缓存是 {4=4, 3=3}
lRUCache.get(1);    // 返回 -1 (未找到)
lRUCache.get(3);    // 返回 3
lRUCache.get(4);    // 返回 4
```

**提示：**

- `1 <= capacity <= 3000`
- `0 <= key <= 10000`
- `0 <= value <= 105`
- 最多调用 `2 * 105` 次 `get` 和 `put`

## 题目解析
### 解法1
**方法名**：哈希表 + 双向链表

**原理：**
使用哈希表实现 O(1) 时间的查找，使用双向链表维护访问顺序，最近使用的节点放在链表尾部，最久未使用的节点在链表头部。

**步骤：**
1. **Get 操作**：
   - 检查哈希表中是否存在 key，不存在则返回 -1
   - 存在则将该节点移动到链表尾部（表示最近使用）
   - 返回节点的值

2. **Put 操作**：
   - 检查哈希表中是否存在 key
   - 存在则更新值并将节点移动到链表尾部
   - 不存在则创建新节点，添加到哈希表和链表尾部
   - 如果容量超出限制，删除链表头部的节点（最久未使用）并从哈希表中移除

3. **moveToTail 辅助操作**：
   - 如果节点已在尾部，无需操作
   - 如果节点在头部，更新头部指针
   - 断开节点与前后节点的连接
   - 将节点连接到链表尾部，更新尾部指针

**示例：**
以 capacity=2 为例：
- put(1, 1) → 链表: [1], 哈希表: {1:node1}
- put(2, 2) → 链表: [1, 2], 哈希表: {1:node1, 2:node2}
- get(1) → 访问节点1，移动到尾部 → 链表: [2, 1], 返回 1
- put(3, 3) → 超出容量，删除头部节点2，添加节点3到尾部 → 链表: [1, 3], 哈希表: {1:node1, 3:node3}
- get(2) → 不存在，返回 -1

```embed-go
PATH: "vault://leetcode/0101-0200/0146_lru_cache/solution.go"
TITLE: "leetcode 146. LRU Cache"
```
