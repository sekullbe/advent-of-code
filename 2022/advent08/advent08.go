package main

import (
	_ "embed"
	"fmt"
	"github.com/sekullbe/advent/parsers"
)

//go:embed input.txt
var inputText string

func main() {
	lines := parsers.SplitByLines(inputText)
	fmt.Printf("Magic number: %d\n", run1(lines))
	fmt.Println("-------------")
	fmt.Printf("Magic number: %d\n", run2(lines))
}

func run1(lines []string) int {

	maxX := len(lines[0]) - 1
	maxY := len(lines) - 1
	w := newWorld(maxX, maxY)
	for y, line := range lines {
		for x, treeHeightRune := range line {
			w.addTree(x, y, int(treeHeightRune-'0'))
		}
	}
	w.computeVisibility()
	//fmt.Println(w.display())
	visibleTrees := w.forest.countVisibleTrees()

	return visibleTrees
}

func run2(lines []string) int {

	return 0
}
