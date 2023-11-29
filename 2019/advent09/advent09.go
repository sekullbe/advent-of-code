package main

import (
	_ "embed"
	"fmt"

	"github.com/sekullbe/advent/2019/computer"
	"github.com/sekullbe/advent/parsers"
)

//go:embed input.txt
var inputText string

func main() {
	fmt.Printf("Magic number: %d\n", run1(inputText))
	fmt.Println("-------------")
	fmt.Printf("Magic number: %d\n", run2(inputText))
}

func run1(inputText string) int {
	c := computer.NewComputer(parsers.StringsWithCommasToInt64Slice(inputText), []int64{1})
	c.Run()
	for _, o := range c.GetOutputs() {
		fmt.Println(o)
	}

	return 0
}

func run2(inputText string) int {
	c := computer.NewComputer(parsers.StringsWithCommasToInt64Slice(inputText), []int64{2})
	c.Run()
	for _, o := range c.GetOutputs() {
		fmt.Println(o)
	}

	return 0
}
