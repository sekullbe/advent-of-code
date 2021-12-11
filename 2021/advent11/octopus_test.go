package main

import "testing"

func Test_octopus_flash(t *testing.T) {
	type fields struct {
		energy  int
		flashed bool
	}
	tests := []struct {
		name        string
		fields      fields
		want        int
		finalEnergy int
	}{
		{
			name:        "1f",
			fields:      fields{energy: 2, flashed: false},
			want:        0,
			finalEnergy: 2,
		},
		{
			name:        "9f",
			fields:      fields{energy: 10, flashed: false},
			want:        1,
			finalEnergy: 10,
		},
		{
			name:        "9t",
			fields:      fields{energy: 10, flashed: true},
			want:        0,
			finalEnergy: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &octopus{
				energy:  tt.fields.energy,
				flashed: tt.fields.flashed,
			}
			if got := o.flash(); got != tt.want {
				t.Errorf("flash() = %d, want %d", got, tt.want)
			}
			if o.energy != tt.finalEnergy {
				t.Errorf("finalenergy = %d, want %d", o.energy, tt.finalEnergy)
			}
		})
	}
}
