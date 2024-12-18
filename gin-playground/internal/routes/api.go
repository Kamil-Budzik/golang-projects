package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kamil-budzik/gin-playground/internal/handlers"
)

func (app *app) SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/ping", handlers.PingHandler)

	return r
}
