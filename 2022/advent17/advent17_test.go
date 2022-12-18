package main

import "testing"

var testinput = ">>><<><>><<<>><>>><<<>>><<<><<<>><>><<>>"

func Test_run1(t *testing.T) {
	type args struct {
		inputText string
		rockCount int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "examplepart1", args: args{testinput, 2022}, want: 3068},
		{name: "examplepart2", args: args{testinput, 1_000_000_000_000}, want: 1514285714288},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run1(tt.args.inputText, tt.args.rockCount); got != tt.want {
				t.Errorf("run1() = %v, want %v", got, tt.want)
			}
		})
	}
}
