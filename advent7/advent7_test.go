package advent7

import (
	"reflect"
	"testing"
)

func Test_parseContainerFromRule(t *testing.T) {
	type args struct {
		rule string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name : "example1",
			args: args{rule: "plaid fuchsia bags contain 5 light violet bags, 1 light yellow bag."},
			want: "plaid fuchsia",
		},
	}
		for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseContainerFromRule(tt.args.rule); got != tt.want {
				t.Errorf("parseContainerFromRule() = %v, want %v", got, tt.want)
			}
		})
	}
}


func Test_parseContentsFromRule1(t *testing.T) {
	type args struct {
		rule string
	}
	tests := []struct {
		name string
		args args
		want []contents
	}{
		{
			name: "example1",
			args: args{rule: "plaid fuchsia bags contain 5 light violet bags, 1 light yellow bag."},
			want: []contents{
				{bagNumber: 5, bagName: "light violet"},
				{bagNumber: 1, bagName: "light yellow"},
			},
		},
		{
			name: "example2",
			args: args{rule: "striped purple bags contain 4 dark silver bags, 4 vibrant gray bags, 2 dim bronze bags, 2 clear aqua bags."},
			want: []contents{
				{bagNumber: 4, bagName: "dark silver"},
				{bagNumber: 4, bagName: "vibrant gray"},
				{bagNumber: 2, bagName: "dim bronze"},
				{bagNumber: 2, bagName: "clear aqua"},
			},
		},
		{
			name: "no other bags",
			args: args{rule: "striped purple bags contain no other bags."},
			want: []contents{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseContentsFromRule(tt.args.rule); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseContentsFromRule() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRun2_doit(t *testing.T) {
	type args struct {
		theRules string
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "example1",
			args: args{theRules:"light red bags contain 1 bright white bag, 2 muted yellow bags.\ndark orange bags contain 3 bright white bags, 4 muted yellow bags.\nbright white bags contain 1 shiny gold bag.\nmuted yellow bags contain 2 shiny gold bags, 9 faded blue bags.\nshiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.\ndark olive bags contain 3 faded blue bags, 4 dotted black bags.\nvibrant plum bags contain 5 faded blue bags, 6 dotted black bags.\nfaded blue bags contain no other bags.\ndotted black bags contain no other bags."},
			want: 32,
		},
		{
			name: "example2",
			args: args{theRules:"shiny gold bags contain 2 dark red bags.\ndark red bags contain 2 dark orange bags.\ndark orange bags contain 2 dark yellow bags.\ndark yellow bags contain 2 dark green bags.\ndark green bags contain 2 dark blue bags.\ndark blue bags contain 2 dark violet bags.\ndark violet bags contain no other bags."},
			want: 126,
		},
	}
		for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Run2_doit(tt.args.theRules); got != tt.want {
				t.Errorf("Run2_doit() = %v, want %v", got, tt.want)
			}
		})
	}
}
