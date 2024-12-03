package main

import (
	"reflect"
	"testing"
)

func Test_run1(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "sampletext", args: args{input: "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"}, want: 161},
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
		{name: "sampletext", args: args{input: "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"}, want: 48},
		{name: "sampletext2", args: args{input: "xmul(2,4)&mul[3,7]!^do()don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"}, want: 48},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run2(tt.args.input); got != tt.want {
				t.Errorf("run2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_extractMuls(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want []pair
	}{
		{name: "sampletext", args: args{input: "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"}, want: []pair{{2, 4}, {5, 5}, {11, 8}, {8, 5}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := extractMuls(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("extractMuls() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_extractEnabled(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{name: "trivial", args: args{input: "xyzzydon't()plugh"}, want: []string{"xyzzy"}},
		{name: "sampletext", args: args{input: "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"}, want: []string{"xmul(2,4)&mul[3,7]!^", "?mul(8,5))"}},
		{name: "sampletextwithnewlines", args: args{input: "xmul(2,4)&\nmul[3,7]!^don't()\nmul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"}, want: []string{"xmul(2,4)&mul[3,7]!^", "?mul(8,5))"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := extractEnabled(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("extractEnabled() = %v, want %v", got, tt.want)
			}
		})
	}
}
