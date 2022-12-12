package main

import (
	"github.com/beefsack/go-astar"
	"math"
)

type point struct {
	x int
	y int
}

// astar examples use x,y instead of point, but I like using maps as grids
type world map[point]*Tile

// Tile gets the tile at the given coordinates in the world.
func (w world) Tile(x, y int) *Tile {
	t := w[point{x, y}]
	return t
}

// astar uses this
type Tile struct {
	p      point
	height rune // it's interchangeable with int, why not
	w      world
}

func (t *Tile) PathNeighbors() []astar.Pather {
	neighbors := []astar.Pather{}
	for _, offset := range [][]int{
		{-1, 0},
		{1, 0},
		{0, -1},
		{0, 1},
	} {
		// Blockers are determined on the fly based on height difference, not intrinsic property of the tile
		if n := t.w.Tile(t.p.x+offset[0], t.p.y+offset[1]); n != nil && n.height <= t.height+1 {
			neighbors = append(neighbors, n)
		}
	}
	return neighbors
}

func (t *Tile) PathNeighborCost(to astar.Pather) float64 {
	//all moves are either blocked or allowed, so movement cost is constant
	return 1
}

func (t *Tile) PathEstimatedCost(to astar.Pather) float64 {
	toT := to.(*Tile)
	return math.Abs(float64(t.p.x-toT.p.x)) + math.Abs(float64(t.p.y-toT.p.y))
}

func parseWorld(lines []string) (world, *Tile, *Tile) {
	w := make(world)
	var from, to *Tile
	for x, line := range lines {
		for y, r := range line {
			p := point{x, y}
			t := Tile{
				p:      p,
				height: r - 'a',
				w:      w,
			}
			switch r {
			case 'S':
				t.height = 0
				from = &t
			case 'E':
				t.height = 'z' - 'a'
				to = &t
			}
			w[p] = &t
		}
	}
	return w, from, to
}
