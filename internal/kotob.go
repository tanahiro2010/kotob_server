package internal

import (
	"kotob_server/internal/handler"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func KotobRouter() {
	err := godotenv.Load()
	if err != nil {
		println("No .env file found, relying on environment variables")
	}

	client := gin.Default()
	client.Use(handler.Middleware)
	client.GET("/api/translate", handler.Translate)
	port := ":8080"
	if os.Getenv("KOTOB_PORT") != "" {
		port = os.Getenv("KOTOB_PORT")
	}
	println("Server is running on port", port)
	err = client.Run(port)
	if err != nil {
		return
	}
}
