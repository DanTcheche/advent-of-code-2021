package main

import (
	"advent-of-code-2021/day1"
	"advent-of-code-2021/day2"
	"advent-of-code-2021/day3"
	"fmt"
)

func main() {
	fmt.Println(day1.TotalIncreases())
	fmt.Println(day1.TotalGroupedIncreases(3))
	fmt.Println(day2.CalculateFinalPosition())
	fmt.Println(day2.CalculateFinalPositionWithAim())
	fmt.Println(day3.GetRates())
}
