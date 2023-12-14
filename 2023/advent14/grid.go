package main

// getting towards a generic grid implementation, not there yet
// keep the diagonals in this one, just don't use them
import (
	"errors"
	"fmt"
	"github.com/sekullbe/advent/tools"
	"image"
	"io"
	"os"
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

// rocks
const (
	SPACE = iota // .
	ROUND        // O
	CUBE         // #
)

// At some point I should make a generalized grid implementation with generic contents and a parser
// have to think about what the interface would be
type grid map[Point]*tile

type board struct {
	grid
	maxX, maxY int // min is always 0
	northest   []int
}

// generalized grid impl will provide this
type baseSquare struct {
	id    int
	Point Point
}

// and a user will add something like this
type tile struct {
	baseSquare
	rock int
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

func newTile(p Point, rock int) tile {
	return tile{
		baseSquare: baseSquare{Point: p},
		rock:       rock,
	}
}

func parseRock(r rune) int {
	switch r {
	case 'O':
		return ROUND
	case '#':
		return CUBE
	case '.':
		return SPACE
	default:
		panic("don't know what kind of a rock this is:" + string(r))
	}
}

func parseBoard(lines []string) *board {

	bb := board{
		grid:     make(grid),
		maxX:     0,
		maxY:     0,
		northest: []int{},
	}
	b := &bb

	for y, line := range lines {
		if len(line) == 0 {
			break // catch extraneous blank lines
		}
		b.northest = append(b.northest, -1)
		b.maxY = max(b.maxY, y)
		for x, r := range line {
			b.maxX = max(b.maxX, x)
			pt := Pt(x, y)

			t := newTile(pt, parseRock(r))
			b.grid[pt] = &t
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
	b.fprintBoard(os.Stdout)
}

func (b *board) fprintBoard(w io.Writer) {
	for y := 0; y <= b.maxY; y++ {
		for x := 0; x <= b.maxX; x++ {
			switch b.At(x, y).rock {
			case ROUND:
				fmt.Fprint(w, "O")
			case SPACE:
				fmt.Fprint(w, "·")
			case CUBE:
				fmt.Fprintf(w, "#")
			}
		}
		fmt.Fprintln(w)
	}
}

func ManhattanDistance(p1, p2 Point) int {
	return tools.AbsInt(p1.X-p2.X) + tools.AbsInt(p1.Y-p2.Y)
}

func (b *board) AtPoint(p Point) *tile {
	return b.grid[p]
}
func (b *board) At(x, y int) *tile {
	return b.grid[Pt(x, y)]
}

func (b *board) moveRock(from, to Point) error {

	rock := b.grid[Pt(from.X, from.Y)].rock

	if b.AtPoint(to).rock != SPACE && from != to {
		return errors.New("target tile is occupied")
	}
	if rock == CUBE {
		return errors.New("can't move a cube rock")
	}
	if rock == SPACE {
		return errors.New("can't move an empty space")
	}

	b.grid[Pt(from.X, from.Y)].rock = SPACE
	b.grid[Pt(to.X, to.Y)].rock = rock
	return nil
}

/*
could use something like this that calls a fn for each point?
	for y := 0; y <= b.maxY; y++ {
		for x := 0; x < b.maxX; x++ {
          fn(b.At(x,y))
		}
	}
*/
