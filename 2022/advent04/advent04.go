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
	pairs := parsers.SplitByLines(inputText)
	var overlaps int
	for _, pair := range pairs {
		var e1a, e1b, e2a, e2b int
		// a pair looks like 2-4,6-8
		fmt.Sscanf(pair, "%d-%d,%d-%d", &e1a, &e1b, &e2a, &e2b)
		if e2a >= e1a && e2b <= e1b || e1a >= e2a && e1b <= e2b {
			overlaps++
		}
	}
	return overlaps
}

func run2(inputText string) int {
	pairs := parsers.SplitByLines(inputText)
	var overlaps int
	for _, pair := range pairs {
		var e1a, e1b, e2a, e2b int
		fmt.Sscanf(pair, "%d-%d,%d-%d", &e1a, &e1b, &e2a, &e2b)
		if e1a <= e2b && e2a <= e1b {
			overlaps++
		}
	}
	return overlaps
}
