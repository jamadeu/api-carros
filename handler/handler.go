package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ShowCarHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "POST car",
	})
}
func ListCarsHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "POST car",
	})
}

func CreateCarHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "POST car",
	})
}

func UpdateCarHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "POST car",
	})
}

func DeleteCarHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "POST car",
	})
}
