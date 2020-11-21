package main

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sesquipedalian-dev/advent2019/pkg/core"
)

func main() {
	days := core.ConfigureDays()

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/days", daysHandler(days))

	// Start server
	e.Logger.Fatal(e.Start(":9999"))
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func daysHandler(days *core.Days) func(echo.Context) error {
	return func(c echo.Context) error {
		dayIds := make([]string, len(days.Days))
		for _, day := range days.Days {
			dayIds = append(dayIds, day.ID)
		}

		jsonBytes, err := json.Marshal(dayIds)
		if err != nil {
			return c.String(http.StatusInternalServerError, "Can't marshal response!")
		}

		jsonStr := string(jsonBytes)
		return c.String(http.StatusOK, jsonStr)
	}
}
