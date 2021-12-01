package advent14

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

func Run() {
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
			fmt.Printf("Updating mask to %s\n", line[7:])
		case "mem[":
			addr,val := parseAddress(line)
			val = mask.applyBitmaskToValue(val)
			memory[addr] = val
			fmt.Printf("mem[%d]=%d\n", addr, val)
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

func run2(inputText string) {
	memory := make(memory)
	mask := createEmptyBitmask()

	for _,line := range strings.Split(inputText,"\n") {
		if len(line) < 4 {
			continue
		}
		switch line[0:4] {
		case "mask":
			mask.update(line[7:])
			fmt.Printf("Updating mask to %s\n", line[7:])
		case "mem[":
			addr,val := parseAddress(line)
			addr[] = mask.applyBitmaskToAddr(addr)
			memory[addr] = val
			fmt.Printf("mem[%d]=%d\n", addr, val)
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

func (mask bitmask) applyBitmaskToAddr(addr int) []int {
	var addrs = []int{}
	for i := 0; i < 36; i++ {
		addr := 0
		switch mask[i] {
		case ZERO:
			// set the bit to zero
			// OR with 1 to set it high then XOR with 1 to set it low
			addr |= 1 << i
			addr ^= 1 << i
		case ONE:
			// set the bit to one
			addr |= 1 << i
		case X:
			// make a list of all floating indices
		}
	}
	// if 0 floating incides return addr only
	// for all floaters, permute the 2^N possibilities
	//perms := []int{addr}


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


