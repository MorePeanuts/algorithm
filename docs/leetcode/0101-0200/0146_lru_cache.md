---
link: https://leetcode.com/problems/lru-cache/
tags:
  - Medium
  - Design
  - Hash_Table
  - Linked_List
  - Doubly-Linked_List
---
## Description
Design a data structure that follows the constraints of a **[Least Recently Used (LRU) cache](https://en.wikipedia.org/wiki/Cache_replacement_policies#LRU)**.

Implement the `LRUCache` class:

- `LRUCache(int capacity)` Initialize the LRU cache with **positive** size `capacity`.
- `int get(int key)` Return the value of the `key` if the key exists, otherwise return `-1`.
- `void put(int key, int value)` Update the value of the `key` if the `key` exists. Otherwise, add the `key-value` pair to the cache. If the number of keys exceeds the `capacity` from this operation, **evict** the least recently used key.

The functions `get` and `put` must each run in `O(1)` average time complexity.

**Example 1:**

```
Input
["LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"]
[[2], [1, 1], [2, 2], [1], [3, 3], [2], [4, 4], [1], [3], [4]]
Output
[null, null, null, 1, null, -1, null, -1, 3, 4]

Explanation
LRUCache lRUCache = new LRUCache(2);
lRUCache.put(1, 1); // cache is {1=1}
lRUCache.put(2, 2); // cache is {1=1, 2=2}
lRUCache.get(1);    // return 1
lRUCache.put(3, 3); // LRU key was 2, evicts key 2, cache is {1=1, 3=3}
lRUCache.get(2);    // returns -1 (not found)
lRUCache.put(4, 4); // LRU key was 1, evicts key 1, cache is {4=4, 3=3}
lRUCache.get(1);    // return -1 (not found)
lRUCache.get(3);    // return 3
lRUCache.get(4);    // return 4
```

**Constraints:**

- `1 <= capacity <= 3000`
- `0 <= key <= 104`
- `0 <= value <= 105`
- At most `2 * 105` calls will be made to `get` and `put`.

## Solution
### Approach 1
**Method Name**: Hash Table + Doubly Linked List

**Principle:**
Use a hash table for O(1) time lookups and a doubly linked list to maintain access order, with the most recently used node at the tail and the least recently used node at the head.

**Steps:**
1. **Get Operation**:
   - Check if the key exists in the hash table, return -1 if not found
   - If found, move the node to the tail of the linked list (marking it as recently used)
   - Return the node's value

2. **Put Operation**:
   - Check if the key exists in the hash table
   - If exists, update the value and move the node to the tail
   - If not exists, create a new node and add it to both the hash table and the tail of the linked list
   - If capacity is exceeded, remove the node at the head (least recently used) and remove it from the hash table

3. **moveToTail Helper Operation**:
   - No operation needed if the node is already at the tail
   - If the node is at the head, update the head pointer
   - Disconnect the node from its previous and next nodes
   - Connect the node to the tail of the linked list and update the tail pointer

**Example:**
With capacity=2:
- put(1, 1) → linked list: [1], hash table: {1:node1}
- put(2, 2) → linked list: [1, 2], hash table: {1:node1, 2:node2}
- get(1) → access node1, move to tail → linked list: [2, 1], return 1
- put(3, 3) → capacity exceeded, delete head node2, add node3 to tail → linked list: [1, 3], hash table: {1:node1, 3:node3}
- get(2) → not found, return -1

```embed-go
PATH: "vault://leetcode/0101-0200/0146_lru_cache/solution.go"
TITLE: "leetcode 146. LRU Cache"
```
