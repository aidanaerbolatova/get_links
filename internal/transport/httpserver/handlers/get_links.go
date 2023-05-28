package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"test/internal/models"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetLink(ctx *gin.Context) {
	page := ctx.Param("page")
	if len(page) == 0 {
		page = "1"
	}
	pageInt, err := strconv.Atoi(page)
	if err != nil || pageInt < 1 || strings.TrimSpace(page) == "" {
		ctx.JSON(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}
	links, err := h.service.GetLinks(pageInt)
	if err != nil {
		log.Println("Error while get links", err)
		ctx.JSON(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}
	_, err = json.Marshal(links)
	if err != nil {
		log.Println("Error while marshaling", err)
		ctx.JSON(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}
	ctx.JSON(http.StatusOK, links)
}

func (h *Handler) CreateLink(ctx *gin.Context) {
	var request models.Data
	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Println("Error while reding create link request body")
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	err = json.Unmarshal(body, &request)
	if err != nil {
		log.Println("Error while unmarshaling", err)
	}
}
