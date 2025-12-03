package main

import (
	_ "embed"
	"fmt"
	"time"

	"github.com/samber/lo"
	"github.com/sekullbe/advent/parsers"
	"github.com/sekullbe/advent/tools"
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
	banks := parseBanks(input)
	sum := 0
	for _, bank := range banks {
		sum += maxOutputFromBank2(bank, 2)

	}

	return sum
}

func run2(input string) int {
	defer tools.Track(time.Now(), "Part 2 Time")

	banks := parseBanks(input)
	sum := 0
	for _, bank := range banks {
		sum += maxOutputFromBank2(bank, 12)

	}

	return sum
}

func parseBanks(input string) [][]int {

	banks := [][]int{}
	lines := parsers.SplitByLines(input)
	for _, line := range lines {
		bankRunes := []rune(line)
		bank := lo.Map(bankRunes, func(r rune, i int) int {
			return int(r - '0')
		})
		banks = append(banks, bank)
	}
	return banks
}

// part 1 version
func maxOutputFromBank(bank []int) int {
	// find the highest number that is not at the end- the leftmost appearance of that is our first digit
	// then the highest number _after_ that
	firstDigit := 0
	firstDigitLoc := 0
	secondDigit := 0
	for _, d := range []int{9, 8, 7, 6, 5, 4, 3, 2, 1} {
		found, loc := tools.FirstIntLoc(bank[:len(bank)-1], d)
		if found {
			firstDigit = d
			firstDigitLoc = loc
			break
		}
	}
	for _, d := range []int{9, 8, 7, 6, 5, 4, 3, 2, 1} {
		found, _ := tools.FirstIntLoc(bank[firstDigitLoc+1:], d)
		if found {
			secondDigit = d
			break
		}
	}
	return firstDigit*10 + secondDigit
}

// part 2 version of max output
// put that in a loop from 11..0
// get the digit times 10^x and add that to the sum
// this can also solve part1 if the max power is parameterized
func maxOutputFromBank2(bank []int, howMany int) int {
	sum := 0
	digitLoc := 0
	for pow := howMany - 1; pow >= 0; pow-- {
		for i := 9; i >= 0; i-- {
			found, loc := tools.FirstIntLoc(bank[digitLoc:len(bank)-pow], i)
			if found {
				digitLoc += (loc + 1)
				sum += tools.PowInt(10, pow) * i
				break
			}
		}

	}
	return sum
}
