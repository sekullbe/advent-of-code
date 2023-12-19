package main

import (
	"reflect"
	"testing"
)

const sampleText = `
px{a<2006:qkq,m>2090:A,rfg}
pv{a>1716:R,A}
lnx{m>1548:A,A}
rfg{s<537:gd,x>2440:R,A}
qs{s>3448:A,lnx}
qkq{x<1416:A,crn}
crn{x>2662:A,R}
in{s<1351:px,qqz}
qqz{s>2770:qs,m<1801:hdj,R}
gd{a>3333:R,R}
hdj{m>838:A,pv}

{x=787,m=2655,a=1222,s=2876}
{x=1679,m=44,a=2067,s=496}
{x=2036,m=264,a=79,s=2244}
{x=2461,m=1339,a=466,s=291}
{x=2127,m=1623,a=2188,s=1013}`

func Test_run1(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "sample",
			args: args{
				input: sampleText,
			},
			want: 19114,
		},
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

func Test_parseSingleRule(t *testing.T) {
	type args struct {
		rs string
	}
	tests := []struct {
		name string
		args args
		want rule
	}{
		{name: "simple rule", args: args{rs: "a<2006:qkq"}, want: rule{rating: "a", operator: "<", operand: 2006, next: "qkq"}},
		{name: "default", args: args{rs: "rfg"}, want: rule{rating: "", operator: "", operand: 0, next: "rfg"}},
		{name: "default AR", args: args{rs: "A"}, want: rule{rating: "", operator: "", operand: 0, next: "A"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseSingleRule(tt.args.rs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseSingleRule() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseSingleWorkflow(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name string
		args args
		want workflow
	}{
		{
			name: "sample1",
			args: args{
				line: "px{a<2006:qkq,m>2090:A,rfg}",
			},
			want: workflow{
				name: "px",
				rules: []rule{
					rule{rating: "a", operator: "<", operand: 2006, next: "qkq"},
					rule{rating: "m", operator: ">", operand: 2090, next: "A"},
					rule{rating: "", operator: "", operand: 0, next: "rfg"},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseSingleWorkflow(tt.args.line); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseSingleWorkflow() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseSinglePart(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name string
		args args
		want part
	}{
		{
			name: "simple",
			args: args{
				line: "{x=787,m=2655,a=1222,s=2876}",
			},
			want: part{"x": 787, "m": 2655, "a": 1222, "s": 2876},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseSinglePart(tt.args.line); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseSinglePart() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_processWorkflowStep(t *testing.T) {
	rules := []rule{
		rule{rating: "a", operator: "<", operand: 2006, next: "qkq"},
		rule{rating: "m", operator: ">", operand: 2090, next: "A"},
		rule{rating: "", operator: "", operand: 0, next: "rfg"}}

	type args struct {
		w workflow
		p part
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "simple1",
			args: args{
				w: workflow{
					name:  "in",
					rules: rules,
				},
				p: part{"x": 787, "m": 2655, "a": 1222, "s": 2876},
			},
			want: "qkq",
		},
		{
			name: "simple1",
			args: args{
				w: workflow{
					name:  "in",
					rules: rules,
				},
				p: part{"x": 787, "m": 2655, "a": 3000, "s": 2876},
			},
			want: "A",
		},
		{
			name: "simple1",
			args: args{
				w: workflow{
					name:  "in",
					rules: rules,
				},
				p: part{"x": 787, "m": 1000, "a": 3000, "s": 2876},
			},
			want: "rfg",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := processWorkflowStep(tt.args.w, tt.args.p); got != tt.want {
				t.Errorf("processWorkflowStep() = %v, want %v", got, tt.want)
			}
		})
	}
}
