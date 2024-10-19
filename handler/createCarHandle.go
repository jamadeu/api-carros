package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jamadeu/api-cars/schemas"
)

// @BasePath /api/v1

// @Summary Create car
// @Descriptaion Create a new car
// @Tags Cars
// @Accept json
// @Produce json
// @Param request body CreateCarRequest true "Request body"
// @Success 200 {object} CreateCarResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /car [post]
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
