package main

import "testing"

const sampleInput = `
R 6 (#70c710)
D 5 (#0dc571)
L 2 (#5713f0)
D 2 (#d2c081)
R 2 (#59c680)
D 2 (#411b91)
L 5 (#8ceee2)
U 2 (#caa173)
L 1 (#1b58a2)
U 2 (#caa171)
R 2 (#7807d2)
U 3 (#a77fa3)
L 2 (#015232)
U 2 (#7a21e3)`

func Test_run1(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "sample",
			args: args{
				input: sampleInput,
			},
			want: 62,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run1better(tt.args.input); got != tt.want {
				t.Errorf("run1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_run2(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "sample",
			args: args{
				input: sampleInput,
			},
			want: 952408144115,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run2(tt.args.input); got != tt.want {
				t.Errorf("run2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parsePart2Instruction(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name     string
		args     args
		wantDir  int
		wantDist int
	}{
		{name: "simple", args: args{line: "R 6 (#70c710)"}, wantDir: RIGHT, wantDist: 461937},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotDir, gotDist := parsePart2Instruction(tt.args.line)
			if gotDir != tt.wantDir {
				t.Errorf("parsePart2Instruction() gotDir = %v, want %v", gotDir, tt.wantDir)
			}
			if gotDist != tt.wantDist {
				t.Errorf("parsePart2Instruction() gotDist = %v, want %v", gotDist, tt.wantDist)
			}
		})
	}
}
