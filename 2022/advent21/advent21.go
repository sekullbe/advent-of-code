package main

import (
	_ "embed"
	"fmt"
	"github.com/sekullbe/advent/parsers"
	"github.com/sekullbe/advent/tools"
	"log"
)

//go:embed input.txt
var inputText string

func main() {
	fmt.Printf("Magic number: %d\n", run1(inputText))
	fmt.Println("-------------")
	fmt.Printf("Magic number: %d\n", run2(inputText))
}

type formula struct {
	yeller, op1, op2 string
	operator         rune
	knowOneOperand   bool
	knownOperand     int
}

type formulas map[string]*formula
type numbers map[string]int

func parseMonkeys(lines []string) (formulas, numbers, *formula) {

	fmap := make(formulas)
	nmap := make(numbers)
	var froot *formula
	for _, line := range lines {
		var my, m1, m2 string
		var op rune
		var mnum int
		n, _ := fmt.Sscanf(line, "%4s: %4s %c %4s", &my, &m1, &op, &m2)
		if n != 4 {
			n, _ = fmt.Sscanf(line, "%4s: %d", &my, &mnum)
			if tools.KeyExists(fmap, m1) || tools.KeyExists(fmap, m2) {
				log.Panicf("crap, we see %s or %s as operands n more than one formula", m1, m2)
			}
			nmap[my] = mnum
		} else {
			f := formula{yeller: my, op1: m1, op2: m2, operator: op}
			fmap[m1] = &f
			fmap[m2] = &f
			if my == "root" {
				froot = &f
			}
		}
	}
	return fmap, nmap, froot
}

func operate(f *formula, yeller string, num int) int {
	res := 0
	switch f.operator {
	case '+':
		res = num + f.knownOperand
	case '-':
		if yeller == f.op1 {
			res = num - f.knownOperand
		} else {
			res = f.knownOperand - num
		}
	case '*':
		res = num * f.knownOperand
	case '/':
		if yeller == f.op1 {
			res = num / f.knownOperand
		} else {
			res = f.knownOperand / num
		}
	case '=':
		res = 0
		log.Println("operating on root?")
	}
	return res
}

func solve(fm formulas, nm numbers) int {
	root := 0
keepsearching:
	for {
		for yeller, num := range nm {
			// It should be an operand in exactly one formula
			f := fm[yeller]
			// Catch the case where we've solved and deleted a formula
			if f == nil {
				continue
			}
			// can we solve that formula now?
			if f.knowOneOperand {
				// we can, hooray!
				res := operate(f, yeller, num)
				nm[f.yeller] = res
				// the formula is solved, we don't need it any longer
				delete(fm, f.op2)
				delete(fm, f.op1)
				//log.Printf("knowing %s=%d and previous operand %d, we solved '%s: %s %c %s' and stored '%s: %d'", yeller, num, f.knownOperand, f.yeller, f.op1, f.operator, f.op2, f.yeller, nm[f.yeller])
			} else {
				// not yet, but we will next time!
				f.knowOneOperand = true
				f.knownOperand = num
				// once we've used it, don't need it any longer
				delete(nm, yeller)
			}

			// are we done?
			mayberoot, ok := nm["root"]
			if ok {
				root = mayberoot
				break keepsearching
			}
		}
	}
	return root

}

func run1(inputText string) int {
	// first load all the numbers and formulae
	fm, nm, _ := parseMonkeys(parsers.SplitByLines(inputText))
	return solve(fm, nm)

}

func run2(inputText string) int {

	fm, nm, froot := parseMonkeys(parsers.SplitByLines(inputText))
	delete(nm, "humn") // part 2 rules- I'm the humn
	froot.operator = '='
	//nm["humn"] = 301 // testing
	_ = fm

	return 0
}
