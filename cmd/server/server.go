package main

import (
	"encoding/json"
	"fmt"
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
	e.GET("/days/:dayID/challenges", challengesHandler(days))
	e.GET("/days/:dayID/challenges/:challengeID", challengeHandler(days))

	// Start server
	e.Logger.Fatal(e.Start(":9999"))
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func daysHandler(days *core.Days) func(echo.Context) error {
	return func(c echo.Context) error {
		dayIDs := make([]string, len(days.Days))
		for dayID := range days.Days {
			dayIDs = append(dayIDs, dayID)
		}

		jsonBytes, err := json.Marshal(dayIDs)
		if err != nil {
			return c.String(http.StatusInternalServerError, "Can't marshal response!")
		}

		jsonStr := string(jsonBytes)
		return c.String(http.StatusOK, jsonStr)
	}
}

func challengesHandler(days *core.Days) func(echo.Context) error {
	return func(c echo.Context) error {
		dayID := "1" // TODO get it from echo

		foundDay, found := days.Days[dayID]
		if !found {
			return c.String(http.StatusBadRequest, fmt.Sprintf("No day with that ID %s", dayID))
		}

		challengeIDs := make([]string, len(foundDay.Challenges))
		for _, challenge := range foundDay.Challenges {
			challengeIDs = append(challengeIDs, challenge.ID)
		}

		jsonBytes, err := json.Marshal(challengeIDs)
		if err != nil {
			return c.String(http.StatusInternalServerError, "Can't marshal response!")
		}

		jsonStr := string(jsonBytes)
		return c.String(http.StatusOK, jsonStr)
	}
}

func challengeHandler(days *core.Days) func(echo.Context) error {
	return func(c echo.Context) error {
		dayID := "1"       // TODO get it from echo
		challengeID := "1" // TODO get it from echo

		foundDay, found := days.Days[dayID]
		if !found {
			return c.String(http.StatusBadRequest, fmt.Sprintf("No day with that ID %s", dayID))
		}

		foundChallenge, found := foundDay.Challenges[challengeID]
		if !found {
			return c.String(http.StatusBadRequest, fmt.Sprintf("No challenge with that ID %s", dayID))
		}

		return foundChallenge.Handler(c)
	}
}
