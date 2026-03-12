package handler

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/kotob-project/kotob/pkg/translate"
)

func Translate(ctx *gin.Context) {
	content := ctx.Query("content")
	model := ctx.Query("model")
	src := ctx.Query("src")
	dst := ctx.Query("dst")
	if content == "" {
		ctx.JSON(400, gin.H{"error": "Content is required"})
	}
	if model == "" {
		model = "gemini-2.5-flash"
	}
	if src == "" {
		src = "ja"
	}
	if dst == "" {
		dst = "en"
	}

	background := context.Background()
	apiKey, _ := ctx.Get("GEMINI_API_KEY")
	if apiKey == nil {
		ctx.JSON(500, gin.H{"error": "API key not found"})
		return
	}

	client, err := translate.NewClient(background, apiKey.(string), model)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to create translation client"})
		return
	}

	result, err := client.Translate(background, content, src, dst, "")
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Translation failed"})
		return
	}

	ctx.JSON(200, gin.H{"translation": result})
}
