package main

import (
	"fmt"
	"github.com/sekullbe/advent/parsers"
)

type square struct {
	num    int
	marked bool
}

// map from num to square, so all squares of the same num can be set at once
type squaremap map[int]*square

type bingoboard map[int]map[int]*square

func (board *bingoboard) isBingo() bool {
	// check rows
	for r := 0; r < 5; r++ {
		bingo := true
		//row := (*board)[r]
		for col := 0; col < 5; col++ {
			bingo = bingo && (*board)[r][col].marked
		}
		if bingo {
			return true
		}
	}
	// check cols
	for c := 0; c < 5; c++ {
		bingo := true
		for r := 0; r < 5; r++ {
			bingo = bingo && (*board)[r][c].marked
		}
		if bingo {
			return true
		}
	}
	return false
}

func (board *bingoboard) calculateScore() (score int) {
	for r := 0; r < 5; r++ {
		for c := 0; c < 5; c++ {
			sq := (*board)[r][c]
			if !sq.marked {
				score += sq.num
			}
		}
	}
	return
}

func (board *bingoboard) toString() string {
	var out string
	for r := 0; r < 5; r++ {
		for c := 0; c < 5; c++ {
			sq := (*board)[r][c]
			sqstr := fmt.Sprintf("%2d", sq.num)
			if sq.marked {
				sqstr += "* "
			} else {
				sqstr += "  "
			}
			out += sqstr
		}
		out += "\n"
	}
	return out
}

func parse(inputText string) (called []int, boards []bingoboard, squares squaremap) {

	lines := parsers.SplitByLines(inputText)
	called = parsers.StringsWithCommasToIntSlice(lines[0])
	boards = []bingoboard{}
	squares = make(squaremap)

	// then a blank line

	// then 5x5 pattern, then a blank line
	// so boards start at 2,8,14
	for i := 2; i < len(lines); i += 6 {
		boards = append(boards, parseBoard(lines[i:i+6], squares))

	}
	return
}

func parseBoard(lines []string, squares squaremap) bingoboard {
	board := make(bingoboard)
	for row, line := range lines {
		board[row] = make(map[int]*square)
		lineNums := parsers.StringsToIntSlice(line)
		for col, num := range lineNums {
			// get square if known
			sq, known := squares[num]
			if !known {
				sq = &square{num: num, marked: false}
			}
			board[row][col] = sq
			squares[num] = sq
		}
	}
	return board
}
