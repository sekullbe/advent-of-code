package main

import (
	"image"
	"log"
)

const (
	EAST = iota
	SOUTH
	WEST
	NORTH
)

func np(x, y int) image.Point {
	return image.Point{X: x, Y: y}
}

type tile struct {
	image.Point
	contents  rune
	neighbors [4]*tile
}

// TODO at some point make a generalized grid implementation
type grid map[image.Point]*tile

type board struct {
	grove      grid
	maxX, maxY int // min is always 0
	moverLoc   *tile
	moverDir   int
	sideSize   int
}

/*
Sample
1: W-3S N-2S E-6W S-4S
2: W-6N E-3E S-5N N-1N
3: W-2W N-1E E-4E S-5E
4: W-3W S-5S N-1N E-6S
5: W-3N N-4N E-6E S-2N
6: W-5W N-4W E-1W S-2E

Real
1: E-2E S-3S W-4E N-6E
2: E-5W S-3W W-1W N-6N
3: E-2N S-5S W-4S N-1N
4: E-5E S-6S W-1E N-3E
5: E-2W S-6W W-4W N-3N
6: E-5N S-2S W-1S N-4N

*/

// maybe this should be a method on tile, which would then need a *board
func (b *board) look(t *tile) *tile {
	var t2 *tile
	var ok bool
	// First check the cache
	cacheTile := t.neighbors[b.moverDir]
	if cacheTile != nil {
		return cacheTile
	}
	switch b.moverDir {
	case EAST:
		p2 := np(t.X+1, t.Y)
		if t2, ok = b.grove[p2]; !ok {
			for x := p2.X; ; x++ {
				if x > b.maxX {
					x = 1
				}
				if t2, ok = b.grove[np(x, p2.Y)]; ok {
					break
				}
			}
		}
		t.neighbors[EAST] = t2
		t2.neighbors[WEST] = t
	case SOUTH:
		p2 := np(t.X, t.Y+1)
		if t2, ok = b.grove[p2]; !ok {
			for y := p2.Y; ; y++ {
				if y > b.maxY {
					y = 1
				}
				if t2, ok = b.grove[np(p2.X, y)]; ok {
					break
				}
			}
		}
		t.neighbors[SOUTH] = t2
		t2.neighbors[NORTH] = t
	case WEST:
		p2 := np(t.X-1, t.Y)
		if t2, ok = b.grove[p2]; !ok {
			for x := p2.X; ; x-- {
				if x < 1 {
					x = b.maxX
				}
				if t2, ok = b.grove[np(x, p2.Y)]; ok {
					break
				}
			}
		}
		t.neighbors[WEST] = t2
		t2.neighbors[EAST] = t
	case NORTH:
		p2 := np(t.X, t.Y-1)
		if t2, ok = b.grove[p2]; !ok {
			for y := p2.Y; ; y-- {
				if y < 0 {
					y = b.maxY
				}
				if t2, ok = b.grove[np(p2.X, y)]; ok {
					break
				}
			}
		}
		t.neighbors[NORTH] = t2
		t2.neighbors[SOUTH] = t
	default:
		log.Panicf("wtf direction %d", b.moverDir)
	}
	return t2
}

// compute which cube side you're on given coordinates
func coordsToSide(p image.Point) int {
	test := false // gross
	// these MUST be tested in order - if tile coords are both <= the x,y values, it's in that side
	sidesTest := [7]image.Point{{0, 0}, {16, 4}, {4, 8}, {8, 8}, {12, 8}, {12, 12}, {12, 16}}
	sidesReal := [7]image.Point{{0, 0}, {100, 50}, {150, 50}, {100, 100}, {50, 150}, {100, 150}, {50, 200}}
	var sides [7]image.Point
	if test {
		sides = sidesTest
	} else {
		sides = sidesReal
	}
	for s := 1; s <= 6; s++ {
		if p.X <= sides[s].X && p.Y <= sides[s].Y {
			return s
		}
	}
	log.Panicf("could not find a side for point %v", p)
	return -1
}

