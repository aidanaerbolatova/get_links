package handlers

import (
	"test/internal/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRoute() *gin.Engine {
	router := gin.Default()

	router.GET("/admin/redirects", h.GetLink)
	router.POST("/admin/redirects", h.CreateLink)
	router.GET("/admin/redirects/:id", h.GetLinkByID)
	router.PATCH("/admin/redirects/:id", h.UpdateLink)
	router.DELETE("/admin/redirects/:id", h.DeleteLinkById)

	return router
}
