package leetcode0212

type TrieNode struct {
	end  bool
	next [26]*TrieNode
}

func (t *TrieNode) Insert(word string) {
	for _, r := range word {
		if t.next[r-'a'] == nil {
			t.next[r-'a'] = &TrieNode{}
		}
		t = t.next[r-'a']
	}
	t.end = true
}

func (t *TrieNode) Find(word string) (found, end bool) {
	for _, r := range word {
		if t.next[r-'a'] == nil {
			return false, true
		}
		t = t.next[r-'a']
	}
	defer func() {
		t.end = false
	}()
	return t.end, false
}

func findWords3(board [][]byte, words []string) []string {
	res := make([]string, 0)
	trie := &TrieNode{}
	for _, word := range words {
		trie.Insert(word)
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
		if oor(i, j) || visit[i][j] {
			return
		}
		*bf = append(*bf, board[i][j])
		visit[i][j] = true
		s := string(*bf)
		found, end := trie.Find(s)
		if end {
			*bf = (*bf)[:len(*bf)-1]
			visit[i][j] = false
			return
		}
		if found {
			res = append(res, s)
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
