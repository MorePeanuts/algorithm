---
link: https://leetcode.com/problems/group-anagrams/
tags:
  - Medium
  - Array
  - Hash_Table
  - String
  - Sorting
---
## Description
Given an array of strings `strs`, group the anagrams together. You can return the answer in **any order**.

**Example 1:**

**Input:** strs = ["eat","tea","tan","ate","nat","bat"]

**Output:** [["bat"],["nat","tan"],["ate","eat","tea"]]

**Explanation:**

- There is no string in strs that can be rearranged to form `"bat"`.
- The strings `"nat"` and `"tan"` are anagrams as they can be rearranged to form each other.
- The strings `"ate"`, `"eat"`, and `"tea"` are anagrams as they can be rearranged to form each other.

**Example 2:**

**Input:** strs = [""]

**Output:** [[""]]

**Example 3:**

**Input:** strs = ["a"]

**Output:** [["a"]]

**Constraints:**

- `1 <= strs.length <= 104`
- `0 <= strs[i].length <= 100`
- `strs[i]` consists of lowercase English letters.

## Solution

### Approach 1

**Character Counting Method**: The essence of anagrams is having the same character types and counts, just in different order.

**Principle:**
Use an array of length 26 to record the occurrence count of each letter. Identical count arrays correspond to the same anagram group.

**Steps:**
1. Iterate through the string, count each character's occurrence, store in a `[26]byte` array.
2. Use this array as the HashMap key, group strings accordingly.

**Example:**
- `"eat"` -> `[1,0,0,0,1,...]` (a=1, e=1, t=1)
- `"tea"` -> `[1,0,0,0,1,...]` (same key)
- `"tan"` -> `[1,0,0,0,0,...,1,...]` (a=1, n=1, t=1, different key)

```embed-go
PATH: "vault://leetcode/0001-0100/0049_group_anagrams/solution1.go"
TITLE: "leetcode 49. Group Anagrams"
```

### Approach 2

**Fundamental Theorem of Arithmetic**: **Every integer greater than 1 can be uniquely represented as a product of prime numbers**.

**Principle:**
Multiplication is commutative (`a * b = b * a`), which naturally solves the "order-independent" problem. Using the property of primes, different letter combinations will always produce different products.

**Steps:**
1. Create a mapping table that maps 26 letters to 26 different **prime numbers**.
    - `a` -> 2, `b` -> 3, `c` -> 5, `d` -> 7, ...
2. Iterate through the string, multiply the corresponding prime for each character.
3. The product result is the unique key.

**Example:**
- `a` = 2, `e` = 11, `t` = 71 (assumed)
- Key for `"eat"` = $11 \times 2 \times 71 = 1562$
- Key for `"ate"` = $2 \times 71 \times 11 = 1562$
- Key for `"aba"` = $2 \times 3 \times 2 = 12$

> **Integer Overflow Risk**: The problem states `strs[i].length <= 100`. For 100 'z's, the product would be `(26th prime)^100`, causing integer overflow, so big integer types are needed.

```embed-go
PATH: "vault://leetcode/0001-0100/0049_group_anagrams/solution2.go"
TITLE: "leetcode 49. Group Anagrams"
```

### Approach 3

**Sorting Method**: The most intuitive approach—anagrams become identical after sorting.

**Principle:**
Sort the characters in each string, and all anagrams will become the same string, which serves as the grouping key.

**Steps:**
1. Sort the characters of each string.
2. Use the sorted string as the HashMap key, group original strings accordingly.

**Example:**
- `"eat"` -> sort -> `"aet"`
- `"tea"` -> sort -> `"aet"` (same key)
- `"tan"` -> sort -> `"ant"` (different key)

> **Complexity**: Sorting each string takes $O(k \log k)$, total time complexity $O(n \cdot k \log k)$, where $k$ is the average string length.

```embed-go
PATH: "vault://leetcode/0001-0100/0049_group_anagrams/solution3.go"
TITLE: "leetcode 49. Group Anagrams"
```
