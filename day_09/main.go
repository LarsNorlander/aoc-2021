package main

import (
	"container/list"
	_ "embed"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

//go:embed inputs.txt
var rawInputs string
var inputs [][]Cell

type Cell struct {
	Location Location
	Value    int
}

type Location struct {
	Row, Col int
}

// Prepare the rawInputs into a format more desirable for the problem at hand.
func init() {
	for rowIdx, rowStr := range strings.Fields(rawInputs) {
		var row []Cell
		for colIdx, cell := range rowStr {
			value, err := strconv.Atoi(string(cell))
			if err != nil {
				panic(err)
			}
			row = append(row, Cell{
				Location: Location{
					Row: rowIdx,
					Col: colIdx,
				},
				Value: value})
		}
		inputs = append(inputs, row)
	}
}

// The final solution could definitely be de-cluttered, but we'll save that for another time.
func main() {
	fmt.Println("Part 1 Answer: ", PartOne(inputs))
	fmt.Println("Part 2 Answer: ", PartTwo(inputs))
}

func PartOne(inputs [][]Cell) int {
	var riskScore int

	for i := 0; i < len(inputs); i++ {
		for j := 0; j < len(inputs[i]); j++ {
			lowestAdjacentValue := math.MaxInt
			for _, cell := range getAdjacentCells(inputs, inputs[i][j].Location) {
				if cell.Value < lowestAdjacentValue {
					lowestAdjacentValue = cell.Value
				}
			}
			if inputs[i][j].Value < lowestAdjacentValue {
				riskScore += 1 + inputs[i][j].Value
			}
		}
	}

	return riskScore
}

// This function is written with the assumption that it'll never be passed a jagged
// two-dimensional array, and that there's at least one row and one column
func getAdjacentCells(grid [][]Cell, loc Location) []Cell {
	maxRowIdx := len(grid) - 1
	maxColIdx := len(grid[0]) - 1

	var adjacentValues []Cell

	// get north
	if loc.Row != 0 {
		adjacentValues = append(adjacentValues, grid[loc.Row-1][loc.Col])
	}
	// get south
	if loc.Row != maxRowIdx {
		adjacentValues = append(adjacentValues, grid[loc.Row+1][loc.Col])
	}
	// get west
	if loc.Col != 0 {
		adjacentValues = append(adjacentValues, grid[loc.Row][loc.Col-1])
	}
	// get east
	if loc.Col != maxColIdx {
		adjacentValues = append(adjacentValues, grid[loc.Row][loc.Col+1])
	}

	return adjacentValues
}

type Set map[Cell]bool

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

func SetOf(s ...Cell) Set {
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

func Append(set Set, elem ...Cell) Set {
	return Union(set, SetOf(elem...))
}

func (set Set) Contains(elem Cell) bool {
	return set[elem]
}

func PartTwo(inputs [][]Cell) int {
	var lowestPoints []Cell
	// Look for lowest points
	for i := 0; i < len(inputs); i++ {
		for j := 0; j < len(inputs[i]); j++ {
			lowestAdjacentValue := math.MaxInt
			for _, cell := range getAdjacentCells(inputs, inputs[i][j].Location) {
				if cell.Value < lowestAdjacentValue {
					lowestAdjacentValue = cell.Value
				}
			}
			if inputs[i][j].Value < lowestAdjacentValue {
				lowestPoints = append(lowestPoints, inputs[i][j])
			}
		}
	}

	// For each low point, figure out how big the basin is
	var basins []Set
	for _, point := range lowestPoints {
		basins = append(basins, floodFill(inputs, point))
	}

	sort.SliceStable(basins, func(i, j int) bool {
		return len(basins[i]) > len(basins[j])
	})

	return len(basins[0]) * len(basins[1]) * len(basins[2])
}

func floodFill(inputs [][]Cell, cell Cell) Set {
	basin := SetOf()
	q := list.New()
	q.PushBack(cell)
	for q.Len() != 0 {
		n := q.Front().Value.(Cell)
		q.Remove(q.Front())
		if n.Value != 9 && !basin.Contains(n) {
			basin = Append(basin, n)
			adjacentCells := getAdjacentCells(inputs, n.Location)
			for _, adjCell := range adjacentCells {
				q.PushBack(adjCell)
			}
		}
	}
	return basin
}
