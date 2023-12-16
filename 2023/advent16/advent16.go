package main

import (
	_ "embed"
	"fmt"
	mapset "github.com/deckarep/golang-set/v2"
	"github.com/sekullbe/advent/parsers"
	"log"
)

//go:embed input.txt
var inputText string

func main() {
	fmt.Printf("Magic number: %d\n", run1(inputText))
	fmt.Println("-------------")
	fmt.Printf("Magic number: %d\n", run2(inputText))
}

func run1(input string) int {
	b := parseBoard(parsers.SplitByLines(input))
	//b.printBoard()
	fmt.Println()
	energized := b.laser(Pt(0, 0), EAST)
	//b.fprintBoardEnergized(os.Stdout, energized)
	//fmt.Println()
	return energized.Cardinality()
}

func run2(input string) int {

	return 0
}

func (b *board) laser(p Point, dir int) mapset.Set[Point] {

	if !b.inRange(p) {
		return mapset.NewSet[Point]()
	}

	t := b.AtPoint(p)
	m := t.contents
	e := mapset.NewSet(p)

	if t.dirs.Contains(dir) {
		return e // we already know what happens from here
	}

	t.dirs.Add(dir)
	b.grid[p] = t // I don't think the helpers work here
	switch m {
	case SPACE:
		e2 := e.Union(b.laser(neighborInDirection(p, dir), dir))
		return e2
	case MIRROR_LEFT:
		newDir := mirrorLeft(dir)
		return e.Union(b.laser(neighborInDirection(p, newDir), newDir))
	case MIRROR_RIGHT:
		newDir := mirrorRight(dir)
		return e.Union(b.laser(neighborInDirection(p, newDir), newDir))
	case SPLIT_VERT:
		nd1, nd2 := split(SPLIT_VERT, dir)
		if nd2 == 0 {
			return e.Union(b.laser(neighborInDirection(p, nd1), nd1))
		}
		return e.Union(b.laser(neighborInDirection(p, nd1), nd1).Union(b.laser(neighborInDirection(p, nd2), nd2)))
	case SPLIT_HORIZ:
		nd1, nd2 := split(SPLIT_HORIZ, dir)
		if nd2 == 0 {
			return e.Union(b.laser(neighborInDirection(p, nd1), nd1))
		}
		return e.Union(b.laser(neighborInDirection(p, nd1), nd1).Union(b.laser(neighborInDirection(p, nd2), nd2)))
	default:
		panic("don't know what to do with this tile")
	}
}

//return e.Union(b.laser(neighborInDirection(p, dir), dir).Union(b.laser(p,dir)))

// could probably do this with mod math but I don't care
// input is the way the beam is going on the way in
func mirrorRight(dir int) int {
	switch dir {
	case NORTH:
		return EAST
	case EAST:
		return NORTH
	case WEST:
		return SOUTH
	case SOUTH:
		return WEST
	default:
		log.Panicf("bad direction: %d", dir)
	}
	return 0
}

func mirrorLeft(dir int) int {
	switch dir {
	case NORTH:
		return WEST
	case EAST:
		return SOUTH
	case WEST:
		return NORTH
	case SOUTH:
		return EAST
	default:
		log.Panicf("bad direction: %d", dir)
	}
	return 0
}

// returns the laser(s) out of the splitter; 2nd is 0 if no split
func split(splitter int, dir int) (int, int) {

	if splitter == SPLIT_VERT {
		switch dir {
		case WEST:
			return NORTH, SOUTH
		case EAST:
			return NORTH, SOUTH
		default:
			return dir, 0
		}
	} else if splitter == SPLIT_HORIZ {
		switch dir {
		case NORTH:
			return EAST, WEST
		case SOUTH:
			return EAST, WEST
		default:
			return dir, 0
		}
	}
	log.Fatalf("bad splitter %d", splitter)
	return 0, 0
}
