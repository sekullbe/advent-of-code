package main

import (
	"testing"
)

func Test_isSymbol(t *testing.T) {
	type args struct {
		r rune
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "blank", args: args{r: '.'}, want: false},
		{name: "number", args: args{r: '1'}, want: false},
		{name: "*", args: args{r: '*'}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSymbol(tt.args.r); got != tt.want {
				t.Errorf("isSymbol() = %v, want %v", got, tt.want)
			}
		})
	}
}
