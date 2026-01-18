# 📚 Algorithm
>  一个基于 Go 语言、结合 LeetCode 实践 与 《算法导论（CLRS，第 4 版）》 的数据结构与算法学习仓库。

[![Go](https://img.shields.io/badge/Go-1.25+-blue) ![[License](https://img.shields.io/badge/license-MIT-green.svg)](../LICENSE)

本项目结合了经典的理论教材 《算法导论》第 4 版（CLRS） 与实战平台 LeetCode，旨在通过“理论学习 + 编码实践 + 文档沉淀”三位一体的方式，构建完整的算法知识体系。

项目文档采用 Obsidian 组织，支持全库双链跳转，方便在理论分析与代码实现之间无缝切换。

## 📖 项目简介
- CLRS 实践：实现《算法导论》中的伪代码，挑战课后习题，将抽象算法具象化。
- LeetCode 刷题：精选高频算法题，使用 Go 语言实现，并附带详尽的单元测试。
- Obsidian 知识库：在 docs 目录下记录解题思路、复杂度分析与学习笔记，打造个人算法百科。
- AI 赋能：规划中的智能翻译与 RAG 答疑助手，利用现代技术提升学习效率。

## 🚀 快速开始
### 环境依赖
- **Go**: 1.25+
- **Obsidian**: 推荐安装以获得最佳文档阅读体验

### 运行与测试
1. **克隆仓库**
```bash
$ git clone https://github.com/MorePeanuts/algorithm.git
$ cd algorithm
```
    
2. **运行测试** 每个 LeetCode 题目或 CLRS 示例均配有测试文件：
```bash
# 测试特定题目
$ go test ./leetcode/0001-0100/0001_Two_Sum/...

# 运行所有测试
$ go test ./...
```

## 📁 仓库结构
```plain
.
├── clrs/               # 《算法导论》第四版代码实现
├── leetcode/           # LeetCode 刷题代码
│   └── 0001-0100/      # 按题号区间分组
│       └── 0001_Two_Sum/
│           ├── solution.go       # 核心算法
│           └── solution_test.go  # 单元测试
├── docs/               # Obsidian 文档库根目录
│   ├── CLRS/           # 算法导论各章节学习笔记与习题思路
│   └── leetcode/       # 题目描述、解法分析（链接至源码）
├── cmd/                # 自研辅助工具（爬虫、翻译脚本等）
├── go.mod              # Go 模块配置
└── README.md
```

文档存储在 `docs/` 目录下，遵循以下原则：
1. 结构化：按照 `题号_题目名称.md` 命名，保持与代码目录对应。
2. 代码链接：在 Markdown 中链接至 Go 源码文件。
## 🛠️ 后续功能规划

- 中英文档自动翻译智能体：开发一个 LLM Agent，自动监听 docs 变动并将中文笔记翻译为英文，保持双语同步。
- 基于 RAG 的智能答疑助手：基于本地 docs 库构建知识索引；实现“理论 + 实践”结合的问答，例如：“哈希表如何实现？LeetCode 中有哪些相关题目？”。
- 自动题目爬虫：在 `cmd/` 下开发 Go 工具，输入 LeetCode URL 自动在 docs 和 leetcode 路径下生成模板。
## 📜 License
本项目采用 MIT License，详见 [LICENSE](../LICENSE)。