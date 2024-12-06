package main

import (
	_ "embed"
	"fmt"
	"github.com/sekullbe/advent/geometry"
	"github.com/sekullbe/advent/grid"
	"github.com/sekullbe/advent/parsers"
)

//go:embed input.txt
var inputText string

func main() {
	fmt.Printf("Magic number: %d\n", run1(inputText))
	fmt.Println("-------------")
	fmt.Printf("Magic number: %d\n", run2(inputText))
}

func run1(input string) int {

	b := grid.ParseBoard(parsers.SplitByLines(input))

	// find the guard
	var guard geometry.Point
	var guardFacing = grid.NORTH
	var count = 1
	for point, tile := range b.Grid {
		if tile.Contents == '^' {
			guard = point
			b.AtPoint(guard).Contents = ' '
			b.AtPoint(guard).Traversed = true
			break
		}
	}
	// start moving until guard goes offboard
	for b.InRange(guard) {
		for {
			next := grid.NeighborInDirection(guard, guardFacing)
			if grid.IsBlank(b.AtPoint(next).Contents) {
				if !b.AtPoint(guard).Traversed {
					count++
				}
				b.AtPoint(guard).Traversed = true
				guard = next
				break
			}
			guardFacing = grid.Clockwise(guardFacing, 2)
		}
	}

	return count
}

func run2(input string) int {

	b := grid.ParseBoard(parsers.SplitByLines(input))

	// find the guard
	var guard geometry.Point
	var guardFacing = grid.NORTH
	for point, tile := range b.Grid {
		if tile.Contents == '^' {
			guard = point
			b.AtPoint(guard).Contents = ' '
			b.AtPoint(guard).Traversed = true
			b.AtPoint(guard).Counter = 1
			break
		}
	}
	initialGuard := guard
	initialGuardFacing := guardFacing

	type visit struct {
		loc    geometry.Point
		facing int
	}
	var visits map[visit]bool = make(map[visit]bool)

	var loops = 0

	for _, newObstacleTile := range b.Grid { // go go gadget brute force
		guard = initialGuard
		guardFacing = initialGuardFacing
		if newObstacleTile.Point == initialGuard || !grid.IsBlank(newObstacleTile.Contents) {
			continue // invalid place to put obstacle
		}
		b.Grid[newObstacleTile.Point].Contents = '#' // can't just assign to newObstacleTile
		// for debugging
		//b.Grid[grid.Pt(3, 6)].Contents = '#'
		clear(visits)
		for b.InRange(guard) {
			next := grid.NeighborInDirection(guard, guardFacing)
			// if we've moved to a place and direction we've already been, then we are in a loop
			if _, visited := visits[visit{guard, guardFacing}]; visited {
				loops++
				break
			}
			visits[visit{guard, guardFacing}] = true
			// if the guard can move, move them; if they can't, turn them in place
			if grid.IsBlank(b.AtPoint(next).Contents) {
				guard = next
			} else {
				guardFacing = grid.Clockwise(guardFacing, 2)
			}
		}
		b.Grid[newObstacleTile.Point].Contents = grid.EMPTY
	}
	return loops
}
