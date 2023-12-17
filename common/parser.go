package common

import (
	"errors"
	"strconv"
	"strings"
)

var parseError = errors.New("Need 2 arguments, first is a number, second is char range [b, k, m, g, t]")

func parseValue(s string) (int, error) {
	sSplited := strings.Split(s, " ")
	if len(sSplited) != 2 {
		return 0, parseError
	}

	intVal, err := strconv.Atoi(sSplited[0])
	if err != nil {
		return 0, parseError
	}
	mesure, err := convertMesure(sSplited[1])
	if err != nil {
		return 0, parseError
	}

	return intVal * mesure, nil
}

func convertMesure(s string) (int, error) {
	switch s {
	case "b":
		return 1, nil
	case "k":
		return 1024, nil
	case "m":
		return 1024 * 1024, nil
	case "g":
		return 1024 * 1024 * 1024, nil
	case "t":
		return 1024 * 1024 * 1024 * 1024, nil
	default:
		return 0, parseError
	}
}
