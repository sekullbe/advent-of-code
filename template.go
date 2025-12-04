package main

import (
	_ "embed"
	"fmt"
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

	return 0
}

func run2(input string) int {
	defer tools.Track(time.Now(), "Part 2 Time")

	return 0
}
