package handlers

import (
	"errors"
	"net/http"
	"test/internal/service"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CheckLink(ctx *gin.Context) {
	link := ctx.Query("link")
	statusCode, err := h.service.Check(link)
	if err != nil {
		if errors.Is(err, service.ErrRedirectPage) {
			ctx.JSON(301, link)
			return
		}
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
	}
	ctx.JSON(statusCode, "Success!")
}
