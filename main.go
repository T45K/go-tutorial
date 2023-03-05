package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"reflect"
	"strconv"
)

func main() {
	route := gin.Default()
	route.GET("/ping", handlePing)
	route.GET("/query", handleQuery)
	route.GET("/param/:name/:id", handlePathParam)
	route.POST("/ping", receivePing)
	route.POST("/ping2", receivePing2)
	route.Run()
}

func handlePing(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func handleQuery(c *gin.Context) {
	name := c.Query("name")
	unknown := c.Query("unknown")
	c.JSON(200, gin.H{"name": name, "unknown": unknown})
}

func handlePathParam(c *gin.Context) {
	name := c.Param("name")
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatus(400)
		return
	}
	c.JSON(200, gin.H{"name": name, "id": id})
}

func receivePing(c *gin.Context) {
	var body PingRequestBody
	if err := c.ShouldBindJSON(&body); err != nil {
		c.AbortWithError(400, err)
		return
	}
	fmt.Println(body)
	c.JSON(200, gin.H{"message": "success"})
}

type PingRequestBody struct {
	Name string
	Id   int
}

// もう一つの方法
func receivePing2(c *gin.Context) {
	var body map[string]interface{}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.AbortWithError(400, err)
		return
	}
	fmt.Println(body)
	fmt.Println(reflect.TypeOf(body["name"]))
	c.JSON(200, gin.H{"message": "success"})
}
