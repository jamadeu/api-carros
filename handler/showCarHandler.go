package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jamadeu/api-cars/schemas"
)

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
