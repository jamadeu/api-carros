package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jamadeu/api-cars/schemas"
)

// @BasePath /api/v1

// @Summary Show car
// @Descriptaion Show a car
// @Tags Cars
// @Accept json
// @Produce json
// @Param id query string true "Car identification"
// @Success 200 {object} ShowCarResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /car [get]
func ShowCarHandler(ctx *gin.Context) {
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
	sendSuccess(ctx, "show-car", car)
}
