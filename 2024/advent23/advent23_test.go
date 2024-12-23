package main

import (
	_ "embed"
	"github.com/stretchr/testify/assert"
	"testing"
)

//go:embed sample.txt
var sampleText string

func Test_createNetwork(t *testing.T) {
	n := createNetwork([]string{"ab-cd", "cd-ef", "ab-ef", "gh-ij"})
	assert.Equal(t, 5, len(n))
	assert.Equal(t, 2, n["ab"].links.Cardinality())
	assert.True(t, n["ab"].links.Contains(n["cd"]))
	assert.True(t, n["cd"].links.Contains(n["ab"]))
	assert.False(t, n["cd"].links.Contains(n["gh"]))

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
		{name: "sample", args: args{input: sampleText}, want: 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, run1(tt.args.input), "run1(%v)", tt.args.input)
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
		want string
	}{
		{name: "sample", args: args{input: sampleText}, want: "co,de,ka,ta"},
		{name: "basic", args: args{input: "ab-cd\ncd-ef\nab-ef\ngh-ij"}, want: "ab,cd,ef"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, run2(tt.args.input), "run2(%v)", tt.args.input)
		})
	}
}
