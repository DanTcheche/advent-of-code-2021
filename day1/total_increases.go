package day1

import (
	"advent-of-code-2021/utils"
	"strconv"
)

var input = utils.GetFileData("./day1/input.txt")

func TotalIncreases() int {
	totalIncreases := 0
	previousValue := -1
	for _, value := range input {
		intValue, err := strconv.Atoi(value)
		if err != nil {
			panic("Invalid input")
		}
		if previousValue > 0 {
			if intValue > previousValue {
				totalIncreases += 1
			}
		}
		previousValue = intValue
	}
	return totalIncreases
}