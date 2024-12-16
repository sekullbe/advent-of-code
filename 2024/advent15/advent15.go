package main

import (
	_ "embed"
	"fmt"
	"github.com/sekullbe/advent/geometry"
	"github.com/sekullbe/advent/grid"
	"github.com/sekullbe/advent/parsers"
	"log"
)

//go:embed input.txt
var inputText string

func main() {
	fmt.Printf("Magic number: %d\n", run1(inputText))
	fmt.Println("-------------")
	fmt.Printf("Magic number: %d\n", run2(inputText))
}

func run1(input string) int {

	inputSegments := parsers.SplitByEmptyNewlineToSlices(input)
	b := grid.ParseBoard(inputSegments[0])
	instructions := decodeInstructions(inputSegments[1])

	//b.PrintBoard()

	robotPoint := findRobot(b)

	for _, instruction := range instructions {
		//	log.Printf("moving %d", instruction)
		robotPoint = moveRobot(b, robotPoint, instruction)
		//	b.PrintBoard()
	}

	return score(b)
}

// what's different here
// after reading the board, doubleize it
// a box is a [] not O
// a [] moves l/r just like an O except both have to move
// moving up down the same except one box can push two
func run2(input string) int {
	inputSegments := parsers.SplitByEmptyNewlineToSlices(input)
	b := grid.ParseBoard(inputSegments[0])
	doubleize(b)
	//b.PrintBoard()
	instructions := decodeInstructions(inputSegments[1])
	robotPoint := findRobot(b)
	for _, instruction := range instructions {
		log.Printf("moving %d", instruction)
		robotPoint = moveRobotWithWideBoxes(b, robotPoint, instruction)
		//	b.PrintBoard()
	}
	//b.PrintBoard()
	return score(b)
}

func doubleize(b *grid.Board) {

	newGrid := make(grid.Grid)
	for y := 0; y <= b.MaxY; y++ {
		for x := 0; x <= b.MaxX; x++ {
			pl := geometry.NewPoint2(2*x, y)
			pr := geometry.NewPoint2(2*x+1, y)
			cl := b.At(x, y).Contents
			cr := 'X'
			switch cl {
			case 'O':
				cl, cr = '[', ']'
			case '.':
				cl, cr = '.', '.'
			case '@':
				cl, cr = '@', '.'
			case '#':
				cl, cr = '#', '#'
			default:
				log.Panicf("dunno what this is: %c", cl)
			}
			tl := grid.NewTile(pl, cl)
			newGrid[pl] = &tl
			tr := grid.NewTile(pr, cr)
			newGrid[pr] = &tr
		}
	}
	b.Grid = newGrid
	b.MaxX = (b.MaxX * 2) + 1
}

func score(b *grid.Board) int {
	score := 0
	for point, tile := range b.Grid {
		if tile.Contents == 'O' || tile.Contents == '[' {
			score += 100*point.Y + point.X
		}

	}
	return score
}

func findRobot(b *grid.Board) geometry.Point2 {
	for point, tile := range b.Grid {
		if tile.Contents == '@' {
			return point
		}
	}
	panic("no robot")
}

func decodeInstructions(dirs []string) []int {
	instructions := []int{}
	for _, s := range dirs {
		dirRunes := []rune(s)
		for _, dirRune := range dirRunes {
			inst, ok := grid.DirRunes[dirRune]
			if ok {
				instructions = append(instructions, inst)
			}
		}
	}
	return instructions
}

// return the new point where the robot is
func moveRobot(b *grid.Board, rp geometry.Point2, dir int) geometry.Point2 {
	// because the board is bounded we don't need to check for offboard
	target := b.AtPoint(grid.NeighborInDirection(rp, dir))
	if grid.IsEmpty(target.Contents) {
		b.SlideTile(rp, dir)
		return target.Point
	} else if target.Contents == 'O' {
		// keep looking ahead for a space
		// if there is one, move target box into that space then move robot to target
		// if not, can't move

		nt := target
		for {
			nt = b.AtPoint(grid.NeighborInDirection(nt.Point, dir))
			if nt.Contents == '#' {
				break // can't move
			} else if nt.Contents == '.' { // move box to empty(nt), robot to target
				nt.Contents = 'O'
				target.Contents = '.'
				b.SlideTile(rp, dir)
				return target.Point
			} else { // it's a box 'O'
				continue // this is going to happen anyway but be explicit, also a place to but a breakpoint if needed
			}
		}
	} else if target.Contents == '#' {
		log.Println("ow my nose")
		return rp
	} else {
		log.Fatalf("that robot ain't right")
	}
	return rp
}

