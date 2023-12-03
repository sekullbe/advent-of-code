package main

import (
	_ "embed"
	"fmt"
	"github.com/sekullbe/advent/parsers"
	"github.com/sekullbe/advent/tools"
)

//go:embed input.txt
var inputText string

func main() {
	fmt.Printf("Magic number: %d\n", run1(inputText))
	fmt.Println("-------------")
	fmt.Printf("Magic number: %d\n", run2(inputText))
}

func run1(inputText string) int {

	// read the whole thing into a grid
	b := parseBoard(parsers.SplitByLines(inputText))
	// scan once, mark every grid square adjacent to a symbol
	for point, _ := range b.grid {
		b.checkNeighborsForSymbol(point)
	}
	// scan again, looking for numbers, and if any of their squares are in the number is in
	sum := 0
	for y := 0; y <= b.maxY; y++ {
		var buildingNumber string
		var anySymbols bool
		for x := 0; x <= b.maxX+1; x++ {
			p, ok := b.grid[Pt(x, y)] // need to handle the case where we've finished the row
			if !ok || !isNumber(p.contents) {
				// we either have not started or have just finished a number
				if buildingNumber != "" && anySymbols {
					sum += tools.Atoi(buildingNumber)
				}
				buildingNumber = ""
				anySymbols = false
			} else {
				anySymbols = anySymbols || p.adjacentToSymbol
				buildingNumber += string(p.contents)
			}
		}
	}

	// OR do an adjacency scan for each number element

	return sum
}

func run2(inputText string) int {

	return 0
}
