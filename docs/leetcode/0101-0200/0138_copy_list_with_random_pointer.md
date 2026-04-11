---
link: https://leetcode.com/problems/copy-list-with-random-pointer/
tags:
  - Medium
  - Hash_Table
  - Linked_List
---
## Description
A linked list of length `n` is given such that each node contains an additional random pointer, which could point to any node in the list, or `null`.

Construct a [**deep copy**](https://en.wikipedia.org/wiki/Object_copying#Deep_copy) of the list. The deep copy should consist of exactly `n` **brand new** nodes, where each new node has its value set to the value of its corresponding original node. Both the `next` and `random` pointer of the new nodes should point to new nodes in the copied list such that the pointers in the original list and copied list represent the same list state. **None of the pointers in the new list should point to nodes in the original list**.

For example, if there are two nodes `X` and `Y` in the original list, where `X.random --> Y`, then for the corresponding two nodes `x` and `y` in the copied list, `x.random --> y`.

Return *the head of the copied linked list*.

The linked list is represented in the input/output as a list of `n` nodes. Each node is represented as a pair of `[val, random_index]` where:

- `val`: an integer representing `Node.val`
- `random_index`: the index of the node (range from `0` to `n-1`) that the `random` pointer points to, or `null` if it does not point to any node.

Your code will **only** be given the `head` of the original linked list.

**Example 1:**

![](https://assets.leetcode.com/uploads/2019/12/18/e1.png)

```
Input: head = [[7,null],[13,0],[11,4],[10,2],[1,0]]
Output: [[7,null],[13,0],[11,4],[10,2],[1,0]]
```

**Example 2:**

![](https://assets.leetcode.com/uploads/2019/12/18/e2.png)

```
Input: head = [[1,1],[2,1]]
Output: [[1,1],[2,1]]
```

**Example 3:**

**![](https://assets.leetcode.com/uploads/2019/12/18/e3.png)**

```
Input: head = [[3,null],[3,0],[3,null]]
Output: [[3,null],[3,0],[3,null]]
```

**Constraints:**

- `0 <= n <= 1000`
- `-104 <= Node.val <= 104`
- `Node.random` is `null` or is pointing to some node in the linked list.

## Solution
### Approach 1
**Method Name**: Hash Table Mapping with Node Indices

**Principle:**
Use a hash table to record the mapping between original nodes and their indices, along with the index each node's random pointer points to. Then construct the new linked list using this index information.

**Steps:**
1. First pass through the original list, record each node's index in the hash table
2. Second pass through the hash table, record the index each node's random pointer points to
3. Create new node array based on index information in the hash table
4. Set next and random pointers for new nodes

**Example:**
Input list: `[[7,null],[13,0],[11,4],[10,2],[1,0]]`
- Hash table records: node0→[0], node1→[1], node2→[2], node3→[3], node4→[4]
- Add random indices: node1→[1,0], node2→[2,4], node3→[3,2], node4→[4,0]
- Create new node array and set pointers

```embed-go
PATH: "vault://leetcode/0101-0200/0138_copy_list_with_random_pointer/solution.go"
TITLE: "leetcode 138. Copy List with Random Pointer"
```

### Approach 2
**Method Name**: Recursion + Hash Table

**Principle:**
Use depth-first recursion to traverse the linked list, while using a hash table to record already created nodes to avoid duplicates. Recursively copy both next and random pointers.

**Steps:**
1. Define hash table old2new to map original nodes to new nodes
2. Define recursive function deepCopy:
   - If node is nil, return nil
   - If node exists in hash table, return the already created new node
   - Otherwise create new node and store in hash table
   - Recursively copy next and random pointers
3. Call recursive function to return new list head

**Example:**
Input list: `[[1,1],[2,1]]`
- Copy node1, create new node1, recursively copy next and random
- Copy node2, create new node2, recursively copy next and random
- Recursion returns, setting pointer relationships

```embed-go
PATH: "vault://leetcode/0101-0200/0138_copy_list_with_random_pointer/solution2.go"
TITLE: "leetcode 138. Copy List with Random Pointer"
```

### Approach 3
**Method Name**: In-place Modification + Splitting

**Principle:**
Insert each new node directly after the original node, forming an old→new→old→new... structure. Then use this structure to set random pointers, and finally split the old and new lists. No additional hash table space needed.

**Steps:**
1. First pass: Insert a new node after each original node
2. Second pass: Set random pointers for new nodes (next of original node's random)
3. Third pass: Split old and new lists, restore original list, get new list

**Example:**
Input list: `[[3,null],[3,0],[3,null]]`
- After insertion: 3→3'→3→3'→3→3'
- Set random: second 3' random points to first 3'
- After splitting get new list: 3'→3'→3'

> **Note**: When splitting, need to restore the original list structure at the same time to avoid damaging the original list.

```embed-go
PATH: "vault://leetcode/0101-0200/0138_copy_list_with_random_pointer/solution3.go"
TITLE: "leetcode 138. Copy List with Random Pointer"
```
