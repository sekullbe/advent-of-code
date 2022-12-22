package main

import (
	_ "embed"
	"fmt"
	"github.com/sekullbe/advent/parsers"
	"github.com/sekullbe/advent/tools"
	"log"
	"math"
)

//go:embed input.txt
var inputText string

func main() {
	fmt.Printf("Magic number: %d\n", run1(inputText))
	fmt.Println("-------------")
	//fmt.Printf("Magic number: %d\n", run2(inputText))
	fmt.Println(run2_eval(inputText))
}

type formula struct {
	yeller, op1, op2 string
	operator         rune
	knowOneOperand   bool
	knownOperand     int
	solved           bool
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
		//if num == f.knownOperand {
		return num
		//	} else {
		//	log.Printf("tried root, %d != %d", num, f.knownOperand)
		//		return -math.MaxInt
		//	}
	}
	return res // in part 1, root always succeeds
}

func solve(fm formulas, nm numbers) int {
	root := 0
keepsearching:
	for {
		for yeller, num := range nm {
			// It should be an operand in exactly one formula
			f := fm[yeller]
			// Catch the case where we've solved and deleted a formula
			if f == nil || f.solved {
				continue
			}
			// can we solve that formula now?
			if f.knowOneOperand {
				// we can, hooray!
				if f.operator == '=' {
					log.Printf("Root comparing %s=%s. Know %s=%d and %s=%d.", f.op1, f.op2, f.op1, nm[f.op1], f.op2, nm[f.op2])
					if num != f.knownOperand {
						return -math.MaxInt
					} else {
						return num
					}
				}
				res := operate(f, yeller, num)
				nm[f.yeller] = res
				// the formula is solved, we don't need it any longer
				delete(fm, f.op2)
				delete(fm, f.op1)
				//fmt.Printf("knowing %s=%d and other operand %d, we solved '%s: %s %c %s' and stored '%s: %d'\n", yeller, num, f.knownOperand, f.yeller, f.op1, f.operator, f.op2, f.yeller, nm[f.yeller])
				//fmt.Printf("%s: %s %c %s = %d\n", f.yeller, f.op1, f.operator, f.op2, nm[f.yeller])
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
	sol := solve(fm, nm)
	return sol

}

// Couldn't get this to work so I revisited it as evaluation
func run2(inputText string) int {

	// right, so how do we guess?... plugging in the number from the eval solution works
	// but it's easy to guess when you know the answer
	for try := 3_882_224_466_000; try <= 3_882_224_466_200; try++ {
		//for try := 1_000_000_000_000; try <= 5_000_000_000_000; try++ {
		if try%1_000_000_000 == 0 {
			fmt.Print(".")
		}
		// reload, because we muck with the maps
		fm, nm, froot := parseMonkeys(parsers.SplitByLines(inputText))
		delete(nm, "humn") // part 2 rules- I'm the humn
		froot.operator = '='

		nm["humn"] = try
		//log.Printf("Guessing %d:", try)
		res := solve(fm, nm)
		// what I'm seeing here is that approaching correct try the numbers go up
		// 301 and 302 both work because of integer division
		if res != -math.MaxInt {
			return try
		}
	}
	return 0
}

// --- redoing part2 as evaluation
type monkey struct {
	val, left, right, op string
}

func run2_eval(inputText string) int {

	monkeys := parseMonkeys2(parsers.SplitByLines(inputText))
	evalme := esolve(monkeys, "root")
	fmt.Println("Part1: Evaluate this in python or whatever:")
	fmt.Println(evalme)
	fmt.Println("---")
	fmt.Println("Part2: Evaluate this in python and print x.real/x.complex:")
	// use complex notation, thanks u/Anton31Kah
	monkeys["humn"] = monkey{val: "-1j"} // - makes the final sign come out correct
	// since you're looking for equals, subtract both sides of the root to move toward 0
	r := monkeys["root"]
	r.op = "-"
	monkeys["root"] = r
	evalme = esolve(monkeys, "root")
	fmt.Println(evalme)
	return 0
}

func esolve(monkeys map[string]monkey, mname string) string {
	targetMonkey := monkeys[mname]
	if targetMonkey.val != "" {
		return targetMonkey.val
	}
	return "(" + esolve(monkeys, targetMonkey.left) + targetMonkey.op + esolve(monkeys, targetMonkey.right) + ")"
}

func parseMonkeys2(lines []string) map[string]monkey {
	monkeys := make(map[string]monkey)
	for _, line := range lines {
		var my, m1, m2 string
		var op rune
		var mnum int
		n, _ := fmt.Sscanf(line, "%4s: %4s %c %4s", &my, &m1, &op, &m2)
		var v monkey
		if n != 4 {
			n, _ = fmt.Sscanf(line, "%4s: %d", &my, &mnum)
			v = monkey{val: fmt.Sprint(mnum)}
		} else {
			v = monkey{left: m1, op: string(op), right: m2}
		}
		monkeys[my] = v
	}
	return monkeys
}
