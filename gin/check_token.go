package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type WebhookPayload struct {
	RequestID string `json:"request_id" binding:"required"`
	Status    string `json:"status" binding:"required"`
	Info      string `json:"info"`
}

const (
	TOKEN = "abc"
)

func main() {

	router := gin.New()

	// router.Use()
	router.POST("/v1/webhook", func(ctx *gin.Context) {

		token := ctx.Request.Header.Get("Authorization")
		fmt.Println("token:", token)

		if token != TOKEN {
			ctx.JSON(http.StatusBadRequest, "Invalid token")
			return
		}

		var json WebhookPayload
		if err := ctx.ShouldBindJSON(&json); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payload"})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"status": "OK",
		})
	})

	router.Run(":8000")
}
