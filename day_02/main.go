package main

import (
	_ "embed"
	"log"
	"strconv"
	"strings"
)

//go:embed inputs.txt
var rawInputs string
var inputs []command

type command struct {
	direction string
	value     int
}

// Prepare the rawInputs into a format more desirable for the problem at hand
func init() {
	values := strings.Fields(rawInputs)
	for i := 0; i < len(values)-1; i += 2 {
		value, err := strconv.Atoi(values[i+1])
		if err != nil {
			log.Fatalln(err)
		}
		inputs = append(inputs, command{
			direction: values[i],
			value:     value,
		})
	}
}

func main() {
	log.Println("Part 1 Answer: ", partOne())
	log.Println("Part 2 Answer: ", partTwo())
}

func partOne() int {
	var hPos, depth int

	for _, cmd := range inputs {
		switch cmd.direction {
		case "forward":
			hPos += cmd.value
			break
		case "down":
			depth += cmd.value
			break
		case "up":
			depth -= cmd.value
			break
		}
	}

	return hPos * depth
}

func partTwo() int {
	var hPos, depth, aim int

	for _, cmd := range inputs {
		switch cmd.direction {
		case "forward":
			hPos += cmd.value
			depth += aim * cmd.value
			break
		case "down":
			aim += cmd.value
			break
		case "up":
			aim -= cmd.value
			break
		}
	}

	return hPos * depth
}
