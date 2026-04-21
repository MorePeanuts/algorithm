---
link: https://leetcode.cn/problems/implement-trie-prefix-tree/
tags:
  - 中等
  - 设计
  - 字典树
  - 哈希表
  - 字符串
---
## 题目描述
**[Trie](https://baike.baidu.com/item/字典树/9825209?fr=aladdin)**（发音类似 "try"）或者说 **前缀树** 是一种树形数据结构，用于高效地存储和检索字符串数据集中的键。这一数据结构有相当多的应用情景，例如自动补全和拼写检查。

请你实现 Trie 类：

- `Trie()` 初始化前缀树对象。
- `void insert(String word)` 向前缀树中插入字符串 `word` 。
- `boolean search(String word)` 如果字符串 `word` 在前缀树中，返回 `true`（即，在检索之前已经插入）；否则，返回 `false` 。
- `boolean startsWith(String prefix)` 如果之前已经插入的字符串 `word` 的前缀之一为 `prefix` ，返回 `true` ；否则，返回 `false` 。

**示例：**

```
输入
["Trie", "insert", "search", "search", "startsWith", "insert", "search"]
[[], ["apple"], ["apple"], ["app"], ["app"], ["app"], ["app"]]
输出
[null, null, true, false, true, null, true]

解释
Trie trie = new Trie();
trie.insert("apple");
trie.search("apple");   // 返回 True
trie.search("app");     // 返回 False
trie.startsWith("app"); // 返回 True
trie.insert("app");
trie.search("app");     // 返回 True
```

**提示：**

- `1 <= word.length, prefix.length <= 2000`
- `word` 和 `prefix` 仅由小写英文字母组成
- `insert`、`search` 和 `startsWith` 调用次数 **总计** 不超过 `3 * 104` 次

## 题目解析
### 解法1
**方法名**：数组实现前缀树

**原理：**
前缀树是一种树形数据结构，每个节点包含一个标记表示是否为单词结尾，以及一个长度为26的子节点数组对应小写英文字母。通过从根节点开始逐字符遍历字符串，实现单词的插入、搜索和前缀查询。

**步骤：**
1. **节点结构**：每个节点包含 `end` 标记（是否是单词结尾）和 `children` 数组（26个字母的子节点）
2. **插入操作**：从根节点开始，逐个字符处理，若对应子节点不存在则创建，移动到子节点，最后标记单词结尾
3. **搜索操作**：从根节点开始，逐个字符查找，若某字符不存在则返回 false，遍历完成后检查 `end` 标记
4. **前缀查询**：与搜索操作类似，但不需要检查 `end` 标记，只要能遍历完所有前缀字符即可

**示例：**
- 插入 "apple"：创建路径 a→p→p→l→e，在 e 节点标记 end=true
- 搜索 "apple"：沿路径遍历到 e，end=true → 返回 true
- 搜索 "app"：沿路径遍历到第二个 p，end=false → 返回 false
- 前缀查询 "app"：能遍历完所有字符 → 返回 true

```embed-go
PATH: "vault://leetcode/0201-0300/0208_implement_trie_prefix_tree/solution.go"
TITLE: "leetcode 208. Implement Trie Prefix Tree"
```
