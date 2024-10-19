package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jamadeu/api-cars/schemas"
)

// @BasePath /api/v1

// @Summary List cars
// @Descriptaion List cars
// @Tags Cars
// @Accept json
// @Produce json
// @Success 200 {object} ListCarsResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /cars [get]
func ListCarsHandler(ctx *gin.Context) {
	cars := []schemas.Car{}

	if err := db.Find(&cars).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error to find cars")
		return
	}
	sendSuccess(ctx, "list-cars", cars)
}
