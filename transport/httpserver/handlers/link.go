package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

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
	if err := h.service.UpdateLink(idInt, link); err != nil {
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
