package main

import (
	_ "embed"
	"errors"
	"fmt"
	"github.com/sekullbe/advent/parsers"
	"github.com/sekullbe/advent/tools"
	"strings"
)

//go:embed input.txt
var inputText string

func main() {
	fmt.Printf("Magic number: %d\n", run1(inputText))
	fmt.Println("-------------")
	fmt.Printf("Magic number: %d\n", run2(inputText))
}

/*
Button A: X+92, Y+24
Button B: X+13, Y+94
Prize: X=8901, Y=8574
*/

const costA = 3
const costB = 1

type machine struct {
	// should this use Points? only if I need the math
	aX, aY int
	bX, bY int
	pX, pY int
}

func run1(input string) int {

	cost := 0
	machines := parseMachineFile(input)
	for _, m := range machines {
		c, err := solveOneMachineCramer(m)
		if err == nil {
			cost += c
		}
	}

	return cost
}

func run2(input string) int {
	cost := 0
	machines := parseMachineFile(input)
	for _, m := range machines {
		m.pX += 10000000000000
		m.pY += 10000000000000
		c, err := solveOneMachineCramer(m)
		if err == nil {
			cost += c
		}
	}

	return cost
}

// returns pA and pB such that a*pA + b*pb = prize, or err if that can't be done
// this solves the sample and part1 correctly but is high on part 2
func solveOneMachine(m machine) (cost int, err error) {
	// brute force would work because of the limit of 100 pushes per button, but that's likely to change in part 2
	// https://math.stackexchange.com/questions/20717/how-to-find-solutions-of-linear-diophantine-ax-by-c
	gcdX := tools.GCD(m.aX, m.bX)
	gcdY := tools.GCD(m.aY, m.bY)
	// this catches some but not all unsolvable machines
	// removing it gives me a number on part 2 that is too high, but it passes my unit tests either way
	if m.pX%gcdX != 0 || m.pY%gcdY != 0 {
		return 0, errors.New("unsolvable by gcd check")
	}

	// go go gadget euclidean- find aC such that aC * aP = some rhs (aP = number of A pushes)
	aC := m.aX*m.bY - m.aY*m.bX
	rhs := m.pX*m.bY - m.pY*m.bX
	if rhs%aC != 0 { // this gets the rest
		return 0, errors.New("unsolvable by coeff check")
	}

	aP := rhs / aC
	bP := (m.pX - m.aX*aP) / m.bX
	// this isn't necessary in part 1 and is ignored in part 2
	if aP > 100 || bP > 100 {
		//return 0, errors.New("too many")
	}
	// this never fires on normal input but is a good catch that something has gone wrong
	if aP < 0 || bP < 0 {
		return 0, errors.New("can't have negative presses")
	}
	return aP*costA + bP*costB, nil
}

// very strangely, my original algorithm worked ok for part1 and every test
// but came in about 10% too high on part 2
// so let's try this Cramer's Rule solution
// again in tests it behaves identically, but solves part 2 correctly.
func solveOneMachineCramer(m machine) (cost int, err error) {
	aP := (m.pX*m.bY - m.pY*m.bX) / (m.aX*m.bY - m.aY*m.bX)
	bP := (m.aX*m.pY - m.aY*m.pX) / (m.aX*m.bY - m.aY*m.bX)
	// suppose I could check that those divisions were remainderless, but easier to just test the solution actually multiplies out
	if m.aX*aP+m.bX*bP != m.pX || m.aY*aP+m.bY*bP != m.pY {
		return 0, errors.New("unsolvable")
	}
	return aP*costA + bP*costB, nil
}

func parseMachineFile(input string) []machine {
	machineLines := parsers.SplitByEmptyNewlineToSlices(input)
	machines := []machine{}
	for _, ml := range machineLines {
		if len(ml) > 2 {
			machines = append(machines, parseMachine(ml))
		}
	}
	return machines
}

func parseMachine(mlines []string) machine {
	m := machine{}
	fmt.Fscanf(strings.NewReader(mlines[0]), "Button A: X+%d, Y+%d", &m.aX, &m.aY)
	fmt.Fscanf(strings.NewReader(mlines[1]), "Button B: X+%d, Y+%d", &m.bX, &m.bY)
	fmt.Fscanf(strings.NewReader(mlines[2]), "Prize: X=%d, Y=%d", &m.pX, &m.pY)
	return m
}
