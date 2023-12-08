package main

import (
	"reflect"
	"testing"
)

const sampleInput1 = `RL

AAA = (BBB, CCC)
BBB = (DDD, EEE)
CCC = (ZZZ, GGG)
DDD = (DDD, DDD)
EEE = (EEE, EEE)
GGG = (GGG, GGG)
ZZZ = (ZZZ, ZZZ)`

const sampleInput2 = `LLR

AAA = (BBB, BBB)
BBB = (AAA, ZZZ)
ZZZ = (ZZZ, ZZZ)`

func Test_parseOneNode(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 node
	}{
		{name: "simple", args: args{input: "AAA = (BBB, CCC)"}, want: "AAA", want1: node{left: "BBB", right: "CCC"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := parseOneNode(tt.args.input)
			if got != tt.want {
				t.Errorf("parseOneNode() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("parseOneNode() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_nextStep(t *testing.T) {
	type args struct {
		instructions string
		index        int
	}
	tests := []struct {
		name string
		args args
		want rune
	}{
		{name: "simple", args: args{instructions: "LLLRRR", index: 0}, want: 'L'},
		{name: "end", args: args{instructions: "LLLRRR", index: 5}, want: 'R'},
		{name: "wraparound", args: args{instructions: "LLLRRR", index: 6}, want: 'L'},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nextStep(tt.args.instructions, tt.args.index); got != tt.want {
				t.Errorf("nextStep() = %v, want %v", got, tt.want)
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
		{name: "sample", args: args{input: sampleInput1}, want: 2},
		{name: "sample", args: args{input: sampleInput2}, want: 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run1(tt.args.input); got != tt.want {
				t.Errorf("run1() = %v, want %v", got, tt.want)
			}
		})
	}
}
