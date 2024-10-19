package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jamadeu/api-cars/schemas"
)

func CreateCarHandler(ctx *gin.Context) {
	request := CreateCarRequest{}
	ctx.BindJSON(&request)
	if err := request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}
	car := schemas.Car{
		CarModel:     request.CarModel,
		Manufacturer: request.Manufacturer,
		Color:        request.Color,
		Value:        request.Value,
	}
	if err := db.Create(&car).Error; err != nil {
		logger.Errorf("error creating car: %v", err)
		sendError(ctx, http.StatusInternalServerError, "creating car on database")
		return
	}
	sendSuccess(ctx, "create-car", car)
}
