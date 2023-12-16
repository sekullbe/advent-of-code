package main

// getting towards a generic grid implementation, not there yet
// keep the diagonals in this one, just don't use them
import (
	"fmt"
	mapset "github.com/deckarep/golang-set/v2"
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

// mirrors
const (
	SPACE        = iota // .
	MIRROR_RIGHT        // /
	MIRROR_LEFT         // \
	SPLIT_VERT          // |
	SPLIT_HORIZ         // -
)

// At some point I should make a generalized grid implementation with generic contents and a parser
// have to think about what the interface would be
type grid map[Point]*tile

type board struct {
	grid
	maxX, maxY int // min is always 0

}

// generalized grid impl will provide this
type baseSquare struct {
	id    int
	Point Point
}

// and a user will add something like this
type tile struct {
	baseSquare
	contents int
	dirs     mapset.Set[int]
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

func newTile(p Point, content int) tile {
	return tile{
		baseSquare: baseSquare{Point: p},
		contents:   content,
		dirs:       mapset.NewSet[int](),
	}
}

func parseContents(r rune) int {
	switch r {
	case '.':
		return SPACE
	case '/':
		return MIRROR_RIGHT
	case '\\':
		return MIRROR_LEFT
	case '|':
		return SPLIT_VERT
	case '-':
		return SPLIT_HORIZ
	default:
		panic("don't know what kind of a space this is:" + string(r))
	}
}

func parseBoard(lines []string) *board {

	bb := board{
		grid: make(grid),
		maxX: 0,
		maxY: 0,
	}
	b := &bb
	b.maxX = len(lines[0]) - 1
	b.maxY = len(lines) - 1

	for y, line := range lines {
		if len(line) == 0 {
			break // catch extraneous blank lines
		}
		for x, tc := range line {
			pt := Pt(x, y)
			t := newTile(pt, parseContents(tc))
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
			switch b.At(x, y).contents {
			case SPACE:
				fmt.Fprint(w, "·")
			case MIRROR_LEFT:
				fmt.Fprint(w, `\`)
			case MIRROR_RIGHT:
				fmt.Fprint(w, `/`)
			case SPLIT_VERT:
				fmt.Fprint(w, `|`)
			case SPLIT_HORIZ:
				fmt.Fprint(w, `-`)
			default:
				panic("don't know how to print this")
			}
		}
		fmt.Fprintln(w)
	}
}

func (b *board) fprintBoardEnergized(w io.Writer, energized mapset.Set[Point]) {
	for y := 0; y <= b.maxY; y++ {
		for x := 0; x <= b.maxX; x++ {
			if energized.Contains(Pt(x, y)) {
				fmt.Fprint(w, `#`)
			} else {
				switch b.At(x, y).contents {
				case SPACE:
					fmt.Fprint(w, "·")
				case MIRROR_LEFT:
					fmt.Fprint(w, `\`)
				case MIRROR_RIGHT:
					fmt.Fprint(w, `/`)
				case SPLIT_VERT:
					fmt.Fprint(w, `|`)
				case SPLIT_HORIZ:
					fmt.Fprint(w, `-`)
				default:
					panic("don't know how to print this")
				}
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

/*
could use something like this that calls a fn for each point?
	for y := 0; y <= b.maxY; y++ {
		for x := 0; x < b.maxX; x++ {
          fn(b.At(x,y))
		}
	}
*/
