package main

import "testing"

func Test_run1(t *testing.T) {
	exampleInput := "5483143223\n2745854711\n5264556173\n6141336146\n6357385478\n4167524645\n2176841721\n6882881134\n4846848554\n5283751526"

	type args struct {
		inputText string
		steps     int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "trivial 0",
			args: args{inputText: "11111\n11111\n11111\n11111\n11111", steps: 1},
			want: 0,
		},
		{
			name: "trivial 1",
			args: args{inputText: "98", steps: 1},
			want: 2,
		},
		{
			name: "simpleexample",
			args: args{inputText: "11111\n19991\n19191\n19991\n11111", steps: 1},
			want: 9,
		},
		{
			name: "example 1",
			args: args{inputText: exampleInput, steps: 1},
			want: 0,
		},
		{
			name: "example 10",
			args: args{inputText: exampleInput, steps: 10},
			want: 204,
		},
		{
			name: "realdeal",
			args: args{inputText: inputText, steps: 100},
			want: 1741,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run1(tt.args.inputText, tt.args.steps); got != tt.want {
				t.Errorf("run1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_run2(t *testing.T) {
	exampleInput := "5483143223\n2745854711\n5264556173\n6141336146\n6357385478\n4167524645\n2176841721\n6882881134\n4846848554\n5283751526"

	type args struct {
		inputText string
		steps     int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example 1",
			args: args{inputText: exampleInput, steps: 1000},
			want: 195,
		},
		{
			name: "realdeal",
			args: args{inputText: inputText, steps: 10000},
			want: 440,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run2(tt.args.inputText, tt.args.steps); got != tt.want {
				t.Errorf("run2() = %v, want %v", got, tt.want)
			}
		})
	}
}
