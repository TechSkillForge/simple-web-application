package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Item struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var items = []Item{}
var nextID = 1

func main() {
	router := gin.Default()

	// Ping test
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// CRUD
	router.GET("/items", getItems)
	router.POST("/items", createItem)
	router.GET("/items/:id", getItem)
	router.PUT("/items/:id", updateItem)
	router.DELETE("/items/:id", deleteItem)

	router.Run(":8080")
}
