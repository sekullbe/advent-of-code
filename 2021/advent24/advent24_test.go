package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_parseLineToStep(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name  string
		args  args
		wantS step
	}{
		{
			name:  "reg",
			args:  args{line: "foo x y"},
			wantS: step{instruction: "foo", rTo: X, rFrom: Y},
		},
		{
			name:  "one arg",
			args:  args{line: "inp x"},
			wantS: step{instruction: "inp", rTo: X},
		},
		{
			name:  "value",
			args:  args{line: "foo x -134"},
			wantS: step{instruction: "foo", rTo: X, value: -134},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.wantS, parseLineToStep(tt.args.line), "parseLineToStep(%v)", tt.args.line)
		})
	}
}

func TestRunExample(t *testing.T) {

	steps := parseInput("inp w\nadd z w\nmod z 2\ndiv w 2\nadd y w\nmod y 2\ndiv w 2\nadd x w\nmod x 2\ndiv w 2\nmod w 2")
	inputs := []int{14}
	a := newAlu()
	for _, s := range steps {
		runStep(a, s, &inputs)
	}
	assert.Equal(t, 1, a.get(W))
	assert.Equal(t, 1, a.get(X))
	assert.Equal(t, 1, a.get(Y))
	assert.Equal(t, 0, a.get(Z))
}

func TestCheckLargest(t *testing.T) {
	steps := parseInput(inputText)
	ok := checkModelNumber(94399898949959, steps) // that is correct, why is my code wrong
	assert.True(t, ok)
}

func TestCheckSmallest(t *testing.T) {
	steps := parseInput(inputText)
	ok := checkModelNumber(21176121611511, steps) // that is correct, why is my code wrong
	assert.True(t, ok)
}
