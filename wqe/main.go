package main

import "github.com/gin-gonic/gin"

func main() {

	engine := gin.New()
	engine.GET("/ping", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "pong",
		})
	})
	engine.Run(":8080")
}
