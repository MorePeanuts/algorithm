---
link: https://leetcode.com/problems/implement-trie-prefix-tree/
tags:
  - Medium
  - Design
  - Trie
  - Hash_Table
  - String
---
## Description
A [**trie**](https://en.wikipedia.org/wiki/Trie) (pronounced as "try") or **prefix tree** is a tree data structure used to efficiently store and retrieve keys in a dataset of strings. There are various applications of this data structure, such as autocomplete and spellchecker.

Implement the Trie class:

- `Trie()` Initializes the trie object.
- `void insert(String word)` Inserts the string `word` into the trie.
- `boolean search(String word)` Returns `true` if the string `word` is in the trie (i.e., was inserted before), and `false` otherwise.
- `boolean startsWith(String prefix)` Returns `true` if there is a previously inserted string `word` that has the prefix `prefix`, and `false` otherwise.

**Example 1:**

```
Input
["Trie", "insert", "search", "search", "startsWith", "insert", "search"]
[[], ["apple"], ["apple"], ["app"], ["app"], ["app"], ["app"]]
Output
[null, null, true, false, true, null, true]

Explanation
Trie trie = new Trie();
trie.insert("apple");
trie.search("apple");   // return True
trie.search("app");     // return False
trie.startsWith("app"); // return True
trie.insert("app");
trie.search("app");     // return True
```

**Constraints:**

- `1 <= word.length, prefix.length <= 2000`
- `word` and `prefix` consist only of lowercase English letters.
- At most `3 * 104` calls **in total** will be made to `insert`, `search`, and `startsWith`.

## Solution
### Approach 1
**Method Name**: Array-based Trie Implementation

**Principle:**
A trie is a tree data structure where each node contains a flag indicating if it's the end of a word, and an array of 26 child nodes corresponding to lowercase English letters. By traversing the string character by character from the root node, we can implement word insertion, search, and prefix query operations.

**Steps:**
1. **Node Structure**: Each node contains an `end` flag (marking word end) and a `children` array (26 letters' child nodes)
2. **Insert Operation**: Start from root, process each character, create child node if it doesn't exist, move to child node, finally mark word end
3. **Search Operation**: Start from root, find each character, return false if any character doesn't exist, check `end` flag after traversal
4. **Prefix Query**: Similar to search operation, but no need to check `end` flag, just need to traverse all prefix characters

**Example:**
- Insert "apple": Create path a→p→p→l→e, mark end=true at e node
- Search "apple": Traverse path to e, end=true → return true
- Search "app": Traverse path to second p, end=false → return false
- Prefix query "app": Can traverse all characters → return true

```embed-go
PATH: "vault://leetcode/0201-0300/0208_implement_trie_prefix_tree/solution.go"
TITLE: "leetcode 208. Implement Trie Prefix Tree"
```
