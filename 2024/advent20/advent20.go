package main

import (
	_ "embed"
	"fmt"
	"github.com/sekullbe/advent/grid"
	"github.com/sekullbe/advent/tools"
	"time"
)

//go:embed input.txt
var inputText string

func main() {
	fmt.Printf("Magic number: %d\n", run1(inputText, 100))
	fmt.Println("-------------")
	fmt.Printf("Magic number: %d\n", run2(inputText))
}

func run1(input string, mustImprove int) int {
	defer tools.Track(time.Now(), "Part 1 Time")
	b := grid.ParseBoardString(input)
	start, _ := b.Find('S')
	end, _ := b.Find('E')
	basePath, baseline, havePath := b.FindPath(start, end)
	if !havePath {
		panic("no path in initial maze")
	}
	// Brute force!
	goodCheats := 0
	interiorWallPoints := b.FindAll(grid.WALL, true)
	for _, point := range interiorWallPoints {
		cxns := 0
		for _, np := range b.GetSquareNeighbors(point) {
			// Only continue if the cheat point connects two points on the baseline path
			// that optimizes 10s (45 to 35s) but it's still very slow
			if tools.Contains(basePath, np) {
				cxns++
			}
		}
		if cxns < 2 {
			continue
		}
		b.Grid[point].Contents = grid.EMPTY
		_, score, havePath := b.FindPath(start, end)
		if havePath && score <= baseline-mustImprove {
			//log.Printf("Cheat at %v saves %d", point, baseline-score)
			goodCheats++
		}
		b.Grid[point].Contents = grid.WALL
	}

	return goodCheats
}

func run2(input string) int {
	defer tools.Track(time.Now(), "Part 2 Time:")

	// it's not really a maze- there is only one path

	return 0
}
