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

func main() {
	fmt.Printf("Magic number: %d\n", run1(inputText))
	fmt.Println("-------------")
	fmt.Printf("Magic number: %d\n", run2(inputText))
}

func run1(inputText string) int {
	// don't care about the spaces to the right- we just need to the left, to get the grid correct
	lines := parsers.SplitByLinesNoTrim(inputText)
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
	// so now we know that the blank line is lines[b.maxY] and the instructions are at lines[b.maxY] +1
	instructionLine := lines[b.maxY+1]
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
	// do we really need to store them? why not just execute them? eh, that's what I've done already
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
	mv := b.move(int(m))
	//log.Printf("Moving %d/%v", mv, m)
	return mv
}

func run2(inputText string) int {

	return 0
}
