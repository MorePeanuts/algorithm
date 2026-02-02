---
description: 为 LeetCode 解法代码生成题目解析文档
argument-hint: [题号] [--lang zh|en]（可选，默认生成中英文两个版本）
allowed-tools: Read, Glob, Edit
---

根据提供的 LeetCode 题号 $ARGUMENTS，为已有的解法代码生成题目解析。

**参数解析：**
- 题号：必填，如 `0049` 或 `49`
- `--lang zh`：只生成中文版本
- `--lang en`：只生成英文版本
- 不指定 `--lang`：同时生成中英文两个版本

**执行步骤：**

1. **定位文件**
   - 将题号补齐为4位（如 49 -> 0049）
   - 确定题号范围目录（如 0049 属于 0001-0100）
   - 使用 Glob 查找：
     - 中文文档：`docs-zh/leetcode/*/XXXX_*.md`
     - 英文文档：`docs/leetcode/*/XXXX_*.md`
     - 代码文件：`leetcode/*/XXXX_*/solution*.go`（排除 `*_test.go`）

2. **读取信息**
   - 读取题目文档，提取：
     - frontmatter tags 中的难度（简单/中等/困难 或 Easy/Medium/Hard）
     - 题目描述内容
   - 读取所有 solution*.go 文件内容

3. **生成题目解析**
   根据难度调整详细程度：
   - **简单/Easy**：原理 + 步骤
   - **中等/Medium**：原理 + 步骤 + 示例
   - **困难/Hard**：原理 + 步骤 + 示例 + 复杂度分析 + 注意事项

   为每个 solution 文件生成解法，根据语言版本：

   **中文版本（docs-zh）：**
   - 标题格式：`### 解法N`
   - 包含：**原理：**、**步骤：**、**示例：**

   **英文版本（docs）：**
   - 标题格式：`### Approach N`
   - 包含：**Principle:**、**Steps:**、**Example:**

4. **更新文档**
   根据 `--lang` 参数决定更新哪些文档：
   - 中文：更新 `docs-zh/leetcode/` 下的文档，追加"## 题目解析"
   - 英文：更新 `docs/leetcode/` 下的文档，追加"## Solution"

   如果已存在对应章节，询问用户是否覆盖。

**embed-go 引用格式：**
```embed-go
PATH: "vault://leetcode/范围目录/题目目录/solutionN.go"
TITLE: "leetcode 题号.题目名"
```

**中文生成格式参考：**
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

**英文生成格式参考：**
```markdown
## Solution
### Approach 1
**Method Name**: Brief description of the core idea.

**Principle:**
One sentence explaining the core principle of the algorithm.

**Steps:**
1. First step
2. Second step

**Example:** (for Medium/Hard difficulty)
- Input example -> Processing -> Output

> **Note**: Special case explanation (for Hard difficulty)

​```embed-go
PATH: "vault://leetcode/XXXX-XXXX/XXXX_problem_name/solution1.go"
TITLE: "leetcode N. Problem Name"
​```
```
