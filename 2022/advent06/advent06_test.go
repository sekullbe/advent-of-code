package main

import "testing"

func Test_run1(t *testing.T) {
	type args struct {
		inputText string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "Example1", args: args{inputText: "mjqjpqmgbljsphdztnvjfqwrcgsmlb"}, want: 7},
		{name: "Example2", args: args{inputText: "bvwbjplbgvbhsrlpgdmjqwftvncz"}, want: 5},
		{name: "Example3", args: args{inputText: "nppdvjthqldpwncqszvftbrmjlhg"}, want: 6},
		{name: "Example4", args: args{inputText: "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg"}, want: 10},
		{name: "Example5", args: args{inputText: "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw"}, want: 11},
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
		{name: "Example1", args: args{inputText: "mjqjpqmgbljsphdztnvjfqwrcgsmlb"}, want: 19},
		{name: "Example2", args: args{inputText: "bvwbjplbgvbhsrlpgdmjqwftvncz"}, want: 23},
		{name: "Example3", args: args{inputText: "nppdvjthqldpwncqszvftbrmjlhg"}, want: 23},
		{name: "Example4", args: args{inputText: "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg"}, want: 29},
		{name: "Example5", args: args{inputText: "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw"}, want: 26},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run2(tt.args.inputText); got != tt.want {
				t.Errorf("run2() = %v, want %v", got, tt.want)
			}
		})
	}
}
