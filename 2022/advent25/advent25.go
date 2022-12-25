package main

import (
	_ "embed"
	"fmt"
	"github.com/sekullbe/advent/parsers"
	"github.com/sekullbe/advent/tools"
	"math"
)

//go:embed input.txt
var inputText string

func main() {
	fmt.Printf("Magic number: %s\n", run1(inputText))
	fmt.Println("-------------")
	fmt.Printf("Magic number: %d\n", run2(inputText))
}

var snafitToDigit = map[string]int{"0": 0, "1": 1, "2": 2, "1=": 3, "1-": 4}

func snafuToInt(snafu string) int {
	var digits []int
	for _, r := range snafu {
		switch r {
		case '0', '1', '2':
			digits = append(digits, int(r-'0'))
		case '-':
			digits = append(digits, -1)
			//pendingDigit = -1
		case '=':
			digits = append(digits, -2)
			//pendingDigit = -1
		}
	}
	total := 0
	for power, digit := range tools.Reverse(digits) {
		total += digit * int(math.Pow(5.0, float64(power)))
	}
	return total
}

func log5(n float64) float64 {
	// log5 a = log10 a / log10 5
	return math.Log10(float64(n)) / math.Log10(5.0)

}

func intToSnafu(n int) (snafits string) {
	//var maxPower5 = int(log5(float64(n)))
	// convert to base 5 ie 5->10, 8->13, 9->14
	// this is reverse order, 1s, 5s, 25s, 125s, etc
	// eg 27 base 5 -> {2,0,1}
	nBase5 := tools.BaseConvert(n, 5)
	// for 8 -> 13, convert 3 to 1=, carry the 1... 2=
	// for 9 -> 14, convert 4 to 1-, carry the 1... 2-
	var carry = false
	for _, fivit := range nBase5 {
		if carry {
			fivit++
			carry = false
		}
		if fivit == 5 {
			snafits += "0"
			carry = true
		}
		if fivit == 3 {
			snafits += "="
			carry = true
		}
		if fivit == 4 {
			snafits += "-"
			carry = true
		}
		if fivit < 3 {
			snafits += fmt.Sprint(fivit)
		}
	}
	if carry {
		snafits += "1"
	}

	return tools.ReverseString(snafits)
}

func run1(inputText string) string {

	sum := 0
	for _, snafu := range parsers.SplitByLines(inputText) {
		sum += snafuToInt(snafu)
	}
	return intToSnafu(sum)
}

func run2(inputText string) int {

	return 0
}
