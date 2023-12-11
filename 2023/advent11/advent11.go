package main

import (
	_ "embed"
	"fmt"
	"github.com/sekullbe/advent/parsers"
)

//go:embed input.txt
var inputText string

func main() {
	fmt.Printf("Magic number: %d\n", run(inputText, 2))
	fmt.Println("-------------")
	fmt.Printf("Magic number: %d\n", run(inputText, 1_000_000))
}

func run(input string, expansionConstant int) int {
	// parse into a grid - can use the grid from day 10
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
	// we can just move each star down and to the right based on how many empty rows/cols there are
	expandedStars := []star{}
	for _, s := range b.stars {
		expandedStars = append(expandedStars, star{
			num: s.num,
			loc: Point{
				// remember we don't *add* N rows, we replace, so add *N-1* rows
				X: s.loc.X + (expansionConstant-1)*b.emptyColsLeft(s.loc.X),
				Y: s.loc.Y + (expansionConstant-1)*b.emptyRowsAbove(s.loc.Y),
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
