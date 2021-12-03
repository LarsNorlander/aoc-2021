package main

import (
	_ "embed"
	"log"
	"strconv"
	"strings"
)

//go:embed inputs.txt
var rawInputs string
var inputs []int

// Prepare the rawInputs into a format more desirable for the problem at hand
func init() {
	for _, str := range strings.Fields(rawInputs) {
		atoi, err := strconv.Atoi(str)
		if err != nil {
			log.Fatalln(err)
		}
		inputs = append(inputs, atoi)
	}
}

func main() {
	log.Println("Part 1 Answer: ", partOne())
	log.Println("Part 2 Answer: ", partTwo())
}

func partOne() int {
	var increases int
	var prev = inputs[0]
	for i := 1; i < len(inputs); i++ {
		if inputs[i] > prev {
			increases++
		}
		prev = inputs[i]
	}
	return increases
}

func partTwo() int {
	var increases int
	var prev = sum(inputs[0:3])
	for i := 1; i < len(inputs)-2; i++ {
		current := sum(inputs[i : i+3])
		if current > prev {
			increases++
		}
		prev = current
	}
	return increases
}

func sum(values []int) int {
	var result int
	for _, val := range values {
		result += val
	}
	return result
}
