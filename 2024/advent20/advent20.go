package main

import (
	_ "embed"
	"fmt"
	"github.com/sekullbe/advent/geometry"
	"github.com/sekullbe/advent/grid"
	"github.com/sekullbe/advent/tools"
	"time"
)

//go:embed input.txt
var inputText string

func main() {
	fmt.Printf("Magic number: %d\n", run1(inputText, 100))
	fmt.Println("-------------")
	fmt.Printf("Magic number: %d\n", run2(inputText, 100))
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

type cheat struct {
	begin geometry.Point2
	end   geometry.Point2
}

func run2(input string, mustImprove int) int {
	defer tools.Track(time.Now(), "Part 2 Time:")

	// it's not really a maze- there is only one path
	b := grid.ParseBoardString(input)
	start, _ := b.Find('S')
	end, _ := b.Find('E')
	basePath, baseLine, havePath := b.FindPath(start, end)
	if !havePath {
		panic("no path in initial maze")
	}
	_ = baseLine

	goodCheats := make(map[cheat]bool)
	countCheats := make(map[int]int)

	// for each point in basepath
	// look at all points _ahead_ of this in basepath where manhattan distance <=20
	// compute the savings - i.e. in example it jumps from csi 0 to cdi 82 in 6 steps, so total cost is 8 and savings are 76 (82-6-0)
	// store that in the cheat map	if it's good enough
	for cheatStartIdx, cheatStart := range basePath {
		for cheatDestIdx, cheatDest := range basePath[cheatStartIdx+1:] {
			cheatJumpDist := grid.ManhattanDistance(cheatStart, cheatDest)
			if cheatJumpDist <= 20 { // not 19!!
				savings := cheatDestIdx - cheatJumpDist + 1
				if savings >= mustImprove {
					goodCheats[cheat{cheatStart, cheatDest}] = true
					countCheats[savings]++
				}
			}
		}
	}
	//fmt.Println(goodCheats)
	//fmt.Println(countCheats)
	return len(goodCheats)
}
