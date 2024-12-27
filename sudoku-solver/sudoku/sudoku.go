package sudoku

type Board [9][9]int

func isValid(board *Board, row, col, num int) bool {
	for i := 0; i < 9; i++ {
		if board[row][i] == num || board[i][col] == num ||
			board[row/3*3+i/3][col/3*3+i%3] == num {
			return false
		}
	}
	return true
}

func dfs(board *Board) bool {
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			if board[row][col] == 0 {
				for num := 1; num <= 9; num++ {
					if isValid(board, row, col, num) {
						board[row][col] = num
						if dfs(board) {
							return true
						}
						board[row][col] = 0
					}
				}
				return false
			}
		}
	}
	return true
}

func SolveSudoku(board *Board) bool {
	return dfs(board)
}
