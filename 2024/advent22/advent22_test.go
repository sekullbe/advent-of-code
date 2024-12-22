package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_iterateNextNum(t *testing.T) {
	type args struct {
		num   int
		steps int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "example1", args: args{num: 123, steps: 1}, want: 15887950},
		{name: "example2", args: args{num: 123, steps: 2}, want: 16495136},
		{name: "example10", args: args{num: 123, steps: 10}, want: 5908254},
		{name: "sample1", args: args{num: 1, steps: 2000}, want: 8685429},
		{name: "sample2", args: args{num: 10, steps: 2000}, want: 4700978},
		{name: "sample3", args: args{num: 100, steps: 2000}, want: 15273692},
		{name: "sample4", args: args{num: 2024, steps: 2000}, want: 8667524},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := iterateNextNum(tt.args.num, tt.args.steps); got != tt.want {
				t.Errorf("iterateNextNum() = %v, want %v", got, tt.want)
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
		{name: "example1", args: args{input: "1 10 100 2024"}, want: 37327623},
		{name: "reddit1", args: args{input: "2021 5017 19751"}, want: 18183557},
		{name: "reddit2", args: args{input: "5053 10083 11263"}, want: 8876699},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run1(tt.args.input); got != tt.want {
				t.Errorf("run1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sequence_push(t *testing.T) {
	s := sequence{2, 3, 4, 5}
	s.push(9)
	assert.Equal(t, 4, len(s))
	assert.Equal(t, 3, s[0])
	assert.Equal(t, 9, s[3])
	assert.Equal(t, sequence{3, 4, 5, 9}, s)
	s.push(0)
	assert.Equal(t, sequence{4, 5, 9, 0}, s)

	// reminding myself how to copy the array
	a := s
	s.push(5)
	assert.NotEqual(t, a, s)
	assert.Equal(t, sequence{4, 5, 9, 0}, a)
	assert.Equal(t, sequence{5, 9, 0, 5}, s)

}

func Test_processOneBuyer(t *testing.T) {
	s2p := processOneBuyer(123, 10)
	fmt.Println(s2p)
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
		{name: "example1", args: args{input: "1 2 3 2024"}, want: 23},
		{name: "reddit1", args: args{input: "2021 5017 19751"}, want: 27},
		{name: "reddit2", args: args{input: "5053 10083 11263"}, want: 27},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, run2(tt.args.input), "run2(%v)", tt.args.input)
		})
	}
}
