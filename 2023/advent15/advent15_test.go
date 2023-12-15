package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const sampleInput = `rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7`

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
				input: "rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7",
			},
			want: 1320,
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

func Test_BytesWorkTheWayIThinkTheyDo(t *testing.T) {
	str := "HASH"
	b := []byte(str)
	assert.Equal(t, 72, int(b[0]))
	assert.Equal(t, 65, int(b[1]))
	assert.Equal(t, 83, int(b[2]))
	assert.Equal(t, 72, int(b[3]))

	i := str[0] // this is a uint8 not an int
	assert.Equal(t, 72, i)
}

func Test_processStep(t *testing.T) {
	type args struct {
		step string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "HASH", args: args{step: "HASH"}, want: 52},
		{name: "rn=1", args: args{step: "rn=1"}, want: 30},
		{name: "cm-", args: args{step: "cm-"}, want: 253},
		{name: "qp=3", args: args{step: "qp=3"}, want: 97},
		{name: "cm=2", args: args{step: "cm=2"}, want: 47},
		{name: "qp-", args: args{step: "qp-"}, want: 14},
		{name: "pc=4", args: args{step: "pc=4"}, want: 180},
		{name: "ot=9", args: args{step: "ot=9"}, want: 9},
		{name: "ab=5", args: args{step: "ab=5"}, want: 197},
		{name: "pc-", args: args{step: "pc-"}, want: 48},
		{name: "pc=6", args: args{step: "pc=6"}, want: 214},
		{name: "ot=7", args: args{step: "ot=7"}, want: 231},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, processStep(tt.args.step), "processStep(%v)", tt.args.step)
		})
	}
}
