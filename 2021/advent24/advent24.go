package main

import (
	_ "embed"
	"fmt"
	"github.com/sekullbe/advent/parsers"
	"github.com/sekullbe/advent/tools"
	"log"
	"regexp"
	"strings"
)

//go:embed input.txt
var inputText string

func main() {
	fmt.Printf("Magic number: %d\n", run1(inputText))
	fmt.Println("-------------")
	fmt.Printf("Magic number: %d\n", run2(inputText))
}
func run1(inputText string) int {
	steps := parseInput(inputText)
	ok := checkModelNumber(94399898949959, steps)
	if ok {
		return 1
	}
	return 0
}

func run2(inputText string) int {
	steps := parseInput(inputText)
	ok := checkModelNumber(21176121611511, steps)
	if ok {
		return 1
	}
	return 0
}

func runstupid(inputText string) int {
	steps := parseInput(inputText)
	// let's see if we can be smarter about the number
	// first try cycling digits
	// and can we parallelize it?
	for mn := 99999999999999; mn >= 11111111111111; mn -= 1 {
		ok := checkModelNumber(mn, steps)
		if ok {
			return mn
		}
	}
	return 0
}

func checkModelNumber(mn int, steps []step) bool {
	mns := fmt.Sprintf("%d", mn)
	if strings.ContainsAny(mns, "0") {
		return false
	}
	inputs := []int{}
	for _, r := range mns {
		inputs = append(inputs, int(r-'0'))
	}
	a := newAlu()
	for _, s := range steps {
		runStep(a, s, &inputs)
	}
	//fmt.Println(a.toString())
	if a.get(Z) == 0 {
		return true
	}
	return false
}

const W = 'w'
const X = 'x'
const Y = 'y'
const Z = 'z'
const NOREG = rune(0)

// could use constants for instructions, but it's either a switch on input to decode them
// or a switch on use to dispatch them- doesn't really matter

// only one or the other of these will be set
type step struct {
	instruction string
	rTo         rune
	rFrom       rune
	value       int
}

// maybe make this a method too
func runStep(a *alu, s step, inputs *[]int) {
	switch s.instruction {
	case "inp":
		a.inp(s, inputs)
	case "add":
		a.add(s)
	case "mul":
		a.mul(s)
	case "div":
		a.div(s)
	case "mod":
		a.mod(s)
	case "eql":
		a.eql(s)
	default:
		log.Panicf("stop! panic time! can't parse this: %s", s.instruction)
	}
}

func parseInput(input string) []step {
	var steps []step
	// an input line is xxx F T?
	// F is always a register
	// T is optional and can be a register or int value
	lines := parsers.SplitByLines(input)
	for _, line := range lines {
		s := parseLineToStep(line)
		steps = append(steps, s)
	}
	return steps
}

func parseLineToStep(line string) (s step) {
	re := regexp.MustCompile(`(...) (.) ?(.*?)?$`)
	matches := re.FindStringSubmatch(line)
	s.instruction = matches[1]
	s.rTo = rune(matches[2][0])
	value := matches[3]
	if value == "" {
		return
	}
	if strings.ContainsAny(value, "01234567890") {
		s.value = tools.Atoi(value)
	} else {
		s.rFrom = rune(value[0])
	}
	return
}
