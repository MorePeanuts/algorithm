---
link: https://leetcode.com/problems/serialize-and-deserialize-binary-tree/
tags:
  - Hard
  - Tree
  - Depth-First_Search
  - Breadth-First_Search
  - Design
  - String
  - Binary_Tree
---
## Description
Serialization is the process of converting a data structure or object into a sequence of bits so that it can be stored in a file or memory buffer, or transmitted across a network connection link to be reconstructed later in the same or another computer environment.

Design an algorithm to serialize and deserialize a binary tree. There is no restriction on how your serialization/deserialization algorithm should work. You just need to ensure that a binary tree can be serialized to a string and this string can be deserialized to the original tree structure.

**Clarification:** The input/output format is the same as [how LeetCode serializes a binary tree](https://support.leetcode.com/hc/en-us/articles/32442719377939-How-to-create-test-cases-on-LeetCode#h_01J5EGREAW3NAEJ14XC07GRW1A). You do not necessarily need to follow this format, so please be creative and come up with different approaches yourself.

**Example 1:**

![](https://assets.leetcode.com/uploads/2020/09/15/serdeser.jpg)

```
Input: root = [1,2,3,null,null,4,5]
Output: [1,2,3,null,null,4,5]
```

**Example 2:**

```
Input: root = []
Output: []
```

**Constraints:**

- The number of nodes in the tree is in the range `[0, 104]`.
- `-1000 <= Node.val <= 1000`

## Solution
### Approach 1
**Method Name**: Level Order Traversal (BFS) + Unicode Encoding

**Principle:**
Use breadth-first search to traverse the binary tree level by level, encoding node values and null nodes with special Unicode characters to achieve compact string serialization.

**Steps:**
1. **Serialization**:
   - Use a queue for level order traversal, including null nodes
   - Non-null nodes: convert value + 1000 to Unicode character
   - Null nodes: use Unicode character 5000 as marker
   - Concatenate all characters into a string

2. **Deserialization**:
   - Parse the string character by character as Unicode
   - Character 5000 represents null node, other characters minus 1000 give node values
   - Use a queue to reconstruct the binary tree in level order

**Example:**
- Input: `[1,2,3,null,null,4,5]`
- Serialization process:
  - Level order traversal yields: 1, 2, 3, null, null, 4, 5
  - Encoded as Unicode: 1001, 1002, 1003, 5000, 5000, 1004, 1005
  - Final string: composed of these Unicode characters
- Deserialization reverses the process to restore the original binary tree

**Complexity Analysis:**
- Time Complexity: O(n), where n is the number of nodes. We need to traverse all nodes once for both serialization and deserialization.
- Space Complexity: O(n). The queue stores up to O(n) nodes, and the serialization result also requires O(n) space.

> **Note**:
> - This method leverages the constraint that Node.val is in the range [-1000, 1000], using a +1000 offset to avoid conflict with the null node marker 5000
> - Using Unicode character encoding makes the serialization result very compact, with each node occupying only one character
> - Level order traversal allows direct reconstruction of the binary tree structure without additional information

```embed-go
PATH: "vault://leetcode/0201-0300/0297_serialize_and_deserialize_binary_tree/solution.go"
TITLE: "leetcode 297. Serialize and Deserialize Binary Tree"
```
