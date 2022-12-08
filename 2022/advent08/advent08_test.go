package main

import (
	"github.com/sekullbe/advent/parsers"
	"github.com/stretchr/testify/assert"
	"testing"
)

var testInput string = `30373
25512
65332
33549
35390`

func Test_run1(t *testing.T) {

	type args struct {
		lines []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "example", args: args{lines: parsers.SplitByLines(testInput)}, want: 21},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, run1(tt.args.lines), "run1(%v)", tt.args.lines)
		})
	}
}
