package main

import (
	"reflect"
	"testing"
)

func Test_diffsequence(t *testing.T) {
	type args struct {
		seq []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "simple1", args: args{seq: []int{0, 3, 6, 9, 12, 15}}, want: []int{3, 3, 3, 3, 3}},
		{name: "simple2", args: args{seq: []int{0, 3, 6, 9, 12, 15, 18}}, want: []int{3, 3, 3, 3, 3, 3}},
		{name: "simple1 repeated", args: args{seq: []int{0, 3, 6, 9, 12, 15}}, want: []int{3, 3, 3, 3, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := diffsequence(tt.args.seq); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("diffsequence() = %v, want %v", got, tt.want)
			}
		})
	}
}

// uses repeatdiff so isn't _quite_ a unit test
func Test_extrapolate(t *testing.T) {
	type args struct {
		sequence []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "example 1", args: args{sequence: []int{0, 3, 6, 9, 12, 15}}, want: 18},
		{name: "example 2", args: args{sequence: []int{1, 3, 6, 10, 15, 21}}, want: 28},
		{name: "example 3", args: args{sequence: []int{10, 13, 16, 21, 30, 45}}, want: 68},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			s2 := repeatdiff(tt.args.sequence)
			if got := extrapolate(s2); got != tt.want {
				t.Errorf("extrapolate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_repeatdiff(t *testing.T) {
	type args struct {
		seq []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{name: "example1", args: args{seq: []int{0, 3, 6, 9, 12, 15}},
			want: [][]int{[]int{0, 3, 6, 9, 12, 15}, []int{3, 3, 3, 3, 3}, []int{0, 0, 0, 0}}},
		{name: "example2", args: args{seq: []int{1, 3, 6, 10, 15, 21}},
			want: [][]int{[]int{1, 3, 6, 10, 15, 21}, []int{2, 3, 4, 5, 6}, []int{1, 1, 1, 1}, []int{0, 0, 0}}},
		{name: "example3", args: args{seq: []int{10, 13, 16, 21, 30, 45, 68}},
			want: [][]int{[]int{10, 13, 16, 21, 30, 45, 68}, []int{3, 3, 5, 9, 15, 23}, []int{0, 2, 4, 6, 8}, []int{2, 2, 2, 2}, []int{0, 0, 0}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := repeatdiff(tt.args.seq); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("repeatdiff() = %v, want %v", got, tt.want)
			}
		})
	}
}
