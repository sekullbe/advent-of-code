package main

import (
	_ "embed"
	"fmt"
	"github.com/sekullbe/advent/parsers"
	"log"
)

//go:embed input.txt
var inputText string

func main() {
	fmt.Printf("Magic number: %d\n", run1(inputText))
	fmt.Println("-------------")
	fmt.Printf("Magic number: %d\n", run2(inputText))
}

func run1(inputText string) int {
	return doit(parsers.SplitByLines(inputText), 2)
}

func run2(inputText string) int {
	return doit(parsers.SplitByLines(inputText), 10)
}

func doit(lines []string, ropeLength int) int {
	b := newBoard(ropeLength)

	for _, line := range lines {
		var dir rune
		var dist int
		n, err := fmt.Sscanf(line, "%c %d", &dir, &dist)
		if n != 2 || err != nil {
			log.Panicf("line broke the parser (%d)(%s): %s", n, err, line)
		}
		for i := 0; i < dist; i++ {
			b.move(dir)
		}
	}
	// now count up how many spots in the grid are true
	return len(b.tailVisited)
}
