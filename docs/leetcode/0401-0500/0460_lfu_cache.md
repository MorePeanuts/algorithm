---
link: https://leetcode.com/problems/lfu-cache/
tags:
  - Hard
  - Design
  - Hash_Table
  - Linked_List
  - Doubly-Linked_List
---
## Description
Design and implement a data structure for a [Least Frequently Used (LFU)](https://en.wikipedia.org/wiki/Least_frequently_used) cache.

Implement the `LFUCache` class:

- `LFUCache(int capacity)` Initializes the object with the `capacity` of the data structure.
- `int get(int key)` Gets the value of the `key` if the `key` exists in the cache. Otherwise, returns `-1`.
- `void put(int key, int value)` Update the value of the `key` if present, or inserts the `key` if not already present. When the cache reaches its `capacity`, it should invalidate and remove the **least frequently used** key before inserting a new item. For this problem, when there is a **tie** (i.e., two or more keys with the same frequency), the **least recently used** `key` would be invalidated.

To determine the least frequently used key, a **use counter** is maintained for each key in the cache. The key with the smallest **use counter** is the least frequently used key.

When a key is first inserted into the cache, its **use counter** is set to `1` (due to the `put` operation). The **use counter** for a key in the cache is incremented either a `get` or `put` operation is called on it.

The functions `get` and `put` must each run in `O(1)` average time complexity.

**Example 1:**

```
Input
["LFUCache", "put", "put", "get", "put", "get", "get", "put", "get", "get", "get"]
[[2], [1, 1], [2, 2], [1], [3, 3], [2], [3], [4, 4], [1], [3], [4]]
Output
[null, null, null, 1, null, -1, 3, null, -1, 3, 4]

Explanation
// cnt(x) = the use counter for key x
// cache=[] will show the last used order for tiebreakers (leftmost element is  most recent)
LFUCache lfu = new LFUCache(2);
lfu.put(1, 1);   // cache=[1,_], cnt(1)=1
lfu.put(2, 2);   // cache=[2,1], cnt(2)=1, cnt(1)=1
lfu.get(1);      // return 1
                 // cache=[1,2], cnt(2)=1, cnt(1)=2
lfu.put(3, 3);   // 2 is the LFU key because cnt(2)=1 is the smallest, invalidate 2.
                 // cache=[3,1], cnt(3)=1, cnt(1)=2
lfu.get(2);      // return -1 (not found)
lfu.get(3);      // return 3
                 // cache=[3,1], cnt(3)=2, cnt(1)=2
lfu.put(4, 4);   // Both 1 and 3 have the same cnt, but 1 is LRU, invalidate 1.
                 // cache=[4,3], cnt(4)=1, cnt(3)=2
lfu.get(1);      // return -1 (not found)
lfu.get(3);      // return 3
                 // cache=[3,4], cnt(4)=1, cnt(3)=3
lfu.get(4);      // return 4
                 // cache=[4,3], cnt(4)=2, cnt(3)=3
```

**Constraints:**

- `1 <= capacity <= 104`
- `0 <= key <= 105`
- `0 <= value <= 109`
- At most `2 * 105` calls will be made to `get` and `put`.

## Solution
### Approach 1
**Method Name**: Double Hash Maps + Doubly Linked List

**Principle:**
Use two hash maps to store key-to-node mapping and frequency-to-linked-list mapping, achieving O(1) average time complexity for get and put operations.

**Steps:**
1. Define `node` struct to store key, value, frequency, and bidirectional pointers
2. Define `LinkedList` struct implementing doubly linked list with Push, Pop, and Remove operations
3. `LFUCache` contains:
   - `hashtable`: key-to-node mapping for O(1) lookup
   - `queues`: frequency-to-linked-list mapping, nodes with same frequency ordered by usage
   - `capacity`: cache capacity
4. `Get` operation: find node, increment frequency, move to corresponding frequency list
5. `Put` operation:
   - Key exists: update value, increment frequency, move node
   - Key doesn't exist: create new node, remove LFU node if full, insert new node
6. `visitNode`: handle node visit - remove from old frequency list, increment frequency, add to new frequency list
7. `popNode`: find the smallest frequency list, remove the least recently used node

**Example:**
Using the sample input with capacity=2:
- put(1,1) → freq=1, queues[0]=[1]
- put(2,2) → freq=1, queues[0]=[1,2]
- get(1) → return 1, freq=2, queues[0]=[2], queues[1]=[1]
- put(3,3) → evict 2 (smallest freq), queues[0]=[3], queues[1]=[1]

> **Note**:
> - In the same frequency list, head is LRU, tail is most recently used
> - When multiple nodes have same frequency, list order distinguishes recency
> - Minimum frequency could be tracked, but this implementation simplifies by iterating queues to find first non-empty list

```embed-go
PATH: "vault://leetcode/0401-0500/0460_lfu_cache/solution.go"
TITLE: "leetcode 460. LFU Cache"
```
