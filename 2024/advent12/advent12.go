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

type region map[geometry.Point]bool

func run1(input string) int {

	b := grid.ParseBoard(parsers.SplitByLines(input))
	regions := []region{}
	processedPoints := make(map[geometry.Point]bool)

	for point, _ := range b.Grid {
		if _, seen := processedPoints[point]; !seen {
			r := floodfillMatchingLetter(b, point)
			regions = append(regions, r)
			for g := range r {
				processedPoints[g] = true
			}
		}
	}
	totalScore := 0
	for _, r := range regions {
		area := len(r)
		perimeter := 0
		var letter rune
		for point, _ := range r {
			letter = b.AtPoint(point).Contents
			for _, np := range b.GetSquareNeighborsNoChecks(point) {
				if !b.InRange(np) || b.AtPoint(np).Contents != letter {
					perimeter += 1
				}
			}
		}
		totalScore += area * perimeter
		//log.Printf("Region %c: area %d perimeter %d, score=%d", letter, area, perimeter, area*perimeter)
	}
	return totalScore
}

func run2(input string) int {

	b := grid.ParseBoard(parsers.SplitByLines(input))
	regions := []region{}
	processedPoints := make(map[geometry.Point]bool)

	for point, _ := range b.Grid {
		if _, seen := processedPoints[point]; !seen {
			r := floodfillMatchingLetter(b, point)
			regions = append(regions, r)
			for g := range r {
				processedPoints[g] = true
			}
		}
	}
	totalScore := 0
	for _, r := range regions {
		area := len(r)
		sides := computeSidesOfRegion(b, r)
		totalScore += area * sides

		//log.Printf("Region: area %d sides %d, score=%d", area, sides, area*sides)
	}
	return totalScore
}

func floodfillMatchingLetter(b *grid.Board, start geometry.Point) region {

	region := make(region)
	floodFillFind(b, start, region)
	return region
}

// starting at Point, find all reachable tiles with the letter in start point
func floodFillFind(b *grid.Board, start geometry.Point, region region) {
	if region[start] {
		return
	}
	startTile := b.AtPoint(start)
	letter := startTile.Contents
	region[start] = true

	for _, n := range b.GetSquareNeighbors(start) {
		if b.AtPoint(n).Contents == letter {
			floodFillFind(b, n, region)
		}
	}
}

func computeSidesOfRegion(b *grid.Board, r region) int {

	// a region has the same number of sides as corners
	// for each tile look N&E,S&E,S&W,N&W
	// eg N&E
	// tile is a corner if N != T && E != T --OR-- N=T && E=T && NE != T
	/*
	  These are both corners, looking northeast
	   ?A?	   ?XA
	   ?XA	   ?XX
	  Not a corner looking northeast, because the edge continues to another tile
	   ?A?  ?X?
	   ?XX  ?XA
	*/
	corners := 0
	for point, _ := range r {
		tile := b.AtPoint(point)
		letter := tile.Contents
		for _, direction := range grid.FourDirections {
			t1 := b.AtPoint(grid.NeighborInDirection(point, direction))                    //eg N
			td := b.AtPoint(grid.NeighborInDirection(point, grid.Clockwise(direction, 1))) // eg NE
			t2 := b.AtPoint(grid.NeighborInDirection(point, grid.Clockwise(direction, 2))) // eg E
			// could simplify but i'll probably need to look at this in debugger
			corner := (t1.Contents != letter && t2.Contents != letter) || (t1.Contents == letter && t2.Contents == letter && td.Contents != letter)
			if corner {
				corners++
			}
		}

	}
	return corners
}
