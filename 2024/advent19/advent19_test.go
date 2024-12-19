package main

import (
	"reflect"
	"testing"
)

const sampleText = `r, wr, b, g, bwu, rb, gb, br

brwrr
bggr
gbbr
rrbgbr
ubwu
bwurrg
brgr
bbrgwb`

func Test_run1(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "sample", args: args{input: sampleText}, want: 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run1(tt.args.input); got != tt.want {
				t.Errorf("run1() = %v, want %v", got, tt.want)
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
		{name: "sample", args: args{input: sampleText}, want: 16},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run2(tt.args.input); got != tt.want {
				t.Errorf("run2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseOnsen(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want onsen
	}{
		{name: "sample", args: args{sampleText},
			want: onsen{patterns: []string{"r", "wr", "b", "g", "bwu", "rb", "gb", "br"}, designs: []string{"brwrr", "bggr", "gbbr", "rrbgbr", "ubwu", "bwurrg", "brgr", "bbrgwb"}}}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseOnsen(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseOnsen() = %v, want %v", got, tt.want)
			}
		})
	}
}
