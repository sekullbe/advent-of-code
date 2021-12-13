package main

import (
	_ "embed"
	"fmt"
	"github.com/sekullbe/advent/parsers"
	"regexp"
	"strconv"
)

//go:embed input.txt
var inputText string

func main() {
	fmt.Printf("Magic number: %d\n", run1(inputText))
	fmt.Println("-------------")
	fmt.Printf("Magic number: %d\n", run2(inputText))
}

func run1(inputText string) int {
	return doTheThing(inputText, 1)
}

func run2(inputText string) int {
	return doTheThing(inputText, 999)
}

func doTheThing(inputText string, steps int) int {
	coordMode := true
	var g grid = newGrid()
	foldCount := 0

	for _, line := range parsers.SplitByLines(inputText) {
		if line == "" {
			coordMode = false
			continue
		}
		if coordMode {
			g.addCoordString(line)
		} else {
			// "fold along x=655" ... presumably this is a balanced fold- check for that and log just in case
			re := regexp.MustCompile(`fold along ([xy])=(\d+)`)
			matches := re.FindStringSubmatch(line)
			axis := matches[1]
			axisNum, _ := strconv.Atoi(matches[2])
			foldCount++
			g.fold(axis, axisNum)
			//log.Printf("after %d folds there are %d visible stars", foldCount, g.countPoints())
			if foldCount == steps {
				return g.countPoints()
			}
		}

	}
	fmt.Println(g.display())
	return g.countPoints()
}
