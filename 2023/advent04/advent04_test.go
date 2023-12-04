package main

import (
	"reflect"
	"testing"
)

var sampleInput string = `Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11`

func Test_parseCard(t *testing.T) {
	type args struct {
		card string
	}
	tests := []struct {
		name        string
		args        args
		wantCardNum int
		wantWinners []int
		wantNumbers []int
	}{
		{
			name: "test",
			args: args{
				card: "Card 11: 41 48 83 86 17 | 83 86  6 31 17  9 48 53",
			},
			wantCardNum: 11,
			wantWinners: []int{41, 48, 83, 86, 17},
			wantNumbers: []int{83, 86, 6, 31, 17, 9, 48, 53},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCardNum, gotWinners, gotNumbers := parseCard(tt.args.card)
			if gotCardNum != tt.wantCardNum {
				t.Errorf("parseCard() gotCardNum = %v, want %v", gotCardNum, tt.wantCardNum)
			}
			if !reflect.DeepEqual(gotWinners, tt.wantWinners) {
				t.Errorf("parseCard() gotWinners = %v, want %v", gotWinners, tt.wantWinners)
			}
			if !reflect.DeepEqual(gotNumbers, tt.wantNumbers) {
				t.Errorf("parseCard() gotNumbers = %v, want %v", gotNumbers, tt.wantNumbers)
			}
		})
	}
}

func Test_evaluateOneCard(t *testing.T) {
	type args struct {
		card string
	}
	tests := []struct {
		name      string
		args      args
		wantScore int
	}{
		{name: "sample 1", args: args{card: "Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53"}, wantScore: 8},
		{name: "sample 2", args: args{card: "Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19"}, wantScore: 2},
		{name: "sample 3", args: args{card: "Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1"}, wantScore: 2},
		{name: "sample 4", args: args{card: "Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83"}, wantScore: 1},
		{name: "sample 5", args: args{card: "Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36"}, wantScore: 0},
		{name: "sample 6", args: args{card: "Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11"}, wantScore: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotScore := evaluateOneCard(tt.args.card); gotScore != tt.wantScore {
				t.Errorf("evaluateOneCard() = %v, want %v", gotScore, tt.wantScore)
			}
		})
	}
}

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
			name: "sample",
			args: args{
				inputText: sampleInput,
			},
			want: 13,
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
