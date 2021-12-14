package main

import (
	"github.com/sekullbe/advent/parsers"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func Test_initializePairs(t *testing.T) {
	type args struct {
		template string
	}
	tests := []struct {
		name string
		args args
		want pairCount
	}{
		{
			name: "example1",
			args: args{template: "NNCB"},
			want: pairCount{elementPair{"N", "N"}: 1, elementPair{"N", "C"}: 1, elementPair{"C", "B"}: 1},
		},
		{
			name: "example2",
			args: args{template: "NCNBCHB"},
			want: pairCount{
				elementPair{"N", "C"}: 1,
				elementPair{"C", "N"}: 1,
				elementPair{"N", "B"}: 1,
				elementPair{"B", "C"}: 1,
				elementPair{"C", "H"}: 1,
				elementPair{"H", "B"}: 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := initializePairs(tt.args.template); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("initializePairs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_polymerize(t *testing.T) {
	type args struct {
		template string
		rules    rules
		steps    int
	}
	tests := []struct {
		name string
		args args
		want pairCount
	}{
		{
			name: "norules1",
			args: args{template: "NNCB", rules: rules{}, steps: 1},
			want: pairCount{elementPair{"N", "N"}: 1, elementPair{"N", "C"}: 1, elementPair{"C", "B"}: 1},
		},
		{
			name: "norules2",
			args: args{template: "NCNBCHB", rules: rules{}, steps: 1},
			want: pairCount{elementPair{"N", "C"}: 1, elementPair{"C", "N"}: 1, elementPair{"N", "B"}: 1, elementPair{"B", "C"}: 1, elementPair{"C", "H"}: 1, elementPair{"H", "B"}: 1},
		},
		{
			name: "example 1 rule",
			args: args{template: "NNCB", rules: rules{elementPair{"N", "N"}: "A"}, steps: 1},
			want: pairCount{
				elementPair{"N", "A"}: 1, elementPair{"A", "N"}: 1,
				elementPair{"N", "C"}: 1, elementPair{"C", "B"}: 1,
			},
		},
		{
			name: "example 1 rule adds existing elementPair and repeats replaced elementPair: NNCB -> NNNCB",
			args: args{template: "NNCB", rules: rules{elementPair{"N", "C"}: "N"}, steps: 1},
			want: pairCount{elementPair{"N", "N"}: 2,
				elementPair{"N", "C"}: 1, elementPair{"C", "B"}: 1,
			},
		},
		{
			name: "example,  full ruleset, 1 step NCNBCHB",
			args: args{template: "NNCB", rules: parseRules(parsers.SplitByLines("CH -> B\nHH -> N\nCB -> H\nNH -> C\nHB -> C\nHC -> B\nHN -> C\nNN -> C\nBH -> H\nNC -> B\nNB -> B\nBN -> B\nBB -> N\nBC -> B\nCC -> N\nCN -> C")),
				steps: 1},
			want: pairCount{
				elementPair{"N", "C"}: 1,
				elementPair{"C", "N"}: 1,
				elementPair{"N", "B"}: 1,
				elementPair{"B", "C"}: 1,
				elementPair{"C", "H"}: 1,
				elementPair{"H", "B"}: 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := polymerize(tt.args.template, tt.args.rules, tt.args.steps); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("initializePairs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_run(t *testing.T) {
	type args struct {
		inputText string
		steps     int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example 0 ",
			args: args{inputText: "NNCB\n\nCH -> B\nHH -> N\nCB -> H\nNH -> C\nHB -> C\nHC -> B\nHN -> C\nNN -> C\nBH -> H\nNC -> B\nNB -> B\nBN -> B\nBB -> N\nBC -> B\nCC -> N\nCN -> C",
				steps: 0},
			want: 1,
		},
		{
			name: "example1",
			args: args{inputText: "NNCB\n\nCH -> B\nHH -> N\nCB -> H\nNH -> C\nHB -> C\nHC -> B\nHN -> C\nNN -> C\nBH -> H\nNC -> B\nNB -> B\nBN -> B\nBB -> N\nBC -> B\nCC -> N\nCN -> C", steps: 1},
			want: 1,
		},
		{
			name: "example 10 steps",
			args: args{inputText: "NNCB\n\nCH -> B\nHH -> N\nCB -> H\nNH -> C\nHB -> C\nHC -> B\nHN -> C\nNN -> C\nBH -> H\nNC -> B\nNB -> B\nBN -> B\nBB -> N\nBC -> B\nCC -> N\nCN -> C",
				steps: 10},
			want: 1588,
		},
		{
			name: "real data 0 steps OKSBBKHFBPVNOBKHBPCO",
			args: args{inputText: inputText, steps: 0},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, run(tt.args.inputText, tt.args.steps), "run1(%v, %v)", tt.args.inputText, tt.args.steps)
		})
	}
}

func Test_countElementsInPairs(t *testing.T) {
	type args struct {
		pc    pairCount
		first string
	}
	tests := []struct {
		name string
		args args
		want map[string]int
	}{
		{
			name: "example NNCB",
			args: args{pc: pairCount{elementPair{"N", "N"}: 1, elementPair{"N", "C"}: 1, elementPair{"C", "B"}: 1},
				first: "N",
			},
			want: map[string]int{"N": 2, "C": 1, "B": 1},
		},
		{
			name: "example NBCCNBBBCBHCB",
			args: args{pc: initializePairs("NBCCNBBBCBHCB"), first: "N"},
			want: map[string]int{"N": 2, "C": 4, "B": 6, "H": 1},
		},
		{
			name: "real template OKSBBKHFBPVNOBKHBPCO - did not count last O",
			args: args{pc: initializePairs("OKSBBKHFBPVNOBKHBPCO"), first: "O"},
			want: map[string]int{"B": 5, "C": 1, "F": 1, "H": 2, "K": 3, "N": 1, "O": 3, "P": 2, "S": 1, "V": 1},
		},
		{
			name: "first and last match undercounts NBCN",
			args: args{pc: initializePairs("NBCN"), first: "N"},
			want: map[string]int{"N": 2, "C": 1, "B": 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, countElementsInPairs(tt.args.pc, tt.args.first), "countElementsInPairs(%v)", tt.args.pc)
		})
	}
}
