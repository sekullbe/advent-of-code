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

func run1(inputText string) int {
	var games = parsers.SplitByLines(inputText)
	var totalscore int
	for _, game := range games {
		move, counter := parseOneGame(game)
		gameScore := scoreMove(counter) + scoreGame(move, counter)
		totalscore += gameScore
	}

	return totalscore
}

func run2(inputText string) int {
	var games = parsers.SplitByLines(inputText)
	var totalscore int
	for _, game := range games {
		move, desiredState := parseOneGame(game)
		// now the 'counter' is the required state not the required move; X=lose,Y=draw,Z=win
		counter := computeCounterFor(move, desiredState)

		totalscore += scoreMove(counter) + scoreGame(move, counter)
	}
	return totalscore
}
func parseOneGame(game string) (string, string) {
	var move, counter string
	n, err := fmt.Sscan(game, &move, &counter)
	if n != 2 || err != nil {
		panic(err)
	}
	return move, counter
}

func computeCounterFor(move, desiredState string) string {

	switch desiredState {
	case "X": //lose
		switch move {
		case "A":
			return "Z"
		case "B":
			return "X"
		case "C":
			return "Y"
		}
	case "Y": //draw
		switch move {
		case "A":
			return "X"
		case "B":
			return "Y"
		case "C":
			return "Z"
		}
	case "Z": //win
		switch move {
		case "A":
			return "Y"
		case "B":
			return "Z"
		case "C":
			return "X"
		}
	}
	return ""
}

func scoreMove(move string) int {
	switch move {
	case "A", "X":
		return 1
	case "B", "Y":
		return 2
	case "C", "Z":
		return 3
	}
	panic("bad move")
}

// given opponent's move, score my counter
func scoreGame(move, counter string) int {
	if (move == "A" && counter == "X") || (move == "B" && counter == "Y") || (move == "C" && counter == "Z") {
		return 3
	}
	// winners
	if (move == "A" && counter == "Y") || (move == "B" && counter == "Z") || (move == "C" && counter == "X") {
		return 6
	}
	// must have lost
	return 0
}
