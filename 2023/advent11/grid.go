package main

// getting towards a generic grid implementation, not there yet
// keep the diagonals in this one, just don't use them
import (
	"fmt"
	"github.com/sekullbe/advent/tools"
	"image"
)

type Point image.Point

const (
	IDLE = iota
	NORTH
	NORTHEAST
	SOUTH
	SOUTHEAST
	WEST
	SOUTHWEST
	EAST
	NORTHWEST
)

// At some point I should make a generalized grid implementation with generic contents and a parser
// have to think about what the interface would be
type grid map[Point]*tile

type board struct {
	grid
	maxX, maxY   int // min is always 0
	emptyRows    []int
	emptyColumns []int
	stars        []star
}

// generalized grid impl will provide this
type baseSquare struct {
	id    int
	Point Point
}

// and a user will add something like this
type tile struct {
	baseSquare
	hasStar bool
}

type star struct {
	num int
	loc Point
}

func Pt(x, y int) Point {
	return Point{X: x, Y: y}
}

func neighborInDirection(p Point, dir int) (neighbor Point) {
	switch dir {
	case NORTH:
		return Pt(p.X, p.Y-1)
	case NORTHEAST:
		return Pt(p.X+1, p.Y-1)
	case EAST:
		return Pt(p.X+1, p.Y)
	case SOUTHEAST:
		return Pt(p.X+1, p.Y+1)
	case SOUTH:
		return Pt(p.X, p.Y+1)
	case SOUTHWEST:
		return Pt(p.X-1, p.Y+1)
	case WEST:
		return Pt(p.X-1, p.Y)
	case NORTHWEST:
		return Pt(p.X-1, p.Y-1)
	}
	return
}

func newTile(p Point, star bool) tile {
	return tile{
		baseSquare: baseSquare{Point: p},
		hasStar:    star,
	}
}

func parseBoard(lines []string) *board {

	bb := board{
		grid:         make(grid),
		maxX:         0,
		maxY:         0,
		emptyColumns: []int{},
		emptyRows:    []int{},
		stars:        []star{},
	}
	b := &bb

	starnum := 1

	for y, line := range lines {
		emptyRow := true
		if len(line) == 0 {
			break // catch extraneous blank lines
		}
		b.maxY = max(b.maxY, y)
		for x, r := range line {
			b.maxX = max(b.maxX, x)
			pt := Pt(x, y)

			t := newTile(pt, r == '#')
			if r == '#' {
				b.stars = append(b.stars, star{num: starnum, loc: pt})
				starnum++
				emptyRow = false
			}
			b.grid[pt] = &t
		}
		if emptyRow {
			b.emptyRows = append(b.emptyRows, y)
		}
	}
	return b
}

func (b *board) inRange(pt Point) bool {
	return pt.X >= 0 && pt.X <= b.maxX && pt.Y >= 0 && pt.Y <= b.maxY
}

// generalized implementation here would take a function and use that
// or I could produce an iterator of neighbors - see https://bitfieldconsulting.com/golang/iterators
func (b *board) checkNeighbors(p Point) {
	for dir := NORTH; dir <= NORTHWEST; dir++ { // kind of gross that this depends on the order of the constants
		dp := neighborInDirection(p, dir)
		n, ok := b.grid[dp]
		//...
		_ = n
		_ = ok
	}
}

func (b *board) printBoard() {
	for y := 0; y <= b.maxY; y++ {
		for x := 0; x <= b.maxX; x++ {
			_ = x
		}
		fmt.Println()
	}
}

func ManhattanDistance(p1, p2 Point) int {
	return tools.AbsInt(p1.X-p2.X) + tools.AbsInt(p1.Y-p2.Y)
}
