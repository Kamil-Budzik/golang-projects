package main

import "fmt"

type Board [9][9]int

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

	for _, row := range board {
		for _, num := range row {
			fmt.Printf("%d ", num)
		}
		fmt.Println()
	}
}
