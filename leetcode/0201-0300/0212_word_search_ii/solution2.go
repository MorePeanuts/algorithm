package leetcode0212

func findWords2(board [][]byte, words []string) []string {
	res := make([]string, 0)
	dictionary := make(map[string]bool)
	maxLength := 0
	for _, word := range words {
		dictionary[word] = true
		maxLength = max(maxLength, len(word))
	}
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

	var dfs func(i, j int, bf *[]byte)
	dfs = func(i, j int, bf *[]byte) {
		if oor(i, j) || visit[i][j] || len(*bf) >= maxLength {
			return
		}
		*bf = append(*bf, board[i][j])
		visit[i][j] = true
		s := string(*bf)
		if _, exist := dictionary[s]; exist {
			res = append(res, s)
			delete(dictionary, s)
		}
		dfs(i-1, j, bf)
		dfs(i+1, j, bf)
		dfs(i, j-1, bf)
		dfs(i, j+1, bf)
		*bf = (*bf)[:len(*bf)-1]
		visit[i][j] = false
	}

	for i := range board {
		for j := range board[i] {
			dfs(i, j, &[]byte{})
		}
	}
	return res
}
