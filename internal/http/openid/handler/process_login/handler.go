package process_login

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func Handle(c echo.Context) error {
	err := c.Request().ParseForm()
	if err != nil {
		return err
	}
	username := c.Request().PostFormValue("username")
	password := c.Request().PostFormValue("password")
	return c.JSON(http.StatusOK, map[string]string{"username": username, "password": password})
}
