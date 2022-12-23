package main

import (
	_ "embed"
	"fmt"
	"github.com/sekullbe/advent/parsers"
	"github.com/sekullbe/advent/tools"
)

//go:embed input.txt
var inputText string

//go:embed sample.txt
var sampleText string

type Instruction interface {
	execute(b *board) int
}
type MoveInst int
type TurnInst rune

var moveOnCube = false

func main() {
	fmt.Printf("Magic number: %d\n", run1(inputText))
	fmt.Println("-------------")
	fmt.Printf("Magic number: %d\n", run2(inputText))
	fmt.Println("-------------")

}

func parseBoard(lines []string) *board {
	bb := board{
		grove:    make(grid),
		maxX:     0,
		maxY:     0,
		moverLoc: nil,
		moverDir: 0,
	}
	b := &bb
	// start with 1,1
	for i, line := range lines {
		if len(line) == 0 {
			break
		}
		y := i + 1
		b.maxY = tools.MaxInt(b.maxY, y) // massive waste, I don't care
		for j, tc := range line {
			x := j + 1
			b.maxX = tools.MaxInt(b.maxX, x) // massive waste, I don't care
			if tc == ' ' {
				continue
			}
			p := np(x, y)
			t := tile{Point: p, contents: tc, neighbors: [4]*tile{}}
			b.grove[p] = &t
		}
	}
	// place the mover
	for x := 1; x < b.maxX; x++ {
		if t, ok := b.grove[np(x, 1)]; ok {
			b.moverLoc = t
			break
		}
	}
	// compute cube side size (pt 2 only)
	b.sideSize = tools.MinInt(b.maxX, b.maxY) / 3 // cube unfold always fits inside a grid with one side 3

	return b
}

func parseInstructions(instructionLine string) []Instruction {
	// let's do something clever with interfaces
	var instructions []Instruction
	var buildingMove string
	for _, ir := range instructionLine {
		if ir == 'R' || ir == 'L' {
			if buildingMove != "" {
				instructions = append(instructions, MoveInst(tools.Atoi(buildingMove)))
				buildingMove = ""
			}
			instructions = append(instructions, TurnInst(ir))
		} else {
			buildingMove = buildingMove + string(ir)
		}
	}
	if buildingMove != "" {
		instructions = append(instructions, MoveInst(tools.Atoi(buildingMove)))
	}
	return instructions
}

func run1(inputText string) int {
	lines := parsers.SplitByLinesNoTrim(inputText)
	b := parseBoard(lines)
	moveOnCube = false
	// so now we know that the blank line is lines[b.maxY] and the instructions are at lines[b.maxY] +1
	instructions := parseInstructions(lines[b.maxY+1])
	for _, instruction := range instructions {
		n := instruction.execute(b)
		_ = n
	}
	password := b.moverLoc.Y*1000 + b.moverLoc.X*4 + b.moverDir
	return password
}

func (t TurnInst) execute(b *board) int {
	//log.Printf("Turning %c", t)
	return b.turn(rune(t))
}

func (m MoveInst) execute(b *board) int {
	mv := b.move(int(m), moveOnCube)
	//log.Printf("Moving %d/%v", mv, m)
	return mv
}

func run2(inputText string) int {
	lines := parsers.SplitByLinesNoTrim(inputText)
	b := parseBoard(lines)
	moveOnCube = true
	// so now we know that the blank line is lines[b.maxY] and the instructions are at lines[b.maxY] +1
	instructions := parseInstructions(lines[b.maxY+1])
	// Same as part 1 except we need a new 'look' function to look around the edge of a cube
	for _, instruction := range instructions {
		n := instruction.execute(b)
		_ = n
	}
	password := b.moverLoc.Y*1000 + b.moverLoc.X*4 + b.moverDir
	return password
	//120299 too high... is my direction wrong? it's 3=north by the map
	//74288 ..  x= 74 y = 72 d = 0 is correct
	// I get X=74 Y =120 d=3 double check my Ys and directions
}
