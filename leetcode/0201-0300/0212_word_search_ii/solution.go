// Package leetcode0212 solves LeetCode 212. Word Search II
package leetcode0212

func findWords(board [][]byte, words []string) []string {
	res := []string{}
	visit := make([][]bool, len(board))
	for i := range visit {
		visit[i] = make([]bool, len(board[0]))
	}

	oor := func(i, j int) bool {
		if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) {
			return true
		}
		return false
	}

	var dfs func(i, j int, word string) bool
	dfs = func(i, j int, word string) bool {
		if len(word) == 0 {
			return true
		}
		if oor(i, j) || visit[i][j] || board[i][j] != word[0] {
			return false
		}
		visit[i][j] = true
		if dfs(i-1, j, word[1:]) || dfs(i+1, j, word[1:]) || dfs(i, j-1, word[1:]) || dfs(i, j+1, word[1:]) {
			visit[i][j] = false
			return true
		}
		visit[i][j] = false
		return false
	}

	for _, word := range words {
	outer:
		for i := range board {
			for j := range board[i] {
				if dfs(i, j, word) {
					res = append(res, word)
					break outer
				}
			}
		}
	}
	return res
}
