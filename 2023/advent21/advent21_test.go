package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_run1(t *testing.T) {
	type args struct {
		input string
		steps int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "sample",
			args: args{
				input: sample,
				steps: 6,
			},
			want: 16,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run1(tt.args.input, tt.args.steps); got != tt.want {
				t.Errorf("run1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_run2WithBFS(t *testing.T) {
	type args struct {
		input string
		steps int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "sample", args: args{input: sample, steps: 6}, want: 16},
		{name: "sample", args: args{input: sample, steps: 22}, want: 261},
		{name: "sample", args: args{input: sample, steps: 50}, want: 1594},
		{name: "sample", args: args{input: sample, steps: 100}, want: 6536},
		{name: "sample", args: args{input: sample, steps: 500}, want: 167004},
		//{name: "sample", args: args{input: sample, steps: 1000}, want: 668697}, // too slow
		//{name: "sample", args: args{input: sample, steps: 5000}, want: 16733044}, // too slow
		// and these special numbers for testing the quadratic solution
		{name: "sample", args: args{input: sample, steps: 0*11 + 6}, want: 16},
		{name: "sample", args: args{input: sample, steps: 1*11 + 6}, want: 145},
		{name: "sample", args: args{input: sample, steps: 2*11 + 6}, want: 460},
		{name: "sample", args: args{input: sample, steps: 3*11 + 6}, want: 944},
		{name: "sample", args: args{input: sample, steps: 4*11 + 6}, want: 1594},
		{name: "sample", args: args{input: sample, steps: 5*11 + 6}, want: 2406},
		{name: "sample", args: args{input: sample, steps: 6*11 + 6}, want: 3380},
		{name: "sample", args: args{input: sample, steps: 7*11 + 6}, want: 4516},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, run2WithBFS(tt.args.input, tt.args.steps), "run2(%v, %v)", tt.args.input, tt.args.steps)
		})
	}
}

func Test_run2(t *testing.T) {
	type args struct {
		input string
		steps int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "sample", args: args{input: sample, steps: 0*11 + 6}, want: 16},
		{name: "sample", args: args{input: sample, steps: 1*11 + 6}, want: 145},
		{name: "sample", args: args{input: sample, steps: 2*11 + 6}, want: 460},
		{name: "sample", args: args{input: sample, steps: 3*11 + 6}, want: 944},
		{name: "sample", args: args{input: sample, steps: 4*11 + 6}, want: 1594},
		{name: "sample", args: args{input: sample, steps: 5*11 + 6}, want: 2406},
		{name: "sample", args: args{input: sample, steps: 6*11 + 6}, want: 3380},
		{name: "sample", args: args{input: sample, steps: 7*11 + 6}, want: 4516},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, run2WithBFS(tt.args.input, tt.args.steps), "run2(%v, %v)", tt.args.input, tt.args.steps)
		})
	}
}
