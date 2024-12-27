package main

import "fmt"

type Board [9][9]int

func isValid(board *Board, row, col, num int) bool {
	for i := 0; i < 9; i++ {
		// Check the row, column, and 3x3 sub-grid
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
func main() {
	board := Board{
		{0, 0, 0, 1, 0, 0, 3, 0, 0},
		{0, 8, 2, 0, 6, 0, 4, 7, 0},
		{0, 0, 0, 0, 4, 0, 0, 1, 0},
		{0, 2, 0, 4, 0, 1, 7, 0, 8},
		{8, 5, 0, 7, 0, 0, 2, 0, 0},
		{7, 0, 3, 9, 8, 2, 0, 0, 0},
		{2, 7, 0, 5, 9, 4, 0, 0, 3},
		{4, 0, 6, 8, 1, 0, 5, 0, 7},
		{3, 0, 0, 0, 2, 0, 0, 9, 0},
	}

	if dfs(&board) {
		fmt.Println("Solved Sudoku:")
		for _, row := range board {
			fmt.Println(row)
		}
	} else {
		fmt.Println("The Sudoku cannot be solved.")
	}
}
