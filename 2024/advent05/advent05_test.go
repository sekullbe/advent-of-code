package main

import (
	_ "embed"
	"reflect"
	"testing"
)

//go:embed sample.txt
var sampleText string

func Test_run1(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "sampletext", args: args{input: sampleText}, want: 143},
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
		{name: "sampletext", args: args{input: sampleText}, want: 123},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run2(tt.args.input); got != tt.want {
				t.Errorf("run2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseRules(t *testing.T) {
	type args struct {
		ruleLines []string
	}
	tests := []struct {
		name string
		args args
		want []rule
	}{
		{name: "basic", args: args{ruleLines: []string{"1|2", "2|3", "10|20"}}, want: []rule{{1, 2}, {2, 3}, {10, 20}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseRules(tt.args.ruleLines); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseRules() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseUpdates(t *testing.T) {
	type args struct {
		updateLines []string
	}
	tests := []struct {
		name string
		args args
		want []update
	}{
		{name: "simple", args: args{updateLines: []string{"75,47,61,53,29", "97,61,53,29,13"}}, want: []update{{75, 47, 61, 53, 29}, {97, 61, 53, 29, 13}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseUpdates(tt.args.updateLines); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseUpdates() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_testRuleWithRegex(t *testing.T) {
	type args struct {
		rule   rule
		update string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "pass", args: args{rule: rule{1, 2}, update: "3,22,5,1,6,2,99"}, want: true},
		{name: "fail", args: args{rule: rule{1, 2}, update: "3,22,4,2,5,1,99"}, want: false},
		{name: "start", args: args{rule: rule{1, 2}, update: "1,22,4,23,5,2,99"}, want: true},
		{name: "endsT", args: args{rule: rule{1, 2}, update: "1,3,22,4,5,43,99,2"}, want: true},
		{name: "endsF", args: args{rule: rule{1, 2}, update: "2,3,22,4,5,99,1"}, want: false},
		{name: "repeats", args: args{rule: rule{1, 2}, update: "3,5,1111,2,22,4,22222,5,99,1,222222"}, want: false},
		{name: "real1", args: args{rule: rule{97, 75}, update: "75,97,47,61,53"}, want: false},
		{name: "real2", args: args{rule: rule{75, 13}, update: "97,13,75,29,47"}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := testRuleWithRegex(tt.args.rule, tt.args.update); got != tt.want {
				t.Errorf("testRuleWithRegex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_extractMiddleOfUpdate(t *testing.T) {
	type args struct {
		update string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "5", args: args{update: "75,47,61,53,29"}, want: 61},
		{name: "3", args: args{update: "75,47,29"}, want: 47},
		{name: "1", args: args{update: "75"}, want: 75},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := extractMiddleOfUpdate(tt.args.update); got != tt.want {
				t.Errorf("extractMiddleOfUpdate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fixUpdate(t *testing.T) {
	type args struct {
		r rule
		u string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "ex1", args: args{r: rule{97, 75}, u: "75,97,47,61,53"}, want: "97,75,47,61,53"},
		{name: "ex2", args: args{r: rule{29, 13}, u: "61,13,29"}, want: "61,29,13"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fixUpdate(tt.args.r, tt.args.u); got != tt.want {
				t.Errorf("fixUpdate() = %v, want %v", got, tt.want)
			}
		})
	}
}
