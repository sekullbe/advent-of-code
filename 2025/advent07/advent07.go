package main

import (
	_ "embed"
	"fmt"
	"time"

	"github.com/sekullbe/advent/grid"
	"github.com/sekullbe/advent/tools"
)

//go:embed input.txt
var inputText string

const BEAM = '|'
const SPLITTER = '^'

func main() {
	fmt.Printf("Magic number: %d\n", run1(inputText))
	fmt.Println("-------------")
	fmt.Printf("Magic number: %d\n", run2(inputText))
}

func run1(input string) int {
	defer tools.Track(time.Now(), "Part 1 Time")

	b := grid.ParseBoardString(input)
	start, _ := b.Find(grid.START)
	b.AtPoint(start).Contents = BEAM
	splits := 0
	for y := 1; y <= b.MaxY; y++ {
		splits += tick(b, y)
		// b.PrintBoard()
	}

	return splits
}

func run2(input string) int {
	defer tools.Track(time.Now(), "Part 2 Time")
	b := grid.ParseBoardString(input)
	start, _ := b.Find(grid.START)
	b.AtPoint(start).Value = 1
	b.AtPoint(start).Contents = BEAM
	for y := 1; y <= b.MaxY; y++ {
		tick2(b, y)
		//b.FprintBoardValues(os.Stdout)
	}
	// sum the values of the last row
	timelines := 0
	for x := 0; x <= b.MaxX; x++ {
		timelines += b.At(x, b.MaxY).Value
	}

	return timelines
}

func tick(b *grid.Board, y int) int {
	upY := y - 1
	splits := 0
	// find all the | in previous row
	for x := 0; x <= b.MaxX; x++ {
		t := b.At(x, y)
		upT := b.At(x, upY)
		if upT.Contents == BEAM {
			if grid.IsEmpty(t.Contents) {
				t.Contents = BEAM
			} else if t.Contents == SPLITTER {
				splits++
				b.At(x-1, y).Contents = BEAM
				b.At(x+1, y).Contents = BEAM
			}
		}
	}
	return splits
}

// count the number of timelines each beam is on
// when splitting, just copy the number
// when merging 2 beams, add them
// see https://www.reddit.com/r/adventofcode/comments/1pgb377/2025_day_7_part_2_hint/ for a visualization
func tick2(b *grid.Board, y int) {
	upY := y - 1
	// find all the | in previous row
	for x := 0; x <= b.MaxX; x++ {
		t := b.At(x, y)
		upT := b.At(x, upY)
		if upT.Contents == BEAM {
			if grid.IsEmpty(t.Contents) || t.Contents == BEAM {
				t.Contents = BEAM
				t.Value += upT.Value
			} else if t.Contents == SPLITTER {
				tl := b.At(x-1, y)
				tr := b.At(x+1, y)
				tl.Contents = BEAM
				tl.Value += upT.Value
				tr.Contents = BEAM
				tr.Value += upT.Value
			}
		}
	}
}
