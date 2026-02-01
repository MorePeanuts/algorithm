// Package leetcode0036 solves LeetCode 36. Valid Sudoku
package leetcode0036

func isValidSudoku(board [][]byte) bool {
	row := make(map[byte]struct{})
	col := make(map[byte]struct{})
	squ := make(map[byte]struct{})
	for i := range 9 {
		for j := range 9 {
			if _, exist := row[board[i][j]]; board[i][j] != '.' && exist {
				return false
			}
			if _, exist := col[board[j][i]]; board[j][i] != '.' && exist {
				return false
			}
			if _, exist := squ[board[j/3+3*(i/3)][j%3+3*(i%3)]]; board[j/3+3*(i/3)][j%3+3*(i%3)] != '.' && exist {
				return false
			}
			row[board[i][j]] = struct{}{}
			col[board[j][i]] = struct{}{}
			squ[board[j/3+3*(i/3)][j%3+3*(i%3)]] = struct{}{}
		}
		clear(row)
		clear(col)
		clear(squ)
	}
	return true
}
