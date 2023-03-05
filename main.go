package main

import "github.com/gin-gonic/gin"

func main() {
	route := gin.Default()
	route.GET("/ping", handlePing)
	route.Run()
}

func handlePing(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
