package main

import (
	_ "embed"
	"fmt"
	"github.com/sekullbe/advent/parsers"
	"image"
	"image/color"
)

//go:embed input.txt
var inputText string

func main() {
	fmt.Printf("Magic number: %d\n", run1(inputText))
	fmt.Println("-------------")
	fmt.Printf("Magic number: %d\n", run2(inputText))
}

type trench struct {
	dir    int
	length int
	color  color.RGBA64
}

func run1(input string) int {

	// we're not parsing a board, so just create a blank one
	// can we know the size of the board in advance? I don't think so
	// that's ok, we can make tiles on the fly
	lines := parsers.SplitByLines(input)
	trenches := []trench{}
	// a line looks like this: `R 6 (#70c710)`
	diggerLoc := Pt(0, 0)
	b := NewBoard()
	initialTile := NewTile(diggerLoc, color.RGBA64{0, 0, 0, 0}) // A=0 is a dug trench
	b.Grid[diggerLoc] = &initialTile
	polygonPoints := []image.Point{diggerLoc}
	for nl, line := range lines {
		t := parseInstruction(line)
		trenches = append(trenches, t) // ???
		// now dig the trench, starting at diggerLoc and going L spaces in direction D
		for i := 1; i <= t.length; i++ {
			newDigPt := NeighborInDirection(diggerLoc, t.dir)
			nt := NewTile(newDigPt, t.color)
			if _, ok := b.Grid[newDigPt]; ok && nl != len(lines)-1 { // ignore closing the loop
				fmt.Printf("New trench '%s' crosses existing trench at point %v\n", line, newDigPt)
			}
			b.Grid[newDigPt] = &nt
			b.MaxX = max(newDigPt.X, b.MaxX)
			b.MaxY = max(newDigPt.Y, b.MaxY)
			b.MinX = min(newDigPt.X, b.MinX) // not sure if we go negative, but let's handle it just in case
			b.MinY = min(newDigPt.Y, b.MinY)
			diggerLoc = newDigPt
		}
		// add an edge to the polygon
		polygonPoints = append(polygonPoints, diggerLoc)
	}

	// now use the library from advent10 to see if each point is in or out of the polygon
	// what color should a new dug tile be?  does it matter?

	inside := 0
	digs := len(b.Grid)
	pf := NewPolyfence(polygonPoints)
	for x := b.MinX; x <= b.MaxX; x++ {
		for y := b.MinY; y <= b.MaxY; y++ {
			// unsparseify the grid
			if _, ok := b.Grid[Pt(x, y)]; !ok {
				nt := NewTile(Pt(x, y), color.RGBA64{0, 0, 0, 255}) // A= 255 is no trench at all
				b.Grid[Pt(x, y)] = &nt
			}
			if pf.Inside(Pt(x, y)) && b.Grid[Pt(x, y)].Contents.A != 0 {
				t := b.Grid[Pt(x, y)]
				t.Contents.A = 1 // A=1 is an interior tile that is dug out
				b.Grid[Pt(x, y)] = t
				inside += 1
			}
		}
	}
	//b.printBoard()
	// if an edge is right or up it is inside, left or down it is outside, so count trenches and insides separately

	return inside + digs
}

func parseInstruction(line string) trench {
	t := trench{}
	var ds string
	n, err := fmt.Sscanf(line, "%s %d (#%02x%02x%02x)", &ds, &t.length, &t.color.R, &t.color.G, &t.color.B)
	if err != nil || n != 5 {
		panic("can't parse line:" + line)
	}
	t.dir = parseDirection(ds)
	return t
}

func parseDirection(d string) int {
	switch d {
	case "U":
		return UP
	case "R":
		return RIGHT
	case "D":
		return DOWN
	case "L":
		return LEFT
	}
	panic("can't parse direction: " + d)
}

func run2(input string) int {

	// betcha we use that color thing here

	return 0
}
