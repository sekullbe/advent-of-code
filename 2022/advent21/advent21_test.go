package main

import (
	"reflect"
	"testing"
)

var testmonkeys = `root: pppw + sjmn
dbpl: 5
cczh: sllz + lgvd
zczc: 2
ptdq: humn - dvpt
dvpt: 3
lfqf: 4
humn: 5
ljgn: 2
sjmn: drzm * dbpl
sllz: 4
pppw: cczh / lfqf
lgvd: ljgn * ptdq
drzm: hmdt - zczc
hmdt: 32`

/*
	{name: "1 number", args: args{lines: []string{"abcd: 17"}}, want: formulas{}, want1: numbers{"abcd": 17}},
	{name: "1 formula", args: args{lines: []string{"zxcv: qwer + tyui"}},
		want: formulas{"qwer": &formula{yeller: "zxcv", op1: "qwer", op2: "tyui", operator: '+'},
			"tyui": &formula{yeller: "zxcv", op1: "qwer", op2: "tyui", operator: '+'}},
		want1: numbers{}},
	{name: "1 both", args: args{lines: []string{"abcd: 17", "zxcv: qwer + tyui"}},
		want: formulas{"qwer": &formula{yeller: "zxcv", op1: "qwer", op2: "tyui", operator: '+'},
			"tyui": &formula{yeller: "zxcv", op1: "qwer", op2: "tyui", operator: '+'}},
		want1: numbers{"abcd": 17}},
	//scaffolding test to prove that each monkey only appears in one formula
	//{name: "quick duplicate test", args: args{lines: parsers.SplitByLines(inputText)}, want: formulas{}, want1: numbers{}},
*/

func Test_run1(t *testing.T) {
	type args struct {
		inputText string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "sample", args: args{inputText: testmonkeys}, want: 152},
		{name: "live", args: args{inputText: inputText}, want: 21208142603224},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run1(tt.args.inputText); got != tt.want {
				t.Errorf("run1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseMonkeys(t *testing.T) {
	type args struct {
		lines []string
	}
	tests := []struct {
		name  string
		args  args
		want  formulas
		want1 numbers
		want2 *formula
	}{
		{name: "1 number", args: args{lines: []string{"abcd: 17"}}, want: formulas{}, want1: numbers{"abcd": 17}},
		{name: "1 formula", args: args{lines: []string{"zxcv: qwer + tyui"}},
			want: formulas{"qwer": &formula{yeller: "zxcv", op1: "qwer", op2: "tyui", operator: '+'},
				"tyui": &formula{yeller: "zxcv", op1: "qwer", op2: "tyui", operator: '+'}},
			want1: numbers{}},
		{name: "root", args: args{lines: []string{"root: qwer + tyui"}},
			want: formulas{"qwer": &formula{yeller: "root", op1: "qwer", op2: "tyui", operator: '+'},
				"tyui": &formula{yeller: "root", op1: "qwer", op2: "tyui", operator: '+'}},
			want1: numbers{},
			want2: &formula{yeller: "root", op1: "qwer", op2: "tyui", operator: '+'}}, // fix this equal thing
		{name: "1 both", args: args{lines: []string{"abcd: 17", "zxcv: qwer + tyui"}},
			want: formulas{"qwer": &formula{yeller: "zxcv", op1: "qwer", op2: "tyui", operator: '+'},
				"tyui": &formula{yeller: "zxcv", op1: "qwer", op2: "tyui", operator: '+'}},
			want1: numbers{"abcd": 17}},
		//scaffolding test to prove that each monkey only appears in one formula
		//{name: "quick duplicate test", args: args{lines: parsers.SplitByLines(inputText)}, want: formulas{}, want1: numbers{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2 := parseMonkeys(tt.args.lines)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseMonkeys() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("parseMonkeys() got1 = %v, want %v", got1, tt.want1)
			}
			if !reflect.DeepEqual(got2, tt.want2) {
				t.Errorf("parseMonkeys() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}
