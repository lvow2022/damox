package main

import "github.com/gin-gonic/gin"

func main() {
	server := InitWebServer()
	server.GET("/hello", func(ctx *gin.Context) {
		ctx.JSON(200, "hello world")
	})
	server.Run(":8081")
}
