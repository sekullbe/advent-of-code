package main

import (
	_ "embed"
	"fmt"
	mapset "github.com/deckarep/golang-set/v2"
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
	//b.printBoard()
	load := b.computeLoad()
	return load
}

func (b *board) computeLoad() int {
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

// BUG: when rolled all the way to the edge, the *est arrays are off by one
// so either reset them each time, or when a ROUND leaves a spot, reset its *est

func (b *board) rollWest() {
	for x := 0; x <= b.maxX; x++ {
		for y := 0; y <= b.maxY; y++ {
			switch b.grid[Pt(x, y)].rock {
			case CUBE:
				b.westest[y] = x
			case ROUND:
				if b.westest[y]+1 <= x {
					//move the rock, even 0 steps
					err := b.moveRock(Pt(x, y), Pt(b.westest[y]+1, y))
					if err != nil {
						panic(err)
					}
					b.westest[y] = b.westest[y] + 1
				}
			default:
				// do nothing
			}
		}
	}
}

func (b *board) rollSouth() {
	for y := b.maxY; y >= 0; y-- {
		for x := 0; x <= b.maxX; x++ {
			switch b.grid[Pt(x, y)].rock {
			case CUBE:
				b.southest[x] = y
			case ROUND:
				if b.southest[x]-1 >= y {
					//move the rock, even 0 steps
					err := b.moveRock(Pt(x, y), Pt(x, b.southest[x]-1))
					if err != nil {
						panic(err)
					}
					b.southest[x] = b.southest[x] - 1
				}
			default:
				// do nothing
			}
		}
	}
}

func (b *board) rollEast() {
	for x := b.maxX; x >= 0; x-- {
		for y := 0; y <= b.maxY; y++ {
			switch b.grid[Pt(x, y)].rock {
			case CUBE:
				b.eastest[y] = x
			case ROUND:
				if b.eastest[y]-1 >= x {
					//move the rock, even 0 steps
					err := b.moveRock(Pt(x, y), Pt(b.eastest[y]-1, y))
					if err != nil {
						panic(err)
					}
					b.eastest[y] = b.eastest[y] - 1
				}
			default:
				// do nothing
			}
		}
	}
}

func run2(input string) int {

	s := mapset.NewSet[int]()

	b := parseBoard(parsers.SplitByLines(input))
	for i := 1; i <= 200; i++ {
		b.cycle()
		//b.printBoard()
		l := b.computeLoad()
		added := s.Add(l)
		if !added {
			fmt.Printf("Beginning cycle on step %d\n", i)
			//
		}
		fmt.Printf("%d: %d\n", i, l)
	}
	//	fmt.Println()
	//b.printBoard()
	/*
			   With the sample data it starts cycling on step 3 with period 7
			   87 69
			   69 69 65 64 65 63 68
			   69 69 65 64 65 63 68
		       etc
			   (1_000_000_000 - 2) mod

			   In the live ive data
			   we start seeing constant repeats at step 191
			   first time we've seen step 191 load  is step 89
			   (1_000_000_000 - 88) mod (191-81) + 88 = 160
			   step 160 answer is my answer
	*/

	return 0
}

func run2_hard_way(input string) int {
	b := parseBoard(parsers.SplitByLines(input))
	for i := 0; i < 1_000_000_000; i++ {
		b.cycle()
		if i%1_000 == 0 {
			fmt.Print(".")
		}
	}
	return b.computeLoad()
}

func (b *board) cycle() {
	b.rollNorth()
	b.resetEsts()
	b.rollWest()
	b.resetEsts()
	b.rollSouth()
	b.resetEsts()
	b.rollEast()
	b.resetEsts()
}
