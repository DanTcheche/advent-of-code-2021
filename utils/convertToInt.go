package utils

import (
	"strconv"
)

func ConvertToInt(value string) int {
	intValue, err := strconv.Atoi(value)
	if err != nil {
		panic("Invalid input")
	}
	return intValue
}