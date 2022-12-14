package main

import (
	_ "embed"
	"fmt"
	"github.com/sekullbe/advent/parsers"
)

//go:embed input.txt
var inputText string

func main() {
	fmt.Printf("Magic number: %d\n", run1(inputText))
	fmt.Println("-------------")
	fmt.Printf("Magic number: %d\n", run2(inputText))
}

func run1(inputText string) int {

	// load the rocks from the grid
	w := parseRockVeins(parsers.SplitByLines(inputText))
	for w.dropSand(point{500, 0}) {
	}
	//w.printWorld()
	return w.sandCount
}

func run2(inputText string) int {
	dropPoint := point{500, 0}
	w := parseRockVeins(parsers.SplitByLines(inputText))
	w.lowestRock = w.lowestRock + 2
	// The width here should be just enough for a pyramid to form based around the drop point
	for x := dropPoint.x - w.lowestRock; x <= dropPoint.x+w.lowestRock; x++ {
		w.grid[point{x, w.lowestRock}] = '#'
	}
	for w.dropSand(dropPoint) {
	}
	//w.printWorld()
	return w.sandCount // 1005 too low
}
