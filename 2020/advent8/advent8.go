package advent8

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var inputText string

type memory map[int]*instruction

func Run() {
	Run2Doit(inputText)
}

func Run1Doit(inputs string) {

	memory := populateMemory(inputs)
	accumulator := 0
	memPtr := 0

	for true {
		newAcc, memPtrOffset, executionCount := memory[memPtr].execute(accumulator)
		// if we just hit an instruction for the 2nd time, don't apply its results
		if executionCount == 2 {
			break
		}
		accumulator = newAcc
		memPtr += memPtrOffset

	}

	fmt.Printf("Completed. Accumulator=%d\n", accumulator)
}

func Run2Doit(inputs string) (int, bool) {

	memory := populateMemory(inputs) // waste to calculate it but meh
	success := false
	accumulator := 0
	frobPtr := 0

	for i := 0; i < len(memory) ; i++ {
		memPtr := 0
		accumulator = 0
		frobPtr = i
		memory := populateMemory(inputs)
		memory[frobPtr].frob()
		for true {
			newAcc, memPtrOffset, executionCount := memory[memPtr].execute(accumulator)
			if memory[memPtr].opcode == "end" {
				success = true
				break
			}
			if executionCount > 1 {
				// that was not the correct fix, reset memory and go again
				memory = populateMemory(inputs)
				break
			}
			accumulator = newAcc
			memPtr += memPtrOffset
		}
		if success {
			break
		}
	}

	fmt.Printf("Completed. Accumulator=%d\n", accumulator)
	fmt.Printf("Invalid instruction pointer = %d\n", frobPtr)

	return accumulator, success

}

func populateMemory(inputs string) memory {
	mem := make(memory)
	for i, s := range strings.Split(inputs, "\n") {
		mem[i] = newInstructionFromString(s)
	}
	return mem
}
