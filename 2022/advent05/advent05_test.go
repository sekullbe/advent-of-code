package main

import (
	"github.com/sekullbe/advent/parsers"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func Test_parseStacks(t *testing.T) {
	type args struct {
		inputLines []string
	}
	tests := []struct {
		name       string
		args       args
		wantStacks stacks
	}{
		{
			name: "Example1", args: args{parsers.SplitByLinesNoTrim("    [D]    \n[N] [C]    \n[Z] [M] [P]\n 1   2   3 \n\nmove 1 from 1 to 2")},
			wantStacks: stacks{0: {}, 1: {'Z', 'N'}, 2: {'M', 'C', 'D'}, 3: {'P'}, 4: {}, 5: {}, 6: {}, 7: {}, 8: {}, 9: {}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotStacks := parseStacks(tt.args.inputLines); !reflect.DeepEqual(gotStacks, tt.wantStacks) {
				t.Errorf("parseStacks() = %v, want %v", gotStacks, tt.wantStacks)
			}
		})
	}
}

func Test_oneMove(t *testing.T) {
	type args struct {
		stacks stacks
		from   int
		to     int
	}
	tests := []struct {
		name string
		args args
		want stacks
	}{
		{
			name: "example",
			args: args{stacks: stacks{0: {'N', 'Z'}, 1: {'D', 'C', 'M'}, 2: {'P'}}, from: 1, to: 0},
			want: stacks{0: {'N', 'Z', 'M'}, 1: {'D', 'C'}, 2: {'P'}},
		},
		{
			name: "to empty",
			args: args{stacks: stacks{0: {'N', 'Z'}, 1: {'D', 'C', 'M'}, 2: {}}, from: 1, to: 2},
			want: stacks{0: {'N', 'Z'}, 1: {'D', 'C'}, 2: {'M'}},
		},
		{
			name: "from one",
			args: args{stacks: stacks{0: {'N', 'Z'}, 1: {'D', 'C', 'M'}, 2: {'P'}}, from: 2, to: 0},
			want: stacks{0: {'N', 'Z', 'P'}, 1: {'D', 'C', 'M'}, 2: {}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			oneMove(tt.args.stacks, tt.args.from, tt.args.to)
			if !reflect.DeepEqual(tt.args.stacks, tt.want) {
				t.Fail()
			}
		})
	}
}

func Test_parseMove(t *testing.T) {
	type args struct {
		move string
	}
	tests := []struct {
		name        string
		args        args
		wantHowmany int
		wantFrom    int
		wantTo      int
	}{
		{name: "basic", args: args{move: "move 1 from 2 to 1"}, wantHowmany: 1, wantFrom: 2, wantTo: 1},
		{name: "double digits", args: args{move: "move 3 from 10 to 3"}, wantHowmany: 3, wantFrom: 10, wantTo: 3},
		{name: "more digits", args: args{move: "move 12 from 2 to 10"}, wantHowmany: 12, wantFrom: 2, wantTo: 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotHowmany, gotFrom, gotTo := parseMove(tt.args.move)
			if gotHowmany != tt.wantHowmany {
				t.Errorf("parseMove() gotHowmany = %v, want %v", gotHowmany, tt.wantHowmany)
			}
			if gotFrom != tt.wantFrom {
				t.Errorf("parseMove() gotFrom = %v, want %v", gotFrom, tt.wantFrom)
			}
			if gotTo != tt.wantTo {
				t.Errorf("parseMove() gotTo = %v, want %v", gotTo, tt.wantTo)
			}
		})
	}
}

func Test_getTopOfStacks(t *testing.T) {
	type args struct {
		stacks stacks
	}
	tests := []struct {
		name        string
		args        args
		wantToppers string
	}{
		{
			name:        "basic",
			args:        args{stacks: stacks{0: {}, 1: {'N', 'Z'}, 2: {'D', 'C', 'M'}, 3: {'A'}}},
			wantToppers: "ZMA      ",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotToppers := getTopOfStacks(tt.args.stacks); gotToppers != tt.wantToppers {
				t.Errorf("getTopOfStacks() = %v, want %v", gotToppers, tt.wantToppers)
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
		want string
	}{
		{
			name: "Example",
			args: args{inputText: "    [D]    \n[N] [C]    \n[Z] [M] [P]\n 1   2   3 \n\nmove 1 from 2 to 1\nmove 3 from 1 to 3\nmove 2 from 2 to 1\nmove 1 from 1 to 2\n"},
			want: "CMZ      ",
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
		want string
	}{
		{
			name: "Example",
			args: args{inputText: "    [D]    \n[N] [C]    \n[Z] [M] [P]\n 1   2   3 \n\nmove 1 from 2 to 1\nmove 3 from 1 to 3\nmove 2 from 2 to 1\nmove 1 from 1 to 2\n"},
			want: "MCD      ",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run2(tt.args.inputText); got != tt.want {
				t.Errorf("run1() = %v, want %v", got, tt.want)
			}
		})
	}
}

// proof of concept to avoid off-by-one errors in my indexes
func Test_movingStructsAround(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	b := []int{6, 7, 8, 9}
	howmany := 3
	b = append(b, a[len(a)-howmany:]...)
	a = a[0 : len(a)-howmany]
	assert.Len(t, a, 2)
	assert.Equal(t, []int{1, 2}, a)
	assert.Len(t, b, 7)
	assert.Equal(t, []int{6, 7, 8, 9, 3, 4, 5}, b)
}

func Test_multiMove(t *testing.T) {
	type args struct {
		stacks  stacks
		howmany int
		from    int
		to      int
	}
	tests := []struct {
		name string
		args args
		want stacks
	}{
		{
			name: "basic",
			args: args{stacks: stacks{0: {}, 1: {'N', 'Z'}, 2: {'D', 'C', 'M'}, 3: {'A'}}, howmany: 2, from: 2, to: 3},
			want: stacks{0: {}, 1: {'N', 'Z'}, 2: {'D'}, 3: {'A', 'C', 'M'}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			multiMove(tt.args.stacks, tt.args.howmany, tt.args.from, tt.args.to)
			if !reflect.DeepEqual(tt.args.stacks, tt.want) {
				t.Fail()
			}
		})
	}
}
