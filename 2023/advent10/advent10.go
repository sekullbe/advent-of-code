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

func run1(input string) int {

	steps := 1

	b := parseBoard(parsers.SplitByLines(input))
	prev := b.start
	curr, _ := b.pipeNeighbors(prev) // choose one arbitrarily
	for {
		steps++
		//b.setSteps(next, steps) // don't really need this?
		next := b.followPipe(curr, prev)
		if next == b.start {
			break
		}
		// take the step
		prev = curr
		curr = next
		//fmt.Printf("Step: %v -> %v\n", prev, curr)
	}
	// now we know how may steps in the loop, so just divide by 2

	return steps / 2
}

func run2(input string) int {

	return 0
}
