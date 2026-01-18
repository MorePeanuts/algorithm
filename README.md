# 📚 Algorithm

> A data structures and algorithms learning repository based on Go language, combining LeetCode practice and "Introduction to Algorithms (CLRS, 4th Edition)."

![Go](https://img.shields.io/badge/Go-1.25+-blue) [![License](https://img.shields.io/badge/license-MIT-green.svg)](./LICENSE)

This project integrates the classic theoretical textbook "Introduction to Algorithms, 4th Edition (CLRS)" with the practical platform LeetCode. It aims to build a comprehensive algorithm knowledge system through a trinity approach of "theoretical learning + coding practice + documentation accumulation."

The project documentation is organized using Obsidian, supporting full-repository bidirectional linking for seamless switching between theoretical analysis and code implementation.

## 📖 Project Introduction

- **CLRS Practice**: Implement pseudocode from "Introduction to Algorithms," tackle exercises, and materialize abstract algorithms.

- **LeetCode Problem Solving**: Select high-frequency algorithm problems, implement them in Go, and include comprehensive unit tests.

- **Obsidian Knowledge Base**: Record problem-solving ideas, complexity analysis, and study notes in the `docs` directory to build a personal algorithm encyclopedia.

- **AI Empowerment**: Planned intelligent translation and RAG-based Q&A assistant to enhance learning efficiency using modern technology.

## 🚀 Quick Start

### Environment Dependencies

- **Go**: 1.25+
- **Obsidian**: Recommended for the best documentation reading experience.

### Running and Testing

1. **Clone the Repository**

```bash
$ git clone https://github.com/MorePeanuts/algorithm.git
$ cd algorithm
```

2. **Run Tests** Each LeetCode problem or CLRS example comes with test files:

```bash
# Test a specific problem
$ go test ./leetcode/0001-0100/0001_Two_Sum/...

# Run all tests
$ go test ./...
```

## 📁 Repository Structure

```plain
.
├── clrs/   # Code implementations for "Introduction to Algorithms, 4th Edition"
├── leetcode/                    # LeetCode problem-solving code
│   └── 0001-0100/               # Grouped by problem number range
│       └── 0001_Two_Sum/
│           ├── solution.go      # Core algorithm
│           └── solution_test.go # Unit tests
├── docs/                        # Obsidian documentation library root
│   ├── CLRS/     # Study notes and exercise solutions for each chapter of CLRS
│   └── leetcode/ # Problem descriptions and solution analyses (linked to source code)
├── cmd/   # Self-developed auxiliary tools (crawlers, translation scripts, etc.)
├── go.mod # Go module configuration
└── README.md
```

Documents are stored in the `docs/` directory, following these principles:

1. **Structured**: Named as `problem_number_problem_name.md`, corresponding to the code directory.
2. **Code Linking**: Link to Go source files in Markdown.

## 🛠️ Future Feature Plans

- **Intelligent Chinese-English Documentation Translation Agent**: Develop an LLM Agent to automatically monitor changes in `docs` and translate Chinese notes into English, maintaining bilingual synchronization.

- **RAG-Based Intelligent Q&A Assistant**: Build a knowledge index based on the local `docs` library; implement Q&A combining theory and practice, e.g., "How is a hash table implemented? What related problems are there on LeetCode?"

- **Automatic Problem Crawler**: Develop a Go tool under `cmd/` to automatically generate templates in `docs` and `leetcode` paths by inputting a LeetCode URL.

## 📜 License

This project is licensed under the MIT License. See [LICENSE](./LICENSE) for details.
