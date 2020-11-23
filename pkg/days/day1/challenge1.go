package day1

import (
	"bufio"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func challenge1Handler(reqBody string) (responseBody string, httpCode int) {
	fuelMass, err := scanBody(reqBody, func(nextMass int) int {
		return (nextMass / 3) - 2
	})
	if err != nil {
		return err.Error(), http.StatusBadRequest
	}
	return strconv.Itoa(fuelMass), 0
}

func challenge2Handler(reqBody string) (responseBody string, httpCode int) {
	fuelMass, err := scanBody(reqBody, func(nextMass int) int {
		return massToFuel(nextMass)
	})
	if err != nil {
		return err.Error(), http.StatusBadRequest
	}
	return strconv.Itoa(fuelMass), 0
}

func scanBody(reqBody string, innerLoop func(int) int) (int, error) {
	runningMass := 0

	sc := bufio.NewScanner(strings.NewReader(reqBody))
	for sc.Scan() {
		nextMassString := sc.Text()
		nextMass, err := strconv.Atoi(nextMassString)
		if err != nil {
			return 0, fmt.Errorf("Invalid line %s", nextMassString)
		}
		runningMass = runningMass + innerLoop(nextMass)
	}

	return runningMass, nil
}

func massToFuel(mass int) int {
	fuel := mass/3 - 2
	if fuel <= 0 {
		return 0
	}

	return fuel + massToFuel(fuel)
}
