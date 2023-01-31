package handler

import (
	"github.com/gin-gonic/gin"
)


// Ping .
func Ping(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "pong",
	})
}