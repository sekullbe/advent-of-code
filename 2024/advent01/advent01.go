package main

import (
	_ "embed"
	"fmt"
	"github.com/sekullbe/advent/parsers"
	"github.com/sekullbe/advent/tools"
)

//go:embed input.txt
var inputText string

func main() {
	fmt.Printf("Magic number: %d\n", run1(inputText))
	fmt.Println("-------------")
	fmt.Printf("Magic number: %d\n", run2(inputText))
}

func run1(input string) int {

	var left, right []int

	for _, s := range parsers.SplitByLines(input) {
		var nL, nR int
		_, err := fmt.Sscanf(s, "%d %d", &nL, &nR)
		if err != nil {
			panic(err)
		}
		left = append(left, nL)
		right = append(right, nR)
	}
	left = tools.Sort(left)
	right = tools.Sort(right)

	sum := 0
	for i, nL := range left {
		sum += tools.AbsInt(nL - right[i])
	}

	return sum
}

func run2(input string) int {

	var left []int
	rightCount := make(map[int]int)
	for _, s := range parsers.SplitByLines(input) {
		var nL, nR int
		_, err := fmt.Sscanf(s, "%d %d", &nL, &nR)
		if err != nil {
			panic(err)
		}
		left = append(left, nL)
		rightCount[nR]++
	}

	sim := 0
	for _, nL := range left {
		sim += nL * rightCount[nL]
	}

	return sim
}
