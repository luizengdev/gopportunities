package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/luizengdev/gopportunities/schemas"
)

func ListOpeningsHandler(ctx *gin.Context) {
	openings := []schemas.Opening{}
	if err := db.Find(&openings).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, fmt.Sprintf("error listing openings: %v", err))
		return
	}
	sendSuccess(ctx, "list-opening", openings)
}
