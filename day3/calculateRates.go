package day3

import (
	"advent-of-code-2021/utils"
)

var input = utils.GetFileData("./day3/input.txt")

func GetRates() int64 {
	var amountOfZeros = getAmountOfZeros()
	totalLines := len(input)
	gammaRate := getGammaRateString(amountOfZeros, totalLines)
	epsilonRate := getEpsilonRateString(gammaRate)
	return utils.BinaryToDecimal(gammaRate) * utils.BinaryToDecimal(epsilonRate)
}

func getGammaRateString(amountOfZeros []int, totalLines int) string {
	gammaRate := ""
	for i := 0; i < len(amountOfZeros); i++ {
		if(amountOfZeros[i] > totalLines/2) {
			gammaRate += "0"
		} else {
			gammaRate += "1"
		}
	}
	return gammaRate
}

func getEpsilonRateString(gammaRate string) string{
	epsilonRate := ""
	for i := 0; i < len(gammaRate); i++ {
		if(gammaRate[i] == '0' ) {
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