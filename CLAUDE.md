# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

A data structures and algorithms learning repository combining LeetCode practice and "Introduction to Algorithms (CLRS, 4th Edition)". Primary implementation language is Go, with Python tools for automation.

## Common Commands

### Go Tests
```bash
# Run all tests
go test ./...

# Test a specific problem
go test ./leetcode/0001-0100/0001_Two_Sum/...

# Test with verbose output
go test -v ./leetcode/0001-0100/0001_Two_Sum/...
```

### Python Tools (uv workspace)
```bash
# Sync workspace dependencies
uv sync

# Run a specific tool
uv run test-gen
uv run auto-docs
uv run algo-chat
```

## Architecture

### LeetCode Solutions
- Location: `leetcode/XXXX-XXXX/XXXX_Problem_Name/`
- Each problem has `solution.go` and `solution_test.go`
- Multiple solution variants can exist (e.g., `groupAnagrams`, `groupAnagrams2`, `groupAnagrams3`)
- Test cases are divided into "examples" (from problem description) and "cases" (comprehensive tests)
- Test functions follow naming: `TestProblemNameExamples`, `TestProblemNameCases`

### Shared Utilities
- `leetcode/utils/testing_utils.go`: Helper functions for test assertions (e.g., `MatchTwo2dStringSlice` for order-independent slice comparison)

### Python Tools Workspace
- Root `pyproject.toml` defines uv workspace with members in `python/*`
- `python/test-gen/`: Automated unit test generation
- `python/auto-docs/`: Documentation automation
- `python/algo-chat/`: RAG-based Q&A assistant (planned)

### Documentation
- `docs/`: Obsidian-based knowledge base with bidirectional linking
- `docs/leetcode/`: Problem descriptions and solution analyses
- `docs/CLRS/`: Study notes for "Introduction to Algorithms"
