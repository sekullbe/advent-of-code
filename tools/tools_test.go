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
