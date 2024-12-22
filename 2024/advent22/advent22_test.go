package main

import "testing"

func Test_iterateNextNum(t *testing.T) {
	type args struct {
		num   int
		steps int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "example1", args: args{num: 123, steps: 1}, want: 15887950},
		{name: "example2", args: args{num: 123, steps: 2}, want: 16495136},
		{name: "example10", args: args{num: 123, steps: 10}, want: 5908254},
		{name: "sample1", args: args{num: 1, steps: 2000}, want: 8685429},
		{name: "sample2", args: args{num: 10, steps: 2000}, want: 4700978},
		{name: "sample3", args: args{num: 100, steps: 2000}, want: 15273692},
		{name: "sample4", args: args{num: 2024, steps: 2000}, want: 8667524},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := iterateNextNum(tt.args.num, tt.args.steps); got != tt.want {
				t.Errorf("iterateNextNum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_run1(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "example1", args: args{input: "1 10 100 2024"}, want: 37327623},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run1(tt.args.input); got != tt.want {
				t.Errorf("run1() = %v, want %v", got, tt.want)
			}
		})
	}
}
