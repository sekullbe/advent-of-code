package main

// getting towards a generic grid implementation, not there yet

import "image"

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

type square struct {
	id int
	image.Point
	contents         rune
	adjacentToSymbol bool
}

func Pt(x, y int) image.Point {
	return image.Point{X: x, Y: y}
}

func neighborInDirection(p image.Point, dir int) (neighbor image.Point) {
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

// TODO at some point make a generalized grid implementation
// with generic contents and a parser
type grid map[image.Point]*square

type board struct {
	grid
	maxX, maxY int // min is always 0
}

func parseBoard(lines []string) *board {

	bb := board{
		grid: make(grid),
		maxX: 0,
		maxY: 0,
	}
	b := &bb
	for y, line := range lines {
		if len(line) == 0 {
			break
		}
		b.maxY = max(b.maxY, y)
		for x, tc := range line {
			b.maxX = max(b.maxX, x)
			if tc == ' ' {
				continue
			}
			p := Pt(x, y)
			t := square{Point: p, contents: tc}
			b.grid[p] = &t
		}
	}
	return b
}

// generalized implementation here would take a function and use that
// or I could produce an iterator of neighbors - see https://bitfieldconsulting.com/golang/iterators
func (b *board) checkNeighborsForSymbol(p image.Point) {
	for dir := NORTH; dir <= NORTHWEST; dir++ { // kind of gross that this depends on the order of the constants
		dp := neighborInDirection(p, dir)
		n, ok := b.grid[dp]
		if ok && isSymbol(n.contents) {
			s := b.grid[p]
			s.adjacentToSymbol = true
			b.grid[p] = s
		}
	}
}

func isSymbol(r rune) bool {
	return r != '.' && (r < '0' || r > '9')
}

func isNumber(r rune) bool {
	return r >= '0' && r <= '9'
}

func isBlank(r rune) bool {
	return r == '.'
}
