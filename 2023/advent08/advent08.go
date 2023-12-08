package main

import (
	_ "embed"
	"fmt"
	"github.com/sekullbe/advent/parsers"
	"github.com/sekullbe/advent/tools"
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
	for {
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
	// observe that once you hit xxZ, you start to repeat
	// i.e. xxA -> 111 -> 222 -> 33Z --> 111 again
	// the time from xxA to 33Z and from 33Z back to 33Z are the same - 3 in this case
	// so the solution is to find the cycle time for each starting xxX position
	// and compute the LCM of all of them
	directions, net := parseInput(input)
	cycles := []int{}
	locations := []string{}
	for loc, _ := range net {
		if loc[2] == 'A' {
			locations = append(locations, loc)
			cycles = append(cycles, 0)
		}
	}

	// effectively, run 'run1' for each of the locations
	for i, location := range locations {
		steps := 0
		for {
			nextStep := nextStep(directions, steps)
			steps++
			nextNode := net.lookupNextNode(nextStep, location)
			location = nextNode.name
			if location[2] == 'Z' {
				break
			}
		}
		cycles[i] = steps
	}
	// and compute LCM if necessary
	if len(cycles) < 2 {
		return cycles[0]
	}

	return tools.LCM_slice(cycles...)

}

// this will find the solution but takes forever. do better.
func run2_slow_solution(input string) int {
	// six nodes begin with A, and six end with Z
	// get all the A's and the direction
	// for each one find its next step
	// repeat
	directions, net := parseInput(input)
	steps := 0
	locations := []string{}
	for loc, _ := range net {
		if loc[2] == 'A' {
			locations = append(locations, loc)
		}
	}
	fmt.Println(locations)
	for {
		nextStep := nextStep(directions, steps)
		steps++
		allAtZ := true
		for i, location := range locations {
			node := net.lookupNextNode(nextStep, location)
			locations[i] = node.name
			if node.name[2] != 'Z' {
				allAtZ = false
			}
		}
		fmt.Println(locations)
		if steps%10000000 == 0 {
			fmt.Print(".")
		}
		if allAtZ {
			break
		}
	}

	return steps

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
