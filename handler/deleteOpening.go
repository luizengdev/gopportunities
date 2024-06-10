package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/luizengdev/gopportunities/schemas"
)

func DeleteOpeningHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}

	opening := schemas.Opening{}
	// Find opening
	if err := db.First(&opening, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("opening with id: %v not found", id))
		return
	}

	// Delete opening
	if err := db.Delete(&opening).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, fmt.Sprintf("error deleting opening: %s", id))
		return
	}

	sendSuccess(ctx, "delete-opening", opening)
}
