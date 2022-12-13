package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"github.com/sekullbe/advent/parsers"
	"log"
	"reflect"
)

//go:embed input.txt
var inputText string

func main() {
	fmt.Printf("Magic number: %d\n", run1(inputText))
	fmt.Println("-------------")
	fmt.Printf("Magic number: %d\n", run2(inputText))
}

// make sure to handle empty packets

func run1(inputText string) int {
	packetPairs := parsers.SplitByEmptyNewline(inputText)
	sum := 0
	for i, packetPair := range packetPairs {
		packets := parsePackets(parsers.SplitByLines(packetPair))
		correctOrder := leftIsLess(packets[0], packets[1])
		if correctOrder {
			sum += i + 1 // +1 because the index is 0 based
		}
	}
	return sum
}

func parsePackets(packetLines []string) [][]any {
	var packets [][]any
	for _, line := range packetLines {
		var packet []any
		err := json.Unmarshal([]byte(line), &packet)
		if err != nil {
			log.Panicf("unmarshal err: %v", err)
		}
		packets = append(packets, packet)
	}
	return packets
}

func leftIsLess(left []any, right []any) bool {
	return compareLeftRight(left, right) < 1
}

// remember we short-circuit out of this as soon as we get any non-equal lineup
// inputs can be either float64 or any nesting of lists of float64
// whoops, tried this pure boolean but no go, because it needs to allow an 'equal' when recursing. so use -1 < 0 = 1 >
func compareLeftRight(left []any, right []any) int {
	for i := 0; i < len(left) && i < len(right); i++ {
		l := left[i]
		r := right[i]

		lnum := reflect.TypeOf(l).Name() == "float64"
		rnum := reflect.TypeOf(r).Name() == "float64"

		// if both are numbers, force conversion and compare the numbers
		if lnum && rnum {
			// compiler thinks l and r are 'any' but I know better; so force it w/ type assertion (not conversion)
			lf := l.(float64)
			rf := r.(float64)
			if lf < rf {
				return -1
			}
			if lf > rf {
				return 1
			}
			// equal; we haven't seen enough to decide yet, so keep going
		} else {
			var lSubPacket []any // wish I could constrain this type but I don't think that's possible
			var rSubPacket []any
			// if they're not both lists, make the one that isn't into a list (slice)
			if lnum {
				lSubPacket = []any{l} // wrap it in listiness
			} else {
				lSubPacket = l.([]any) // it's a 'any' but we know it's a []any so force it
			}
			if rnum { // and the same for right
				rSubPacket = []any{r} // wrap it in listiness
			} else {
				rSubPacket = r.([]any) // it's a 'any' but we know it's a []any so force it
			}

			// now they're lists; either they were lists or we made them lists, so... forward!
			comparison := compareLeftRight(lSubPacket, rSubPacket)
			// if they're equal we don't know anything yet so forge on
			if comparison != 0 {
				return comparison
			}
		}
	}
	// if we're out of matches between the packets, see which ran out first
	if len(left) < len(right) {
		return -1
	}
	if len(left) > len(right) {
		return 1
	}
	// we don't know enough yet, hope the recursion gets something
	return 0
}

func run2(inputText string) int {

	return 0
}
