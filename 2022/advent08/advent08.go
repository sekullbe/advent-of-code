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
	run(lines)
}

func run(lines []string) (int, int) {

	maxX := len(lines[0]) - 1
	maxY := len(lines) - 1
	w := newWorld(maxX, maxY)
	for y, line := range lines {
		for x, treeHeightRune := range line {
			w.addTree(x, y, int(treeHeightRune-'0'))
		}
	}

	//fmt.Println(w.display())
	maxScenicScore := w.computeVisibility()
	visibleTrees := w.forest.countVisibleTrees()

	fmt.Printf("Part 1 Magic number: %d\n", visibleTrees)
	fmt.Println("-------------")
	fmt.Printf("Part 2 Magic number: %d\n", maxScenicScore)
	return visibleTrees, maxScenicScore
}

func run1(lines []string) int {
	vt, _ := run(lines)
	return vt
}

func run2(lines []string) int {
	_, mss := run(lines)
	return mss
}
