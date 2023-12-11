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

func run1(input string) int {

	// parse into a grid - can use the 4-way from day 10
	// could probably do this without an explicit map but it's easier to think of that way
	b := parseBoard(parsers.SplitByLines(input))
	// we know the empty rows, now find the empty columns
	for x := 0; x <= b.maxX; x++ {
		emptyCol := true
		for y := 0; y <= b.maxY; y++ {
			if b.grid[Pt(x, y)].hasStar {
				emptyCol = false
				break
			}
		}
		if emptyCol {
			b.emptyColumns = append(b.emptyColumns, x)
		}
	}
	// make a new map where each such row or column counts double
	// write something that copies every point with X or Y > N by +1
	// scan the map... each point records how many dX and dY it has
	// then make a new map where each point is recorded in its new location
	// this map will be sparse but that is ok
	// does it even need to be a map? a list of stars will do
	expandedStars := []star{}
	for _, s := range b.stars {
		expandedStars = append(expandedStars, star{
			num: s.num,
			loc: Point{
				X: s.loc.X + b.emptyColsLeft(s.loc.X),
				Y: s.loc.Y + b.emptyRowsAbove(s.loc.Y),
			},
		})
	}
	numStars := len(expandedStars)
	sumPaths := 0
	for i := 0; i < numStars; i++ {
		for j := i + 1; j < numStars; j++ {
			sumPaths += ManhattanDistance(expandedStars[i].loc, expandedStars[j].loc)
		}
	}

	// we're just going to count manhattan distance between each pair
	// figure out how to do the pairs thing later- probably nested loops
	// that's a lot of counting but computers are good at that

	return sumPaths
}

func (b *board) emptyRowsAbove(y int) int {
	ra := 0
	for _, row := range b.emptyRows {
		if row < y {
			ra++
		}
	}
	return ra
}

func (b *board) emptyColsLeft(x int) int {
	cl := 0
	for _, col := range b.emptyColumns {
		if col < x {
			cl++
		}
	}
	return cl
}

func run2(input string) int {

	return 0
}
