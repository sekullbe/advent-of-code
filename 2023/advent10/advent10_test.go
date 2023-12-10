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
