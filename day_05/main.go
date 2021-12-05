package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed inputs.txt
var rawInputs string
var inputs []LineSegment
var maxX, maxY int

type Point struct {
	x, y int
}

type LineSegment struct {
	p1, p2 Point
}

type Grid [][]int

func newGrid(xSize, ySize int) Grid {
	grid := make([][]int, ySize)
	for i := 0; i < ySize; i++ {
		grid[i] = make([]int, xSize)
	}
	return grid
}

// Create a string representation of the grid so that the Y-axis goes downwards when
// printed into a terminal.
func (grid Grid) String() string {
	var str string
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			value := grid[i][j]
			if value == 0 {
				str += "."
			} else {
				str += strconv.Itoa(value)
			}
		}
		str += "\n"
	}
	return str
}

// Prepare the rawInputs into a format more desirable for the problem at hand.
func init() {
	data := strings.Fields(rawInputs)
	for i := 0; i < len(data)-2; i = i + 3 {
		point1 := pointFromString(data[i])
		point2 := pointFromString(data[i+2])
		inputs = append(inputs, LineSegment{
			p1: point1,
			p2: point2,
		})

		// search for max x
		if point1.x > maxX {
			maxX = point1.x
		}
		if point2.x > maxX {
			maxX = point2.x
		}

		// search for max y
		if point1.y > maxY {
			maxY = point1.y
		}
		if point2.y > maxY {
			maxY = point2.y
		}
	}
}

func main() {
	fmt.Println("Part 1 Answer: ", partOne())
	fmt.Println("Part 2 Answer: ", partTwo())
}

func partOne() int {
	grid := newGrid(maxX+1, maxY+1)
	for _, lineSeg := range inputs {
		// handle horizontal lines
		if lineSeg.p1.y == lineSeg.p2.y {
			var i, last int
			if lineSeg.p1.x < lineSeg.p2.x {
				i = lineSeg.p1.x
				last = lineSeg.p2.x
			} else {
				i = lineSeg.p2.x
				last = lineSeg.p1.x
			}

			for ; i <= last; i++ {
				grid[lineSeg.p1.y][i]++
			}

		}
		// handle vertical lines
		if lineSeg.p1.x == lineSeg.p2.x {
			var i, last int
			if lineSeg.p1.y < lineSeg.p2.y {
				i = lineSeg.p1.y
				last = lineSeg.p2.y
			} else {
				i = lineSeg.p2.y
				last = lineSeg.p1.y
			}

			for ; i <= last; i++ {
				grid[i][lineSeg.p1.x]++
			}
		}
	}

	fmt.Print(grid.String())

	return calculateDangerousPoints(grid)
}

func partTwo() int {
	grid := newGrid(maxX+1, maxY+1)

	for _, lineSeg := range inputs {
		if lineSeg.p1.y == lineSeg.p2.y {
			// handle horizontal lines
			var i, last int
			if lineSeg.p1.x < lineSeg.p2.x {
				i = lineSeg.p1.x
				last = lineSeg.p2.x + 1
			} else {
				i = lineSeg.p2.x
				last = lineSeg.p1.x + 1
			}

			for ; i < last; i++ {
				grid[lineSeg.p1.y][i]++
			}
		} else if lineSeg.p1.x == lineSeg.p2.x {
			// handle vertical lines
			var i, last int
			if lineSeg.p1.y < lineSeg.p2.y {
				i = lineSeg.p1.y
				last = lineSeg.p2.y + 1
			} else {
				i = lineSeg.p2.y
				last = lineSeg.p1.y + 1
			}

			for ; i < last; i++ {
				grid[i][lineSeg.p1.x]++
			}
		} else {
			// handle diagonal lines
			diagLineSeg := lineSeg

			// flips line segment points so that p1.x is always the smaller value
			// this allows us to only care if the line goes upwards or downward in
			// relation to p1
			if lineSeg.p1.x > lineSeg.p2.x {
				diagLineSeg = LineSegment{p1: lineSeg.p2, p2: lineSeg.p1}
			}

			distance := diagLineSeg.p2.x - diagLineSeg.p1.x

			for i := 0; i <= distance; i++ {
				if diagLineSeg.p1.y < diagLineSeg.p2.y {
					// line goes upwards
					grid[diagLineSeg.p1.y+i][diagLineSeg.p1.x+i]++
				} else {
					// line goes downwards
					grid[diagLineSeg.p1.y-i][diagLineSeg.p1.x+i]++
				}
			}
		}
	}

	fmt.Print(grid.String())

	return calculateDangerousPoints(grid)
}

func calculateDangerousPoints(grid Grid) int {
	var dangerousPoints int
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] > 1 {
				dangerousPoints++
			}
		}
	}
	return dangerousPoints
}

func pointFromString(s string) Point {
	values := strings.Split(s, ",")
	x, err := strconv.Atoi(values[0])
	if err != nil {
		panic(err)
	}
	y, err := strconv.Atoi(values[1])
	if err != nil {
		panic(err)
	}
	return Point{x: x, y: y}
}
