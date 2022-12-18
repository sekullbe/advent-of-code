package main

import (
	_ "embed"
	"fmt"
	"github.com/sekullbe/advent/parsers"
	"github.com/sekullbe/advent/tools"
)

//go:embed input.txt
var inputText string

func main() {
	fmt.Printf("Magic number: %d\n", run1(inputText))
	fmt.Println("-------------")
	fmt.Printf("Magic number: %d\n", run2(inputText))
}

type point3 struct {
	x, y, z int
}

type voxel struct {
	loc       point3
	openSides int
}

func adjacent(a, b point3) bool {
	xok := a.y == b.y && a.z == b.z && tools.AbsInt(a.x-b.x) <= 1
	yok := a.x == b.x && a.z == b.z && tools.AbsInt(a.y-b.y) <= 1
	zok := a.x == b.x && a.y == b.y && tools.AbsInt(a.z-b.z) <= 1
	return xok || yok || zok
}

func adjacent2(a, b point3) bool {
	return tools.AbsInt(a.x-b.x)+tools.AbsInt(a.y-b.y)+tools.AbsInt(a.z-b.z) <= 1
}

func run1(inputText string) int {

	droplet := make(map[point3]voxel)
	for _, dline := range parsers.SplitByLines(inputText) {
		coords := parsers.StringsWithCommasToIntSlice(dline)
		newVoxel := voxel{
			loc:       point3{coords[0], coords[1], coords[2]},
			openSides: 6,
		}
		// try to do it while loading
		// look at all the other points in this droplet
		for evLoc, existingVoxel := range droplet {
			// if they're adjacent, drop each one's opensides by one
			// two drops are adjacent if they share 2 coordinates and the third only differs by one
			if adjacent(newVoxel.loc, evLoc) {
				newVoxel.openSides--
				existingVoxel.openSides-- // can't modify in place, either use pointers or resave it
				droplet[existingVoxel.loc] = existingVoxel
			}
		}
		droplet[newVoxel.loc] = newVoxel
	}

	totalArea := 0
	for _, v := range droplet {
		totalArea += v.openSides
	}

	return totalArea
}

func run2(inputText string) int {

	return 0
}
