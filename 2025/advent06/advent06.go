package main

import (
	_ "embed"
	"fmt"
	"strings"
	"time"

	"github.com/samber/lo"
	"github.com/sekullbe/advent/grid"
	"github.com/sekullbe/advent/parsers"
	"github.com/sekullbe/advent/tools"
)

//go:embed input.txt
var inputText string

func main() {
	fmt.Printf("Magic number: %d\n", run1(inputText))
	fmt.Println("-------------")
	fmt.Printf("Magic number: %d\n", run2(inputText))
}

func run1(input string) int {
	defer tools.Track(time.Now(), "Part 1 Time")

	problems := parseMathInputsPart1(input)
	total := 0
	for _, p := range problems {
		total += solveProblem(p)
	}

	return total
}

func run2(input string) int {
	defer tools.Track(time.Now(), "Part 2 Time")

	// I have this library, might as well use it
	b := grid.ParseBoardString(input)
	b.PrintBoard()
	numStartColumns := []int{0}
	for x := 0; x <= b.MaxX; x++ {
		if b.IsColumnEmpty(x) {
			numStartColumns = append(numStartColumns, x+1)
		}
	}
	numStartColumns = append(numStartColumns, b.MaxX+2) // cheat so we can scan the last column
	// this was +2 so the last column was included in the _previous_ column group

	// for each NSC number, scan down from y=0 to MaxY-1 and get a number
	// stick that in a problem{} - get the operator if it's there when looking at each column
	// when you hit the next NSC, solve the problem and add it to the sum
	total := 0
	for i, xsc := range numStartColumns {
		if xsc > b.MaxX {
			continue
		}
		p := problem{}
		if !grid.IsEmpty(b.At(xsc, b.MaxY).Contents) {
			p.operator = string(b.At(xsc, b.MaxY).Contents)
		}
		for x := xsc; x < numStartColumns[i+1]-1; x++ {
			var numstr string
			for y := 0; y <= b.MaxY-1; y++ {
				r := b.At(x, y).Contents
				if !grid.IsEmpty(r) {
					numstr = numstr + string(r)
				}
			}
			n := tools.Atoi(numstr)
			//log.Println(n)
			p.operands = append(p.operands, n)
		}
		//log.Printf("operating on %v, solution is %d", p, solveProblem(p))
		total += solveProblem(p)
	}

	return total
}

type problem struct {
	operands []int
	operator string
}

func parseMathInputsPart1(input string) []problem {
	lines := parsers.SplitByLines(input)
	problems := []problem{}
	for i := 0; i < len(strings.Fields(lines[0])); i++ {
		p := problem{[]int{}, ""}
		for ln, line := range lines {
			f := strings.Fields(line)
			if ln < len(lines)-1 {
				p.operands = append(p.operands, tools.Atoi(f[i]))
			} else {
				p.operator = f[i]
			}
		}
		problems = append(problems, p)
	}
	return problems
}

func solveProblem(p problem) int {

	switch p.operator {
	case "+":
		return tools.SumSlice(p.operands)
	case "*":
		return lo.Reduce(p.operands, func(total, a, b int) int { return total * a }, 1)
	}

	return -1

}
