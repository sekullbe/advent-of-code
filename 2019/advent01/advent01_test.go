package main

import (
	_ "embed"
	"testing"
)

func Test_calculateFuel(t *testing.T) {
	type args struct {
		mass int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "simple1", args: args{mass: 12}, want: 2},
		{name: "simple2", args: args{mass: 14}, want: 2},
		{name: "simple3", args: args{mass: 1969}, want: 654},
		{name: "simple4", args: args{mass: 100756}, want: 33583},
		{name: "small", args: args{mass: 2}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateFuel(tt.args.mass); got != tt.want {
				t.Errorf("calculateFuel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_calculateFuelAccountingForFuel(t *testing.T) {
	type args struct {
		modMass int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "simple1", args: args{modMass: 12}, want: 2},
		{name: "simple2", args: args{modMass: 14}, want: 2},
		{name: "simple3", args: args{modMass: 1969}, want: 966},
		{name: "simple4", args: args{modMass: 100756}, want: 50346},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateFuelAccountingForFuel(tt.args.modMass); got != tt.want {
				t.Errorf("calculateFuelAccountingForFuel() = %v, want %v", got, tt.want)
			}
		})
	}
}
