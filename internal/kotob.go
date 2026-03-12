package internal

import (
    "github.com/gin-gonic/gin"
	"kotob_server/internal/handler"
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
