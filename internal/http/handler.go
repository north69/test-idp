package http

import (
	"github.com/labstack/echo/v4"
	"test-idp/internal/config"
	"test-idp/internal/http/openid/handler/configuration"
	"test-idp/internal/http/status"
)

type Handler struct {
	server *echo.Echo
	cnf    *config.Config
}

func NewHandler(s *echo.Echo, cnf *config.Config) *Handler {
	return &Handler{
		server: s,
		cnf:    cnf,
	}
}

func (h *Handler) InitRoutes() {
	h.server.GET("/status", status.Handler)
	h.server.GET("/auth/openid/.well-known/openid-configuration", configuration.Init(h.cnf).Handle)
	h.server.File("/auth/login", "web/pages/login_page.html")
}
