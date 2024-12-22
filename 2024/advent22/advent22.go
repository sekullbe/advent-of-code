package main

import (
	_ "embed"
	"fmt"
	"github.com/sekullbe/advent/parsers"
	"github.com/sekullbe/advent/tools"
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

func run2(input string) int {
	defer tools.Track(time.Now(), "Part 2 Time")

	return 0
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
