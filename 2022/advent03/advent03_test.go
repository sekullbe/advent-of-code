package main

import (
	"reflect"
	"testing"
)

func Test_divideRucksack(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name   string
		args   args
		wantC1 []rune
		wantC2 []rune
	}{
		{name: "even", args: args{input: "ABCdef"}, wantC1: []rune{'A', 'B', 'C'}, wantC2: []rune{'d', 'e', 'f'}},
		{name: "odd", args: args{input: "ABCdefg"}, wantC1: []rune{'A', 'B', 'C'}, wantC2: []rune{'d', 'e', 'f', 'g'}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotC1, gotC2 := divideRucksack(tt.args.input)
			if !reflect.DeepEqual(gotC1, tt.wantC1) {
				t.Errorf("divideRucksack() gotC1 = %v, want %v", gotC1, tt.wantC1)
			}
			if !reflect.DeepEqual(gotC2, tt.wantC2) {
				t.Errorf("divideRucksack() gotC2 = %v, want %v", gotC2, tt.wantC2)
			}
		})
	}
}

func Test_run1(t *testing.T) {
	type args struct {
		inputText string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example1", args: args{inputText: "vJrwpWtwJgWrhcsFMMfFFhFp\njqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL\nPmmdzqPrVvPwwTWBwg\nwMqvLMZHhHMvwLHjbvcjnnSBnvTQFn\nttgJtRGJQctTZtZT\nCrZsJsPPZsGzwwsLwLmpwMDw\n"}, want: 157,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run1(tt.args.inputText); got != tt.want {
				t.Errorf("run1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_run2(t *testing.T) {
	type args struct {
		inputText string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example1", args: args{inputText: "vJrwpWtwJgWrhcsFMMfFFhFp\njqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL\nPmmdzqPrVvPwwTWBwg\nwMqvLMZHhHMvwLHjbvcjnnSBnvTQFn\nttgJtRGJQctTZtZT\nCrZsJsPPZsGzwwsLwLmpwMDw\n"}, want: 70,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run2(tt.args.inputText); got != tt.want {
				t.Errorf("run2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findSharedItem(t *testing.T) {
	type args struct {
		c1 []rune
		c2 []rune
	}
	tests := []struct {
		name string
		args args
		want rune
	}{
		{name: "d", args: args{c1: []rune{'A', 'B', 'C', 'd'}, c2: []rune{'d', 'e', 'f', 'f'}}, want: 'd'},
		{name: "e", args: args{c1: []rune{'A', 'B', 'e', 'C', 'd'}, c2: []rune{'d', 'D', 'e', 'f', 'f'}}, want: 'e'},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findSharedItem(tt.args.c1, tt.args.c2); got != tt.want {
				t.Errorf("findSharedItem() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_scoreItem(t *testing.T) {
	type args struct {
		item rune
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "a", args: args{item: 'a'}, want: 1},
		{name: "z", args: args{item: 'z'}, want: 26},
		{name: "B", args: args{item: 'B'}, want: 28},
		{name: "Y", args: args{item: 'Y'}, want: 51},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := scoreItem(tt.args.item); got != tt.want {
				t.Errorf("scoreItem() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findSharedItemThree(t *testing.T) {
	type args struct {
		c1 []rune
		c2 []rune
		c3 []rune
	}
	tests := []struct {
		name string
		args args
		want rune
	}{
		{name: "d", args: args{c1: []rune{'A', 'B', 'd'}, c2: []rune{'d', 'B', 'f', 'f'}, c3: []rune{'A', 'f', 'd'}}, want: 'd'},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findSharedItemThree(tt.args.c1, tt.args.c2, tt.args.c3); got != tt.want {
				t.Errorf("findSharedItemThree() = %v, want %v", got, tt.want)
			}
		})
	}
}
