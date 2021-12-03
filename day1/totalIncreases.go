package day1

import (
	"advent-of-code-2021/utils"
)

var input = utils.GetFileData("./day1/input.txt")

func TotalIncreases() int {
	totalIncreases := 0
	previousValue := -1
	for _, line := range input {
		intValue := utils.ConvertToInt(line)
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
			valueToAdd := utils.ConvertToInt(input[i+j])
			totalSum += valueToAdd
		}
		if totalSum > previousValue {
			totalIncreases++
		}
		previousValue = totalSum	
	}
	return totalIncreases
}