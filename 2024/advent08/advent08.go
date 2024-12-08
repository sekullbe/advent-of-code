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

// grid.Tile.Contents is the antenna freq

func run1(input string) int {

	// this is going to be a bit loopy...
	// for each antenna
	// for each other antenna of the same frequency
	// compute the x/y offset
	// apply it to one and then the other reversed

	b := grid.ParseBoard(parsers.SplitByLines(input))
	antennas := make(map[geometry.Point]rune)
	antinodes := make(map[geometry.Point]bool)
	for _, tile := range b.Grid {
		if !grid.IsBlank(tile.Contents) {
			antennas[tile.Point] = tile.Contents
		}
	}
	for p1, f1 := range antennas {
		for p2, f2 := range antennas {
			if p1 == p2 || f1 != f2 {
				continue
			}
			a1, a2 := calculateAntinodes(p1, p2)
			if b.InRange(a1) {
				antinodes[a1] = true
			}
			if b.InRange(a2) {
				antinodes[a2] = true
			}
		}
	}

	return len(antinodes)
}

func run2(input string) int {
	b := grid.ParseBoard(parsers.SplitByLines(input))
	antennas := make(map[geometry.Point]rune)
	antinodes := make(map[geometry.Point]bool)
	for _, tile := range b.Grid {
		if !grid.IsBlank(tile.Contents) {
			antennas[tile.Point] = tile.Contents
		}
	}
	for p1, f1 := range antennas {
		for p2, f2 := range antennas {
			if p1 == p2 || f1 != f2 {
				continue
			}
			antinodeCandidates := calculateAntinodesAggressively(p1, p2, b)
			for _, candidate := range antinodeCandidates {
				if b.InRange(candidate) {
					antinodes[candidate] = true
				}
			}
		}
	}

	return len(antinodes)
}

func calculateAntinodes(a, b geometry.Point) (geometry.Point, geometry.Point) {
	offX, offY := geometry.CalculateOffsets(a, b)
	a1 := a.MovePoint2(-offX, -offY)
	a2 := b.MovePoint2(offX, offY)
	return a1, a2
}

// part 2 version
func calculateAntinodesAggressively(a, b geometry.Point, board *grid.Board) []geometry.Point {
	offX, offY := geometry.CalculateOffsets(a, b)
	// antinodes are all points on the board in line with a +/- any number of offsets
	antinodes := []geometry.Point{}
	// starting at b means we'll include a as an antinode, and the same the other direction
	for anti := b; board.InRange(anti); anti = anti.MovePoint2(-offX, -offY) {
		antinodes = append(antinodes, anti)
	}
	for anti := a; board.InRange(anti); anti = anti.MovePoint2(-offX, -offY) {
		antinodes = append(antinodes, anti)
	}

	return antinodes
}
