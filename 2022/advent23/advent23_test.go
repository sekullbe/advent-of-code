package main

import (
	_ "embed"
	"testing"
)

var tinySample string = `.....
..##.
..#..
.....
..##.
.....`

//go:embed sample.txt
var sampleText string

func Test_run1(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		//{name: "tinySample", args: args{tinySample}, want: 0},
		{name: "sample", args: args{sampleText}, want: 110},
		{name: "live", args: args{inputText}, want: 3684},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run1(tt.args.input); got != tt.want {
				t.Errorf("run1() = %v, want %v", got, tt.want)
			}
		})
	}
}
