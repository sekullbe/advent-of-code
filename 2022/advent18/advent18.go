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

var exists = struct{}{}

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

// for part1 it was enough to ask about neighbors but for part2 I need to list them
func neighbors(p point3) []point3 {
	return []point3{
		{x: p.x - 1, y: p.y, z: p.z},
		{x: p.x + 1, y: p.y, z: p.z},
		{x: p.x, y: p.y - 1, z: p.z},
		{x: p.x, y: p.y + 1, z: p.z},
		{x: p.x, y: p.y, z: p.z - 1},
		{x: p.x, y: p.y, z: p.z + 1},
	}
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

	// load the points like part1 except without the processing
	droplet := make(map[point3]voxel)
	var maxX, maxY, maxZ int // need to know size of the world
	for _, dline := range parsers.SplitByLines(inputText) {
		coords := parsers.StringsWithCommasToIntSlice(dline)
		newVoxel := voxel{
			loc:       point3{coords[0], coords[1], coords[2]},
			openSides: 6,
		}
		droplet[newVoxel.loc] = newVoxel
		maxX = tools.MaxInt(maxX, newVoxel.loc.x)
		maxY = tools.MaxInt(maxY, newVoxel.loc.y)
		maxZ = tools.MaxInt(maxZ, newVoxel.loc.z)
	}

	totalArea := markExternallyAccessible(point3{0, 0, 0}, droplet, make(map[point3]any), maxX, maxY, maxZ)

	return totalArea
}

// idea from moshan1997 on reddit- when you see a lava voxel you know you came to it from
// a neighbor that was external air, because the func only continues if it was an air voxel
func markExternallyAccessible(p point3, droplet map[point3]voxel, knownExterior map[point3]any, maxX, maxY, maxZ int) int {
	// we already know this is exterior and not lava so it doesn't contribute to area
	if tools.KeyExists(knownExterior, p) {
		return 0
	}
	// bounds check
	if p.x < -1 || p.x > maxX+1 || p.y < -1 || p.y > maxY+1 || p.z < -1 || p.z > maxZ+1 {
		return 0
	}
	if tools.KeyExists(droplet, p) {
		return 1 // we found one exterior face of a lava voxel
	}
	knownExterior[p] = exists // we only care about the existence of the key, not the value
	var externalArea int
	for _, neighbor := range neighbors(p) { // for each neighbor,
		externalArea += markExternallyAccessible(neighbor, droplet, knownExterior, maxX, maxY, maxZ)
	}
	return externalArea
}
