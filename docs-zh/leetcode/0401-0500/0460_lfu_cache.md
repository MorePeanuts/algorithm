---
link: https://leetcode.cn/problems/lfu-cache/
tags:
  - 困难
  - 设计
  - 哈希表
  - 链表
  - 双向链表
---
## 题目描述
请你为 [最不经常使用（LFU）](https://baike.baidu.com/item/%E7%BC%93%E5%AD%98%E7%AE%97%E6%B3%95)缓存算法设计并实现数据结构。

实现 `LFUCache` 类：

- `LFUCache(int capacity)` - 用数据结构的容量 `capacity` 初始化对象
- `int get(int key)` - 如果键 `key` 存在于缓存中，则获取键的值，否则返回 `-1` 。
- `void put(int key, int value)` - 如果键 `key` 已存在，则变更其值；如果键不存在，请插入键值对。当缓存达到其容量 `capacity` 时，则应该在插入新项之前，移除最不经常使用的项。在此问题中，当存在平局（即两个或更多个键具有相同使用频率）时，应该去除 **最久未使用** 的键。

为了确定最不常使用的键，可以为缓存中的每个键维护一个 **使用计数器** 。使用计数最小的键是最久未使用的键。

当一个键首次插入到缓存中时，它的使用计数器被设置为 `1` (由于 put 操作)。对缓存中的键执行 `get` 或 `put` 操作，使用计数器的值将会递增。

函数 `get` 和 `put` 必须以 `O(1)` 的平均时间复杂度运行。

**示例：**

```
输入：
["LFUCache", "put", "put", "get", "put", "get", "get", "put", "get", "get", "get"]
[[2], [1, 1], [2, 2], [1], [3, 3], [2], [3], [4, 4], [1], [3], [4]]
输出：
[null, null, null, 1, null, -1, 3, null, -1, 3, 4]

解释：
// cnt(x) = 键 x 的使用计数
// cache=[] 将显示最后一次使用的顺序（最左边的元素是最近的）
LFUCache lfu = new LFUCache(2);
lfu.put(1, 1);   // cache=[1,_], cnt(1)=1
lfu.put(2, 2);   // cache=[2,1], cnt(2)=1, cnt(1)=1
lfu.get(1);      // 返回 1
                 // cache=[1,2], cnt(2)=1, cnt(1)=2
lfu.put(3, 3);   // 去除键 2 ，因为 cnt(2)=1 ，使用计数最小
                 // cache=[3,1], cnt(3)=1, cnt(1)=2
lfu.get(2);      // 返回 -1（未找到）
lfu.get(3);      // 返回 3
                 // cache=[3,1], cnt(3)=2, cnt(1)=2
lfu.put(4, 4);   // 去除键 1 ，1 和 3 的 cnt 相同，但 1 最久未使用
                 // cache=[4,3], cnt(4)=1, cnt(3)=2
lfu.get(1);      // 返回 -1（未找到）
lfu.get(3);      // 返回 3
                 // cache=[3,4], cnt(4)=1, cnt(3)=3
lfu.get(4);      // 返回 4
                 // cache=[3,4], cnt(4)=2, cnt(3)=3
```

**提示：**

- `1 <= capacity <= 104`
- `0 <= key <= 105`
- `0 <= value <= 109`
- 最多调用 `2 * 105` 次 `get` 和 `put` 方法

## 题目解析
### 解法1
**方法名**：双哈希表 + 双向链表

**原理：**
使用两个哈希表分别存储键到节点的映射，以及频率到双向链表的映射，实现 O(1) 时间复杂度的 get 和 put 操作。

**步骤：**
1. 定义 `node` 结构体存储键、值、频率及双向指针
2. 定义 `LinkedList` 结构体实现双向链表，支持 Push、Pop 和 Remove 操作
3. `LFUCache` 包含：
   - `hashtable`：键到节点的映射，O(1) 查找
   - `queues`：频率到链表的映射，同一频率的节点按使用顺序排列
   - `capacity`：缓存容量
4. `Get` 操作：查找节点，增加频率，移动到对应频率链表
5. `Put` 操作：
   - 键存在：更新值，增加频率，移动节点
   - 键不存在：创建新节点，容量满时移除最不常用节点，插入新节点
6. `visitNode`：处理节点访问，从原频率链表移除，频率加1，加入新频率链表
7. `popNode`：找到最小频率的链表，移除最久未使用的节点

**示例：**
以示例输入为例，容量为2：
- put(1,1) → freq=1, queues[0]=[1]
- put(2,2) → freq=1, queues[0]=[1,2]
- get(1) → 返回1, freq=2, queues[0]=[2], queues[1]=[1]
- put(3,3) → 移除2（freq最小）, queues[0]=[3], queues[1]=[1]

> **注意**：
> - 同一频率链表中，头部是最久未使用的节点，尾部是最近使用的
> - 当多个节点频率相同时，通过链表顺序区分最近使用
> - 需要维护最小频率，但本实现通过遍历 queues 找到第一个非空链表来简化实现

```embed-go
PATH: "vault://leetcode/0401-0500/0460_lfu_cache/solution.go"
TITLE: "leetcode 460.LFU Cache"
```
