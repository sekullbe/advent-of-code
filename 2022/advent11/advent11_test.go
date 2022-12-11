package main

import (
	_ "embed"
	"github.com/sekullbe/advent/parsers"
	"reflect"
	"testing"
)

//go:embed testinput.txt
var testInput string

var sampleMonkey string = `Monkey 0:
  Starting items: 79, 98
  Operation: new = old * 19
  Test: divisible by 23
    If true: throw to monkey 2
    If false: throw to monkey 3
`

func Test_parseOneMonkey(t *testing.T) {
	type args struct {
		lines []string
	}
	tests := []struct {
		name string
		args args
		want monkey
	}{
		{name: "simple", args: args{parsers.SplitByLines(sampleMonkey)}, want: monkey{
			id:          0,
			items:       []int{79, 98},
			op:          '*',
			opArg:       "19",
			testNum:     23,
			targetTrue:  2,
			targetFalse: 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseOneMonkey(tt.args.lines); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseOneMonkey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseMonkeys(t *testing.T) {
	type args struct {
		lines []string
	}
	tests := []struct {
		name string
		args args
		want barrel
	}{
		{name: "sample", args: args{parsers.SplitByLinesNoTrim(testInput)}, want: barrel{
			lcd: 1,
			monkeys: map[int]monkey{
				0: monkey{
					id:          0,
					items:       []int{79, 98},
					op:          '*',
					opArg:       "19",
					testNum:     23,
					targetTrue:  2,
					targetFalse: 3},
				1: monkey{
					id:          1,
					items:       []int{54, 65, 75, 74},
					op:          '+',
					opArg:       "6",
					testNum:     19,
					targetTrue:  2,
					targetFalse: 0},
				2: monkey{
					id:          2,
					items:       []int{79, 60, 97},
					op:          '*',
					opArg:       "old",
					testNum:     13,
					targetTrue:  1,
					targetFalse: 3},
				3: monkey{
					id:          3,
					items:       []int{74},
					op:          '+',
					opArg:       "3",
					testNum:     17,
					targetTrue:  0,
					targetFalse: 1},
			}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseMonkeys(tt.args.lines); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseMonkeys() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_run1(t *testing.T) {
	type args struct {
		inputText string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "sample", args: args{testInput}, want: 10605},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run1(tt.args.inputText); got != tt.want {
				t.Errorf("run1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_run2(t *testing.T) {
	type args struct {
		inputText string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "sample", args: args{testInput}, want: 2713310158},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run2(tt.args.inputText); got != tt.want {
				t.Errorf("run2() = %v, want %v", got, tt.want)
			}
		})
	}
}
