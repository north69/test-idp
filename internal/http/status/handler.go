package status

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func Handler(c echo.Context) error {
	status := Status{Status: "ok"}
	return c.JSON(http.StatusOK, status)
}
