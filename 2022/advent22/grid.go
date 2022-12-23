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
}

// maybe this should be a method on tile, which would then need a *board
func (b *board) look(t *tile, d int) *tile {
	var t2 *tile
	var ok bool
	// First check the cache
	cacheTile := t.neighbors[d]
	if cacheTile != nil {
		return cacheTile
	}
	switch d {
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
		log.Panicf("wtf direction %d", d)
	}
	return t2
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

// returns the number of steps actually moved
func (b *board) move(distance int) int {
	steps := 0
	for n := 0; n < distance; n++ {
		nt := b.look(b.moverLoc, b.moverDir)
		if nt.contents != '#' { // safe to move
			b.moverLoc = nt
			steps++
		} else {
			break // wall, stop
		}
	}
	return steps
}
