package main

import (
	_ "embed"
	"fmt"
	"math/big"
	"strconv"
	"strings"
)

//go:embed inputs.txt
var rawInputs string
var inputs []int

// Prepare the rawInputs into a format more desirable for the problem at hand.
func init() {
	data := strings.Split(strings.TrimSpace(rawInputs), ",")
	for _, val := range data {
		num, err := strconv.Atoi(val)
		if err != nil {
			panic(err)
		}
		inputs = append(inputs, num)
	}
}

func main() {
	fmt.Println("Part 1 Answer: ", partOne())
	fmt.Println("Part 2 Answer: ", partTwo())
}

func partOne() *big.Int {
	return calculatePopulation(inputs, 80)
}

func partTwo() *big.Int {
	return calculatePopulation(inputs, 256)
}

func calculatePopulation(values []int, days int) *big.Int {
	fish := make(map[int]*big.Int, 9)
	for i := 0; i < 9; i++ {
		fish[i] = big.NewInt(0)
	}
	for _, val := range values {
		fish[val].Add(fish[val], big.NewInt(1))
	}

	for i := 0; i < days; i++ {
		fish = map[int]*big.Int{
			0: fish[1],
			1: fish[2],
			2: fish[3],
			3: fish[4],
			4: fish[5],
			5: fish[6],
			6: fish[7].Add(fish[7], fish[0]),
			7: fish[8],
			8: fish[0],
		}
	}

	sum := big.NewInt(0)
	for i := range fish {
		sum.Add(sum, fish[i])
	}
	return sum
}