// return the new point where the robot is
func moveRobotWithWideBoxes(b *grid.Board, rp geometry.Point2, dir int) geometry.Point2 {
	// because the board is bounded we don't need to check for offboard
	target := b.AtPoint(grid.NeighborInDirection(rp, dir))
	if grid.IsEmpty(target.Contents) {
		b.SlideTile(rp, dir)
		return target.Point
	} else if target.Contents == '[' || target.Contents == ']' {
		if dir == grid.EAST || dir == grid.WEST {
			rp = moveBoxHorizontally(b, rp, target, dir)
		} else {
			if recursivelyMoveBoxVertically(b, target.Point, dir, false) {
				recursivelyMoveBoxVertically(b, target.Point, dir, true)
				b.SlideTile(rp, dir) // move the robot too
				rp = grid.NeighborInDirection(rp, dir)
			}
		}

	} else if target.Contents == '#' {
		//log.Println("ow my nose")
		return rp
	} else {
		log.Fatalf("that robot ain't right")
	}
	return rp
}

func moveBoxHorizontally(b *grid.Board, rp geometry.Point2, target *grid.Tile, dir int) geometry.Point2 {
	// look ahead for a space
	// if there is one, move [ (west) or ] (east) into it
	// toggle each [] from rp to there
	// and move the robot
	nt := target
	toggles := []geometry.Point2{}
	for {
		nt = b.AtPoint(grid.NeighborInDirection(nt.Point, dir))
		if nt.Contents == '#' {
			break // can't move
		} else if nt.Contents == '.' { // move half the box into that square
			if dir == grid.EAST {
				nt.Contents = ']'
			} else {
				nt.Contents = '['
			}
			// now toggle each box part between rp and nt
			for _, toggle := range toggles {
				tt := b.AtPoint(toggle)
				if tt.Contents == '[' {
					tt.Contents = ']'
				} else if tt.Contents == ']' {
					tt.Contents = '['
				} else {
					log.Panicf("tried to toggle non-box %v", tt)
				}
			}
			// and finally move the robot
			target.Contents = '.'
			b.SlideTile(rp, dir)
			return target.Point
		} else { // it's a box
			toggles = append(toggles, nt.Point)
		}
	}
	return rp
}

func recursivelyMoveBoxVertically(b *grid.Board, bp geometry.Point2, dir int, commit bool) bool {

	// bpL/r = Box Point Left/Right - where is the box now
	bpL, bpR := bothPointsForBox(b, bp)
	// Target Tile Left/Right -- this is where this block would move
	ttL := b.AtPoint(grid.NeighborInDirection(bpL, dir))
	ttR := b.AtPoint(grid.NeighborInDirection(bpR, dir))

	iCanMove := true

	if ttL.Contents == '#' || ttR.Contents == '#' {
		// wall, cannot move
		return false
	}

	// if there's a box ahead of us, try to move it- these will return true if that box moved

	if ttL.Contents == '[' || ttL.Contents == ']' { // there is a box directly above or to the left above
		iCanMove = recursivelyMoveBoxVertically(b, ttL.Point, dir, commit)
	}
	if ttR.Contents == '[' { // there is a box to the right above
		iCanMove = iCanMove && recursivelyMoveBoxVertically(b, ttR.Point, dir, commit)
	}

	if ttL.Contents == '.' && ttR.Contents == '.' {
		// we can move cleanly
		if commit {
			b.SlideTile(bpL, dir)
			b.SlideTile(bpR, dir)
		}
		return true
	}

	return iCanMove
}

func bothPointsForBox(b *grid.Board, bp geometry.Point2) (geometry.Point2, geometry.Point2) {
	bt := b.AtPoint(bp)
	if bt.Contents == '[' {
		return bp, geometry.NewPoint2(bp.X+1, bp.Y)
	} else {
		return geometry.NewPoint2(bp.X-1, bp.Y), bp
	}
}
