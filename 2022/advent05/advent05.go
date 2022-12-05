package main

import (
	_ "embed"
	"fmt"
	"github.com/sekullbe/advent/parsers"
)

//go:embed input.txt
var inputText string

type stacks map[int][]rune

func main() {
	fmt.Printf("Magic boxes: %s\n", run1(inputText))
	fmt.Println("-------------")
	fmt.Printf("Magic boxes: %s\n", run2(inputText))
}

func run1(inputText string) string {

	stacks := parseStacks(parsers.SplitByLinesNoTrim(inputText))

	for _, line := range parsers.SplitByLines(inputText) {
		howmany, from, to := parseMove(line)
		for i := 0; i < howmany; i++ {
			oneMove(stacks, from, to)
		}
	}

	toppers := getTopOfStacks(stacks)

	return toppers
}

func run2(inputText string) string {
	stacks := parseStacks(parsers.SplitByLinesNoTrim(inputText))

	for _, line := range parsers.SplitByLines(inputText) {
		howmany, from, to := parseMove(line)
		// surprisingly this works when parseMove returns 0,0,0, but still handle it to avoid pointless slicing
		if howmany > 0 {
			multiMove(stacks, howmany, from, to)
		}
	}

	toppers := getTopOfStacks(stacks)
	return toppers
}

func getTopOfStacks(stacks stacks) (toppers string) {

	// can't do range because that gets the keys in random order
	for i := 1; i <= 9; i++ {
		stack := stacks[i]
		if len(stack) == 0 {
			toppers = toppers + " "
		} else {
			toppers = toppers + string(stack[len(stack)-1])
		}
	}
	return
}

func parseStacks(inputLines []string) stacks {
	stacks := make(stacks)
	for i := 0; i <= 9; i++ {
		stacks[i] = make([]rune, 0)
	}

	for _, line := range inputLines {
		for i, box := range line {
			// Catch the 'done' case
			if box == '1' {
				return stacks
			}
			if (i-1)%4 == 0 && box != ' ' {
				col := ((i - 1) / 4) + 1
				// this is an unshift operation to put the box at the beginning
				stacks[col] = append([]rune{box}, stacks[col]...)
			}
		}
	}
	return stacks
}

// this modifies the map in place
func oneMove(stacks stacks, from, to int) {
	// pop from, push to
	var box rune
	box, stacks[from] = stacks[from][len(stacks[from])-1], stacks[from][:len(stacks[from])-1]
	stacks[to] = append(stacks[to], box)
}

func multiMove(stacks stacks, howmany, from, to int) {
	// 	howmany := 3
	//	b = append(b, a[len(a)-howmany:]...)
	//	a = a[0 : len(a)-howmany]
	stacks[to] = append(stacks[to], stacks[from][len(stacks[from])-howmany:]...)
	stacks[from] = stacks[from][0 : len(stacks[from])-howmany]
}

func parseMove(move string) (howmany, from, to int) {
	_, _ = fmt.Sscanf(move, "move %d from %d to %d", &howmany, &from, &to)
	// i don't care about errors, just return zeroes

	return
}
