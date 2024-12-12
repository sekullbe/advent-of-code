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

	return 0
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
