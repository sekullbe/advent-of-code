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

func run1(inputText string) (magic int) {
	called, boards, squaremap := parse(inputText)

	for _, call := range called {
		squaremap[call].marked = true
		fmt.Printf("Calling %d\n", call)
		for _, board := range boards {
			if board.isBingo() {
				fmt.Println(board.toString())

				return board.calculateScore() * call
			}
		}
	}
	panic("no bingo!")
	return 0
}

func run2(inputText string) (magic int) {
	called, boards, squaremap := parse(inputText)

	boardMap := make(map[int]bingoboard)
	for i, board := range boards {
		boardMap[i] = board
	}

	for _, call := range called {
		squaremap[call].marked = true
		fmt.Printf("Calling %d\n", call)
		if len(boardMap) > 1 {
			for k, board := range boards {
				if board.isBingo() {
					delete(boardMap, k)
				}
			}
		} else {
			fmt.Println("Only one left")
			//fmt.Println(boardMap)
			for _, b := range boardMap {
				fmt.Println(b.toString())
				return b.calculateScore() * call
			}
		}

	}
	panic("no bingo!")

	//66649 too high
	//41951 too high
	// 33462 too high
	return
}
