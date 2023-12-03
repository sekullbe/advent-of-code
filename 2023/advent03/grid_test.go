package main

import (
	"image"
	"testing"
)

func Test_modifyInMap(t *testing.T) {
	m := make(map[int]image.Point)
	m[1] = image.Pt(5, 5)
	p := m[1]
	p.X = 10
	m[1] = p

	q := m[1]
	if q.X != 10 {
		t.Errorf("was %d, wanted 10", q.X)
	}

}

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
