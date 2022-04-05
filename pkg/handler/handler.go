package handler

import (
	"github.com/labstack/echo/v4"
	"ozon-fintech/pkg/service"
)

type Handler struct {
	services service.Services
}

func NewHandler(services service.Services) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) InitRotes(router *echo.Echo) {
	api := router.Group("/api")
	{
		links := api.Group("/tokens")
		{
			links.GET("/:token", h.getBase)
			links.POST("", h.createShort)
		}
	}
}
