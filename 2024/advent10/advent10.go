package main

import (
	_ "embed"
	"fmt"
	"github.com/sekullbe/advent/geometry"
	"github.com/sekullbe/advent/grid"
	"github.com/sekullbe/advent/parsers"
)

//go:embed input.txt
var inputText string

func main() {
	fmt.Printf("Magic number: %d\n", run1(inputText))
	fmt.Println("-------------")
	fmt.Printf("Magic number: %d\n", run2(inputText))
}

// input is a 60x60 grid of numbers

func run1(input string) int {
	b := grid.ParseBoard(parsers.SplitByLines(input))

	score := 0

	for point, tile := range b.Grid {
		if tile.Value == 0 {
			peaks := make(map[geometry.Point]int)
			followTrail(b, point, peaks, true)
			score += len(peaks)
		}
	}

	return score
}

func run2(input string) int {
	b := grid.ParseBoard(parsers.SplitByLines(input))

	score := 0

	for point, tile := range b.Grid {
		if tile.Value == 0 {
			peaks := make(map[geometry.Point]int)
			followTrail(b, point, peaks, false)
			score += totalTrails(peaks)
		}
	}

	return score
}

func totalTrails(peaks map[geometry.Point]int) int {
	score := 0
	for _, i := range peaks {
		score += i
	}

	return score
}

func followTrail(b *grid.Board, pt geometry.Point, peaks map[geometry.Point]int, careAboutRevisits bool) {
	heightHere := b.AtPoint(pt).Value
	if !b.InRange(pt) {
		return
	}

	if heightHere == 9 {
		//log.Printf("ding! at %v", pt)
		peaks[pt] = peaks[pt] + 1
		return
	}

	for _, np := range b.GetSquareNeighbors(pt) {
		_, beenThere := peaks[np]
		if b.AtPoint(np).Value == heightHere+1 && !(beenThere && careAboutRevisits) {
			//log.Printf("at %v/%d going to %v/%d", pt, heightHere, np, b.AtPoint(np).Value)
			followTrail(b, np, peaks, careAboutRevisits)
		}
	}
	return
}
