package router

import (
	"github.com/gin-gonic/gin"
	"github.com/jamadeu/api-cars/handler"
)

func initializeRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		v1.GET("/car", handler.ShowCarHandler)
		v1.POST("/car", handler.CreateCarHandler)
		v1.DELETE("/car", handler.DeleteCarHandler)
		v1.PUT("/car", handler.UpdateCarHandler)
		v1.GET("/cars", handler.ListCarsHandler)
	}
}
