package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
	ozon_fintech "ozon-fintech"
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
		h.initLink(api)
	}
}

func (h *Handler) initLink(api *echo.Group) {
	links := api.Group("/tokens")
	{
		links.GET("/:token", h.getBase)
		links.POST("", h.createShort)
	}
}

func (h *Handler) createShort(ctx echo.Context) error {
	input := &ozon_fintech.Link{}
	if err := ctx.Bind(input); err != nil {
		return ctx.JSON(http.StatusBadRequest, NewValidationError("can't bind input link data"))
	}

	if err := service.ValidateBaseURL(input); err != nil {
		return ctx.JSON(http.StatusBadRequest, NewValidationError("validation error"))
	}

	token, err := h.services.CreateShortURL(ctx.Request().Context(), input)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, NewError("something went wrong"))
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}

func (h *Handler) getBase(ctx echo.Context) error {
	input := &ozon_fintech.Link{}
	input.Token = ctx.Param("token")

	if err := service.ValidateToken(input); err != nil {
		return ctx.JSON(http.StatusBadRequest, NewValidationError("validation error"))
	}

	baseURL, err := h.services.GetBaseURL(ctx.Request().Context(), input)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, NewError("something went wrong"))
	}

	if baseURL == "" {
		return ctx.JSON(http.StatusNotFound, NewError("not such token"))
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"baseURL": baseURL,
	})
}
