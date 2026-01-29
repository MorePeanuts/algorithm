---
description: 为 LeetCode 解法代码生成单元测试
argument-hint: [题号，如 0049 或 49]
allowed-tools: Read, Glob, Edit, Write
---

根据提供的 LeetCode 题号 $ARGUMENTS，为已有的解法代码生成单元测试。

**执行步骤：**

1. **定位文件**
   - 将题号补齐为4位（如 49 -> 0049）
   - 确定题号范围目录（如 0049 属于 0001-0100）
   - 使用 Glob 查找：
     - 题目文档：`docs/leetcode/*/XXXX_*.md`
     - 代码文件：`leetcode/*/XXXX_*/solution*.go`（排除 `*_test.go`）

2. **读取并分析信息**
   **从题目文档提取：**
   - 题目描述（理解问题本质，用于生成边缘测试）
   - 示例用例（`示例 N:` 格式的输入/输出）
   - 提示/约束条件（如数据范围、特殊限制）

   **从 solution*.go 文件提取：**
   - 包名（如 `leetcode0049`）
   - 所有解法函数名及其签名（如 `groupAnagrams(strs []string) [][]string`）

3. **分析是否需要辅助函数**
   根据函数签名和测试需求，判断是否需要通用的测试辅助函数。辅助函数应当是可复用的、泛用的工具函数，例如：
   - 无序切片比较（如 `MatchTwo2dStringSlice`）
   - 链表/树等数据结构的构建与比较
   - 浮点数近似比较
   - 其他可能在多个题目中复用的测试工具

   检查 `leetcode/utils/testing_utils.go` 是否已有所需函数：
   - 如有：直接使用
   - 如无：先在 `testing_utils.go` 中生成辅助函数，再使用

4. **生成 TestXXXExamples 函数**
   - 从文档解析示例（`示例 N:` 格式）
   - 生成 table-driven test 结构
   - 测试所有解法变体
   - 测试用例命名：`example_1`, `example_2`, ...

5. **生成 TestXXX 函数**
   根据题目描述和约束条件生成边缘测试：
   - 边界值测试（最小/最大长度、最小/最大值）
   - 特殊情况（空输入、单元素、全相同元素等）
   - 针对题目特点的特殊用例
   - 较大规模数据测试

   测试用例数量：20-50 个，全面覆盖各种边缘情况
   测试用例命名：使用描述性名称，如 `single_element`, `all_same`, `max_length` 等

6. **写入/更新文件**
   - 如果 `solution_test.go` 不存在：创建新文件
   - 如果存在：询问用户是否覆盖

**生成的测试文件格式：**

```go
package leetcodeXXXX

import (
    "testing"

    "github.com/MorePeanuts/algorithm/leetcode/utils"  // 如需辅助函数
)

func TestXXXExamples(t *testing.T) {
    tests := []struct {
        name string
        // 输入参数字段（从函数签名推断）
        want // 返回类型（从函数签名推断）
    }{
        // 从文档解析的示例
    }

    for _, tst := range tests {
        t.Run(tst.name, func(t *testing.T) {
            // 测试所有解法变体
            // 使用适当的比较方式（直接比较/DeepEqual/辅助函数）
        })
    }
}

func TestXXX(t *testing.T) {
    tests := []struct {
        name string
        // 输入参数字段
        want // 返回类型
    }{
        // 边缘情况测试用例（20-50个）
    }

    for _, tst := range tests {
        t.Run(tst.name, func(t *testing.T) {
            // 测试所有解法变体
        })
    }
}
```

**示例参考（0049 题）：**

```go
package leetcode0049

import (
    "testing"

    "github.com/MorePeanuts/algorithm/leetcode/utils"
)

func TestGroupAnagramsExamples(t *testing.T) {
    tests := []struct {
        name string
        strs []string
        want [][]string
    }{
        {
            "example_1",
            []string{"eat", "tea", "tan", "ate", "nat", "bat"},
            [][]string{{"bat"}, {"nat", "tan"}, {"ate", "eat", "tea"}},
        },
        // ...更多示例
    }

    for _, tst := range tests {
        t.Run(tst.name, func(t *testing.T) {
            got := groupAnagrams(tst.strs)
            if !utils.MatchTwo2dStringSlice(got, tst.want) {
                t.Errorf("groupAnagrams(%v) = %v, want %v", tst.strs, got, tst.want)
            }

            got = groupAnagrams2(tst.strs)
            if !utils.MatchTwo2dStringSlice(got, tst.want) {
                t.Errorf("groupAnagrams2(%v) = %v, want %v", tst.strs, got, tst.want)
            }

            got = groupAnagrams3(tst.strs)
            if !utils.MatchTwo2dStringSlice(got, tst.want) {
                t.Errorf("groupAnagrams3(%v) = %v, want %v", tst.strs, got, tst.want)
            }
        })
    }
}

func TestGroupAnagrams(t *testing.T) {
    tests := []struct {
        name string
        strs []string
        want [][]string
    }{
        {"single_empty_string", []string{""}, [][]string{{""}}},
        {"single_char", []string{"a"}, [][]string{{"a"}}},
        {"all_same", []string{"a", "a", "a"}, [][]string{{"a", "a", "a"}}},
        // ...更多边缘测试用例（20-50个）
    }

    for _, tst := range tests {
        t.Run(tst.name, func(t *testing.T) {
            // 测试所有解法变体
        })
    }
}
```
