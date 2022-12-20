package tools

import (
	"reflect"
	"testing"
)

func TestContains(t *testing.T) {
	type args[T comparable] struct {
		s []T
		v T
	}
	tests := []struct {
		name string
		args args[string] // can't see how to make this more generic
		want bool
	}{
		{name: "f", args: args[string]{s: []string{"foo", "bar"}, v: "baz"}, want: false},
		{name: "t", args: args[string]{s: []string{"foo", "bar"}, v: "foo"}, want: true},
		//{name: "int t", args: args[int]{s: []int{1, 2}, v: 1}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Contains(tt.args.s, tt.args.v); got != tt.want {
				t.Errorf("Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSumSlice(t *testing.T) {
	s := []int{1, 2, 3}
	got := SumSlice(s)
	want := 6
	if got != want {
		t.Errorf("SumSlice() = %v, want %v", got, want)
	}

}

func TestSliceSubtract(t *testing.T) {
	type args[T comparable] struct {
		a []T
		b []T
	}
	type testCase[T comparable] struct {
		name string
		args args[T]
		want []T
	}
	tests := []testCase[string]{
		{name: "simple", args: args[string]{a: []string{"A", "B", "C", "D"}, b: []string{"B", "D"}}, want: []string{"A", "C"}},
		{name: "zero", args: args[string]{a: []string{"A", "B", "C", "D"}, b: []string{}}, want: []string{"A", "B", "C", "D"}},
		{name: "all", args: args[string]{a: []string{"A", "B", "C", "D"}, b: []string{"A", "B", "C", "D"}}, want: []string{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SliceSubtract(tt.args.a, tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SliceSubtract() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKeyExists(t *testing.T) {
	type args[K comparable, V any] struct {
		m map[K]V
		k K
	}
	type testCase[K comparable, V any] struct {
		name string
		args args[K, V]
		want bool
	}
	tests := []testCase[string, bool]{
		{name: "exists", args: args[string, bool]{map[string]bool{"foo": true, "bar": false}, "foo"}, want: true},
		{name: "exists, falseval", args: args[string, bool]{map[string]bool{"foo": true, "bar": false}, "bar"}, want: true},
		{name: "does not exist", args: args[string, bool]{map[string]bool{"foo": true, "bar": false}, "baz"}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := KeyExists(tt.args.m, tt.args.k); got != tt.want {
				t.Errorf("KeyExists() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_Triangular(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name         string
		args         args
		wantTriangle int
	}{
		{name: "1", args: args{1}, wantTriangle: 1},
		{name: "2", args: args{2}, wantTriangle: 3},
		{name: "3", args: args{3}, wantTriangle: 6},
		{name: "4", args: args{4}, wantTriangle: 10},
		{name: "5", args: args{5}, wantTriangle: 15},
		{name: "6", args: args{6}, wantTriangle: 21},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotTriangle := Triangular(tt.args.n); gotTriangle != tt.wantTriangle {
				t.Errorf("triangular() = %v, want %v", gotTriangle, tt.wantTriangle)
			}
		})
	}
}
