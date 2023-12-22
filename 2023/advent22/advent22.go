package main

import (
	_ "embed"
	"fmt"
	"github.com/sekullbe/advent/geometry"
	"github.com/sekullbe/advent/parsers"
	"slices"
	"strings"
)

//go:embed input.txt
var inputText string

func main() {
	fmt.Printf("Magic number: %d\n", run1(inputText))
	fmt.Println("-------------")
	fmt.Printf("Magic number: %d\n", run2(inputText))
}

// note: a Brick is always a straight line of cubes- 1x1xN in some axis
// so there's no funny business with larger bricks

func run1(input string) int {
	bricks := []Brick{}
	for _, l := range parsers.SplitByLines(input) {
		b := parseBrick(l)
		bricks = append(bricks, b)
	}

	// Sort the bricks by increasing Z
	slices.SortFunc(bricks, func(a, b Brick) int { return a.End1.Z - b.End1.Z })
	// so we know each Brick in the list can always fall without going through another Brick

	// make them all fall to their lowest position
	fallers := fall(bricks)
	fmt.Printf("Fallers: %d\n", fallers)

	// and see which can be disintegrated- remove it and try to fall() the other bricks
	zappable := 0
	bcopy := make([]Brick, len(bricks))
	for i, _ := range bricks {
		copy(bcopy, bricks)
		//b.zapped = true
		bcopy[i] = Brick{zapped: true}
		changes := fall(bcopy)
		if changes == 0 {
			zappable += 1
		}
	}
	return zappable
}

func run2(input string) int {
	bricks := []Brick{}
	for _, l := range parsers.SplitByLines(input) {
		b := parseBrick(l)
		bricks = append(bricks, b)
	}

	// Sort the bricks by increasing Z
	slices.SortFunc(bricks, func(a, b Brick) int { return a.End1.Z - b.End1.Z })
	// so we know each Brick in the list can always fall without going through another Brick

	// make them all fall to their lowest position
	fallers := fall(bricks)
	fmt.Printf("Fallers: %d\n", fallers)

	// and see which can be disintegrated- remove it and try to fall() the other bricks
	fallers = 0
	bcopy := make([]Brick, len(bricks))
	for i, _ := range bricks {
		copy(bcopy, bricks)
		^bcopy[i] = Brick{zapped: true}
		fallers += fall(bcopy)
	}
	return fallers
}

func parseBrick(line string) Brick {
	//x1,y1,z1~x2,y2,z2
	var x1, x2, y1, y2, z1, z2 int
	fmt.Sscanf(strings.TrimSpace(line), "%d,%d,%d~%d,%d,%d", &x1, &y1, &z1, &x2, &y2, &z2)
	return Brick{End1: geometry.Point3{x1, y1, z1}, End2: geometry.Point3{x2, y2, z2}}
}

// return how many bricks fall
func fall(bricks []Brick) int {
	fellNum := 0

	for i := range bricks {
		b := &bricks[i] // need the real brick, not a copy, since we're going to modify it
		thisBrickFell := false
	CheckForLowerBricks:
		for b.End1.Z > 1 { // while this brick is above the floor
			for j := i - 1; j >= 0; j-- { // for each bricks before this in the list, i.e. on its level or below
				ob := &bricks[j] // ob = Other Brick
				// if we drop b by one Z, will it collide with ob?
				if !ob.zapped { // a zapped brick does not collide
					if b.End2.Z-1 >= ob.End1.Z &&
						b.End1.Z-1 <= ob.End2.Z &&
						b.End2.X >= ob.End1.X &&
						b.End1.X <= ob.End2.X &&
						b.End2.Y >= ob.End1.Y &&
						b.End1.Y <= ob.End2.Y {
						break CheckForLowerBricks
					}
				}
			}
			if !thisBrickFell {
				fellNum++
				thisBrickFell = true
			}
			b.FallBy(1)
		}
	}

	return fellNum
}
