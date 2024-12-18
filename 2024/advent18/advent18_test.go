package main

import (
	"github.com/sekullbe/advent/geometry"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func Test_parseBytes(t *testing.T) {
	type args struct {
		lines []string
	}
	tests := []struct {
		name string
		args args
		want []geometry.Point2
	}{
		{name: "test1", args: args{lines: []string{"1,2", "3,4}"}}, want: []geometry.Point2{{1, 2}, {3, 4}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseBytes(tt.args.lines); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dropByte(t *testing.T) {
	b := initBoard(5, 5)
	bytes := parseBytes([]string{"0,0", "1,2", "3,4"})
	dropBytes(b, bytes)
	assert.Equal(t, '#', b.At(0, 0).Contents)
	assert.Equal(t, '.', b.At(1, 0).Contents)
}
