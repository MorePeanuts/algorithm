package leetcode0036

func isValidSudoku2(board [][]byte) bool {
	var row, col, squ int
	for i := range 9 {
		for j := range 9 {
			if board[i][j] != '.' && (row&(1<<(board[i][j]-'.'))) != 0 {
				return false
			}
			if board[j][i] != '.' && (col&(1<<(board[j][i]-'.'))) != 0 {
				return false
			}
			if board[j/3+3*(i/3)][j%3+3*(i%3)] != '.' && (squ&(1<<(board[j/3+3*(i/3)][j%3+3*(i%3)]-'.'))) != 0 {
				return false
			}
			row |= (1 << (board[i][j] - '.'))
			col |= (1 << (board[j][i] - '.'))
			squ |= (1 << (board[j/3+3*(i/3)][j%3+3*(i%3)] - '.'))
		}
		row, col, squ = 0, 0, 0
	}
	return true
}
