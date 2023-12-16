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

	//	energized := b.laser(Pt(0, 0), EAST)
	//	return energized.Cardinality()

	return b.laserAndCount(Pt(0, 0), EAST)
}

func (b *board) laserAndCount(p Point, dir int) int {
	energized := b.laser(p, dir)
	return energized.Cardinality()
}

func run2(input string) int {

	eMax := 0
	b := parseBoard(parsers.SplitByLines(input))
	// I *expect* this will accelerate as it goes, as the cache kicks in
	for x := 0; x <= b.maxX; x++ {
		b := parseBoard(parsers.SplitByLines(input))
		eMax = max(eMax, b.laserAndCount(Pt(x, 0), SOUTH))
		fmt.Print(".")
	}
	fmt.Println()
	for x := 0; x <= b.maxX; x++ {
		b := parseBoard(parsers.SplitByLines(input))
		eMax = max(eMax, b.laserAndCount(Pt(x, b.maxY), NORTH))
		fmt.Print(".")
	}
	fmt.Println()
	for y := 0; y <= b.maxY; y++ {
		b := parseBoard(parsers.SplitByLines(input))
		eMax = max(eMax, b.laserAndCount(Pt(0, y), EAST))
		fmt.Print(".")
	}
	fmt.Println()
	for y := 0; y <= b.maxY; y++ {
		b := parseBoard(parsers.SplitByLines(input))
		eMax = max(eMax, b.laserAndCount(Pt(b.maxX, y), WEST))
		fmt.Print(".")
	}

	return eMax
	// 7655 too low!
}

func (b *board) laser(p Point, dir int) mapset.Set[Point] {

	if !b.inRange(p) {
		return mapset.NewSet[Point]()
	}

	t := b.AtPoint(p)
	m := t.contents
	e := mapset.NewSet(p)

	// this caching isn't *quite* working, I think because it conflicts with the loop detection
	// and doesn't do that right either.
	// Either a separate cache of pt,dir to set or num would probably work, instead of setting the cache on the tile
	// itself, but it's fast enough.
	//if ee, ok := t.knownFromHere[cacheKey{p, dir}]; ok {
	//	b.fprintBoardEnergized(os.Stdout, ee)
	//	return ee.Union(e)
	//}

	if t.dirs.Contains(dir) {
		return e // we already know what happens from here
	}
	t.dirs.Add(dir)
	b.grid[p] = t // I don't think the helpers work here
	var e2 mapset.Set[Point]
	switch m {
	case SPACE:
		e2 = e.Union(b.laser(neighborInDirection(p, dir), dir))
	case MIRROR_LEFT:
		newDir := mirrorLeft(dir)
		e2 = e.Union(b.laser(neighborInDirection(p, newDir), newDir))
	case MIRROR_RIGHT:
		newDir := mirrorRight(dir)
		e2 = e.Union(b.laser(neighborInDirection(p, newDir), newDir))
	case SPLIT_VERT:
		nd1, nd2 := split(SPLIT_VERT, dir)
		if nd2 == 0 {
			e2 = e.Union(b.laser(neighborInDirection(p, nd1), nd1))
		} else {
			e2 = e.Union(b.laser(neighborInDirection(p, nd1), nd1).Union(b.laser(neighborInDirection(p, nd2), nd2)))
		}
	case SPLIT_HORIZ:
		nd1, nd2 := split(SPLIT_HORIZ, dir)
		if nd2 == 0 {
			e2 = e.Union(b.laser(neighborInDirection(p, nd1), nd1))
		} else {
			e2 = e.Union(b.laser(neighborInDirection(p, nd1), nd1).Union(b.laser(neighborInDirection(p, nd2), nd2)))
		}
	default:
		panic("don't know what to do with this tile")
	}
	//t.knownFromHere[cacheKey{p: p, dir: dir}] = e2
	//b.grid[p] = t // I don't think the helpers work here
	//b.fprintBoardEnergized(os.Stdout, e2)

	return e2
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
