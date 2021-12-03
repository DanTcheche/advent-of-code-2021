package day2

import (
	"advent-of-code-2021/utils"
	"strings"
)

var input = utils.GetFileData("./day2/input.txt")

type movement struct {
	direction string
	units int
}

type position struct {
	horizontal int
	depth int
	aim int
}

func newMovement(direction string, depth int) movement{
	return movement{direction, depth}
}

func newPosition(initialHorizontal int, initialDepth int, initalAim int) position{
	return position{initialHorizontal, initialDepth, initalAim}
}

func getMovements() []movement{
	var movements []movement
	for _, line := range input {
		values := strings.Split(line, " ")
		movement := newMovement(values[0], utils.ConvertToInt(values[1]))
		movements = append(movements, movement)
	}
	return movements
}

func initialize() (position, []movement){
	var movements []movement
	position := newPosition(0, 0, 0)
	movements = getMovements()
	return position, movements
}

func CalculateFinalPosition() int {
	position, movements := initialize()
	for _, movement := range movements {
		if movement.direction == "forward" {
			position.horizontal += movement.units
		}
		if movement.direction == "up" {
			position.depth -= movement.units
		}
		if movement.direction == "down" {
			position.depth += movement.units
		}
	}
	return position.depth * position.horizontal
}

func CalculateFinalPositionWithAim() int {
	position, movements := initialize()
	for _, movement := range movements {
		if movement.direction == "forward" {
			position.horizontal += movement.units
			position.depth += position.aim * movement.units
		}
		if movement.direction == "up" {
			position.aim -= movement.units
		}
		if movement.direction == "down" {
			position.aim += movement.units
		}
	}
	return position.depth * position.horizontal
}