package main

import (
	_ "embed"
	"fmt"
	"github.com/sekullbe/advent/parsers"
	"strconv"
)

//go:embed input.txt
var inputText string

type farm []*octopus

func main() {
	fmt.Printf("Magic number: %d\n", run1(inputText, 100))
	fmt.Println("-------------")
	fmt.Printf("Magic number: %d\n", run2(inputText, 2000))
}

func run1(inputText string, steps int) int {

	totalFlashes := 0
	farm := parseFarm(inputText)
	// set up the neighbor links
	for _, o := range farm {
		o.attachNeighbors(farm)
	}

	for step := 1; step <= steps; step++ {
		totalFlashes += farm.step()
	}
	return totalFlashes
}

func run2(inputText string, steps int) int {
	farm := parseFarm(inputText)
	// set up the neighbor links
	for _, o := range farm {
		o.attachNeighbors(farm)
	}
	for step := 1; step <= steps; step++ {
		stepFlashes := farm.step()
		//log.Printf("step %d: %d", step, stepFlashes)
		if stepFlashes == len(farm) {
			return step
		}
	}

	return 0
}

func (f farm) energize() {
	for _, o := range f {
		o.energize()
	}
}

func (f farm) reset() {
	for _, o := range f {
		o.reset()
	}
}

func (f farm) step() int {
	f.energize()
	flashes := 0
	for _, o := range f {
		flashes += o.flash()
	}
	f.reset()
	return flashes

}

func parseFarm(inputText string) farm {
	farm := farm{}
	lines := parsers.SplitByLines(inputText)
	for row, line := range lines {
		if len(line) == 0 {
			continue
		}
		for col, es := range line {
			e, err := strconv.Atoi(string(es))
			if err != nil {
				panic("bad input row: " + line)
			}
			o := newOctopus(e, row, col)
			farm = append(farm, &o)
		}
	}

	return farm
}
