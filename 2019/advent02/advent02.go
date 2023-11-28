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

	c := computer.NewComputer(parsers.StringsWithCommasToIntSlice(inputText))
	c.Set(1, 12)
	c.Set(2, 2)
	m := c.Run()

	return m[0]
}

func run2(inputText string) int {
	inp := parsers.StringsWithCommasToIntSlice(inputText)
	for n := 0; n < 99; n++ {
		for v := 0; v < 99; v++ {
			c := computer.NewComputer(inp)
			c.Set(1, n)
			c.Set(2, v)
			c.Run()
			if c.Get(0) == 19690720 {
				return 100*n + v
			}
		}
	}
	return 0
}
