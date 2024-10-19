package router

import (
	"github.com/gin-gonic/gin"
	"github.com/jamadeu/api-cars/docs"
	"github.com/jamadeu/api-cars/handler"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func initializeRoutes(router *gin.Engine) {
	// Initialize handler
	handler.InitializeHandler()
	basePath := "/api/v1"
	docs.SwaggerInfo.BasePath = basePath
	v1 := router.Group(basePath)
	{
		v1.GET("/car", handler.ShowCarHandler)
		v1.POST("/car", handler.CreateCarHandler)
		v1.DELETE("/car", handler.DeleteCarHandler)
		v1.PUT("/car", handler.UpdateCarHandler)
		v1.GET("/cars", handler.ListCarsHandler)
	}
	// Initialize Swagger
	router.GET("/swagger/*any",
		ginSwagger.WrapHandler(swaggerfiles.Handler))
}
