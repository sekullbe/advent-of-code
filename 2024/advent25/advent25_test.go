package main

import (
	_ "embed"
	"reflect"
	"testing"
)

//go:embed sample.txt
var sample1 string

func Test_parseLock(t *testing.T) {
	type args struct {
		block []string
	}
	tests := []struct {
		name string
		args args
		want lock
	}{
		{name: "simple", args: args{block: []string{"#####", ".####", ".####", ".####", ".#.#.", ".#...", "....."}}, want: lock{0, 5, 3, 4, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseLock(tt.args.block); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseLock() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseKey(t *testing.T) {
	type args struct {
		block []string
	}
	tests := []struct {
		name string
		args args
		want key
	}{
		{name: "simple", args: args{block: []string{".....", "#....", "#....", "#...#", "#.#.#", "#.###", "#####"}}, want: key{5, 0, 2, 1, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseKey(tt.args.block); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseKey() = %v, want %v", got, tt.want)
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
		{name: "sample1", args: args{input: sample1}, want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run1(tt.args.input); got != tt.want {
				t.Errorf("run1() = %v, want %v", got, tt.want)
			}
		})
	}
}
