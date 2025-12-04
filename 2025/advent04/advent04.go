package main

import (
	_ "embed"
	"fmt"
	"time"

	"github.com/sekullbe/advent/grid"
	"github.com/sekullbe/advent/tools"
)

//go:embed input.txt
var inputText string

func main() {
	fmt.Printf("Magic number: %d\n", run1(inputText))
	fmt.Println("-------------")
	fmt.Printf("Magic number: %d\n", run2(inputText))
}

func run1(input string) int {
	defer tools.Track(time.Now(), "Part 1 Time")

	// read it into a grid
	b := grid.ParseBoardString(input)
	return checkBoardOnceAndMarkRemovals(b)
}

func checkBoardOnceAndMarkRemovals(b *grid.Board) int {
	accessible := 0
	for point, _ := range b.Grid {
		if b.AtPoint(point).Contents != '@' && b.AtPoint(point).Contents != 'x' {
			continue
		}
		observedRolls := b.CheckNeighborsForSymbol(point, '@')
		observedRolls = append(observedRolls, b.CheckNeighborsForSymbol(point, 'x')...)
		if len(observedRolls) < 4 {
			accessible++
			b.AtPoint(point).Contents = 'x'
		}
	}
	return accessible
}

func run2(input string) int {
	defer tools.Track(time.Now(), "Part 2 Time")

	b := grid.ParseBoardString(input)
	totalRemovals := 0
	for {
		removals := checkBoardOnceAndMarkRemovals(b)
		if removals == 0 {
			break
		}
		totalRemovals += removals
		// now actually do the removals
		b.ReplaceAll('x', grid.EMPTY, false)

	}

	return totalRemovals
}
