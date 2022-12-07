package main

import (
	"fmt"
	"github.com/sekullbe/advent/parsers"
)

var inputText = "16,12,1,0,15,7,11"

func main() {
	fmt.Printf("Magic number: %d\n", run1(inputText))
	fmt.Println("-------------")
	fmt.Printf("Magic number: %d\n", run2(inputText))
}

func run1(inputText string) int {
	return playGame(parsers.StringsWithCommasToIntSlice(inputText), 2020)
}

func playGame(starters []int, lastTurn int) int {

	var last int
	var spokenBefore bool

	memory := make(map[int]int)
	for i, starter := range starters {
		_, spokenBefore = memory[starter]
		memory[starter] = i + 1 // store the turn number, which is 1-based
		last = starter
	}
	for turn := len(starters) + 1; turn <= lastTurn; turn++ {
		// when it was spoken, was that the first time it was spoken?
		var speak int
		if !spokenBefore {
			speak = 0
		} else {
			lastTurnSeen, _ := memory[last]
			speak = turn - 1 - lastTurnSeen // diff between the turn it was just seen (t-1) and the previous time it was seen
		}
		memory[last] = turn - 1 // now store that we remember it
		_, spokenBefore = memory[speak]
		//log.Printf("On turn %d speak %d", turn, speak)
		last = speak

	}

	return last
}

func run2(inputText string) int {
	return playGame(parsers.StringsWithCommasToIntSlice(inputText), 30000000)
}
