---
link: https://leetcode.com/problems/word-search-ii/
tags:
  - Hard
  - Trie
  - Array
  - String
  - Backtracking
  - Matrix
---
## Description
Given an `m x n` `board` of characters and a list of strings `words`, return *all words on the board*.

Each word must be constructed from letters of sequentially adjacent cells, where **adjacent cells** are horizontally or vertically neighboring. The same letter cell may not be used more than once in a word.

**Example 1:**

![](https://assets.leetcode.com/uploads/2020/11/07/search1.jpg)

```
Input: board = [["o","a","a","n"],["e","t","a","e"],["i","h","k","r"],["i","f","l","v"]], words = ["oath","pea","eat","rain"]
Output: ["eat","oath"]
```

**Example 2:**

![](https://assets.leetcode.com/uploads/2020/11/07/search2.jpg)

```
Input: board = [["a","b"],["c","d"]], words = ["abcb"]
Output: []
```

**Constraints:**

- `m == board.length`
- `n == board[i].length`
- `1 <= m, n <= 12`
- `board[i][j]` is a lowercase English letter.
- `1 <= words.length <= 3 * 104`
- `1 <= words[i].length <= 10`
- `words[i]` consists of lowercase English letters.
- All the strings of `words` are unique.

## Solution

### Approach 1: Brute Force DFS (Check Each Word)

**Method Name**: Depth-first search for each word in the matrix

**Principle:**
Iterate through each word and perform DFS starting from every position in the matrix to check if the current path can form the word. Mark visited cells to avoid reusing the same cell in a single path.

**Steps:**
1. Iterate through each word in the words list
2. For each word, try starting from every position (i,j) in the matrix
3. Use DFS: if current character matches the first letter of the word, mark as visited and continue searching four directions
4. If word length reduces to 0, the complete word is found
5. Backtrack: unmark the current cell as visited for other paths

**Example:**
Take the word "oath" as an example:
- Start with 'o' at (0,0)
- Go down to 'e' at (1,0) (doesn't match 'a', backtrack)
- Go right to 'a' at (0,1), matches the second character
- Continue down to 't' at (1,1), matches the third character
- Then down to 'h' at (2,1), matches the fourth character
- Complete word matched, add to result

> **Note**: This method is inefficient when the word list is large because each word requires traversing the entire matrix.

**Complexity Analysis:**
- **Time Complexity**: O(W × M × N × 4<sup>L</sup>), where W is the number of words, M×N is the matrix size, and L is the maximum length of words. Each word requires M×N starting points, and each starting point has at most 4<sup>L</sup> searches.
- **Space Complexity**: O(M×N + L). Requires M×N visited matrix, and recursion stack depth is at most L.

```embed-go
PATH: "vault://leetcode/0201-0300/0212_word_search_ii/solution.go"
TITLE: "leetcode 212.Word Search II"
```

### Approach 2: Hash Map + DFS (Generate Words from Matrix)

**Method Name**: DFS from each matrix position to generate possible words, using hash map for quick lookup

**Principle:**
Store all words in a hash map, then perform DFS from each position in the matrix to generate all possible path strings. For each generated string, check if it exists in the hash map. If found, add to result and remove from hash map (to avoid duplicates).

**Steps:**
1. Store all words in hash map `dictionary`, also record the maximum word length `maxLength`
2. Start DFS from each position (i,j) in the matrix
3. Accumulate characters during search to form current string
4. If current string is in hash map, add to result and remove from hash map
5. Stop searching when path length reaches `maxLength`
6. Backtrack: remove last character and unmark visited

**Example:**
Starting from 'o' at (0,0):
- Generate "o" (not in hash map)
- Go right to (0,1) generate "oa" (not in hash map)
- Go right to (0,2) generate "oaa" (not in hash map)
- Go down to (1,2) generate "oaae" (exceeds length limit, backtrack)
- ...
- Another path: (0,0)→(0,1)→(1,1)→(2,1) generates "oath" (in hash map, add to result)

> **Note**: This method uses hash map for quick word lookup and avoids duplicates by removing found words, but may still generate many invalid paths in worst case.

**Complexity Analysis:**
- **Time Complexity**: O(M × N × 4<sup>L</sup>), where M×N is the matrix size, and L is the maximum length of words. Requires M×N starting points, each with at most 4<sup>L</sup> searches. Hash map lookup and deletion are O(1).
- **Space Complexity**: O(W + M×N + L). Requires hash map storing W words, M×N visited matrix, and recursion stack depth at most L.

```embed-go
PATH: "vault://leetcode/0201-0300/0212_word_search_ii/solution2.go"
TITLE: "leetcode 212.Word Search II"
```

### Approach 3: Trie + DFS (Prefix Pruning Optimization)

**Method Name**: Use Trie to store word prefixes, prune invalid paths in real-time during DFS

**Principle:**
Build a Trie (prefix tree) from all words, where each node represents a prefix. During DFS search, check at each step if the current prefix exists in the Trie. If not, immediately prune that path to avoid searching invalid paths.

**Steps:**
1. Insert all words into Trie, each node marks if it's the end of a word
2. Start DFS from each position (i,j) in the matrix
3. Accumulate characters during search to form current string
4. Look up current string in Trie:
   - If prefix doesn't exist, immediately prune and return
   - If prefix exists and is a complete word, add to result
5. Continue searching four directions (if prefix still exists)
6. Backtrack: remove last character and unmark visited

**Example:**
Trie contains "oath", "pea", "eat", "rain":
```
root
├── o -> o
│   └── a -> oa
│       └── t -> oat
│           └── h -> oath (end=true)
├── p -> p
│   └── e -> pe
│       └── a -> pea (end=true)
├── e -> e
│   └── a -> ea
│       └── t -> eat (end=true)
└── r -> r
    └── a -> ra
        └── i -> rai
            └── n -> rain (end=true)
```

Starting from (0,0) search "o" → "oa" → "oat" → "oath" (found), while encountering prefix "oaa" which doesn't exist in Trie, immediately prune.

> **Note**: This method achieves efficient pruning through Trie prefix matching, significantly reducing invalid searches. It's the optimal among the three approaches. Setting end mark to false after finding a word avoids adding the same word multiple times.

**Complexity Analysis:**
- **Time Complexity**: O(M × N × 4<sup>L</sup>), where M×N is the matrix size, and L is the maximum length of words. Although theoretical upper bound is same as Approach 2, actual runtime is significantly reduced through Trie pruning, especially with large word lists.
- **Space Complexity**: O(W×L + M×N + L). Requires Trie storing W words (each word has at most L nodes), M×N visited matrix, and recursion stack depth at most L.

```embed-go
PATH: "vault://leetcode/0201-0300/0212_word_search_ii/solution3.go"
TITLE: "leetcode 212.Word Search II"
```
