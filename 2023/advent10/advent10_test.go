package main

import "testing"

func Test_run1(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "sample1", args: args{input: sample1}, want: 4},
		{name: "sample1b", args: args{input: sample1b}, want: 4},
		{name: "sample2", args: args{input: sample2}, want: 8},
		{name: "sample2b", args: args{input: sample2b}, want: 8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run1(tt.args.input); got != tt.want {
				t.Errorf("run1() = %v, want %v", got, tt.want)
			}
		})
	}
}

const samplePart21 = `...........
.S-------7.
.|F-----7|.
.||.....||.
.||.....||.
.|L-7.F-J|.
.|..|.|..|.
.L--J.L--J.
...........`

const samplePart21b = `..........
.S------7.
.|F----7|.
.||....||.
.||....||.
.|L-7F-J|.
.|II||II|.
.L--JL--J.
..........`

const samplePart22 = `.F----7F7F7F7F-7....
.|F--7||||||||FJ....
.||.FJ||||||||L7....
FJL7L7LJLJ||LJ.L-7..
L--J.L7...LJS7F-7L7.
....F-J..F7FJ|L7L7L7
....L7.F7||L7|.L7L7|
.....|FJLJ|FJ|F7|.LJ
....FJL-7.||.||||...
....L---J.LJ.LJLJ...`

func Test_run2(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "sample1", args: args{input: samplePart21}, want: 4},
		{name: "sample1", args: args{input: samplePart21b}, want: 4},
		{name: "sample2", args: args{input: samplePart22}, want: 8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run2(tt.args.input); got != tt.want {
				t.Errorf("run2() = %v, want %v", got, tt.want)
			}
		})
	}
}
