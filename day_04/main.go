package main

import (
	_ "embed"
	"log"
	"strconv"
	"strings"
)

//go:embed inputs.txt
var rawInputs string
var inputs struct {
	draws  []int
	boards []bingoBoard
}

type bingoBoard struct {
	cells [5][5]cell
	won   bool
}

type cell struct {
	value  int
	marked bool
}

// Prepare the rawInputs into a format more desirable for the problem at hand
func init() {
	data := strings.Fields(rawInputs)
	drawsStr := strings.Split(data[0], ",")
	for _, val := range drawsStr {
		something, err := strconv.Atoi(val)
		if err != nil {
			log.Fatalln(err)
		}
		inputs.draws = append(inputs.draws, something)
	}
	for i, val := range data[1:] {
		if i%25 == 0 {
			inputs.boards = append(inputs.boards, bingoBoard{cells: [5][5]cell{}})
		}
		valInt, err := strconv.Atoi(val)
		if err != nil {
			log.Fatalln(err)
		}
		inputs.boards[i/25].cells[(i/5 - i/25*5)][i%5] = cell{value: valInt}
	}
}

func main() {
	log.Println("Part 1 Answer: ", partOne())
	log.Println("Part 2 Answer: ", partTwo())
}

func partOne() int {
	boards := inputs.boards
	draws := inputs.draws

	var lastDraw int
	winningBoardIdx := -1

	for _, draw := range draws {
		for i := range boards {
			for j := range boards[i].cells {
				for k := range boards[i].cells[j] {
					if boards[i].cells[j][k].value == draw {
						boards[i].cells[j][k].marked = true
					}
				}
			}
		}

		for i := range boards {
			if isWinningBoard(boards[i]) {
				winningBoardIdx = i
			}
		}
		if winningBoardIdx != -1 {
			lastDraw = draw
			break
		}
	}

	return boardScore(boards[winningBoardIdx], lastDraw)
}

func partTwo() int {
	boards := inputs.boards
	draws := inputs.draws

	var lastDraw int
	var lastWinningBoard bingoBoard

	for _, draw := range draws {
		for i := range boards {
			if boards[i].won {
				continue
			}
			for j := range boards[i].cells {
				for k := range boards[i].cells[j] {
					if boards[i].cells[j][k].value == draw {
						boards[i].cells[j][k].marked = true
					}
				}
			}
		}

		for i := range boards {
			if !boards[i].won && isWinningBoard(boards[i]) {
				lastWinningBoard = boards[i]
				boards[i].won = true
				lastDraw = draw
			}
		}
	}

	return boardScore(lastWinningBoard, lastDraw)
}

func isWinningBoard(b bingoBoard) bool {
	// check all rows
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if !b.cells[i][j].marked {
				break
			}
			if j == 4 {
				return true
			}
		}
	}
	// check all columns
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if !b.cells[j][i].marked {
				break
			}
			if j == 4 {
				return true
			}
		}
	}
	return false
}

func boardScore(b bingoBoard, mult int) int {
	var sum int
	for i := range b.cells {
		for j := range b.cells[i] {
			if !b.cells[i][j].marked {
				sum += b.cells[i][j].value
			}
		}
	}
	return sum * mult
}
