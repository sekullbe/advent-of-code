package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_run1(t *testing.T) {
	type args struct {
		inputText string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example",
			args: args{inputText: "3,4,3,1,2"},
			want: 5934,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run1(tt.args.inputText); got != tt.want {
				t.Errorf("run1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fish_iterate(t *testing.T) {

	a := assert.New(t)

	var f fish = 8
	newborn := f.iterate()
	a.Nil(newborn)
	a.Equal(fish(7), f)

	newborn = f.iterate()
	a.Nil(newborn)
	a.Equal(fish(6), f)

	newborn = f.iterate()
	a.Nil(newborn)
	a.Equal(fish(5), f)

	newborn = f.iterate()
	a.Nil(newborn)
	a.Equal(fish(4), f)

	newborn = f.iterate()
	a.Nil(newborn)
	a.Equal(fish(3), f)

	newborn = f.iterate()
	a.Nil(newborn)
	a.Equal(fish(2), f)

	newborn = f.iterate()
	a.Nil(newborn)
	a.Equal(fish(1), f)

	newborn = f.iterate()
	a.Nil(newborn)
	a.Equal(fish(0), f)

	newborn = f.iterate()
	a.NotNil(newborn)
	nb := *newborn
	a.Equal(fish(8), nb)
	a.Equal(fish(6), f)

}
