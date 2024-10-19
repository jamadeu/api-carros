package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jamadeu/api-cars/schemas"
)

func ListCarsHandler(ctx *gin.Context) {
	cars := []schemas.Car{}

	if err := db.Find(&cars).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error to find cars")
		return
	}
	sendSuccess(ctx, "list-cars", cars)
}
