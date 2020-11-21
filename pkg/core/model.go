package core

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Days struct {
	Days []Day
}

type Day struct {
	ID         string
	Challenges []Challenge
}

type Challenge struct {
	ID      string
	Handler func(echo.Context) error
}

func ConfigureDays() *Days {
	return &Days{}
}

func defaultHandler(c echo.Context) error {
	return c.String(http.StatusNotImplemented, "Not Implemented")
}
