package main

import (
	"encoding/json"

	"github.com/gofiber/fiber/v3"
	"github.com/kamil-budzik/sudoku-solver/sudoku"
)

type SolveRequest struct {
	Board [9][9]int `json:"board"`
}

func main() {
	app := fiber.New()

	app.Post("/solve", func(c fiber.Ctx) error {
		var req SolveRequest
		if err := json.Unmarshal(c.Body(), &req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Invalid input. Please send a valid Sudoku grid.",
			})
		}
		var board sudoku.Board

		copy(board[:], req.Board[:])

		if sudoku.SolveSudoku(&board) {
			return c.Status(fiber.StatusOK).JSON(fiber.Map{
				"board": board,
			})
		} else {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Sudoku cannot be solved.",
			})
		}
	})

	app.Listen(":3000")
}
