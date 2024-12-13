package main

import (
	"github.com/sekullbe/advent/parsers"
	"reflect"
	"testing"
)

const oneMachine = `Button A: X+94, Y+34
Button B: X+22, Y+67
Prize: X=8400, Y=5400`
const sampleText = `Button A: X+94, Y+34
Button B: X+22, Y+67
Prize: X=8400, Y=5400

Button A: X+26, Y+66
Button B: X+67, Y+21
Prize: X=12748, Y=12176

Button A: X+17, Y+86
Button B: X+84, Y+37
Prize: X=7870, Y=6450

Button A: X+69, Y+23
Button B: X+27, Y+71
Prize: X=18641, Y=10279`

func Test_run1(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "sample1", args: args{input: sampleText}, want: 480},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run1(tt.args.input); got != tt.want {
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
		{name: "sample1", args: args{input: sampleText}, want: 875318608908},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run2(tt.args.input); got != tt.want {
				t.Errorf("run2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseMachine(t *testing.T) {
	type args struct {
		mlines []string
	}
	tests := []struct {
		name string
		args args
		want machine
	}{
		{name: "one",
			args: args{mlines: parsers.SplitByLines(oneMachine)},
			want: machine{
				aX: 94,
				aY: 34,
				bX: 22,
				bY: 67,
				pX: 8400,
				pY: 5400,
			}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseMachine(tt.args.mlines); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseMachine() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_solveOneMachine(t *testing.T) {
	type args struct {
		m machine
	}
	tests := []struct {
		name     string
		args     args
		wantCost int
		wantErr  bool
	}{
		{name: "unsolvable0", args: args{machine{aX: 4, aY: 1, bX: 10, bY: 1, pX: 17, pY: 1}}, wantCost: 0, wantErr: true},
		{name: "solvable 1", args: args{machine{aX: 94, aY: 34, bX: 22, bY: 67, pX: 8400, pY: 5400}}, wantCost: 280, wantErr: false},
		{name: "unsolvable 2", args: args{machine{aX: 26, aY: 66, bX: 67, bY: 21, pX: 12748, pY: 12176}}, wantCost: 0, wantErr: true},
		{name: "solvable 3", args: args{machine{aX: 17, aY: 86, bX: 84, bY: 37, pX: 7870, pY: 6450}}, wantCost: 200, wantErr: false},
		{name: "unsolvable 4", args: args{machine{aX: 69, aY: 23, bX: 27, bY: 71, pX: 18641, pY: 10279}}, wantCost: 0, wantErr: true},

		{name: "unsolvable part2 1", args: args{machine{aX: 94, aY: 34, bX: 22, bY: 67, pX: 10000000008400, pY: 10000000005400}}, wantCost: 0, wantErr: true},
		{name: "solvable part2 2", args: args{machine{aX: 26, aY: 66, bX: 67, bY: 21, pX: 10000000012748, pY: 10000000012176}}, wantCost: 459236326669, wantErr: false},
		{name: "unsolvable part2 3", args: args{machine{aX: 17, aY: 86, bX: 84, bY: 37, pX: 10000000007870, pY: 10000000006450}}, wantCost: 0, wantErr: true},
		{name: "solvable part2 4", args: args{machine{aX: 69, aY: 23, bX: 27, bY: 71, pX: 10000000018641, pY: 10000000010279}}, wantCost: 416082282239, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCost, err := solveOneMachine(tt.args.m)
			if (err != nil) != tt.wantErr {
				t.Errorf("solveOneMachine() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotCost != tt.wantCost {
				t.Errorf("solveOneMachine() gotCost = %v, want %v", gotCost, tt.wantCost)
			}
		})
	}
}

func Test_solveOneMachineCramer(t *testing.T) {
	type args struct {
		m machine
	}
	tests := []struct {
		name     string
		args     args
		wantCost int
		wantErr  bool
	}{
		{name: "unsolvable0", args: args{machine{aX: 4, aY: 1, bX: 10, bY: 1, pX: 17, pY: 1}}, wantCost: 0, wantErr: true},
		{name: "solvable 1", args: args{machine{aX: 94, aY: 34, bX: 22, bY: 67, pX: 8400, pY: 5400}}, wantCost: 280, wantErr: false},
		{name: "unsolvable 2", args: args{machine{aX: 26, aY: 66, bX: 67, bY: 21, pX: 12748, pY: 12176}}, wantCost: 0, wantErr: true},
		{name: "solvable 3", args: args{machine{aX: 17, aY: 86, bX: 84, bY: 37, pX: 7870, pY: 6450}}, wantCost: 200, wantErr: false},
		{name: "unsolvable 4", args: args{machine{aX: 69, aY: 23, bX: 27, bY: 71, pX: 18641, pY: 10279}}, wantCost: 0, wantErr: true},

		{name: "unsolvable part2 1", args: args{machine{aX: 94, aY: 34, bX: 22, bY: 67, pX: 10000000008400, pY: 10000000005400}}, wantCost: 0, wantErr: true},
		{name: "solvable part2 2", args: args{machine{aX: 26, aY: 66, bX: 67, bY: 21, pX: 10000000012748, pY: 10000000012176}}, wantCost: 459236326669, wantErr: false},
		{name: "unsolvable part2 3", args: args{machine{aX: 17, aY: 86, bX: 84, bY: 37, pX: 10000000007870, pY: 10000000006450}}, wantCost: 0, wantErr: true},
		{name: "solvable part2 4", args: args{machine{aX: 69, aY: 23, bX: 27, bY: 71, pX: 10000000018641, pY: 10000000010279}}, wantCost: 416082282239, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCost, err := solveOneMachineCramer(tt.args.m)
			if (err != nil) != tt.wantErr {
				t.Errorf("solveOneMachineCramer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotCost != tt.wantCost {
				t.Errorf("solveOneMachineCramer() gotCost = %v, want %v", gotCost, tt.wantCost)
			}
		})
	}
}
