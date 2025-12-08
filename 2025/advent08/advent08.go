package main

import (
	"cmp"
	_ "embed"
	"fmt"
	"maps"
	"slices"
	"time"

	"github.com/sekullbe/advent/geometry"
	"github.com/sekullbe/advent/parsers"
	"github.com/sekullbe/advent/tools"
)

//go:embed input.txt
var inputText string

type connection struct {
	a    geometry.Point3
	b    geometry.Point3
	dist float64
}

func main() {
	fmt.Printf("Magic number: %d\n", run1(inputText, 1000))
	fmt.Println("-------------")
	fmt.Printf("Magic number: %d\n", run2(inputText))
}

func run1(input string, numCxns int) int {
	defer tools.Track(time.Now(), "Part 1 Time")

	lines := parsers.SplitByLines(input)
	boxes := []geometry.Point3{}
	for _, line := range lines {
		coords := parsers.StringsWithCommasToIntSlice(line)
		boxes = append(boxes, geometry.Point3{coords[0], coords[1], coords[2]})
	}
	cxns := []connection{}
	boxToCircuit := make(map[geometry.Point3]int)
	// instead of using ranges, use indices, because once we've calculated box[5]-box[0] we don't want to calculate 0-5 agai
	nboxes := len(boxes)
	for i, b1 := range boxes {
		boxToCircuit[b1] = 0 - i
		for n := i + 1; n < nboxes; n++ {
			b2 := boxes[n]
			c := connection{b1, b2, b1.Dist(b2)}
			cxns = append(cxns, c)
		}
	}
	// sort by value to find the shortest N numCxns
	slices.SortFunc(cxns, func(a, b connection) int {
		return cmp.Compare(a.dist, b.dist)
	})

	// map from coords to an identifier for a circuit
	// each one already has an identifier which is negative to show it is unconnected
	nextCircuitId := 1
	circuits := make(map[int][]geometry.Point3)
	// pull off the top N and make circuits
	for i := 0; i < numCxns; i++ {
		c := cxns[i]
		// if both boxes are already in circuits
		if boxToCircuit[c.a] > 0 && boxToCircuit[c.b] > 0 {
			cAid := boxToCircuit[c.a] // this will be the new ID
			cBid := boxToCircuit[c.b]
			if cAid == cBid {
				continue // they're already both in the same circuit, no-op
			}
			// set the circuit ID for B and all in its circuit to 'cAid'
			for _, bc := range circuits[boxToCircuit[c.b]] {
				boxToCircuit[bc] = cAid
				circuits[cAid] = append(circuits[cAid], bc)
			}
			delete(circuits, cBid) // the old circuit no longer exists and all of its boxes are now in another circuit
		} else if boxToCircuit[c.a] > 0 {
			// add B to A's circuit
			boxToCircuit[c.b] = boxToCircuit[c.a]
			circuits[boxToCircuit[c.a]] = append(circuits[boxToCircuit[c.a]], c.b)
		} else if boxToCircuit[c.b] > 0 {
			// add A to B's circuit
			boxToCircuit[c.a] = boxToCircuit[c.b]
			circuits[boxToCircuit[c.b]] = append(circuits[boxToCircuit[c.b]], c.a)
		} else {
			// neither is in a circuit
			boxToCircuit[c.a] = nextCircuitId
			boxToCircuit[c.b] = nextCircuitId
			circuits[nextCircuitId] = []geometry.Point3{c.a, c.b}
			nextCircuitId++
		}
		//fmt.Println(circuits)
	}
	// get the 3 largest circuits
	circuitSlice := slices.Collect(maps.Values(circuits))
	slices.SortFunc(circuitSlice, func(a, b []geometry.Point3) int {
		return cmp.Compare(len(a), len(b))
	})

	return len(circuitSlice[len(circuitSlice)-1]) * len(circuitSlice[len(circuitSlice)-2]) * len(circuitSlice[len(circuitSlice)-3])
}

func run2(input string) int {
	defer tools.Track(time.Now(), "Part 2 Time")

	// feh I should have pulled some of the part 1 code out into functions
	lines := parsers.SplitByLines(input)
	boxes := []geometry.Point3{}
	for _, line := range lines {
		coords := parsers.StringsWithCommasToIntSlice(line)
		boxes = append(boxes, geometry.Point3{coords[0], coords[1], coords[2]})
	}
	cxns := []connection{}
	boxToCircuit := make(map[geometry.Point3]int)
	// instead of using ranges, use indices, because once we've calculated box[5]-box[0] we don't want to calculate 0-5 agai
	nboxes := len(boxes)
	for i, b1 := range boxes {
		boxToCircuit[b1] = 0 - i
		for n := i + 1; n < nboxes; n++ {
			b2 := boxes[n]
			c := connection{b1, b2, b1.Dist(b2)}
			cxns = append(cxns, c)
		}
	}
	// sort by value to find the shortest N numCxns
	slices.SortFunc(cxns, func(a, b connection) int {
		return cmp.Compare(a.dist, b.dist)
	})

	// map from coords to an identifier for a circuit
	// each one already has an identifier which is negative to show it is unconnected
	nextCircuitId := 1
	circuits := make(map[int][]geometry.Point3)
	loners := nboxes
	for i := 0; i < 10000; i++ { // we'll break out before hitting that point
		c := cxns[i]
		// if both boxes are already in circuits
		if boxToCircuit[c.a] > 0 && boxToCircuit[c.b] > 0 {
			cAid := boxToCircuit[c.a] // this will be the new ID
			cBid := boxToCircuit[c.b]
			if cAid == cBid {
				continue // they're already both in the same circuit, no-op
			}
			// set the circuit ID for B and all in its circuit to 'cAid'
			for _, bc := range circuits[boxToCircuit[c.b]] {
				boxToCircuit[bc] = cAid
				circuits[cAid] = append(circuits[cAid], bc)
			}
			delete(circuits, cBid) // the old circuit no longer exists and all of its boxes are now in another circuit
		} else if boxToCircuit[c.a] > 0 {
			// add B to A's circuit
			boxToCircuit[c.b] = boxToCircuit[c.a]
			circuits[boxToCircuit[c.a]] = append(circuits[boxToCircuit[c.a]], c.b)
			loners--
		} else if boxToCircuit[c.b] > 0 {
			// add A to B's circuit
			boxToCircuit[c.a] = boxToCircuit[c.b]
			circuits[boxToCircuit[c.b]] = append(circuits[boxToCircuit[c.b]], c.a)
			loners--
		} else {
			// neither is in a circuit
			boxToCircuit[c.a] = nextCircuitId
			boxToCircuit[c.b] = nextCircuitId
			circuits[nextCircuitId] = []geometry.Point3{c.a, c.b}
			nextCircuitId++
			loners -= 2
		}
		//fmt.Printf("%d circuits, %d loners: %v\n", len(circuits), loners, circuits)
		// If we now have only one circuit we are done
		if len(circuits) == 1 && loners == 0 { // delay a bit before checking
			fmt.Printf("done! connected %v and %v\n", c.a, c.b)
			return c.a.X * c.b.X
		}

	}
	return -1
}
