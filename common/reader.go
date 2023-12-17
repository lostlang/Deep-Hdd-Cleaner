package common

import (
	"bufio"
	"os"
)

func ReadValue() (int, error) {
	reader := bufio.NewReader(os.Stdin)

	tempStr, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}
	tempStr = tempStr[:len(tempStr)-1]

	ram, err := parseValue(tempStr)
	if err != nil {
		return 0, err
	}

	return ram, nil
}
