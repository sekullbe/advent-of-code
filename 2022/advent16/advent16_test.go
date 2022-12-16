package main

import (
	"github.com/sekullbe/advent/parsers"
	"reflect"
	"testing"
)

var testinput string = `Valve AA has flow rate=0; tunnels lead to valves DD, II, BB
Valve BB has flow rate=13; tunnels lead to valves CC, AA
Valve CC has flow rate=2; tunnels lead to valves DD, BB
Valve DD has flow rate=20; tunnels lead to valves CC, AA, EE
Valve EE has flow rate=3; tunnels lead to valves FF, DD
Valve FF has flow rate=0; tunnels lead to valves EE, GG
Valve GG has flow rate=0; tunnels lead to valves FF, HH
Valve HH has flow rate=22; tunnel leads to valve GG
Valve II has flow rate=0; tunnels lead to valves AA, JJ
Valve JJ has flow rate=21; tunnel leads to valve II`

func Test_parseValve(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name string
		args args
		want valve
	}{
		{name: "valve", args: args{"Valve DC has flow rate=25; tunnel leads to valve ST"},
			want: valve{name: "DC", flowRate: 25, leadsTo: nil, leadsToNames: []string{"ST"}}},
		{name: "valves2", args: args{"Valve RU has flow rate=0; tunnels lead to valves YH, ID"},
			want: valve{name: "RU", flowRate: 0, leadsTo: nil, leadsToNames: []string{"YH", "ID"}}},
		{name: "valves5", args: args{"Valve PF has flow rate=10; tunnels lead to valves WK, MZ, QL, XL, LK"},
			want: valve{name: "PF", flowRate: 10, leadsTo: nil, leadsToNames: []string{"WK", "MZ", "QL", "XL", "LK"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseValve(tt.args.line); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseValve() = %v, want %v", got, tt.want)
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
		{name: "example", args: args{testinput}, want: 1651},
		{name: "real", args: args{inputText: inputText}, want: 2330},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run1(tt.args.inputText); got != tt.want {
				t.Errorf("run1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func xTest_parseAllValves(t *testing.T) {
	type args struct {
		valveLines []string
	}
	tests := []struct {
		name string
		args args
		want volcano
	}{
		// just had this to observe, no test, typing out a whole vol{} would be boring
		{name: "example", args: args{parsers.SplitByLines(testinput)}, want: newVolcano()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseAllValves(tt.args.valveLines); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseAllValves() = %v, want %v", got, tt.want)
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
		{name: "example", args: args{testinput}, want: 1707},
		//{name: "real", args: args{inputText: inputText}, want: 2675},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run2(tt.args.inputText); got != tt.want {
				t.Errorf("run2() = %v, want %v", got, tt.want)
			}
		})
	}
}
