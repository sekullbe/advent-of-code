package main

import (
	_ "embed"
	"fmt"
	"github.com/sekullbe/advent/parsers"
	"github.com/sekullbe/advent/tools"
	"sort"
)

//go:embed input.txt
var inputText string

func main() {
	fmt.Printf("Magic number: %d\n", run1(inputText))
	fmt.Println("-------------")
	fmt.Printf("Magic number: %d\n", run2(inputText))
}

func run1(inputText string) int {

	elves := make(map[int]int)
	var elfChunks = parsers.SplitByEmptyNewline(inputText)
	var mostcal, mostCalIndex int

	for i, chunk := range elfChunks {
		var calories int
		for _, s := range parsers.SplitByLines(chunk) {
			calories += tools.Atoi(s)
		}
		elves[i] = calories

		if calories > mostcal {
			mostCalIndex = i
			mostcal = calories
		}
	}

	fmt.Printf("elf %d has the most, %d calories (check %d)\n", mostCalIndex, mostcal, elves[mostCalIndex])
	return elves[mostCalIndex]
}

func run2(inputText string) int {
	elves := make(map[int]int)
	var elfChunks = parsers.SplitByEmptyNewline(inputText)
	elfIds := make([]int, len(elfChunks))

	for i, chunk := range elfChunks {
		var calories int
		for _, s := range parsers.SplitByLines(chunk) {
			calories += tools.Atoi(s)
		}
		elves[i] = calories
		elfIds[i] = i
	}

	sort.SliceStable(elfIds, func(i, j int) bool {
		return elves[elfIds[i]] > elves[elfIds[j]]
	})

	return elves[elfIds[0]] + elves[elfIds[1]] + elves[elfIds[2]]

}
