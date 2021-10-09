package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)



func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.POST("/ping/post", func(c *gin.Context) {
		json := make(map[string]interface{}) //注意该结构接受的内容
		c.BindJSON(&json)
		fmt.Printf("%v",&json)
		c.JSON(200,json)
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}