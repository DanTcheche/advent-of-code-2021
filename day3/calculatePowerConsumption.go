package day3

import (
	"advent-of-code-2021/utils"
	"strconv"
)

var input = utils.GetFileData("./day3/input.txt")

func CalculatePowerConsumption() int64 {
	var amountOfZeros = getAmountOfZeros()
	totalLines := len(input)
	mostCommonValues := getMostCommonValues(amountOfZeros, totalLines)
	gammaRate := getGammaRateString(mostCommonValues)
	epsilonRate := getEpsilonRateString(mostCommonValues)
	return utils.BinaryToDecimal(gammaRate) * utils.BinaryToDecimal(epsilonRate)
}

func getGammaRateString(mostCommonValues []int) string {
	gammaRate := ""
	for i := 0; i < len(mostCommonValues); i++ {
		gammaRate += strconv.Itoa(mostCommonValues[i])
	}
	return gammaRate
}

func getEpsilonRateString(mostCommonValues []int) string{
	epsilonRate := ""
	for i := 0; i < len(mostCommonValues); i++ {
		if(mostCommonValues[i] == 0 ) {
			epsilonRate += "1"
		} else {
			epsilonRate += "0"
		}
	}
	return epsilonRate
}

func getAmountOfZeros() []int{
	amountOfZeros := make([]int, len(input[0]))
	for _, line := range input {
		for i := 0; i < len(line); i++ {
			if(line[i] == '0') {
				amountOfZeros[i] += 1
			}
		}
	}
	return amountOfZeros
}

func getMostCommonValues(amountOfZeros []int, totalLines int) []int{
	mostCommonValues := make([]int, len(amountOfZeros))
	for i := 0; i < len(amountOfZeros); i++ {
		if(amountOfZeros[i] > totalLines/2) {
			mostCommonValues[i] = 0
		} else {
			mostCommonValues[i] = 1
		}
	}
	return mostCommonValues
}