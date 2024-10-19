package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Initialize() {
	// Initializing router
	router := gin.Default()

	// Define route
	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// Start server
	router.Run(":8080")
}
