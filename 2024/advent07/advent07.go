package main

import (
	_ "embed"
	"fmt"
	"github.com/sekullbe/advent/parsers"
	"github.com/sekullbe/advent/tools"
	"strconv"
	"strings"
)

//go:embed input.txt
var inputText string

func main() {
	fmt.Printf("Magic number: %d\n", run1(inputText))
	fmt.Println("-------------")
	fmt.Printf("Magic number: %d\n", run2(inputText))
}

type equation struct {
	testValue int
	operands  []int
}

func run1(input string) int {

	equations := make([]equation, 0)
	for _, s := range parsers.SplitByLines(input) {
		equations = append(equations, parseEquation(s))
	}
	sum := 0
	for _, e := range equations {
		if solvable(e) {
			sum += e.testValue
		}
	}

	return sum
}

func run2(input string) int {

	equations := make([]equation, 0)
	for _, s := range parsers.SplitByLines(input) {
		equations = append(equations, parseEquation(s))
	}
	sum := 0
	for _, e := range equations {
		if solvableWithConcat(e) {
			sum += e.testValue
		}
	}

	return sum
}

func parseEquation(eqstr string) equation {

	eq := equation{0, []int{}}

	c := strings.Split(eqstr, ":")
	eq.testValue = tools.Atoi(c[0])
	eq.operands = parsers.StringsToIntSlice(c[1])
	return eq
}

// in both cases, start from the end and see if there's an obvious operator

func solvable(e equation) bool {
	last := e.operands[len(e.operands)-1]
	// if there's one number left, we're done
	if len(e.operands) == 1 {
		return e.testValue == last
	}
	// can't ever come back from being too high
	if last > e.testValue {
		return false
	}
	//if the last number is divisible, do that
	if e.testValue%last == 0 && solvable(equation{e.testValue / last, e.operands[0 : len(e.operands)-1]}) {
		return true
	}
	// if the last number is subtractable, do that
	if solvable(equation{e.testValue - last, e.operands[0 : len(e.operands)-1]}) {
		return true
	}

	return false
}

// two functions is better than one with a boolean behavior switch.
func solvableWithConcat(e equation) bool {
	last := e.operands[len(e.operands)-1]
	// if there's one number left, we're done
	if len(e.operands) == 1 {
		return e.testValue == last
	}
	// can't ever come back from being too high
	if last > e.testValue {
		return false
	}
	//if the last number is divisible, do that
	if e.testValue%last == 0 && solvableWithConcat(equation{e.testValue / last, e.operands[0 : len(e.operands)-1]}) {
		return true
	}
	// if the last number is subtractable, do that and recurse the rest
	if solvableWithConcat(equation{e.testValue - last, e.operands[0 : len(e.operands)-1]}) {
		return true
	}
	// if the testresult ends with last, do that, but be careful you don't remove the whole thing
	if e.testValue != last && strings.HasSuffix(strconv.Itoa(e.testValue), strconv.Itoa(last)) && solvableWithConcat(equation{tools.Atoi(strings.TrimSuffix(strconv.Itoa(e.testValue), strconv.Itoa(last))), e.operands[0 : len(e.operands)-1]}) {
		return true
	}

	return false
}
