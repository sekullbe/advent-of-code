package main

import (
	_ "embed"
	"fmt"
	"github.com/ctessum/geom"
	"github.com/sekullbe/advent/parsers"
	"github.com/sekullbe/advent/tools"
	"image"
	"image/color"
	"strconv"
)

//go:embed input.txt
var inputText string

func main() {
	fmt.Printf("Magic number: %d\n", run1(inputText))
	fmt.Printf("Magic number: %d\n", run1better(inputText))
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

	// I can actually do this without a grid
	// start at 0,0 and add points to the polygon as we go
	// iterate over all the points
	// if an edge is right or up it is inside, left or down it is outside, so count trenches and insides separately

	diggerLoc := Pt(0, 0)
	var minX, minY, maxX, maxY int
	outsideDigs := 0

	polygonPoints := []image.Point{diggerLoc}
	for _, line := range parsers.SplitByLines(input) {
		dir, dist := parsePart2Instruction(line)
		diggerLoc = newPoint(diggerLoc, dir, dist)
		maxX = max(diggerLoc.X, maxX)
		maxY = max(diggerLoc.Y, maxY)
		minX = min(diggerLoc.X, minX) // not sure if we go negative, but let's handle it just in case
		minY = min(diggerLoc.Y, minY)
		polygonPoints = append(polygonPoints, diggerLoc)
		// if it's left or down, add the length to outsideDigs
		if dir == LEFT || dir == DOWN {
			outsideDigs += dist
		}
	}

	// this is a much faster algorithm
	gpts := []geom.Point{}
	for _, ipt := range polygonPoints {
		gpts = append(gpts, geom.Point{float64(ipt.X), float64(ipt.Y)})
	}
	gpoly := geom.Polygon{gpts}
	return int(gpoly.Area()) + outsideDigs + 1

	/*
		// eh... this is going to take months
			pf := NewPolyfence(polygonPoints)
			inside := 0
			for x := minX; x <= maxX; x++ {
				fmt.Printf(".")
				for y := minY; y <= maxY; y++ {
					if pf.Inside(Pt(x, y)) {
						inside += 1
					}
				}
			}
			return inside + outsideDigs
	*/

}

func parsePart2Instruction(line string) (dir, dist int) {

	var xdir string
	var xdist int
	var hex string
	n, err := fmt.Sscanf(line, "%s %d %s", &xdir, &xdist, &hex)
	if n != 3 || err != nil {
		panic("can't parse " + line)
	}
	hexdist := hex[2:7]
	hexdir := string(hex[7])
	dist = int(tools.Must(strconv.ParseInt(hexdist, 16, 0)))
	switch hexdir {
	case "0":
		dir = RIGHT
	case "1":
		dir = DOWN
	case "2":
		dir = LEFT
	case "3":
		dir = UP
	default:
		panic("bad dir")
	}
	return dir, dist
}

func newPoint(p image.Point, dir int, dist int) image.Point {

	switch dir {
	case UP:
		return p.Add(Pt(0, -dist))
	case RIGHT:
		return p.Add(Pt(dist, 0))
	case DOWN:
		return p.Add(Pt(0, dist))
	case LEFT:
		return p.Add(Pt(-dist, 0))
	}
	return p
}

func run1better(input string) int {

	// I can actually do this without a grid
	// start at 0,0 and add points to the polygon as we go
	// iterate over all the points
	// if an edge is right or up it is inside, left or down it is outside, so count trenches and insides separately

	diggerLoc := Pt(0, 0)
	var minX, minY, maxX, maxY int
	outsideDigs := 0

	polygonPoints := []image.Point{diggerLoc}
	for _, line := range parsers.SplitByLines(input) {
		dir, dist := parsePart1Instruction(line)
		diggerLoc = newPoint(diggerLoc, dir, dist)
		maxX = max(diggerLoc.X, maxX)
		maxY = max(diggerLoc.Y, maxY)
		minX = min(diggerLoc.X, minX) // not sure if we go negative, but let's handle it just in case
		minY = min(diggerLoc.Y, minY)
		polygonPoints = append(polygonPoints, diggerLoc)
		// if it's left or down, add the length to outsideDigs
		if dir == LEFT || dir == DOWN {
			outsideDigs += dist
		}
	}

	// this is a much faster algorithm
	gpts := []geom.Point{}
	for _, ipt := range polygonPoints {
		gpts = append(gpts, geom.Point{float64(ipt.X), float64(ipt.Y)})
	}
	gpoly := geom.Polygon{gpts}
	return int(gpoly.Area()) + outsideDigs + 1
}

func parsePart1Instruction(line string) (dir, dist int) {

	var ds, hex string
	n, err := fmt.Sscanf(line, "%s %d %s", &ds, &dist, &hex)
	if n != 3 || err != nil {
		panic("can't parse " + line)
	}
	return parseDirection(ds), dist
}
