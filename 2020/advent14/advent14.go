package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

//go:embed input.txt
var inputText string

type memory map[int]int

func main() {
	run1(inputText)
	run2(inputText)
}

func run1(inputText string) int {
	memory := make(memory)
	mask := createEmptyBitmask()

	for _,line := range strings.Split(inputText,"\n") {
		if len(line) < 4 {
			continue
		}
		switch line[0:4] {
		case "mask":
			mask.update(line[7:])
			// fmt.Printf("Updating mask to %s\n", line[7:])
		case "mem[":
			addr,val := parseAddress(line)
			val = mask.applyBitmaskToValue(val)
			memory[addr] = val
			// fmt.Printf("mem[%d]=%d\n", addr, val)
		default:
			fmt.Printf("instruction parse failure: '%s' ", line)
		}
	}
	sum := 0
	for _,v := range memory {
		sum += v
	}
	fmt.Printf("Sum of memory = %d\n", sum)
	return sum
}

func parseAddress(line string) (addr int, val int) {
	re := regexp.MustCompile("^mem\\[(\\d+)\\] = (\\d+)$")
	matches := re.FindStringSubmatch(line)
	addr, _ = strconv.Atoi(matches[1])
	val, _ = strconv.Atoi(matches[2])
	return addr,val
}

func run2(inputText string) (sum int) {
	memory := make(memory)
	mask := createEmptyBitmask()

	for _,line := range strings.Split(inputText,"\n") {
		if len(line) < 4 {
			continue
		}
		switch line[0:4] {
		case "mask":
			mask.update(line[7:])
			// fmt.Printf("Updating mask to %s\n", line[7:])
		case "mem[":
			addr,val := parseAddress(line)
			addrsToWrite := mask.applyToAddr(addr)
			for _, a := range addrsToWrite {
				memory[a] = val
			}

			// fmt.Printf("mem[%d]=%d\n", addr, val)
		default:
			fmt.Printf("instruction parse failure: '%s' ", line)
		}
	}

	for _,v := range memory {
		sum += v
	}
	fmt.Printf("Sum of memory = %d\n", sum)
	return sum
}

const (
	ZERO int = iota
	ONE
	X
)

type bitmask map[int]int

func createEmptyBitmask() bitmask {
	b := make(bitmask)
	for i := 0; i < 36; i++ {
		b[i] = X
	}
	return b
}

func createBitmask(mask string) bitmask {
	b := make(bitmask)
	b.update(mask)
	return b
}


func (mask bitmask) applyBitmaskToValue(value int) int {
	for i := 0; i <36 ; i++ {
		switch mask[i] {
		case ZERO:
			// set the bit to zero
			// OR with 1 then XOR with 1
			value |= 1 << i
			value ^= 1 << i
		case ONE:
			// set the bit to one
			value |= 1 << i
		}
	}
	return value
}

func (mask bitmask) applyToAddr(addr int) []int {
	var floaters []int
	for i := 0; i < 36; i++ {
		switch mask[i] {
		case ZERO:
			// do nothing
		case ONE:
			// set the bit to one
			addr |= 1 << i
		case X:
			// make a list of all floating indices
			floaters = append(floaters, i)
		}
	}
	// put the "base" in- it's ok that it won't be unique
	addrs := []int{addr}
	for _, index := range floaters {
		for _, perm := range addrs {
			// add the 1 and 0
			with1 := perm | 1<<index  // OR with a shifted 1 -> 1
			with0 := with1 ^ 1<<index // with1 XOR'ed with a shifted 1 -> 0
			addrs = append(addrs, with1, with0)
		}
	}

	return addrs
}

func (mask bitmask) update(newMask string) {
	for i, maskElt := range newMask {
		switch maskElt {
		case '0':
			mask[35-i] = ZERO
		case '1':
			mask[35-i] = ONE
		case 'X':
			mask[35-i] = X
		}
	}
}


