---
link: https://leetcode.com/problems/design-add-and-search-words-data-structure/
tags:
  - Medium
  - Depth-First_Search
  - Design
  - Trie
  - String
---
## Description
Design a data structure that supports adding new words and finding if a string matches any previously added string.

Implement the `WordDictionary` class:

- `WordDictionary()` Initializes the object.
- `void addWord(word)` Adds `word` to the data structure, it can be matched later.
- `bool search(word)` Returns `true` if there is any string in the data structure that matches `word` or `false` otherwise. `word` may contain dots `'.'` where dots can be matched with any letter.

**Example:**

```
Input
["WordDictionary","addWord","addWord","addWord","search","search","search","search"]
[[],["bad"],["dad"],["mad"],["pad"],["bad"],[".ad"],["b.."]]
Output
[null,null,null,null,false,true,true,true]

Explanation
WordDictionary wordDictionary = new WordDictionary();
wordDictionary.addWord("bad");
wordDictionary.addWord("dad");
wordDictionary.addWord("mad");
wordDictionary.search("pad"); // return False
wordDictionary.search("bad"); // return True
wordDictionary.search(".ad"); // return True
wordDictionary.search("b.."); // return True
```

**Constraints:**

- `1 <= word.length <= 25`
- `word` in `addWord` consists of lowercase English letters.
- `word` in `search` consist of `'.'` or lowercase English letters.
- There will be at most `2` dots in `word` for `search` queries.
- At most `104` calls will be made to `addWord` and `search`.

## Solution
### Approach 1
**Method Name**: Trie + Depth-First Search

**Principle:**
Use a Trie (prefix tree) to store words. For regular characters, match directly; for wildcard '.' , traverse all possible child nodes with depth-first search.

**Steps:**
1. Define Trie node structure containing `end` marker (whether it's the end of a word) and an array of 26 child nodes for each letter
2. `AddWord` operation: Start from root, iterate through each character, create child node if it doesn't exist, move to child node, finally mark the end of word
3. `Search` operation: Use DFS recursive search. For regular characters, match the corresponding child node; for '.', traverse all non-null child nodes to continue search

**Example:**
- Input: addWord("bad") → builds Trie path b→a→d, d node marked end=true
- Search: search(".ad") → first character '.' traverses all children, finds b then matches a→d, returns true
- Search: search("b..") → matches b, then two '.' traverse children sequentially, finds a→d, returns true

```embed-go
PATH: "vault://leetcode/0201-0300/0211_design_add_and_search_words_data_structure/solution.go"
TITLE: "leetcode 211. Design Add and Search Words Data Structure"
```
