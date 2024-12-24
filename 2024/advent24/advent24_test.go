package main

import (
	_ "embed"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

//go:embed sample1.txt
var sample1 string

//go:embed sample2.txt
var sample2 string

//go:embed sample3.txt
var sample3 string

func Test_parseOneWire(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name string
		args args
		want wire
	}{
		{name: "simple", args: args{line: "x00: 1"}, want: wire{name: "x00", value: 1, valid: false}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseOneWire(tt.args.line); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseOneWire() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_operate(t *testing.T) {
	type args struct {
		op opType
		a  uint8
		b  uint8
	}
	tests := []struct {
		name string
		args args
		want uint8
	}{
		{name: "and1", args: args{op: AND, a: 1, b: 1}, want: 1},
		{name: "and0", args: args{op: AND, a: 1, b: 0}, want: 0},
		{name: "or1", args: args{op: OR, a: 1, b: 1}, want: 1},
		{name: "or1", args: args{op: OR, a: 1, b: 0}, want: 1},
		{name: "or1", args: args{op: OR, a: 0, b: 1}, want: 1},
		{name: "or0", args: args{op: OR, a: 0, b: 0}, want: 0},
		{name: "xor1", args: args{op: XOR, a: 1, b: 0}, want: 1},
		{name: "xor1", args: args{op: XOR, a: 0, b: 1}, want: 1},
		{name: "xor0", args: args{op: XOR, a: 1, b: 1}, want: 0},
		{name: "xor0", args: args{op: XOR, a: 0, b: 0}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, operate(tt.args.op, tt.args.a, tt.args.b), "operate(%v, %v, %v)", tt.args.op, tt.args.a, tt.args.b)
		})
	}
}

func Test_wiresToDecimal(t *testing.T) {
	type args struct {
		wires []wire
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "0", args: args{wires: []wire{{value: 0}}}, want: 0},
		{name: "1", args: args{wires: []wire{{value: 1}}}, want: 1},
		{name: "2", args: args{wires: []wire{{value: 1}, {value: 0}}}, want: 2},
		{name: "4", args: args{wires: []wire{{value: 1}, {value: 0}, {value: 0}}}, want: 4},
		{name: "5", args: args{wires: []wire{{value: 1}, {value: 0}, {value: 1}}}, want: 5},
		{name: "6", args: args{wires: []wire{{value: 1}, {value: 1}, {value: 0}}}, want: 6},
		{name: "7", args: args{wires: []wire{{value: 1}, {value: 1}, {value: 1}}}, want: 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, wiresToDecimal(tt.args.wires), "wiresToDecimal(%v)", tt.args.wires)
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
		{name: "sample1", args: args{input: sample1}, want: 4},
		{name: "sample2", args: args{input: sample2}, want: 2024},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, run1(tt.args.input), "run1(%v)", tt.args.input)
		})
	}
}

func Test_parseOneGate(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name string
		args args
		want *gate
	}{
		{name: "simple", args: args{line: "x00 AND y00 -> z00"}, want: &gate{name: "x00 AND y00 -> z00", valid: false, output: "z00", inputs: []string{"x00", "y00"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, parseOneGate(tt.args.line), "parseOneGate(%v)", tt.args.line)
		})
	}
}
