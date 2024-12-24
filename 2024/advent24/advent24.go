package main

import (
	"cmp"
	_ "embed"
	"fmt"
	"github.com/sekullbe/advent/parsers"
	"github.com/sekullbe/advent/tools"
	"slices"
	"strings"
	"time"
)

//go:embed input.txt
var inputText string

const (
	AND = iota
	OR
	XOR
)

type opType = uint8

type wires map[string]wire
type gates []*gate

type gate struct {
	name   string
	op     opType
	valid  bool
	inputs []string
	output string
}

type wire struct {
	name  string
	value uint8
	valid bool
	//	gates []gate
}

func main() {
	fmt.Printf("Magic number: %d\n", run1(inputText))
	fmt.Println("-------------")
	fmt.Printf("Magic number: %d\n", run2(inputText))
}

func run1(input string) int {
	defer tools.Track(time.Now(), "Part 1 Time")

	segments := parsers.SplitByEmptyNewlineToSlices(input)
	wires := parseWires(segments[0])
	gates := parseGates(segments[1])
	// not all the wires that exist are in the wires segment!
	// so go through the gates and create invalid wires for each wire not already known
	completeWires(wires, gates)

	for !tick(wires, gates) {
	}

	return getOutput(wires)
}

func run2(input string) int {
	defer tools.Track(time.Now(), "Part 2 Time")

	return 0
}

func completeWires(wires wires, gates gates) {
	for _, g := range gates {
		iw1n := g.inputs[0]
		iw2n := g.inputs[1]
		own := g.output
		if _, exists := wires[iw1n]; !exists {
			w := wire{name: iw1n, value: 0, valid: false}
			wires[iw1n] = w
		}
		if _, exists := wires[iw2n]; !exists {
			w := wire{name: iw2n, value: 0, valid: false}
			wires[iw2n] = w
		}
		if _, exists := wires[own]; !exists {
			w := wire{name: own, value: 0, valid: false}
			wires[own] = w
		}
	}
}

func getOutput(wires wires) int {
	zwires := []wire{}
	for n, w := range wires {
		if n[0] == 'z' {
			zwires = append(zwires, w)
		}
	}

	// need them in reverse order, eg {z02, z01, z00}, because we process the highest bit first
	slices.SortFunc(zwires, func(a, b wire) int {
		return cmp.Compare(b.name, a.name)
	})
	return wiresToDecimal(zwires)
}

func wiresToDecimal(wires []wire) int {
	out := 0
	for _, w := range wires {
		out <<= 1
		out += int(w.value)
	}

	return out
}

func tick(wires wires, gates gates) bool {
	done := true
	for _, g := range gates {
		if g.valid {
			continue // this one has been completed
		}
		// see if its inputs are valid
		// if they are, set gate valid, set wires to value, set wires valid
		iw1 := wires[g.inputs[0]]
		iw2 := wires[g.inputs[1]]
		ow := wires[g.output]
		if iw1.valid && iw2.valid {
			g.valid = true
			ow.value = operate(g.op, iw1.value, iw2.value)
			ow.valid = true
			wires[g.output] = ow
		} else {
			done = false // there is a gate which cannot be evaluated, so we'll need to loop again
		}
	}
	return done
}

func operate(op opType, a, b uint8) uint8 {
	switch op {
	case AND:
		return a & b
	case OR:
		return a | b
	case XOR:
		return a ^ b
	default:
		panic("invalid operation")
	}
}

func parseWires(lines []string) map[string]wire {
	// x00: 1
	wires := make(map[string]wire, len(lines))

	for _, line := range lines {
		w := parseOneWire(line)
		wires[w.name] = w
	}

	return wires
}

func parseOneWire(line string) wire {
	p := strings.Split(line, " ")
	w := wire{
		name:  strings.TrimSuffix(p[0], ":"),
		value: uint8(tools.Atoi(p[1])),
		valid: true,
	}
	return w
}

func parseGates(lines []string) []*gate {
	gates := make([]*gate, 0, len(lines))
	for _, line := range lines {
		g := parseOneGate(line)
		gates = append(gates, g)
	}
	return gates
}

func parseOneGate(line string) *gate {
	// x00 AND y00 -> z00
	var iw1, opstr, iw2, ow string
	fmt.Sscanf(line, "%s %s %s -> %s", &iw1, &opstr, &iw2, &ow)
	op := AND
	switch opstr {
	case "AND":
		op = AND
	case "OR":
		op = OR
	case "XOR":
		op = XOR
	default:
		panic("Invalid operator")
	}
	g := gate{
		name:   line,
		op:     opType(op),
		valid:  false,
		inputs: []string{iw1, iw2},
		output: ow,
	}
	return &g
}