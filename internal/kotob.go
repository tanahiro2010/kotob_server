package internal

import (
	"kotob_server/internal/handler"

	"github.com/gin-gonic/gin"
)

func KotobRouter() {
	client := gin.Default()
	client.Use(handler.Middleware)
	client.GET("/translate", handler.Translate)

	err := client.Run(":8080")
	if err != nil {
		return
	}
}
