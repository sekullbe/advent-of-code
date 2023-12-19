package main

import (
	_ "embed"
	"fmt"
	"github.com/sekullbe/advent/parsers"
	"github.com/sekullbe/advent/tools"
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

	return 0
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
