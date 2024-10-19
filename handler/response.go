package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jamadeu/api-cars/schemas"
)

func sendError(ctx *gin.Context, code int, msg string) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(code, gin.H{
		"message":   msg,
		"errorCode": code,
	})
}

func sendSuccess(ctx *gin.Context, op string, data interface{}) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("operation from handler: %s successfull", op),
		"data":    data,
	})
}

type ErrorResponse struct {
	Message   string `json:"message"`
	ErrorCode string `json:"errorCode"`
}

type CreateCarResponse struct {
	Message string              `json:"message"`
	Data    schemas.CarResponse `json:"data"`
}

type DeleteCarResponse struct {
	Message string              `json:"message"`
	Data    schemas.CarResponse `json:"data"`
}

type ShowCarResponse struct {
	Message string              `json:"message"`
	Data    schemas.CarResponse `json:"data"`
}

type ListCarsResponse struct {
	Message string                `json:"message"`
	Data    []schemas.CarResponse `json:"data"`
}

type UpdateCarResponse struct {
	Message string              `json:"message"`
	Data    schemas.CarResponse `json:"data"`
}
