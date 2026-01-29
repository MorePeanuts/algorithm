---
description: 为 LeetCode 解法代码生成题目解析文档
argument-hint: [题号，如 0049 或 49]
allowed-tools: Read, Glob, Edit
---

根据提供的 LeetCode 题号 $ARGUMENTS，为已有的解法代码生成题目解析。

**执行步骤：**

1. **定位文件**
   - 将题号补齐为4位（如 49 -> 0049）
   - 确定题号范围目录（如 0049 属于 0001-0100）
   - 使用 Glob 查找：
     - 题目文档：`docs/leetcode/*/XXXX_*.md`
     - 代码文件：`leetcode/*/XXXX_*/solution*.go`（排除 `*_test.go`）

2. **读取信息**
   - 读取题目文档，提取：
     - frontmatter tags 中的难度（简单/中等/困难）
     - 题目描述内容
   - 读取所有 solution*.go 文件内容

3. **生成题目解析**
   根据难度调整详细程度：
   - **简单**：原理 + 步骤
   - **中等**：原理 + 步骤 + 示例
   - **困难**：原理 + 步骤 + 示例 + 复杂度分析 + 注意事项

   为每个 solution 文件生成一个"### 解法N"，包含：
   - 方法名称（加粗）和核心思想
   - **原理：** 说明算法原理
   - **步骤：** 实现步骤列表
   - **示例：** 具体示例（中等/困难难度）
   - embed-go 代码引用块

4. **更新文档**
   使用 Edit 工具将生成的"## 题目解析"部分追加到题目文档末尾。
   如果已存在"## 题目解析"，询问用户是否覆盖。

**embed-go 引用格式：**
```embed-go
PATH: "vault://leetcode/范围目录/题目目录/solutionN.go"
TITLE: "leetcode 题号.题目名"
```

**生成格式参考：**
```markdown
## 题目解析
### 解法1
**方法名**：简述核心思想

**原理：**
用一句话说明算法的核心原理。

**步骤：**
1. 第一步操作
2. 第二步操作

**示例：**（中等/困难难度包含此部分）
- 输入示例 -> 处理过程 -> 输出

> **注意**：特殊情况说明（困难难度包含此部分）

​```embed-go
PATH: "vault://leetcode/XXXX-XXXX/XXXX_problem_name/solution1.go"
TITLE: "leetcode N.题目名"
​```
```
