package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jamadeu/api-cars/schemas"
)

// @BasePath /api/v1

// @Summary Update car
// @Descriptaion Update a car
// @Tags Cars
// @Accept json
// @Produce json
// @Param id query string true "Car identification"
// @Param request body UpdateCarRequest true "Car data to Update"
// @Success 200 {object} UpdateCarResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /car [put]
func UpdateCarHandler(ctx *gin.Context) {
	request := UpdateCarRequest{}
	ctx.BindJSON(&request)
	if err := request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id",
			"queryParameter").Error())
		return
	}
	car := schemas.Car{}
	if err := db.First(&car, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound,
			fmt.Sprintf("car with id: %s not found", id))
		return
	}
	// Update car
	if request.CarModel != "" {
		car.CarModel = request.CarModel
	}
	if request.Manufacturer != "" {
		car.Manufacturer = request.Manufacturer
	}
	if request.Color != "" {
		car.Color = request.Color
	}
	if request.Value > 0 {
		car.Value = request.Value
	}
	if err := db.Save(&car).Error; err != nil {
		logger.Errorf("error updating car: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error updating car")
		return
	}
	sendSuccess(ctx, "update-car", car)

}
