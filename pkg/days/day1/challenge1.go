package day1

import (
	"bufio"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func challenge1Handler(reqBody string) (responseBody string, httpCode int) {
	runningMass := 0

	sc := bufio.NewScanner(strings.NewReader(reqBody))
	for sc.Scan() {
		nextMassString := sc.Text()
		nextMass, err := strconv.Atoi(nextMassString)
		if err != nil {
			return fmt.Sprintf("Invalid line %s", nextMassString), http.StatusBadRequest
		}
		fmt.Printf("Handling input line %s; %d + (%d / 3) - 2\n", nextMassString, runningMass, nextMass)
		runningMass = runningMass + (nextMass / 3) - 2
	}

	return strconv.Itoa(runningMass), 0
}
