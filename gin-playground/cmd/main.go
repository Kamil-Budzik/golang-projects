package main

import "github.com/kamil-budzik/gin-playground/internal/routes"

func main() {
	router := routes.SetupRouter()
	router.Run(":8080")
}
