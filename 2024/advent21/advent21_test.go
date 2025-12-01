package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const sampleText = "029A\n980A\n179A\n456A\n379A"

/*
my problem is that my keysFromTo for any robot don't generate optimal keysFromTo for the robots upstream
i.e. I generate a >^> when I should have >>^
why doesn't my sorter catch that? do I need more manual sequences?
https://www.reddit.com/r/adventofcode/comments/1hja685/2024_day_21_here_are_some_examples_and_hints_for/


*/

func Test_run1(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "0", args: args{input: "0"}, want: 68 * 29},
		{name: "029A", args: args{input: "029A"}, want: 68 * 29},
		{name: "980A", args: args{input: "980A"}, want: 60 * 980},
		{name: "179A", args: args{input: "179A"}, want: 68 * 179},
		{name: "456A", args: args{input: "456A"}, want: 64 * 456},
		{name: "379A", args: args{input: "379A"}, want: 64 * 379},
		{name: "sample", args: args{input: sampleText}, want: 126384},
		{name: "input", args: args{input: inputText}, want: 188398},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run1(tt.args.input); got != tt.want {
				t.Errorf("run1() = %v, want %v", got, tt.want)
			}
		})
	}
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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run2(tt.args.input); got != tt.want {
				t.Errorf("run2() = %v, want %v", got, tt.want)
			}
		})
	}
}

//  94 ^
// 118 v
//  62 >
//  60 <
//  65 A

func Test_keypad_targetKeypadPresses(t *testing.T) {

	numkp := initKeypad(numKeypad)

	// press 029A from the example
	toPressNum := []rune{'0', '2', '9', 'A'}
	dirMoves := numkp.targetKeypadPresses(toPressNum)
	assert.Equal(t, [][]rune{{'<', 'A'}, {'^', 'A'}, {'^', '^', '>', 'A'}, {'v', 'v', 'v', 'A'}}, dirMoves)

	toPressNum = []rune{'0', '5', '9', 'A'}
	dirMoves = numkp.targetKeypadPresses(toPressNum)
	assert.Equal(t, [][]rune{{'<', 'A'}, {'^', '^', 'A'}, {'^', '>', 'A'}, {'v', 'v', 'v', 'A'}}, dirMoves)

	dirkp := initKeypad(dirKeypad)
	toPressDirs := []rune{'<', 'A'}
	dirMoves = dirkp.targetKeypadPresses(toPressDirs)
	assert.Equal(t, [][]rune{{'v', '<', '<', 'A'}, {'>', '>', '^', 'A'}}, dirMoves)

}

func Test_numericPart(t *testing.T) {
	type args struct {
		code string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "simple", args: args{"029"}, want: 29},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, numericPart(tt.args.code), "numericPart(%v)", tt.args.code)
		})
	}
}

/*
sample 029A
<vA<AA>>^AvAA<^A>A<v<A>>^AvA^A<vA>^A<v<A>^A>AAvA^A<v<A>A>^AAAvA<^A>A
me:
<vA<A>>^AvA<^A>Av<<A>>^AvA^Av<<A>>^A<vA>A^A<A>Av<<A>A^>A<Av>A^A

reddit 456
v<<A^>>AAv<A<A^>>AAvAA^<A>Av<A^>A<A>Av<A^>A<A>Av<<A>A^>AAvA^<A>A
v<<A>>^A<vA<A>>^AvA<^A>A<vA^>A<A>A<vA^>A<A>Av<<A>A^>A<Av>A^A
v<<A>>^A<vA<A>>^AvA<^A>A<vA^>A<A>A<vA^>A<A>Av<<A>A^>A<Av>A^A

my 456

*/

func Test_sortSequences(t *testing.T) {
	type args struct {
		steps []rune
	}
	tests := []struct {
		name string
		args args
		want []rune
	}{
		{name: ">^>", args: args{steps: []rune{'>', '^', '>'}}, want: []rune{'>', '>', '^'}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, sortSequences(tt.args.steps), "sortSequences(%v)", tt.args.steps)
		})
	}
}

func Test_press(t *testing.T) {
	nkp := initKeypad(numKeypad)
	presses := nkp.press('4')
	assert.Equal(t, []rune{'^', '^', '<', '<', 'A'}, presses)

	dkp := initKeypad(dirKeypad)
	presses = dkp.press('r')

}
