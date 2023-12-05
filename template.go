package main

import (
	_ "embed"
	"fmt"
)

//go:embed input.txt
var inputText string

func main() {
	fmt.Printf("Magic number: %d\n", run1(inputText))
	fmt.Println("-------------")
	fmt.Printf("Magic number: %d\n", run2(inputText))
}

func run1(input string) int {

	return 0
}

func run2(input string) int {

	return 0
}
