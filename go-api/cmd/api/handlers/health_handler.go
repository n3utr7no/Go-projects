package handlers

import (
	"net/http"

	"github.com/labstack/echo"
)

func HealthCheckerHandler(c echo.Context) error {
	return c.String(http.StatusOK, "OK!")
}
