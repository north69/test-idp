package app

import (
	"github.com/labstack/echo/v4"
	"test-idp/internal/app/openid/handler/configuration"
	"test-idp/internal/app/status"
)

type Handler struct {
	server *echo.Echo
}

func NewHandler(s *echo.Echo) *Handler {
	return &Handler{
		server: s,
	}
}

func (h *Handler) InitRoutes() {
	h.server.GET("/status", status.Handler)
	h.server.GET("/auth/openid/.well-known/openid-configuration", configuration.Handle)
}