---
link: https://leetcode.cn/problems/design-add-and-search-words-data-structure/
tags:
  - 中等
  - 深度优先搜索
  - 设计
  - 字典树
  - 字符串
---
## 题目描述
请你设计一个数据结构，支持 添加新单词 和 查找字符串是否与任何先前添加的字符串匹配 。

实现词典类 `WordDictionary` ：

- `WordDictionary()` 初始化词典对象
- `void addWord(word)` 将 `word` 添加到数据结构中，之后可以对它进行匹配
- `bool search(word)` 如果数据结构中存在字符串与 `word` 匹配，则返回 `true` ；否则，返回  `false` 。`word` 中可能包含一些 `'.'` ，每个 `.` 都可以表示任何一个字母。

**示例：**

```
输入：
["WordDictionary","addWord","addWord","addWord","search","search","search","search"]
[[],["bad"],["dad"],["mad"],["pad"],["bad"],[".ad"],["b.."]]
输出：
[null,null,null,null,false,true,true,true]

解释：
WordDictionary wordDictionary = new WordDictionary();
wordDictionary.addWord("bad");
wordDictionary.addWord("dad");
wordDictionary.addWord("mad");
wordDictionary.search("pad"); // 返回 False
wordDictionary.search("bad"); // 返回 True
wordDictionary.search(".ad"); // 返回 True
wordDictionary.search("b.."); // 返回 True
```

**提示：**

- `1 <= word.length <= 25`
- `addWord` 中的 `word` 由小写英文字母组成
- `search` 中的 `word` 由 '.' 或小写英文字母组成
- 最多调用 `104` 次 `addWord` 和 `search`

## 题目解析
### 解法1
**方法名**：字典树 + 深度优先搜索

**原理：**
使用字典树（Trie）存储单词，对于普通字符直接匹配，对于通配符 '.' 则遍历所有可能的子节点进行深度优先搜索。

**步骤：**
1. 定义字典树节点结构，包含 `end` 标记（是否为单词结尾）和 26 个字母的子节点数组
2. `AddWord` 操作：从根节点开始，逐个字符遍历，若对应子节点不存在则创建，移动到子节点，最后标记单词结尾
3. `Search` 操作：使用 DFS 递归搜索，当前字符为普通字母时直接匹配对应子节点，为 '.' 时遍历所有非空子节点继续搜索

**示例：**
- 输入：addWord("bad") → 构建字典树路径 b→a→d，d 节点标记 end=true
- 搜索：search(".ad") → 第一个字符 '.' 遍历所有子节点，找到 b 后匹配 a→d，返回 true
- 搜索：search("b..") → 匹配 b 后，后续两个 '.' 依次遍历子节点，找到 a→d，返回 true

```embed-go
PATH: "vault://leetcode/0201-0300/0211_design_add_and_search_words_data_structure/solution.go"
TITLE: "leetcode 211. Design Add and Search Words Data Structure"
```
