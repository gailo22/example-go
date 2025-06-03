package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	server := gin.Default()

	server.GET("/api/cloud/key", func(ctx *gin.Context) {
		// decodedBytes, err := base64.StdEncoding.DecodeString("4fbb8c4ac64c4e66")
		// if err != nil {
		// 	fmt.Println("Error decoding:", err)
		// 	return
		// }
		// fmt.Println(decodedBytes)
		// ctx.Data(http.StatusOK, "text/plain", []byte("4fbb8c4ac64c4e66"))
		ctx.Data(http.StatusOK, "binary/octet-stream", []byte("4fbb8c4ac64c4e66"))
	})

	server.Run(":3333")
}
