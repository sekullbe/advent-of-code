package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const sample1 = `broadcaster -> a, b, c
%a -> b
%b -> c
%c -> inv
&inv -> a`

const sample2 = `broadcaster -> a
%a -> inv, con
&inv -> b
%b -> con
&con -> output
output -> output`

func Test_run1(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "sample1", args: args{input: sample1}, want: 32000000},
		{name: "sample2", args: args{input: sample2}, want: 11687500},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, run1(tt.args.input), "run1(%v)", tt.args.input)
		})
	}
}
