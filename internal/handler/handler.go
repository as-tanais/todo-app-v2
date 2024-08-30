package handler

import (
	"github.com/gin-gonic/gin"
	"todo-app-v2/internal/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	items := router.Group("/api/items")

	items.GET("/", h.getAllItems)
	items.PUT("/", h.createItem)
	items.GET("/:id", h.getById)
	items.DELETE("/:id", h.deleteItem)

	return router
}
