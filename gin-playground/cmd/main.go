package main

import (
	"github.com/kamil-budzik/gin-playground/internal/models/sqlite"
	"github.com/kamil-budzik/gin-playground/internal/repositories"
)

type app struct {
	characters *sqlite.CharactersModel
}

func main() {
	db := repositories.InitDB()

	app := app{
		characters: &sqlite.CharactersModel{
			DB: db,
		},
	}

	router := setupRouter()
	router.Run(":8080")

}
