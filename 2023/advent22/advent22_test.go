package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

const sample = `1,0,1~1,2,1
0,0,2~2,0,2
0,2,3~2,2,3
0,0,4~0,2,4
2,0,5~2,2,5
0,1,6~2,1,6
1,1,8~1,1,9`

// I keep wanting to use this and it keeps not _quite_ doing what I want.
// OK this time though.
func Test_Scanf(t *testing.T) {
	input := "1,0,1~1,2,1"
	var x1, x2, y1, y2, z1, z2 int
	n, err := fmt.Sscanf(input, "%d,%d,%d~%d,%d,%d", &x1, &y1, &z1, &x2, &y2, &z2)
	if n != 6 || err != nil {
		t.Errorf("well that didn't work")
	}
	assert.Equal(t, 1, x1)
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
		{
			name: "sample",
			args: args{
				input: sample,
			},
			want: 5,
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
