package main

import (
	"reflect"
	"testing"
)

const sampleText = `123 328  51 64 
 45 64  387 23 
  6 98  215 314
*   +   *   +  `

func Test_parseMathInputsPart1(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want []problem
	}{
		{name: "sample", args: args{sampleText}, want: []problem{
			{[]int{123, 45, 6}, "*"},
			{[]int{328, 64, 98}, "+"},
			{[]int{51, 387, 215}, "*"},
			{[]int{64, 23, 314}, "+"},
		},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseMathInputsPart1(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseMathInputsPart1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_solveProblem(t *testing.T) {
	type args struct {
		p problem
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "sample1", args: args{problem{[]int{123, 45, 6}, "*"}}, want: 33210},
		{name: "sample2", args: args{problem{[]int{328, 64, 98}, "+"}}, want: 490},
		{name: "sample3", args: args{problem{[]int{51, 387, 215}, "*"}}, want: 4243455},
		{name: "sample4", args: args{problem{[]int{64, 23, 314}, "+"}}, want: 401},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solveProblem(tt.args.p); got != tt.want {
				t.Errorf("solveProblem() = %v, want %v", got, tt.want)
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
		{name: "sample", args: args{sampleText}, want: 3263827},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run2(tt.args.input); got != tt.want {
				t.Errorf("run2() = %v, want %v", got, tt.want)
			}
		})
	}
}
