package main

import (
	"github.com/sekullbe/advent/parsers"
	"reflect"
	"testing"
)

func Test_bingoboard_calculateScore(t *testing.T) {
	tests := []struct {
		name      string
		board     bingoboard
		wantScore int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotScore := tt.board.calculateScore(); gotScore != tt.wantScore {
				t.Errorf("calculateScore() = %v, want %v", gotScore, tt.wantScore)
			}
		})
	}
}

func Test_bingoboard_isBingo(t *testing.T) {
	tests := []struct {
		name  string
		board bingoboard
		want  bool
	}{
		{
			name:  "example1",
			board: bingoboard{},
			want:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.board.isBingo(); got != tt.want {
				t.Errorf("isBingo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parse(t *testing.T) {
	type args struct {
		inputText string
	}
	tests := []struct {
		name        string
		args        args
		wantCalled  []int
		wantBoards  []bingoboard
		wantSquares squaremap
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCalled, gotBoards, gotSquares := parse(tt.args.inputText)
			if !reflect.DeepEqual(gotCalled, tt.wantCalled) {
				t.Errorf("parse() gotCalled = %v, want %v", gotCalled, tt.wantCalled)
			}
			if !reflect.DeepEqual(gotBoards, tt.wantBoards) {
				t.Errorf("parse() gotBoards = %v, want %v", gotBoards, tt.wantBoards)
			}
			if !reflect.DeepEqual(gotSquares, tt.wantSquares) {
				t.Errorf("parse() gotSquares = %v, want %v", gotSquares, tt.wantSquares)
			}
		})
	}
}

func Test_parseBoard(t *testing.T) {

	type args struct {
		lines   []string
		squares squaremap
	}
	tests := []struct {
		name string
		args args
		want bingoboard
	}{
		{
			name: "example1",
			args: args{squares: squaremap{}, lines: parsers.SplitByLines("22 13 17 11  0\n 8  2 23  4 24\n21  9 14 16  7\n 6 10  3 18  5\n 1 12 20 15 19\n")},
			want: bingoboard{
				0: map[int]*square{0: {num: 22}, 1: {num: 13}, 2: {num: 17}, 3: {num: 11}, 4: {num: 0}},
				1: map[int]*square{0: {num: 8}, 1: {num: 2}, 2: {num: 23}, 3: {num: 4}, 4: {num: 24}},
				2: map[int]*square{0: {num: 21}, 1: {num: 9}, 2: {num: 14}, 3: {num: 16}, 4: {num: 7}},
				3: map[int]*square{0: {num: 6}, 1: {num: 10}, 2: {num: 3}, 3: {num: 18}, 4: {num: 5}},
				4: map[int]*square{0: {num: 1}, 1: {num: 12}, 2: {num: 20}, 3: {num: 15}, 4: {num: 19}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseBoard(tt.args.lines, tt.args.squares); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseBoard() = %v, want %v", got, tt.want)
			}
		})
	}
}
