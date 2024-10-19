package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateCarHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "POST car",
	})
}
