package main

import "testing"

const sampleText = `###############
#...#...#.....#
#.#.#.#.#.###.#
#S#...#.#.#...#
#######.#.#.###
#######.#.#...#
#######.#.###.#
###..E#...#...#
###.#######.###
#...###...#...#
#.#####.#.###.#
#.#...#.#.#...#
#.#.#.#.#.#.###
#...#...#...###
###############`

func Test_run1(t *testing.T) {
	type args struct {
		input       string
		mustImprove int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "improve 1", args: args{input: sampleText, mustImprove: 1}, want: 44},
		{name: "improve 10", args: args{input: sampleText, mustImprove: 10}, want: 10},
		{name: "improve 20", args: args{input: sampleText, mustImprove: 20}, want: 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run1(tt.args.input, tt.args.mustImprove); got != tt.want {
				t.Errorf("run1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_run2(t *testing.T) {
	type args struct {
		input       string
		mustImprove int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "improve 50", args: args{input: sampleText, mustImprove: 50}, want: 32 + 31 + 29 + 39 + 25 + 23 + 20 + 19 + 12 + 14 + 12 + 22 + 4 + 3},
		{name: "improve 72", args: args{input: sampleText, mustImprove: 72}, want: 29},
		{name: "improve 74", args: args{input: sampleText, mustImprove: 74}, want: 7},
		{name: "improve 76", args: args{input: sampleText, mustImprove: 76}, want: 3},
		{name: "improve 77", args: args{input: sampleText, mustImprove: 77}, want: 0},
		{name: "improve 90", args: args{input: sampleText, mustImprove: 90}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run2(tt.args.input, tt.args.mustImprove); got != tt.want {
				t.Errorf("run2() = %v, want %v", got, tt.want)
			}
		})
	}
}
