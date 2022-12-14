package main

import (
	"fmt"
	"github.com/samber/lo"
	"github.com/sekullbe/advent/parsers"
	"sort"
	"strings"
)

type point struct {
	x int
	y int
}
type pointPair struct {
	from, to point
}

// It doesn't really matter what a point is full of, but store the rune for visualization later
type world struct {
	grid                            map[point]rune
	leftRock, rightRock, lowestRock int
	sandCount                       int
}

// a little gratuitious but lets me change the implementation if I need to later
func newWorld() world {
	return world{
		grid:       make(map[point]rune),
		leftRock:   500,
		rightRock:  500,
		lowestRock: 0,
		sandCount:  0,
	}
}

func parseRockVeins(lines []string) (w world) {
	w = newWorld()
	for _, line := range lines {
		// split on -> into points
		pointStrings := lo.Filter(strings.Fields(line), func(item string, idx int) bool { return item != "->" })
		var points []point
		for _, pointString := range pointStrings {
			coords := parsers.StringsWithCommasToIntSlice(pointString)
			p := point{coords[0], coords[1]}
			points = append(points, p)
			if p.y > w.lowestRock {
				w.lowestRock = p.y
			}
			if p.x < w.leftRock {
				w.leftRock = p.x
			}
			if p.x > w.rightRock {
				w.rightRock = p.x
			}

		}
		// double interior points to get point pairs
		pointPairs := []pointPair{}
		for i := 0; i < len(points)-1; i++ {
			pointPairs = append(pointPairs, pointPair{points[i], points[i+1]})
		}
		// get all the intermediate points for a point pair and set w[p] true
		for _, pair := range pointPairs {
			// kind of dumb, but meh
			// remember to deal if the line is in both directions- just sort the points
			// in every pointpair, either the X or Y coords will be the same
			if pair.from.x == pair.to.x { // vertical line, x constant
				ycs := []int{pair.from.y, pair.to.y}
				sort.Ints(ycs)
				for y := ycs[0]; y <= ycs[1]; y++ {
					w.grid[point{pair.from.x, y}] = '#'
				}
			} else { // horizontal line, y constant
				xcs := []int{pair.from.x, pair.to.x}
				sort.Ints(xcs)
				for x := xcs[0]; x <= xcs[1]; x++ {
					w.grid[point{x, pair.from.y}] = '#'
				}
			}
		}
	}
	return
}

func (w world) printWorld() {
	for y := 0; y <= w.lowestRock; y++ {
		for x := w.leftRock; x <= w.rightRock; x++ {
			p := point{x, y}
			tile, ok := w.grid[p]
			if !ok {
				fmt.Print(".")
			} else {
				fmt.Printf("%c", tile)
			}
		}
		fmt.Println()
	}
}

// simple but keeps from littering the code with math
func (p point) downPoint() point {
	return point{p.x, p.y + 1}
}
func (p point) downLeftPoint() point {
	return point{p.x - 1, p.y + 1}
}
func (p point) downRightPoint() point {
	return point{p.x + 1, p.y + 1}
}

// generalize this, because we might call it recursively
// Returns true if sand was placed successfully
func (w *world) dropSand(dropPoint point) bool {

	if dropPoint.y >= w.lowestRock {
		// Fallen off the bottom. Don't place anything, just return false
		return false
	}

	_, filled := w.grid[dropPoint]
	if filled { // Drop point is blocked!
		return false
	}

	// look down, if we can go down, call that
	_, filled = w.grid[dropPoint.downPoint()]
	if !filled {
		return w.dropSand(dropPoint.downPoint())
	} // look left, if that's clear, call that
	_, filled = w.grid[dropPoint.downLeftPoint()]
	if !filled {
		return w.dropSand(dropPoint.downLeftPoint())
	}
	// look right, if that's clear, call that
	_, filled = w.grid[dropPoint.downRightPoint()]
	if !filled {
		return w.dropSand(dropPoint.downRightPoint())
	}
	//	if none of those are possible, come to rest, set w[p]='o' and return, which should propagate back up the call stack
	w.grid[dropPoint] = 'o'
	w.sandCount++
	return true
}
