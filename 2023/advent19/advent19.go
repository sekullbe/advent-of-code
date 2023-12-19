package main

import (
	_ "embed"
	"fmt"
	"github.com/sekullbe/advent/parsers"
	"github.com/sekullbe/advent/tools"
	"maps"
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

type rule struct {
	rating   string
	operator string
	operand  int
	next     string
}

type workflow struct {
	name  string
	rules []rule
}

type part map[string]int
type workflows map[string]workflow

func run1(input string) int {

	workflows, parts := parseInput(input)
	accepted := []part{}
	for _, p := range parts {
		next := "in"
		for next != "A" && next != "R" {
			w := workflows[next]
			next = processWorkflowStep(w, p)
			if next == "A" {
				accepted = append(accepted, p)
			}
		}
	}
	// sum all the ratings of all accepted parts
	sum := 0
	for _, p := range accepted {
		for _, v := range p {
			sum += v
		}
	}

	return sum
}

func run2(input string) int {

	// can I look through every rule and find all of the states that lead to A?
	/*
				s < 1351 && a >= 2006 && m>2090
			    hdj && m <=838 && a<=1716
			    hdj && m > 838
		...
		reverse that, starting at "in" continue with recursive DFS and keep all the ranges that got you there
		at each step, keep "true" and remove "false" ranges from the set starting at 1..4000
		when you hit A, count the ranges (low*high for each rating x-m-a-s)
		when you hit R, return 0; all those ranges were invalid
	*/

	workflows, _ := parseInput(input)
	ranges := map[string]rng{"x": {1, 4000}, "m": {1, 4000}, "a": {1, 4000}, "s": {1, 4000}}
	total := countValidRanges(ranges, "in", workflows)

	return total
}

type rng struct {
	low, high int
}

func countValidRanges(ranges map[string]rng, wfn string, wfs workflows) int {
	// reimplementation of https://github.com/derailed-dash/Advent-of-Code/blob/master/src/AoC_2023/Dazbo's_Advent_of_Code_2023.ipynb algorithm
	if wfn == "R" {
		return 0
	}
	if wfn == "A" {
		// multiply accepted ranges - ie there are x*m*a*s ways to hit this accepted state
		prod := 1
		for _, r := range ranges {
			prod *= r.high - r.low + 1
		}
		return prod
	}

	w := wfs[wfn]
	total := 0
	for _, r := range w.rules {

		// are we at a terminal rule?
		if r.rating == "" {
			total += countValidRanges(ranges, r.next, wfs)
		} else {
			curRng := ranges[r.rating]
			// find the ranges for the rule's rating where the rule is true or false
			var trueRng, falseRng rng
			// ok I'll grant this switch could be two lines of python :)
			switch r.operator {
			case "<":
				trueRng.low = curRng.low
				trueRng.high = r.operand - 1
				falseRng.low = r.operand
				falseRng.high = curRng.high
			case ">":
				trueRng.low = r.operand + 1
				trueRng.high = curRng.high
				falseRng.low = curRng.low
				falseRng.high = r.operand
			default:
				panic("bad operator")
			}
			// if the ranges are impossible, don't go to the next workflow or next rule in this workflow
			if trueRng.low <= trueRng.high {

				rangeCopy := maps.Clone(ranges)
				rangeCopy[r.rating] = trueRng
				total += countValidRanges(rangeCopy, r.next, wfs)
			}
			if falseRng.low <= falseRng.high {
				ranges[r.rating] = falseRng
			} else {
				// can't continue this workflow
				break
			}
		}
	}
	return total
}

// do it the long way to see if _enough_ brute force works
// on the sample data, this takes about 4s per "m" so 4*4000*4000 = a hecking long time
func run2stupid(input string) int {
	workflows, _ := parseInput(input)
	accepted := 0
	for x := 1; x <= 4000; x++ {
		fmt.Print("x")
		for m := 1; m <= 4000; m++ {
			fmt.Print("m")
			for a := 1; a <= 4000; a++ {
				for s := 1; s <= 4000; s++ {
					p := part{"x": x, "m": m, "a": a, "s": s}
					next := "in"
					for next != "A" && next != "R" {
						w := workflows[next]
						next = processWorkflowStep(w, p)
						if next == "A" {
							accepted++
						}
					}
				}
			}
		}
	}

	return accepted
}

func parseInput(input string) (map[string]workflow, []part) {
	segments := parsers.SplitByEmptyNewlineToSlices(input)
	return parseWorkflows(segments[0]), parseParts(segments[1])

}
func parseWorkflows(lines []string) (workflows map[string]workflow) {
	workflows = make(map[string]workflow)
	// px{a<2006:qkq,m>2090:A,rfg} // operators are always lt or gt
	for _, line := range lines {
		w := parseSingleWorkflow(line)
		workflows[w.name] = w
	}
	return workflows
}

func parseSingleWorkflow(line string) workflow {
	re := regexp.MustCompile(`(\w+){(.+?)}`)
	w := workflow{rules: []rule{}}
	// pull off the prefix and the {}
	matches := re.FindStringSubmatch(line)
	w.name = matches[1]
	wfRules := strings.Split(matches[2], ",")
	for _, wfRule := range wfRules {
		rule := parseSingleRule(wfRule)
		w.rules = append(w.rules, rule)
	}
	return w
}

func parseSingleRule(rs string) rule {
	reWorkflow := regexp.MustCompile(`(\w)([<>])(\d+):(\w+)`)
	rule := rule{}
	// these are `r<1111:target` or `target` where target is a wf name or R|A final state
	matches := reWorkflow.FindStringSubmatch(rs)
	// if we fail the regexp, the rule must just be a default
	if len(matches) == 0 {
		rule.next = rs
	} else {
		rule.rating = matches[1]
		rule.operator = matches[2]
		rule.operand = tools.Atoi(matches[3])
		rule.next = matches[4]
	}
	return rule
}

func parseParts(lines []string) []part {
	parts := []part{}
	for _, line := range lines {
		parts = append(parts, parseSinglePart(line))
	}
	return parts
}

func parseSinglePart(line string) part {
	// they're always in x,m,a,s order, so don't bother parsing the part name
	re := regexp.MustCompile(`{x=(\d+),m=(\d+),a=(\d+),s=(\d+)}`)
	matches := re.FindStringSubmatch(line)
	part := make(part)
	part["x"] = tools.Atoi(matches[1])
	part["m"] = tools.Atoi(matches[2])
	part["a"] = tools.Atoi(matches[3])
	part["s"] = tools.Atoi(matches[4])
	return part
}

// returns the next workflow step or a halt state (A|R)
func processWorkflowStep(w workflow, p part) string {

	for _, r := range w.rules {
		//test: p[r.rating] r.operator r.operand and if true return r.next
		switch r.operator {
		case "<":
			if p[r.rating] < r.operand {
				return r.next
			}
		case ">":
			if p[r.rating] > r.operand {
				return r.next
			}
		case "":
			return r.next
		default:
			panic("unknown operator")
		}
	}
	panic("rule did not resolve next state")
}