// lookOnCube returns the tile in the current looking direction and the modified direction for the new cube face
func (b *board) lookOnCube(t1 *tile) (*tile, int) {
	var t2 *tile
	nd := b.moverDir

	onEastEdge := b.moverLoc.Point.X%b.sideSize == 0
	onSouthEdge := b.moverLoc.Point.Y%b.sideSize == 0
	onWestEdge := (b.moverLoc.Point.X-1)%b.sideSize == 0
	onNorthEdge := (b.moverLoc.Point.Y-1)%b.sideSize == 0

	var p2 image.Point
	side := coordsToSide(b.moverLoc.Point)
	switch side {
	case 1:
		// safe moves are E,S. W -> 4E, N -> 6E
		switch b.moverDir {
		case EAST:
			p2 = np(t1.X+1, t1.Y)
		case SOUTH:
			p2 = np(t1.X, t1.Y+1)
		case WEST:
			if onWestEdge { // coming into the west side of 4, facing east
				nd = EAST
				p2 = np(1, 151-t1.Y) // y=1->y=151 y=50->y=101
			} else {
				p2 = np(t1.X-1, t1.Y)
			}
		case NORTH:
			if onNorthEdge { // coming into the west side of 6, facing east
				nd = EAST
				p2 = np(1, 100+t1.X) //x=51 -> y=151, x=100 y=200
			} else {
				p2 = np(t1.X, t1.Y-1)
			}
		}
	case 2:
		// only safe move is west
		switch b.moverDir {
		case EAST:
			if onEastEdge {
				nd = WEST
				p2 = np(100, 151-t1.Y)
			} else {
				p2 = np(t1.X+1, t1.Y)
			}
		case SOUTH:
			if onSouthEdge {
				nd = WEST
				p2 = np(100, t1.X-50)
			} else {
				p2 = np(t1.X, t1.Y+1)
			}
		case WEST:
			p2 = np(t1.X-1, t1.Y)
		case NORTH:
			if onNorthEdge {
				nd = NORTH
				p2 = np(t1.X-100, 200)
			} else {
				p2 = np(t1.X, t1.Y-1)
			}
		}
	case 3:
		// safe moves are north/south
		switch b.moverDir {
		case EAST:
			if onEastEdge { // goes to bottom of 2
				nd = NORTH
				p2 = np(t1.Y+50, 50)
			} else {
				p2 = np(t1.X+1, t1.Y)
			}
		case SOUTH:
			p2 = np(t1.X, t1.Y+1)
		case WEST:
			if onWestEdge {
				nd = SOUTH
				p2 = np(t1.Y-50, 101)
			} else {
				p2 = np(t1.X-1, t1.Y)
			}
		case NORTH:
			p2 = np(t1.X, t1.Y-1)
		}
	case 4:
		// safe moves are east,south
		switch b.moverDir {
		case EAST:
			p2 = np(t1.X+1, t1.Y)
		case SOUTH:
			p2 = np(t1.X, t1.Y+1)
		case WEST: //to west edge of 1
			if onWestEdge {
				nd = EAST
				p2 = np(51, 151-t1.Y) // y=101->y=50  y=150->y=1
			} else {
				p2 = np(t1.X-1, t1.Y)
			}
		case NORTH:
			if onNorthEdge {
				nd = EAST
				p2 = np(51, t1.X+50) //x=1->y=51 x=50->y=100
			} else {
				p2 = np(t1.X, t1.Y-1)
			}
		}
	case 5:
		// save moves are north and west
		switch b.moverDir {
		case EAST:
			if onEastEdge {
				nd = WEST
				p2 = np(150, 151-t1.Y) // y=101->y=50  y=151->y=1
			} else {
				p2 = np(t1.X+1, t1.Y)
			}
		case SOUTH:
			if onSouthEdge {
				nd = WEST
				p2 = np(50, t1.X+100) // x=51,y=151  x=100, y=200
			} else {
				p2 = np(t1.X, t1.Y+1)
			}
		case WEST:
			p2 = np(t1.X-1, t1.Y)
		case NORTH:
			p2 = np(t1.X, t1.Y-1)
		}
	case 6:
		// only N is safe
		switch b.moverDir {
		case EAST:
			if onEastEdge {
				nd = NORTH
				p2 = np(t1.Y-100, 150) //y=150->x=51 y=200->x=100
			} else {
				p2 = np(t1.X+1, t1.Y)
			}
		case SOUTH:
			if onSouthEdge {
				nd = SOUTH
				p2 = np(t1.X+100, 1)
			} else {
				p2 = np(t1.X, t1.Y+1)
			}
		case WEST:
			if onWestEdge { // coming into the north side of 1, facing south
				nd = SOUTH
				p2 = np(t1.Y-100, 1) // y=151,x=50  y=200,x=100
			} else {
				p2 = np(t1.X-1, t1.Y)
			}
		case NORTH:
			p2 = np(t1.X, t1.Y-1)
		}
	}
	t2 = b.grove[p2]
	return t2, nd
}

func (b *board) turn(dir rune) int {
	if dir == 'R' {
		b.moverDir++
		if b.moverDir == 4 {
			b.moverDir = 0 // I could do some mod math but I don't care
		}
	}
	if dir == 'L' {
		b.moverDir--
		if b.moverDir == -1 {
			b.moverDir = 3
		}
	}
	return b.moverDir
}

// make lookOnCube() return the new direction
// returns the number of steps actually moved
func (b *board) move(distance int, moveOnCube bool) int {
	steps := 0
	for n := 0; n < distance; n++ {
		var nt *tile
		nd := b.moverDir
		if moveOnCube {
			nt, nd = b.lookOnCube(b.moverLoc)
		} else {
			nt = b.look(b.moverLoc)
		}
		if nt.contents != '#' { // safe to move
			b.moverLoc = nt
			b.moverDir = nd
			steps++
		} else {
			break // wall, stop
		}
	}
	return steps
}
