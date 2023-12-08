package main

import (
	_ "embed"
	"fmt"
	"github.com/sekullbe/advent/parsers"
	"regexp"
)

//go:embed input.txt
var inputText string

func main() {
	fmt.Printf("Magic number: %d\n", run1(inputText))
	fmt.Println("-------------")
	fmt.Printf("Magic number: %d\n", run2(inputText))
}

func run1(input string) int {
	directions, net := parseInput(input)
	steps := 0
	location := "AAA"
	for true {
		nextStep := nextStep(directions, steps)
		steps++
		nextNode := net.lookupNextNode(nextStep, location)
		location = nextNode.name
		if location == "ZZZ" {
			break
		}
	}
	return steps
}

func run2(input string) int {

	return 0
}

// do lookups directly with pointers, or just store the keys and do lookups manually?
// doesn't much matter - strings mean it can be parsed in one pass but there's
// one more instruction to look up each node.
// easier to write tests with strings though.
type node struct {
	name  string
	left  string
	right string
}
type network map[string]node

func parseInput(input string) (directions string, net network) {
	inputParts := parsers.SplitByEmptyNewlineToSlices(input)
	directions = inputParts[0][0]
	nodeInput := inputParts[1]
	net = make(map[string]node)
	for _, s := range nodeInput {
		name, newNode := parseOneNode(s)
		net[name] = newNode
	}
	return directions, net
}

func parseOneNode(input string) (string, node) {
	re := regexp.MustCompile(`(...) = \((...), (...)\)`)
	matches := re.FindStringSubmatch(input)
	return matches[1], node{name: matches[1], left: matches[2], right: matches[3]}
}

func (net network) lookupNextNode(direction rune, location string) node {
	node, ok := net[location]
	if !ok {
		panic("looking up nonexistent node: " + location)
	}
	switch direction {
	case 'L':
		return net[node.left]
	case 'R':
		return net[node.right]
	default:
		panic("bad direction: " + string(direction))
	}
}

func nextStep(instructions string, index int) rune {
	return rune(instructions[index%len(instructions)])
}
