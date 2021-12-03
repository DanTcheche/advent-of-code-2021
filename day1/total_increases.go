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
				totalIncreases++
			}
		}
		previousValue = intValue
	}
	return totalIncreases
}

func TotalGroupedIncreases(groupSize int) int {
	totalIncreases := 0
	previousValue := -1
	for i := 1; i < len(input) - (groupSize-1) ; i++ {
		totalSum := 0
		for j := 0; j < groupSize; j++ {
			valueToAdd, err := strconv.Atoi(input[i+j])
			if err != nil {
				panic("Invalid input")
			}
			totalSum += valueToAdd
		}
		if totalSum > previousValue {
			totalIncreases++
		}
		previousValue = totalSum	
	}
	return totalIncreases
}