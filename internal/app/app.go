package app

import (
	"github.com/labstack/echo/v4"
)

type App struct {
	server *echo.Echo
}

func New() (*App, error) {
	a := &App{}
	a.server = echo.New()
	h := NewHandler(a.server)
	h.InitRoutes()
	return a, nil
}

func (a *App) Run() error {
	err := a.server.Start(":8081")
	if err != nil {
		a.server.Logger.Fatal(err)
	}
	return nil
}
