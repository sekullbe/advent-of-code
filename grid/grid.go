package grid

// hard to use this directly- needs to be more generalizable
// but the dream continues.

import (
	"fmt"
	"github.com/sekullbe/advent/parsers"
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
const EMPTY = '.'

type Board struct {
	Grid
	MaxX, MaxY     int // min is always 0
	startX, startY int
}

// working towards a generalized Grid implementation with generic contents and a parser
// have to think about what the interface would be
type Grid map[image.Point]*Tile

// generalized Grid impl will provide this
type BaseTile struct {
	Id    int
	Point image.Point
}

type Tile struct {
	BaseTile
	Contents rune
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

func NewTile(p image.Point, content rune) Tile {
	return Tile{
		BaseTile: BaseTile{Point: p},
		Contents: content,
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

func NeighborInDirection(p image.Point, dir int) (neighbor image.Point) {
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
func (b *Board) InRange(pt image.Point) bool {
	return pt.X >= 0 && pt.X <= b.MaxX && pt.Y >= 0 && pt.Y <= b.MaxY
}

// look in all directions around p and return the list of directions in which the symbol was found
func (b *Board) CheckNeighborsForSymbol(p image.Point, symbol rune) []int {
	observedDirections := []int{}
	for dir := NORTH; dir <= NORTHWEST; dir++ { // kind of gross that this depends on the order of the constants
		dp := neighborInDirection(p, dir)
		n, ok := b.Grid[dp]
		if ok && n.Contents == symbol {
			observedDirections = append(observedDirections, dir)
		}
	}
	return observedDirections
}

// get onboard neighbors in square directions- no diagonals.
func (b Board) GetSquareNeighbors(p image.Point) []image.Point {
	ns := []image.Point{}
	for d := NORTH; d <= WEST; d += 2 {
		np := NeighborInDirection(p, d)
		if b.InRange(np) {
			ns = append(ns, np)
		}
	}
	return ns
}

// Doesn't check to see if a point is offboard
func (b Board) GetSquareNeighborsNoChecks(p image.Point) []image.Point {
	ns := []image.Point{}
	for d := NORTH; d <= WEST; d += 2 {
		np := NeighborInDirection(p, d)
		ns = append(ns, np)
	}
	return ns
}

func (b *Board) printBoard() {
	b.FprintBoard(os.Stdout)
}

func (b *Board) FprintBoard(w io.Writer) {
	for y := 0; y <= b.MaxY; y++ {
		for x := 0; x <= b.MaxX; x++ {
			t := b.At(x, y)
			switch t.Contents {
			case EMPTY:
				fmt.Fprint(w, "·")
			default:
				fmt.Fprintf(w, "%c", t.Contents)
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

// if this point is offboard, wrap until it's onboard
func (b *Board) AtPointWrapped(p image.Point) *Tile {
	np := Pt(wrapmod(p.X, b.MaxX+1), wrapmod(p.Y, b.MaxX+1))
	return b.Grid[np]
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

func wrapmod(a, b int) int {
	return (a%b + b) % b
}

func isSymbol(r rune) bool {
	return !(isBlank(r) || isNumber(r))
}

func isNumber(r rune) bool {
	return r >= '0' && r <= '9'
}

func isBlank(r rune) bool {
	return r == EMPTY || r == ' '
}
