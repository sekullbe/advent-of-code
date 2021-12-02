package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var inputText string

func main() {
	fmt.Printf("Magic number: %d\n", run1(inputText))
	fmt.Println("-------------")
	fmt.Printf("Magic number: %d\n", run2(inputText))
}

func run1(inputText string) int {

	pos := 0
	depth := 0
	for _,command := range strings.Split(inputText, "\n") {

		commandParts := strings.Fields(command)
		if len(commandParts) < 2 {
			continue
		}
		commandNum, _ := strconv.Atoi(commandParts[1])
		switch commandParts[0] {
		case "forward":
			pos += commandNum
		case "up":
			depth -= commandNum
		case "down":
			depth += commandNum
		}
	}

	fmt.Printf("Position: %d\n", pos)
	fmt.Printf("Depth: %d\n", depth)
	
	return pos*depth
}

func run2(inputText string) int {
	pos := 0
	depth := 0
	aim := 0
	for _,command := range strings.Split(inputText, "\n") {

		commandParts := strings.Fields(command)
		if len(commandParts) < 2 {
			continue
		}
		commandNum, _ := strconv.Atoi(commandParts[1])
		switch commandParts[0] {
		case "forward":
			pos += commandNum
			depth += commandNum * aim
		case "up":
			aim -= commandNum
		case "down":
			aim += commandNum
		}
	}

	fmt.Printf("Position: %d\n", pos)
	fmt.Printf("Depth: %d\n", depth)

	return pos*depth

}
