package advent8

import (
	"strconv"
	"strings"
)

type instruction struct {
	opcode         string
	arg            int
	executionCount int
}

func newInstruction(opcode string, addr string) *instruction {
	addrNum, err  := strconv.Atoi(addr)
	if err != nil {
		return &instruction{opcode: "nop", arg: 0}
	}
	return &instruction{opcode: opcode, arg: addrNum}
}

func newInstructionFromString(opText string) *instruction {
	opElements := strings.Fields(opText)
	if len(opElements) != 2 {
		// this catches the newline at the end of the file and is kind of a hack
		return newInstruction("end", "0")
	}
	return newInstruction(opElements[0], opElements[1])
}

func (i *instruction) frob()  {
	switch i.opcode {
	case "jmp":
		i.opcode = "nop"
	case "nop":
		i.opcode = "jmp"
	}
}

// returns accumulator and next address
func (i *instruction) execute(accumulator int) (int, int, int) {

	//fmt.Printf("+++Executing %s: acc=%d\n", i.opcode, accumulator)
	i.executionCount += 1
	switch i.opcode {
	case "acc":
		return accumulator + i.arg, 1, i.executionCount
	case "jmp":
		return accumulator, i.arg, i.executionCount
	case "end":
		return accumulator, 0, -1
	}
	// nop or error
	return accumulator, 1, i.executionCount
}
