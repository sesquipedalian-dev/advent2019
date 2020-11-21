package core

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Days struct {
	Days map[string]Day
}

type Day struct {
	ID         string
	Challenges map[string]Challenge
}

type Challenge struct {
	ID      string
	Handler func(echo.Context) error
}

func ConfigureDays() *Days {
	return &Days{
		Days: map[string]Day{
			"1": Day{
				ID: "1",
				Challenges: map[string]Challenge{
					"1": Challenge{
						ID:      "1",
						Handler: defaultHandler,
					},
				},
			},
		},
	}
}

func defaultHandler(c echo.Context) error {
	reqBodyBytes, ok := ioutil.ReadAll(c.Request().Body)
	if ok != nil {
		return c.String(http.StatusBadRequest, "Can't read request body")
	}
	reqBodyString := string(reqBodyBytes)
	return c.String(http.StatusNotImplemented,
		fmt.Sprintf("Not Implemented %s", reqBodyString))
}
