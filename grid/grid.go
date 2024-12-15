package grid

// hard to use this directly- needs to be more generalizable
// but the dream continues.

import (
	"fmt"
	"github.com/sekullbe/advent/geometry"
	"github.com/sekullbe/advent/parsers"
	"github.com/sekullbe/advent/tools"
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

// common directional instructions
var DirRunes = map[rune]int{'^': NORTH, '>': EAST, 'v': SOUTH, 'V': SOUTH, '<': WEST}

var FourDirections = [...]int{NORTH, EAST, SOUTH, WEST}

// tile contents
const EMPTY = '.'

type Board struct {
	Grid
	MaxX, MaxY     int // min is always 0
	startX, startY int
}

// working towards a generic Grid implementation with generic contents and a parser
// have to think about what the interface would be
// probably woud have methods on Tile so a grid[T] would contain anything implementing something like 'Tiler'
type Grid map[geometry.Point]*Tile

type BaseTile struct {
	Id    int
	Point geometry.Point
}

type Tile struct {
	BaseTile
	Contents rune
	Value    int // we store numbers often enough
	// these two are used often enough, but this really ought to be somehow generic
	Counter   int
	Traversed bool
	Offboard  bool
}

// maybe this belongs in geometry package
func Pt(x, y int) geometry.Point {
	return geometry.Point{X: x, Y: y}
}

func NewTile(p geometry.Point, content rune) Tile {
	return Tile{
		BaseTile: BaseTile{Point: p},
		Contents: content,
		Value:    int(content - '0'),
	}
}

func ParseBoardString(fullBoard string) *Board {
	return ParseBoard(parsers.SplitByLines(fullBoard))
}

func ParseBoard(lines []string) *Board {

	bb := Board{
		Grid: make(Grid),
		MaxX: 0,
		MaxY: 0,
	}
	b := &bb
	for y, line := range lines {
		if len(line) == 0 {
			break
		}
		b.MaxY = max(b.MaxY, y)
		for x, tc := range line {
			b.MaxX = max(b.MaxY, x)
			if tc == ' ' {
				continue
			}
			p := Pt(x, y)
			t := NewTile(p, tc)
			b.Grid[p] = &t
		}
	}
	return b
}

func NeighborInDirection(p geometry.Point, dir int) (neighbor geometry.Point) {
	switch dir % 8 {
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
	default:
		panic("unknown direction") // with the %8 this ought not be possible
	}

}
func (b *Board) InRange(pt geometry.Point) bool {
	return pt.X >= 0 && pt.X <= b.MaxX && pt.Y >= 0 && pt.Y <= b.MaxY
}

// look in all directions around p and return the list of directions in which the symbol was found
func (b *Board) CheckNeighborsForSymbol(p geometry.Point, symbol rune) []int {
	observedDirections := []int{}
	for dir := NORTH; dir <= NORTHWEST; dir++ { // kind of gross that this depends on the order of the constants
		dp := NeighborInDirection(p, dir)
		n, ok := b.Grid[dp]
		if ok && n.Contents == symbol {
			observedDirections = append(observedDirections, dir)
		}
	}
	return observedDirections
}

// get onboard neighbors in square directions- no diagonals.
func (b Board) GetSquareNeighbors(p geometry.Point) []geometry.Point {
	ns := []geometry.Point{}
	for d := NORTH; d <= WEST; d += 2 {
		np := NeighborInDirection(p, d)
		if b.InRange(np) {
			ns = append(ns, np)
		}
	}
	return ns
}

// Doesn't check to see if a point is offboard
func (b Board) GetSquareNeighborsNoChecks(p geometry.Point) []geometry.Point {
	ns := []geometry.Point{}
	for d := NORTH; d <= WEST; d += 2 {
		np := NeighborInDirection(p, d)
		ns = append(ns, np)
	}
	return ns
}

func (b *Board) PrintBoard() {
	b.FprintBoard(os.Stdout)
}

func (b *Board) FprintBoard(w io.Writer) {
	for y := 0; y <= b.MaxY; y++ {
		for x := 0; x <= b.MaxX; x++ {
			t := b.At(x, y)
			if t == nil {
				fmt.Fprint(w, "·"+
					"")
			} else {
				switch t.Contents {
				case EMPTY:
					fmt.Fprint(w, "·")
				default:
					fmt.Fprintf(w, "%c", t.Contents)
				}
			}
		}
		fmt.Fprintln(w)
	}
}

func ManhattanDistance(p1, p2 geometry.Point) int {
	return tools.AbsInt(p1.X-p2.X) + tools.AbsInt(p1.Y-p2.Y)
}

// should it return an error or a tile guaranteed to match nothing?
func (b *Board) AtPoint(p geometry.Point) *Tile {
	if !b.InRange(p) {
		t := NewTile(p, EMPTY)
		t.Offboard = true
		return &t
	}
	return b.Grid[p]
}

// if this point is offboard, wrap until it's onboard
func (b *Board) AtPointWrapped(p geometry.Point) *Tile {
	np := Pt(wrapmod(p.X, b.MaxX+1), wrapmod(p.Y, b.MaxY+1))
	return b.Grid[np]
}

func (b *Board) At(x, y int) *Tile {
	p := geometry.NewPoint2(x, y)
	if !b.InRange(p) {
		t := NewTile(p, EMPTY)
		t.Offboard = true
		return &t
	}
	return b.Grid[Pt(x, y)]
}

// move the contents of the tile at point p in direction dir, only if the target tile is not empty
// return true if move was successful
func (b *Board) SlideTile(p geometry.Point, dir int) bool {
	t := b.AtPoint(p)
	if IsEmpty(t.Contents) {
		return true // moving nothing is a successful no-op
	}
	nt := b.AtPoint(NeighborInDirection(p, dir))
	if !b.InRange(nt.Point) {
		return false // can't move offboard
	}
	if !IsEmpty(nt.Contents) {
		return false // something is already there
	}
	nt.Contents = t.Contents
	t.Contents = EMPTY
	return true
}

func Clockwise(dir int, ticks int) int {
	return (dir + ticks) % 8
}

func CounterClockwise(dir int, ticks int) int {
	return (dir + (8 - ticks)) % 8
}

func Opposite(dir int) int {
	return Clockwise(dir, 4)
}

func ValidDir(dir int) bool {
	return dir >= NORTH && dir <= NORTHWEST
}

func wrapmod(a, b int) int {
	return (a%b + b) % b
}

func IsSymbol(r rune) bool {
	return !(IsBlank(r) || IsNumber(r))
}

func IsNumber(r rune) bool {
	return r >= '0' && r <= '9'
}

func IsBlank(r rune) bool {
	return r == EMPTY || r == ' ' || r == 0 // uninitialized == empty
}

func IsEmpty(r rune) bool {
	return IsBlank(r)
}
