package internal

import (
	"kotob_server/internal/handler"

	"github.com/gin-gonic/gin"
)

func KotobRouter() {
	client := gin.Default()
	client.Use(handler.Middleware)
	client.GET("/translate", handler.Translate)

    port := ":8080"
    if os.ExpandEnv("KOTOB_PORT") != "" {
        port = os.ExpandEnv("KOTOB_PORT")
    }
	err := client.Run(port)
	if err != nil {
		return
	}
}
