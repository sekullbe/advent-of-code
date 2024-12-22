package main

import (
	_ "embed"
	"fmt"
	"github.com/sekullbe/advent/parsers"
	"github.com/sekullbe/advent/tools"
	"log"
	"time"
)

//go:embed input.txt
var inputText string

func main() {
	fmt.Printf("Magic number: %d\n", run1(inputText))
	fmt.Println("-------------")
	fmt.Printf("Magic number: %d\n", run2(inputText))
}

func run1(input string) int {
	defer tools.Track(time.Now(), "Part 1 Time")
	secretNumbers := parsers.StringsToIntSlice(input)
	newSecretNumbers := make([]int, len(secretNumbers))
	for i := 0; i < len(secretNumbers); i++ {
		newSecretNumbers[i] = iterateNextNum(secretNumbers[i], 2000)
	}

	return tools.SumSlice(newSecretNumbers)
}

/*
for each monkey
compute its list of numbers -> prices -> deltas
 store the deltas and the price each would give in a map
 store _that_ in a slice of monkeys

for all possible sequences -9->+9 x 4
sum the bananas of each monkey for that sequence & find the max

*/

func run2(input string) int {
	defer tools.Track(time.Now(), "Part 2 Time")
	secretNumbers := parsers.StringsToIntSlice(input)

	allBuyers := []buyer{}

	// very weird behavior here-
	// before I fixed the bug in processOneBuyer() where I wasn't checking for the _first_ matching
	// sequence, I also had this typoed as for sn := range secretnumbers
	// and the unit tests failed, as expected, but the real data _worked_.
	// That stopped happening when I fixed the underlying bug.
	for _, sn := range secretNumbers {
		allBuyers = append(allBuyers, processOneBuyer(sn, 2000))
	}

	var bestSeq sequence
	bestTotalSales := -999
	// BRUTE FORCE!
	for d1 := -9; d1 <= 9; d1++ {
		for d2 := -9; d2 <= 9; d2++ {
			for d3 := -9; d3 <= 9; d3++ {
				for d4 := -9; d4 <= 9; d4++ {
					seq := sequence{d1, d2, d3, d4}
					sumPrices := 0
					for _, buyer := range allBuyers {
						if price, ok := buyer[seq]; ok {
							sumPrices += price
						}
					}
					if sumPrices > bestTotalSales {
						fmt.Printf("new best price: %d %v\n", sumPrices, seq)
						bestTotalSales = sumPrices
						bestSeq = seq
					}
				}
			}
		}
	}
	// 1908 too high, 1801 too low
	log.Println(bestSeq)
	return bestTotalSales
}

func iterateNextNum(num, steps int) int {
	for i := 0; i < steps; i++ {
		num = nextNum(num)
	}
	return num
}

func nextNum(n int) int {
	n = mix(n*64, n) // this is 5 shifts left
	n = prune(n)
	n = mix(n/32, n) // 3 shifts right
	n = prune(n)
	n = mix(n*2048, n) // 11 shifts left
	n = prune(n)
	return n
}

func prune(n int) int {
	return n % 16777216
}

func mix(i int, n int) int {
	return i ^ n
}

func onesDigit(n int) int {
	return n % 10
}

type sequence [4]int
type buyer map[sequence]int

func processOneBuyer(sn int, steps int) buyer {
	delta := 0
	diffs := sequence{0, 0, 0, 0}
	sequencesToPrice := make(map[sequence]int)
	price := onesDigit(sn)
	for i := 0; i < steps; i++ {
		sn = nextNum(sn)
		newPrice := onesDigit(sn)
		delta = onesDigit(sn) - price
		diffs.push(delta)
		price = newPrice
		if _, ok := sequencesToPrice[diffs]; !ok && i >= 3 {
			sequencesToPrice[diffs] = price
		}
	}
	return sequencesToPrice
}

func (s *sequence) push(n int) {
	s[0] = s[1]
	s[1] = s[2]
	s[2] = s[3]
	s[3] = n
}
