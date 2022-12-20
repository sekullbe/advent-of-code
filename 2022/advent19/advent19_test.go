package main

import (
	"reflect"
	"testing"
)

var testInput = `Blueprint 1: Each ore robot costs 4 ore. Each clay robot costs 2 ore. Each obsidian robot costs 3 ore and 14 clay. Each geode robot costs 2 ore and 7 obsidian.
Blueprint 2: Each ore robot costs 2 ore. Each clay robot costs 3 ore. Each obsidian robot costs 3 ore and 8 clay. Each geode robot costs 3 ore and 12 obsidian.
Blueprint 3: Each ore robot costs 2 ore. Each clay robot costs 3000 ore. Each obsidian robot costs 3 ore and 8 clay. Each geode robot costs 3 ore and 12 obsidian.`

func Test_parseBlueprint(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name   string
		args   args
		wantBp blueprint
	}{
		{name: "live 2",
			args: args{line: "Blueprint 2: Each ore robot costs 4 ore. Each clay robot costs 3 ore. Each obsidian robot costs 2 ore and 5 clay. Each geode robot costs 2 ore and 10 obsidian."},
			wantBp: blueprint{
				number:   2,
				orebot:   resources{4, 0, 0, 0},
				claybot:  resources{3, 0, 0, 0},
				obsbot:   resources{2, 5, 0, 0},
				geodebot: resources{2, 0, 10, 0},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotBp := parseBlueprint(tt.args.line); !reflect.DeepEqual(gotBp, tt.wantBp) {
				t.Errorf("parseBlueprint() = %v, want %v", gotBp, tt.wantBp)
			}
		})
	}
}

func Test_run1(t *testing.T) {
	type args struct {
		inputText string
		ticks     int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "example", args: args{inputText: testInput, ticks: 24}, want: 33},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run1(tt.args.inputText, tt.args.ticks); got != tt.want {
				t.Errorf("run1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_run2(t *testing.T) {
	type args struct {
		inputText string
		ticks     int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "example", args: args{inputText: testInput, ticks: 32}, want: 0}, // fails
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run2(tt.args.inputText, tt.args.ticks); got != tt.want {
				t.Errorf("run2() = %v, want %v", got, tt.want)
			}
		})
	}
}
