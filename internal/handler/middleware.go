package handler

import (
	"os"

	"github.com/gin-gonic/gin"
)

func Middleware(ctx *gin.Context) {
	ctx.Set("GEMINI_API_KEY", os.ExpandEnv("GEMINI_API_KEY"))
	ctx.Next()
}
