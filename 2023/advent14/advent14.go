package main

import (
	_ "embed"
	"fmt"
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

	b := parseBoard(parsers.SplitByLines(input))
	b.rollNorth()

	// compute load
	// for each ROUND, load is maxY-y+1
	b.printBoard()
	load := 0
	for y := 0; y <= b.maxY; y++ {
		for x := 0; x <= b.maxX; x++ {
			rock := b.At(x, y).rock
			if rock == ROUND {
				load += b.maxY - y + 1
			}
		}
	}

	return load
}

func (b *board) rollNorth() {
	for y := 0; y <= b.maxY; y++ {
		for x := 0; x <= b.maxX; x++ {
			switch b.grid[Pt(x, y)].rock {
			case CUBE:
				b.northest[x] = y
			case ROUND:
				if b.northest[x]+1 <= y {
					//move the rock, even 0 steps
					err := b.moveRock(Pt(x, y), Pt(x, b.northest[x]+1))
					if err != nil {
						panic(err)
					}
					b.northest[x] = b.northest[x] + 1
				}
			default:
				// do nothing
			}
		}
	}
}

func run2(input string) int {

	return 0
}
