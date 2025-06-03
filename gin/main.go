package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Hello struct {
	Message string `json:"message"`
}

func main() {
	server := gin.Default()

	server.GET("/hello", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "OK",
		})
	})

	server.POST("/posts", func(ctx *gin.Context) {
		var hello Hello
		ctx.BindJSON(&hello)

		fmt.Println(hello)

		ctx.JSON(http.StatusOK, hello)
	})

	server.GET("/404", func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "unbinded",
		})
	})

	server.Run(":8000")
}
