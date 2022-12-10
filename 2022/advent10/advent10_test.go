package main

import (
	_ "embed"
	"github.com/sekullbe/advent/parsers"
	"testing"
)

//go:embed testinput.txt
var testInput string

func Test_run1(t *testing.T) {
	type args struct {
		lines []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "trivial", args: args{lines: []string{"noop", "addx 3", "addx -5"}}, want: 0},
		{name: "example", args: args{lines: parsers.SplitByLines(testInput)}, want: 13140},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run1(tt.args.lines); got != tt.want {
				t.Errorf("run1() = %v, want %v", got, tt.want)
			}
		})
	}
}
