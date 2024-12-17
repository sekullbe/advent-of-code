package main

import (
	"github.com/sekullbe/advent/tools"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

const sampleText = `Register A: 729
Register B: 0
Register C: 0

Program: 0,1,5,4,3,0`

const sampleText2 = `Register A: 729
Register B: 111
Register C: 222

Program: 0,1,5,4,3,0,555,111`

const sampleText_part2 = `Register A: 2024
Register B: 0
Register C: 0

Program: 0,3,5,4,3,0`

func Test_initialize(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want *computer
	}{
		{name: "sample", args: args{sampleText}, want: &computer{
			program: []int{0, 1, 5, 4, 3, 0},
			rA:      729,
			rB:      0,
			rC:      0,
			instPtr: 0,
		}},
		{name: "sample2", args: args{sampleText2}, want: &computer{
			program: []int{0, 1, 5, 4, 3, 0, 555, 111},
			rA:      729,
			rB:      111,
			rC:      222,
			instPtr: 0,
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := initialize(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("initialize() = %v, want %v", got, tt.want)
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
		want string
	}{
		{name: "sample", args: args{input: sampleText}, want: "4,6,3,5,6,3,5,2,1,0"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run1(tt.args.input); got != tt.want {
				t.Errorf("run1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_instructions1(t *testing.T) {
	cpu := &computer{program: []int{2, 6}, rA: 999, rB: 999, rC: 9, instPtr: 0}
	run(cpu)
	assert.Equal(t, 1, cpu.rB)
}

func Test_instructions2(t *testing.T) {
	cpu := &computer{program: []int{5, 0, 5, 1, 5, 4}, rA: 10, rB: 999, rC: 999, instPtr: 0}
	output := run(cpu)
	assert.Equal(t, "0,1,2", output)
}

func Test_instructions3(t *testing.T) {
	cpu := &computer{program: []int{0, 1, 5, 4, 3, 0}, rA: 2024, rB: 999, rC: 9, instPtr: 0}
	output := run(cpu)
	assert.Equal(t, 0, cpu.rA)
	assert.Equal(t, "4,2,5,6,7,7,7,7,3,1,0", output)
}

func Test_instructions4(t *testing.T) {
	cpu := &computer{program: []int{1, 7}, rA: 10, rB: 29, rC: 999, instPtr: 0}
	run(cpu)
	assert.Equal(t, 26, cpu.rB)
}

func Test_instructions5(t *testing.T) {
	cpu := &computer{program: []int{4, 0}, rA: 999, rB: 2024, rC: 43690, instPtr: 0}
	run(cpu)
	assert.Equal(t, 44354, cpu.rB)
}

func Test_part2_example(t *testing.T) {
	cpu := &computer{program: []int{0, 3, 5, 4, 3, 0}, rA: 117440, rB: 0, rC: 0, instPtr: 0}
	output := run(cpu)
	assert.Equal(t, tools.IntArrayToString(cpu.program), output)

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
		{name: "sample", args: args{input: sampleText_part2}, want: 117440},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, run2(tt.args.input), "run2(%v)", tt.args.input)
		})
	}
}
