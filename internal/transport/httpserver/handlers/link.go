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
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	if err = h.service.Add(request.Active_link, request.History_link); err != nil {
		log.Printf("error while add data to redis: %v", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
}

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

func (h *Handler) GetLinkByID(ctx *gin.Context) {
	id := ctx.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		log.Println("Error while converting id", err)
		ctx.AbortWithError(http.StatusNotFound, err)
		return
	}
	link, err := h.service.GetLinkByID(idInt)
	if err != nil {
		log.Println("Error while get link by ID", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	value, ok, err := h.service.Get(link.Active_link)
	if !ok || err != nil {
		log.Println("Error while get link from redis", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	if len(value) != 0 {
		ctx.JSON(http.StatusOK, value)
	}

	ctx.JSON(http.StatusOK, link)
}

func (h *Handler) UpdateLink(ctx *gin.Context) {
	id := ctx.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		log.Println("Error while converting", err)
		ctx.AbortWithError(http.StatusNotFound, err)
		return
	}
	link, err := h.service.GetLinkByID(idInt)
	if err != nil {
		log.Println("Error while get link by ID", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	if err := h.service.UpdateLink(idInt, *link); err != nil {
		log.Println("Error while update link", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
}

func (h *Handler) DeleteLinkById(ctx *gin.Context) {
	id := ctx.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		log.Println("Error while converting", err)
		ctx.AbortWithError(http.StatusNotFound, err)
		return
	}
	if err := h.service.DeleteLinkById(idInt); err != nil {
		log.Println("Error while delete link", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
}
