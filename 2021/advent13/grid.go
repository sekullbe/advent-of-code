package main

import (
	"errors"
	"fmt"
	"github.com/sekullbe/advent/parsers"
)

// looks like the largest number is 1310
//type grid map[int]map[int]bool
type point struct {
	x int
	y int
}

type grid map[point]bool

func newGrid() grid {
	return make(grid)
}

func (g grid) addCoords(x, y int) point {
	p := point{x, y}
	g.addPoint(p)
	return p
}

func (g grid) addCoordString(s string) point {
	coords := parsers.StringsWithCommasToIntSlice(s)
	return g.addCoords(coords[0], coords[1])
}

func (g grid) addPoint(p point) {
	g[p] = true
}

func (g grid) removePoint(p point) {
	delete(g, p)
}

func (g grid) getPoint(p point) bool {
	if _, present := g[p]; !present {
		return false
	}
	return g[p]
}

func (g grid) movePoint(from, to point) error {
	if _, present := g[from]; !present {
		return errors.New("can't move a point that isn't there")
	}
	delete(g, from)
	g[to] = true
	return nil
}

func (g grid) fold(axis string, foldCoord int) {
	foldx := axis == "x" // else foldLeft
	//log.Printf("Folding on x=%d; maxX = %d, maxY = %d", foldCoord, g.maxX(), g.maxY())
	for p1 := range g {
		if !g[p1] {
			continue // the point was unset
		}
		p2 := point{p1.x, p1.y}
		if foldx {
			if p1.x > foldCoord {
				p2.x = 2*foldCoord - p1.x
			}
		} else {
			if p1.y > foldCoord {
				p2.y = 2*foldCoord - p1.y
			}
		}
		err := g.movePoint(p1, p2)
		if err != nil {
			panic("folding tried to move a point that wasn't there")
		}
	}
}

// this was the most annoying part of the whole thing, to convert the map to a grid without writing over the edge and panicing
func (g grid) display() string {
	// paper is array of rows, so it's [row][col]
	paper := make([][]string, g.maxY()+1)
	maxcol := g.maxX() + 1
	for row := 0; row < g.maxY()+1; row++ {
		paper[row] = make([]string, maxcol)
		for col := 0; col < maxcol; col++ {
			paper[row][col] = ".."
		}
	}
	for p, b := range g {
		if b {
			paper[p.y][p.x] = "##"
		}

	}

	var out string
	for _, row := range paper {
		for _, col := range row {
			out += fmt.Sprintf("%s", col)
		}
		out += fmt.Sprintln()
	}
	return out
}

func (g grid) countPoints() int {
	count := 0
	for k := range g {
		if g[k] {
			count++ // false points should not be there, but just in case
		}
	}
	return count
}

func (g grid) maxX() int {
	max := 0
	for k := range g {
		if k.x > max {
			max = k.x
		}
	}
	return max
}

func (g grid) maxY() int {
	max := 0
	for k := range g {
		if k.y > max {
			max = k.y
		}
	}
	return max
}
