package main

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type AddHistoryWebhook struct {
	Type          string `json:"type" binding:"required"`
	Timestamp     string `json:"timestamp" binding:"required"`
	TransactionID string `json:"transaction_id" binding:"required"`
	SSOID         string `json:"ssoid" binding:"required"`
	Status        string `json:"status" binding:"required"`
	FailCode      string `json:"fail_code,omitempty"`
	FailMessage   string `json:"fail_message,omitempty"`
}

const (
	AUTOADD_WEBHOOK_AUTH_TOKEN = "abc"
)

func main() {
	route := gin.Default()
	route.POST("/v1/webhook", handleWebhook)
	route.Run(":8000")
}

func handleWebhook(ctx *gin.Context) {

	err := validateToken(ctx)
	if err != nil {
		return
	}

	var json AddHistoryWebhook
	if err := ctx.ShouldBindJSON(&json); err != nil {
		validateRequiredFields(ctx, json)
		return
	}

	// publish to IPC_ADDED_DEVICE.updated

	ctx.JSON(http.StatusOK, gin.H{
		"code":    10001,
		"message": "Success",
	})

	// update mongo by transaction id
	go func() {
		updateTransaction(json)
	}()

}

func updateTransaction(json AddHistoryWebhook) {
	// update mongo
	time.Sleep(10 * time.Second)
	fmt.Println("update done!")
}

func validateRequiredFields(ctx *gin.Context, json AddHistoryWebhook) {
	if json.Type == "" {
		ctx.JSON(http.StatusBadRequest, errorMessage("type"))
	} else if json.Timestamp == "" {
		ctx.JSON(http.StatusBadRequest, errorMessage("timestamp"))
	} else if json.TransactionID == "" {
		ctx.JSON(http.StatusBadRequest, errorMessage("transaction_id"))
	} else if json.SSOID == "" {
		ctx.JSON(http.StatusBadRequest, errorMessage("ssoid"))
	} else if json.Status == "" {
		ctx.JSON(http.StatusBadRequest, errorMessage("status"))
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payload"})
	}
}

func errorMessage(field string) any {
	return gin.H{
		"error": gin.H{
			"code":    10021,
			"message": fmt.Sprintf("Error Parameter: %v must be specified", field),
		},
	}
}

func validateToken(ctx *gin.Context) error {
	authToken := ctx.Request.Header.Get("Authorization")
	if authToken != AUTOADD_WEBHOOK_AUTH_TOKEN {
		ctx.JSON(http.StatusBadRequest, "Invalid authorization token")
		return errors.New("validation failed")
	}
	return nil
}
