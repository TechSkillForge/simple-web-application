package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// Echo the user value
	r.GET("/user/:name", func(c *gin.Context) {
		user := c.Params.ByName("name")
		greetingMsg := "Hello, " + user + "!"
		c.String(http.StatusOK, greetingMsg)
	})

	return r
}

func main() {
	r := setupRouter()
	r.Run(":8080")
}
