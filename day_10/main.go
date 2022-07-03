package main

import (
	_ "embed"
	"fmt"
	"github.com/LarsNorlander/aoc-2021/lib/set"
	"github.com/LarsNorlander/aoc-2021/lib/stack"
	"log"
	"sort"
	"strings"
)

//go:embed inputs.txt
var rawInputs string
var inputs []string

type ChunkWrapper struct {
	open, close                 rune
	errorScore, correctionScore int
}

var (
	chunkWrappers = []ChunkWrapper{
		{'(', ')', 3, 1},
		{'[', ']', 57, 2},
		{'{', '}', 1197, 3},
		{'<', '>', 25137, 4},
	}
	openingSet      = set.Of[rune]()
	chunkMapByClose = make(map[rune]*ChunkWrapper)
	chunkMapByOpen  = make(map[rune]*ChunkWrapper)
)

// Prepare the rawInputs into a format more desirable for the problem at hand.
func init() {
	for _, str := range strings.Fields(rawInputs) {
		inputs = append(inputs, str)
	}

	// Pre-Compute some data
	for i := range chunkWrappers {
		openingSet.Add(chunkWrappers[i].open)
		chunkMapByOpen[chunkWrappers[i].open] = &chunkWrappers[i]
		chunkMapByClose[chunkWrappers[i].close] = &chunkWrappers[i]
	}
}

func main() {
	fmt.Println("Part 1 Answer: ", PartOne(inputs))
	fmt.Println("Part 2 Answer: ", PartTwo(inputs))
}

func PartOne(inputs []string) int {
	charStack := stack.Empty[rune]()
	score := 0

	for _, line := range inputs {
		for _, char := range line {
			if openingSet.Contains(char) {
				charStack.Push(char)
				continue
			}
			poppedChar, err := charStack.Pop()
			if err != nil {
				log.Fatalln(err)
			}
			chunkWrapper, ok := chunkMapByClose[char]
			if !ok {
				log.Fatalf("unexpected character %c", char)
			}
			if poppedChar != chunkWrapper.open {
				score += chunkWrapper.errorScore
			}
		}
	}

	return score
}

func PartTwo(inputs []string) int {
	var scores []int

LineLoop:
	for _, line := range inputs {
		charStack := stack.Empty[rune]()
		for _, char := range line {
			if openingSet.Contains(char) {
				charStack.Push(char)
				continue
			}
			poppedChar, err := charStack.Pop()
			if err != nil {
				log.Fatalln(err)
			}
			chunkWrapper, ok := chunkMapByClose[char]
			if !ok {
				log.Fatalf("unexpected character %c", char)
			}
			if poppedChar != chunkWrapper.open {
				continue LineLoop
			}
		}

		var score int
		for charStack.Peek() != 0 {
			char, _ := charStack.Pop()
			chunkwrapper := chunkMapByOpen[char]
			score = score*5 + chunkwrapper.correctionScore
		}
		scores = append(scores, score)
	}
	sort.Ints(scores)
	n := len(scores) / 2
	return scores[n]
}
