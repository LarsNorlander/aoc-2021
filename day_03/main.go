package main

import (
	_ "embed"
	"log"
	"strconv"
	"strings"
)

//go:embed inputs.txt
var rawInputs string
var inputs [][]int

// Prepare the rawInputs into a format more desirable for the problem at hand
func init() {
	for _, value := range strings.Fields(rawInputs) {
		var bits []int
		for _, bitChar := range value {
			bit, err := strconv.Atoi(string(bitChar))
			if err != nil {
				log.Fatalln(err)
			}
			bits = append(bits, bit)
		}
		inputs = append(inputs, bits)
	}
}

func main() {
	log.Println("Part 1 Answer: ", partOne())
	log.Println("Part 2 Answer: ", partTwo())
}

func partOne() int {
	var gamma []int
	for i := 0; i < len(inputs[0]); i++ {
		if mostCommonBit(inputs, i) == 1 {
			gamma = append(gamma, 1)
		} else {
			gamma = append(gamma, 0)
		}
	}
	return bitListToDecimal(gamma) * bitListToDecimal(flipBitList(gamma))
}

func partTwo() int {
	var o2GeneratorRating = reduce(inputs, 0, filterByOxygenGeneratorBitCriteria)
	var co2ScrubberRating = reduce(inputs, 0, filterByCO2ScrubberBitCriteria)
	return bitListToDecimal(o2GeneratorRating) * bitListToDecimal(co2ScrubberRating)
}

// Returns 1 or 0 depending on which is which values shows up more often
// Returns -1 if it's equal
// Assumes there are no unexpected values, e.g. any number that isn't 0 or 1
func mostCommonBit(values [][]int, position int) int {
	var sum int
	for _, value := range values {
		sum += value[position]
	}

	if len(values)%2 == 0 && sum == len(values)/2 {
		return -1
	} else if sum > len(values)/2 {
		return 1
	} else {
		return 0
	}
}

func bitListToDecimal(bits []int) int {
	var valueStr string
	for _, bit := range bits {
		valueStr += strconv.Itoa(bit)
	}
	result, err := strconv.ParseInt(valueStr, 2, 0)
	if err != nil {
		log.Fatalln(err)
	}
	return int(result)
}

func flipBitList(bits []int) []int {
	var result []int
	for _, bit := range bits {
		if bit == 0 {
			result = append(result, 1)
		} else {
			result = append(result, 0)
		}
	}
	return result
}

func reduce(values [][]int, position int, bitCriteriaFilter func(values [][]int, position int) [][]int) []int {
	if len(values) == 1 {
		return values[0]
	}
	filteredValues := bitCriteriaFilter(values, position)
	return reduce(filteredValues, position+1, bitCriteriaFilter)
}

func filterByOxygenGeneratorBitCriteria(values [][]int, position int) [][]int {
	if mostCommonBit(values, position) == 0 {
		return filter(values, position, 0)
	}
	return filter(values, position, 1)
}

func filterByCO2ScrubberBitCriteria(values [][]int, position int) [][]int {
	if mostCommonBit(values, position) == 0 {
		return filter(values, position, 1)
	}
	return filter(values, position, 0)
}

func filter(values [][]int, position int, filterBit int) [][]int {
	var filteredValues [][]int

	for _, value := range values {
		if value[position] == filterBit {
			filteredValues = append(filteredValues, value)
		}
	}

	return filteredValues
}
