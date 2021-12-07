package main

import (
	_ "embed"
	"fmt"
	"math"
	"strconv"
	"strings"
)

//go:embed inputs.txt
var rawInputs string
var inputs []int
var min, max int

// Prepare the rawInputs into a format more desirable for the problem at hand.
func init() {
	data := strings.Split(strings.TrimSpace(rawInputs), ",")
	for _, val := range data {
		num, err := strconv.Atoi(val)
		if err != nil {
			panic(err)
		}
		inputs = append(inputs, num)
		if num > max {
			max = num
		}
		if num < min {
			min = num
		}
	}
}

// Not particularly proud of this solution but it works. Very sure there's a more efficient way to solve this using
// some mathematics
func main() {
	fmt.Println("Part 1 Answer: ", PartOne(inputs))
	fmt.Println("Part 2 Answer: ", PartTwo(inputs))
}

func PartOne(values []int) int {
	minConsumption := math.MaxInt

	for i := min; i <= max; i++ {
		var fuelConsumption int
		for _, val := range values {
			fuelConsumption += abs(val - i)
		}
		if fuelConsumption < minConsumption {
			minConsumption = fuelConsumption
		}
	}

	return minConsumption
}

func PartTwo(values []int) int {
	minConsumption := math.MaxInt
	for pointToMoveTo := min; pointToMoveTo <= max; pointToMoveTo++ {
		var fuelConsumption int
		for _, val := range values {
			distance := abs(val - pointToMoveTo)
			fuelConsumption += calculateFuel(distance)
		}
		if fuelConsumption < minConsumption {
			minConsumption = fuelConsumption
		}
	}
	return minConsumption
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func calculateFuel(distance int) int {
	var consumption int
	for i := 0; i <= distance; i++ {
		consumption += i
	}
	return consumption
}
