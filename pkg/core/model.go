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

type challengeHandler =
// if httpCode == 0 then we'll use http.StatusOK
func(bodyString string) (responseBody string, httpCode int)

func defaultChallengeHandler(bodyString string) (string, int) {
	return fmt.Sprintf("not implemented %s", bodyString), http.StatusNotImplemented
}

func ChallengeHandlerWrapped(handler challengeHandler) func(echo.Context) error {
	return func(c echo.Context) error {
		reqBodyBytes, ok := ioutil.ReadAll(c.Request().Body)
		if ok != nil {
			return c.String(http.StatusBadRequest, "Can't read request body")
		}
		reqBodyString := string(reqBodyBytes)

		responseBody, httpCode := handler(reqBodyString)
		if httpCode == 0 {
			httpCode = http.StatusOK
		}

		return c.String(httpCode, responseBody)
	}
}
