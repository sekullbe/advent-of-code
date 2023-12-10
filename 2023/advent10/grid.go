package main

// getting towards a generic grid implementation, not there yet
// keep the diagonals in this one, just don't use them
import (
	"fmt"
	"image"
)

// so I don't have to keep typing 'image.Point' but I don't want to bring in the whole package as "."
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
	maxX, maxY int // min is always 0
	start      Point
}

// generalized grid impl will provide this
type baseSquare struct {
	id    int
	Point Point
}

// and a user will add something like this
type tile struct {
	baseSquare
	pipe           int
	traversed      bool // probably redundant
	stepsFromStart int

	//neighbors? or just calculate it

}

const (
	PIPE_NS = iota
	PIPE_EW
	PIPE_NE
	PIPE_NW
	PIPE_SW
	PIPE_SE
	PIPE_NONE
	PIPE_START // we don't know the shape of this pipe
)

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

func newTile(p Point, pipe int) tile {
	return tile{
		baseSquare: baseSquare{Point: p},
		pipe:       pipe,
	}
}

var pipeType = map[rune]int{
	'|': PIPE_NS,
	'-': PIPE_EW,
	'L': PIPE_NE,
	'J': PIPE_NW,
	'7': PIPE_SW,
	'F': PIPE_SE,
	'.': PIPE_NONE,
	'S': PIPE_START,
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
			break // catch extraneous blank lines
		}
		b.maxY = max(b.maxY, y)
		for x, r := range line {
			b.maxX = max(b.maxX, x)
			pt := Pt(x, y)
			t := newTile(pt, pipeType[r])
			if t.pipe == PIPE_START {
				b.start = pt
			}
			b.grid[pt] = &t
		}
	}
	return b
}

// we're in 'curr' and came from 'prev', what is the next point?
func (b *board) followPipe(curr, prev Point) Point {
	n1, n2 := b.pipeNeighbors(curr)
	if n1 == prev {
		return n2
	}
	return n1
}

// given a pipe which are its neighbors
func (b *board) pipeNeighbors(pt Point) (one Point, two Point) {
	if !b.inRange(pt) {
		return Pt(-99, -99), Pt(-99, -99)
	}
	pipe := b.grid[pt].pipe
	switch pipe {
	case PIPE_NS:
		return neighborInDirection(pt, NORTH), neighborInDirection(pt, SOUTH)
	case PIPE_EW:
		return neighborInDirection(pt, EAST), neighborInDirection(pt, WEST)
	case PIPE_NE:
		return neighborInDirection(pt, NORTH), neighborInDirection(pt, EAST)
	case PIPE_NW:
		return neighborInDirection(pt, NORTH), neighborInDirection(pt, WEST)
	case PIPE_SW:
		return neighborInDirection(pt, SOUTH), neighborInDirection(pt, WEST)
	case PIPE_SE:
		return neighborInDirection(pt, SOUTH), neighborInDirection(pt, EAST)
	case PIPE_START:
		// call pipeNeighbors for all 4 directions and return the two that pipe to this point
		var possibles [4]struct{ n, p1, p2 Point }
		possibles[0].n = neighborInDirection(pt, NORTH)
		possibles[0].p1, possibles[0].p2 = b.pipeNeighbors(possibles[0].n)
		possibles[1].n = neighborInDirection(pt, SOUTH)
		possibles[1].p1, possibles[1].p2 = b.pipeNeighbors(possibles[1].n)
		possibles[2].n = neighborInDirection(pt, EAST)
		possibles[2].p1, possibles[2].p2 = b.pipeNeighbors(possibles[2].n)
		possibles[3].n = neighborInDirection(pt, WEST)
		possibles[3].p1, possibles[3].p2 = b.pipeNeighbors(possibles[3].n)
		neighbors := []Point{}
		for _, possible := range possibles {
			if b.inRange(possible.p1) && possible.p1 == pt {
				neighbors = append(neighbors, possible.n)
			}
			if b.inRange(possible.p2) && possible.p2 == pt {
				neighbors = append(neighbors, possible.n)
			}
		}
		return neighbors[0], neighbors[1]
	default:
		// not a pipe so it can't neighbor anything; return a garbage point
		return Pt(-99, -99), Pt(-99, -99)
	}
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

func (b *board) setSteps(pt Point, steps int) {
	t := b.grid[pt]
	t.stepsFromStart = steps
	b.grid[pt] = t
}

func (b *board) getSteps(pt Point) (steps int) {
	return b.grid[pt].stepsFromStart
}

func (b *board) printBoard() {
	for y := 0; y <= b.maxY; y++ {
		for x := 0; x <= b.maxX; x++ {
			switch b.grid[Pt(x, y)].pipe {
			case PIPE_NS:
				fmt.Print("│")
			case PIPE_EW:
				fmt.Print("─")
			case PIPE_NE:
				fmt.Print("└")
			case PIPE_NW:
				fmt.Print("┘")
			case PIPE_SW:
				fmt.Print("┐")
			case PIPE_SE:
				fmt.Print("┌")
			case PIPE_START:
				fmt.Print("S")
			case PIPE_NONE:
				fmt.Print("·")

			}
		}
		fmt.Println()
	}
}
