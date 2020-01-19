package handler

import (
	"github.com/labstack/echo"
	"net/http"
)

func Index() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World")
	}
}
