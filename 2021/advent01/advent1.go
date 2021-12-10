package main

import (
	_ "embed"
	"fmt"

	"github.com/sekullbe/advent/parsers"
)

//go:embed input.txt
var inputText string

func main() {
	fmt.Printf("Depth increases:%d\n", run1(inputText))
	fmt.Println("-------------")
	fmt.Printf("Depth increases:%d\n", run2(inputText))
}

func run1(inputText string) int {

	var increases int = -1 // The first depth will count as an 'increase' so ignore it
	var prevDepth int
	for _, depth := range parsers.StringsToIntSlice(inputText) {
		if depth > prevDepth {
			increases++
		}
		prevDepth = depth
	}
	return increases
}

func run2(inputText string) int {

	measurements := parsers.StringsToIntSlice(inputText)
	var increases int
	for i := 0; i < len(measurements); i++ {
		if i < 3 {
			continue // not enough yet
		}
		prevWindow := sumSlice(measurements[i-3:i])
		curWindow := sumSlice(measurements[i-2:i+1])
		if curWindow > prevWindow {
			increases++
		}
	}
	return increases
}


func sumSlice (measurements []int) int {
	sum := 0
	for _, measurement := range measurements {
		sum += measurement
	}
	return sum
}
