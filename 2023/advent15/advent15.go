package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var inputText string

func main() {
	fmt.Printf("Magic number: %d\n", run1(inputText))
	fmt.Println("-------------")
	fmt.Printf("Magic number: %d\n", run2(inputText))
}

func run1(input string) int {

	sum := 0
	// without the trim it picked up a newline at the end of the string
	steps := strings.Split(strings.TrimSpace(input), ",")
	for _, step := range steps {
		sum += processStep(step)
	}

	return sum
}

func run2(input string) int {

	return 0
}

func processStep(step string) int {
	val := 0
	stepInts := stringToAsciiInt(step)
	for _, u := range stepInts {
		val += u
		val *= 17
		val %= 256
	}
	return val
}

func stringToAsciiInt(str string) []int {
	b := []byte(str)
	i := []int{}
	for _, b2 := range b {
		i = append(i, int(b2))
	}
	return i
}
