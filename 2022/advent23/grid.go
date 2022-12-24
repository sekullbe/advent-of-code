package main

import "image"

// intentions
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

var scanlooks = map[int][3]int{
	NORTH: [3]int{NORTHWEST, NORTH, NORTHEAST},
	EAST:  [3]int{NORTHEAST, EAST, SOUTHEAST},
	SOUTH: [3]int{SOUTHEAST, SOUTH, SOUTHWEST},
	WEST:  [3]int{SOUTHWEST, WEST, NORTHWEST},
}

type elf struct {
	id int
	image.Point
	proposalPoint image.Point
	board         *board
}

func np(x, y int) image.Point {
	return image.Point{X: x, Y: y}
}

func neighborInDirection(p image.Point, dir int) (sees image.Point) {
	switch dir {
	case NORTH:
		return np(p.X, p.Y-1)
	case NORTHEAST:
		return np(p.X+1, p.Y-1)
	case EAST:
		return np(p.X+1, p.Y)
	case SOUTHEAST:
		return np(p.X+1, p.Y+1)
	case SOUTH:
		return np(p.X, p.Y+1)
	case SOUTHWEST:
		return np(p.X-1, p.Y+1)
	case WEST:
		return np(p.X-1, p.Y)
	case NORTHWEST:
		return np(p.X-1, p.Y-1)
	}
	return
}

// TODO at some point make a generalized grid implementation
type grid map[image.Point]*elf

type board struct {
	grove       grid
	elves       []*elf
	proposalDir int
}
