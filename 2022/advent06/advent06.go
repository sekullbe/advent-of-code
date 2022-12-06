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

func run1(inputText string) int {

	buf := NewReadBuffer(4)
	for i, r := range inputText {
		buf.add(r)
		if i >= 3 && buf.contentsAllUnique() { // don't check until we've read 4
			return i + 1 // offset because reading the string is zero-based
		}
	}
	return 0
}

func run2(inputText string) int {

	buf := NewReadBuffer(14)
	for i, r := range inputText {
		buf.add(r)
		if i >= 3 && buf.contentsAllUnique() { // don't check until we've read 4
			return i + 1 // offset because reading the string is zero-based
		}
	}
	return 0
}
