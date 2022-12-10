package main

import (
	"github.com/sekullbe/advent/tools"
	"log"
)

type point struct {
	x int
	y int
}

type board struct {
	h, t        point
	tailVisited map[point]any
}

var exists = struct{}{}

func newBoard() board {
	tv := make(map[point]any)
	tv[point{0, 0}] = exists
	return board{h: point{0, 0}, t: point{0, 0}, tailVisited: tv}
}

// Moves head per the direction, and tail to catch up
func (b *board) move(dir rune) {
	b.h = b.h.getNewMovePoint(dir)
	// if touching or adjacent do nothing
	if adjacent(b.h, b.t) {
		return
	}
	if isStraightLine(b.h, b.t) {
		b.t = moveOneCloserRookwise(b.h, b.t)
	} else {
		b.t = moveOneCloserBishopwise(b.h, b.t)
	}
	b.tailVisited[b.t] = exists
}

// Unlike previous implementations, positive Y is UP not DOWN
func (p point) getNewMovePoint(dir rune) point {
	switch dir {
	case 'U':
		return point{x: p.x, y: p.y + 1}
	case 'D':
		return point{x: p.x, y: p.y - 1}
	case 'L':
		return point{x: p.x - 1, y: p.y}
	case 'R':
		return point{x: p.x + 1, y: p.y}
	}
	log.Fatal("Move failed, bad direction")
	return p
}

func adjacent(p1, p2 point) bool {
	//return (p1.x == p2.x && tools.AbsInt(p1.y-p2.y) <= 1) || (p1.y == p2.y && tools.AbsInt(p1.x-p2.x) <= 1)
	return tools.AbsInt(p1.x-p2.x) <= 1 && tools.AbsInt(p1.y-p2.y) <= 1
}

func rectangularDistance(p1, p2 point) int {
	return tools.AbsInt(p1.x-p2.x) + tools.AbsInt(p1.y-p2.y)
}
func isStraightLine(p1, p2 point) bool {
	return p1.x == p2.x || p1.y == p2.y
}

// returns a point that moves p2 one closer to p1
func moveOneCloserRookwise(p1, p2 point) point {
	if p1.x != p2.x && p1.y != p2.y { // precondition check. in real code i'd return an error but this is for debugging
		log.Panic("Shouldn't be moving rookwise if the points are not in line")
	}
	np := point{p2.x, p2.y}
	if p1.x == p2.x {
		if p1.y > p2.y {
			np.y = p2.y + 1
		} else {
			np.y = p2.y - 1
		}
	} else {
		if p1.x > p2.x {
			np.x = p2.x + 1
		} else {
			np.x = p2.x - 1
		}
	}
	return np
}

func moveOneCloserBishopwise(p1, p2 point) point {
	if p1.x == p2.x || p1.y == p2.y { // precondition check
		log.Panic("Shouldn't be moving bishopwise if the points are in line")
	}
	np := point{p2.x, p2.y}
	// move such that both coordinates get one close

	if p1.x > p2.x {
		np.x = p2.x + 1
	} else {
		np.x = p2.x - 1
	}
	if p1.y > p2.y {
		np.y = p2.y + 1
	} else {
		np.y = p2.y - 1
	}
	return np
}
