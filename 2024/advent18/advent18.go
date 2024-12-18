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

func run1(input string) int {

	b := initBoard(70, 70)
	bytes := parseBytes(parsers.SplitByLines(input))
	dropBytes(b, bytes[:1024]) // bytes 0-1023
	//b.PrintBoard()

	// now run a pathfinding algorithm on the board
	// unlike previous versions there are no hard walls so need to check inbounds
	_, steps, found := b.FindPath(grid.Pt(0, 0), grid.Pt(70, 70))
	if !found {
		panic("no solution found")
	}

	return steps
}

func run2(input string) int {

	return 0
}

func initBoard(maxX, maxY int) *grid.Board {
	b := &grid.Board{Grid: make(grid.Grid), MaxX: maxX, MaxY: maxY}

	for y := 0; y <= maxY; y++ {
		for x := 0; x <= maxX; x++ {
			pt := geometry.Point2{X: x, Y: y}
			t := grid.NewTile(pt, grid.EMPTY)
			b.Grid[pt] = &t
		}
	}
	return b
}

func parseBytes(lines []string) []geometry.Point2 {
	bytes := make([]geometry.Point2, len(lines))
	for i, line := range lines {
		pt := geometry.Point2{}
		_, err := fmt.Sscanf(line, "%d,%d", &pt.X, &pt.Y)
		if err != nil {
			continue
		}
		bytes[i] = pt
	}
	return bytes
}

func dropBytes(b *grid.Board, bytes []geometry.Point2) {
	for _, pt := range bytes {
		t := b.AtPoint(pt)
		t.Contents = '#'
		//b.Grid[pt]=t
	}
}
