package main

import (
	_ "embed"
	"fmt"
	"github.com/sekullbe/advent/grid"
)

//go:embed input.txt
var inputText string

func main() {
	fmt.Printf("Magic number: %d\n", run1(inputText))
	fmt.Println("-------------")
	fmt.Printf("Magic number: %d\n", run2(inputText))
}

// man i really should generalize my grid implementation from previous years
// should this one be a map of point or an array?

func run1(input string) int {

	// read it into a grid
	b := grid.ParseBoardString(input)
	matches := 0
	for point, tile := range b.Grid { // remember that the grid will be checked in arbitrary order because we're iterating a map!
		if tile.Contents == 'X' {
			// if it is X, look around it for M's
			dirs := b.CheckNeighborsForSymbol(point, 'M')
			for _, mdir := range dirs {
				mpoint := grid.NeighborInDirection(point, mdir) // we know this point is an M so see if the next two in that direction are A S
				maybeA := grid.NeighborInDirection(mpoint, mdir)
				maybeS := grid.NeighborInDirection(maybeA, mdir)
				if b.InRange(maybeA) && b.InRange(maybeS) && b.AtPoint(maybeA).Contents == 'A' && b.AtPoint(maybeS).Contents == 'S' {
					matches++
				}
			}
		}
	}
	return matches
}

func run2(input string) int {

	// read it into a grid
	b := grid.ParseBoardString(input)
	matches := 0
	for point, tile := range b.Grid { // remember that the grid will be checked in arbitrary order because we're iterating a map!
		if tile.Contents == 'A' {
			couldBeMAS := 0
			// check all 4 directions for M and S
			for d := grid.NORTHEAST; d <= grid.NORTHWEST; d += 2 {
				if b.AtPoint(grid.NeighborInDirection(point, d)).Contents == 'M' && b.AtPoint(grid.NeighborInDirection(point, grid.Opposite(d))).Contents == 'S' {
					couldBeMAS++
				}
			}
			if couldBeMAS >= 2 {
				matches++
			}
		}
	}
	return matches
}
