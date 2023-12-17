package main

// getting towards a generic Grid implementation, not there yet
import (
	"fmt"
	"github.com/sekullbe/advent/tools"
	"image"
	"io"
	"os"
)

const (
	NORTH = iota
	NORTHEAST
	EAST
	SOUTHEAST
	SOUTH
	SOUTHWEST
	WEST
	NORTHWEST
)

// tile contents
const (
	EMPTY = iota
)

// At some point I should make a generalized Grid implementation with generic contents and a parser
// have to think about what the interface would be
type Grid map[image.Point]*Tile

type Board struct {
	Grid
	MaxX, MaxY int // min is always 0

}

// generalized Grid impl will provide this
type BaseSquare struct {
	Id int
	Pt image.Point
}

// and a user will add something like this
type Tile struct {
	BaseSquare
	Contents int // heatloss in this case
}

func Pt(x, y int) image.Point {
	return image.Point{X: x, Y: y}
}

func NeigborInDirection(p image.Point, dir int) (neighbor image.Point) {
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

func NewTile(p image.Point, content int) Tile {
	return Tile{
		BaseSquare: BaseSquare{Pt: p},
		Contents:   content,
	}
}

func parseContents(r rune) int {
	switch r {
	case ' ':
		return EMPTY
	default:
		return int(r - '0')
	}
}

func ParseBoard(lines []string) *Board {

	bb := Board{
		Grid: make(Grid),
		MaxX: 0,
		MaxY: 0,
	}
	b := &bb
	b.MaxX = len(lines[0]) - 1
	b.MaxY = len(lines) - 1

	for y, line := range lines {
		if len(line) == 0 {
			break // catch extraneous blank lines
		}
		for x, tc := range line {
			pt := Pt(x, y)
			t := NewTile(pt, parseContents(tc))
			b.Grid[pt] = &t
		}
	}
	return b
}

func (b *Board) InRange(pt image.Point) bool {
	return pt.X >= 0 && pt.X <= b.MaxX && pt.Y >= 0 && pt.Y <= b.MaxY
}

// generalized implementation here would take a function and use that
// or I could produce an iterator of neighbors - see https://bitfieldconsulting.com/golang/iterators
func (b *Board) CheckNeighbors(p image.Point) {
	for dir := NORTH; dir <= NORTHWEST; dir++ { // kind of gross that this depends on the order of the constants
		dp := NeigborInDirection(p, dir)
		n, ok := b.Grid[dp]
		//...
		_ = n
		_ = ok
	}
}

func (b *Board) printBoard() {
	b.FprintBoard(os.Stdout)
}

func (b *Board) FprintBoard(w io.Writer) {
	for y := 0; y <= b.MaxY; y++ {
		for x := 0; x <= b.MaxX; x++ {
			switch b.At(x, y).Contents {
			case EMPTY:
				fmt.Fprint(w, "·")
			default:
				fmt.Fprint(w, b.At(x, y).Contents)
			}
		}
		fmt.Fprintln(w)
	}
}

func ManhattanDistance(p1, p2 image.Point) int {
	return tools.AbsInt(p1.X-p2.X) + tools.AbsInt(p1.Y-p2.Y)
}

func (b *Board) AtPoint(p image.Point) *Tile {
	return b.Grid[p]
}
func (b *Board) At(x, y int) *Tile {
	return b.Grid[Pt(x, y)]
}

func Clockwise(dir int, ticks int) int {
	return (dir + ticks) % 8
}

func CounterClockwise(dir int, ticks int) int {
	return (dir + (8 - ticks)) % 8
}

func ValidDir(dir int) bool {
	return dir >= NORTH && dir <= NORTHWEST
}

/*
could use something like this that calls a fn for each point?
	for y := 0; y <= b.MaxY; y++ {
		for x := 0; x < b.MaxX; x++ {
          fn(b.At(x,y))
		}
	}
*/
