package main

import (
	_ "embed"
	"fmt"
	"github.com/sekullbe/advent/parsers"
	"github.com/sekullbe/advent/tools"
	"log"
	"strings"
)

//go:embed input.txt
var inputText string

func main() {
	fmt.Printf("Magic string: %s\n", run1(inputText))
	fmt.Println("-------------")
	fmt.Printf("Magic string: %d\n", run2(inputText))
}

type computer struct {
	program    []int
	rA, rB, rC int
	instPtr    int // index into program
}

func run1(input string) string {

	computer := initialize(input)

	output := run(computer)

	return output
}

func run2(input string) int {

	a := 0
	cpu := initialize(input)
	program := cpu.program
	output := "x"
	// the insight is that the last 3 bits of A produce the last number in the output
	// so iterate over all a in that space
	// once you find it, shift it left 3 spaces and count up A in that space again,
	// to produce the last 2 numbers
	// repeat until you've got the whole thing
	for i := len(program) - 1; i >= 0; i-- {
		a <<= 3
		output = "x"
		log.Println(i)
		for {
			cpu.instPtr = 0
			cpu.rA = a
			cpu.rB = 0
			cpu.rC = 0
			output = run(cpu)
			if output == tools.IntArrayToString(program[i:]) {
				break
			}
			a++
		}
	}
	return a
}

func initialize(input string) *computer {
	lines := parsers.SplitByLines(input)
	cpu := computer{instPtr: 0}
	fmt.Sscanf(lines[0], "Register A: %d", &cpu.rA)
	fmt.Sscanf(lines[1], "Register B: %d", &cpu.rB)
	fmt.Sscanf(lines[2], "Register C: %d", &cpu.rC)
	cpu.program = parsers.StringsWithCommasToIntSlice(strings.Split(lines[4], " ")[1])
	return &cpu
}

func run(cpu *computer) string {
	outputs := []int{}
	lastInst := len(cpu.program) - 1
	for cpu.instPtr <= lastInst {
		opcode := cpu.program[cpu.instPtr]
		operand := cpu.program[cpu.instPtr+1]
		combo := decodeOperand(cpu, operand)
		cpu.instPtr += 2
		switch opcode {
		case 0: // adv
			if combo == 7 {
				//panic("reserved operand")
			}
			cpu.rA = cpu.rA / tools.PowInt(2, combo)
		case 1: // bxl
			cpu.rB = cpu.rB ^ operand
		case 2: // bst
			if combo == 7 {
				//panic("reserved operand")
			}
			cpu.rB = combo % 8
		case 3: // jnz
			if cpu.rA > 0 {
				cpu.instPtr = operand
			}
		case 4: // bxc
			cpu.rB = cpu.rB ^ cpu.rC
		case 5: // out
			if combo == 7 {
				//panic("reserved operand")
			}
			outputs = append(outputs, combo%8)
		case 6: // bdv
			if combo == 7 {
				//panic("reserved operand")
			}
			cpu.rB = cpu.rA / tools.PowInt(2, combo)
		case 7: // cdv
			if combo == 7 {
				//panic("reserved operand")
			}
			cpu.rC = cpu.rA / tools.PowInt(2, combo)

		}

	}

	return tools.IntArrayToString(outputs)
}

func decodeOperand(cpu *computer, operand int) int {
	if operand >= 0 && operand <= 3 {
		return operand
	}
	switch operand {
	case 4:
		return cpu.rA
	case 5:
		return cpu.rB
	case 6:
		return cpu.rC
	case 7:
		//log.Println("WARNING: reserved operand")
		return -1

	}
	panic("invalid operand")
}
