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

	return 0
}
