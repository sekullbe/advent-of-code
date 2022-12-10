package main

import (
	_ "embed"
	"fmt"
	"github.com/sekullbe/advent/parsers"
	"log"
)

//go:embed input.txt
var inputText string

func main() {
	lines := parsers.SplitByLines(inputText)
	fmt.Printf("Magic number: %d\n", run1(lines))
	fmt.Println("-------------")
	fmt.Printf("Magic number: %d\n", run2(lines))
}

type processor struct {
	clock            int
	xreg             int // if there are more regs make this a map
	operationTime    int
	scheduledOp      string
	scheduledArg     int
	instructionTimes map[string]int
}

func newProcessor() processor {
	it := make(map[string]int) // ADDITIONAL ticks to complete this argument; if 0 it happens immediately
	it["noop"] = 0
	it["addx"] = 1
	return processor{
		clock:            0,
		xreg:             1,
		operationTime:    0,
		scheduledOp:      "",
		scheduledArg:     0,
		instructionTimes: it,
	}
}

// returns true if the instruction was executed and we need a new one next tick
func (p *processor) tick(nextInstr string) (consumedInstruction bool, xregDuring int) {

	p.clock += 1
	xregDuring = p.xreg
	if p.operationTime > 0 {
		p.operationTime--
		p.executeOp(p.scheduledOp, p.scheduledArg)
		consumedInstruction = true
	} else {
		// read an op and either do it or schedule it
		var op string
		var arg int
		fmt.Sscanf(nextInstr, "%s %d", &op, &arg)

		optime, ok := p.instructionTimes[op]
		if !ok {
			log.Panicf("Unimplemented op: %s", nextInstr)
		}
		if optime > 0 {
			// Given that I don't move the instruction pointer until the slow op completes, I could just read it again.
			// But it's already parsed, so might as well just store it.
			p.scheduledOp = op
			p.scheduledArg = arg
			p.operationTime = optime
			consumedInstruction = false
		} else {
			p.executeOp(op, arg)
			consumedInstruction = true
		}
	}
	return
}

func (p *processor) executeOp(op string, arg int) {
	switch op {
	case "noop":
		return
	case "addx":
		p.xreg += arg
	default:
		log.Panicf("Unimplemented op: %s", op)
	}
}

func run1(lines []string) int {

	instrPtr := 1
	proc := newProcessor()
	sumStrengths := 0
	for instrPtr <= len(lines) {
		line := lines[instrPtr-1]
		completedOp, xregDuring := proc.tick(line)
		if completedOp {
			instrPtr++
		}
		if (proc.clock-20)%40 == 0 {
			signalStrength := proc.clock * xregDuring
			sumStrengths += signalStrength
			log.Printf("Cycle %d: Signal strength %d, sum %d", proc.clock, signalStrength, sumStrengths)
		}
	}

	return sumStrengths
	// 14660 too high
}

func run2(lines []string) int {
	instrPtr := 1
	proc := newProcessor()
	for instrPtr <= len(lines) {
		line := lines[instrPtr-1]
		completedOp, xregDuring := proc.tick(line)
		if completedOp {
			instrPtr++
		}
		spritePos := (proc.clock - 1) % 40
		pixel := ".."
		if xregDuring == spritePos || xregDuring-1 == spritePos || xregDuring+1 == spritePos {
			pixel = "##"
		}
		fmt.Printf("%s", pixel)
		if proc.clock%40 == 0 {
			fmt.Println()
		}
	}

	return 0
}
