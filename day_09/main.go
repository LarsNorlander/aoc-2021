package main

import (
	_ "embed"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

//go:embed inputs.txt
var rawInputs string
var inputs []Input

type Input struct {
	Patterns []string
	Readings []string
}

type Set map[rune]bool

func (set Set) String() string {
	var s string
	s += "("
	for elem, _ := range set {
		s += fmt.Sprintf("%c, ", elem)
	}
	s = strings.TrimSuffix(s, ", ")
	s += ")"
	return s
}

func SetOf(s ...rune) Set {
	set := make(Set)
	for _, val := range s {
		set[val] = true
	}
	return set
}

func Union(sets ...Set) Set {
	union := make(Set)
	for _, set := range sets {
		for elem, _ := range set {
			union[elem] = true
		}
	}
	return union
}

func Difference(setA, setB Set) Set {
	diff := make(Set)
	for elem, _ := range setA {
		diff[elem] = true
	}
	for elem, _ := range setB {
		delete(diff, elem)
	}
	return diff
}

func Intersect(setA, setB Set) Set {
	intersect := make(Set)
	for elem, _ := range setA {
		if setB[elem] {
			intersect[elem] = true
		}
	}
	return intersect
}

func (set Set) Contains(elem rune) bool {
	return set[elem]
}

// Prepare the rawInputs into a format more desirable for the problem at hand.
func init() {
	lines := strings.Split(strings.TrimSpace(rawInputs), "\n")
	for _, line := range lines {
		input := Input{}
		temp := strings.Split(line, " | ")
		for _, enc := range strings.Split(temp[0], " ") {
			input.Patterns = append(input.Patterns, enc)
		}
		for _, reading := range strings.Split(temp[1], " ") {
			input.Readings = append(input.Readings, reading)
		}
		inputs = append(inputs, input)
	}
}

func main() {
	fmt.Println("Part 1 Answer: ", PartOne(inputs))
	fmt.Println("Part 2 Answer: ", PartTwo(inputs))
}

func PartOne(values []Input) int {
	count := 0
	for _, value := range values {
		for _, reading := range value.Readings {
			switch len(reading) {
			case 7, 4, 3, 2:
				count++
				break
			}
		}
	}
	return count
}

func PartTwo(values []Input) int {
	var sumOfReadings int
	for _, value := range values {
		encodings := make(map[int]Set)
		var sixSegmentEncodings []Set
		for _, pattern := range value.Patterns {
			switch len(pattern) {
			case 7:
				encodings[8] = SetOf([]rune(pattern)...)
				break
			case 6:
				sixSegmentEncodings = append(sixSegmentEncodings, SetOf([]rune(pattern)...))
				break
			case 4:
				encodings[4] = SetOf([]rune(pattern)...)
				break
			case 3:
				encodings[7] = SetOf([]rune(pattern)...)
				break
			case 2:
				encodings[1] = SetOf([]rune(pattern)...)
				break
			}
		}
		// Solve for all segments
		A := Difference(encodings[7], encodings[1])
		EG := Difference(encodings[8], Union(encodings[7], encodings[4]))
		BD := Difference(encodings[8], Union(encodings[7], EG))
		CDE := Union(
			Difference(encodings[8], sixSegmentEncodings[0]),
			Difference(encodings[8], sixSegmentEncodings[1]),
			Difference(encodings[8], sixSegmentEncodings[2]),
		)
		G := Difference(EG, CDE)
		E := Difference(EG, G)
		CD := Difference(CDE, E)
		C := Intersect(CD, encodings[1])
		D := Difference(CD, C)
		F := Difference(encodings[1], C)
		B := Difference(BD, D)
		// Assign remaining encodings
		encodings[0] = Union(A, B, C, E, F, G)
		encodings[2] = Union(A, C, D, E, G)
		encodings[3] = Union(A, C, D, F, G)
		encodings[5] = Union(A, B, D, F, G)
		encodings[6] = Union(A, B, D, E, F, G)
		encodings[9] = Union(A, B, C, D, F, G)
		// Loop through Readings
		var readings []int
		for _, reading := range value.Readings {
			readings = append(readings, lookup(SetOf([]rune(reading)...), encodings))
		}
		sumOfReadings += intSliceToDecimal(readings)
	}
	return sumOfReadings
}

func lookup(reading Set, encodings map[int]Set) int {
	for value, encoding := range encodings {
		if reflect.DeepEqual(reading, encoding) {
			return value
		}
	}
	panic("ya goofed")
}

func intSliceToDecimal(values []int) int {
	var str string
	for _, value := range values {
		str += strconv.Itoa(value)
	}
	val, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return val
}
