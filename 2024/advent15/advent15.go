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

	// common directional instructions

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

func run2(input string) int {

	return 0
}

func score(b *grid.Board) int {
	score := 0
	for point, tile := range b.Grid {
		if tile.Contents == 'O' {
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
