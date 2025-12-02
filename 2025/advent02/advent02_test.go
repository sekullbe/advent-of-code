package main

import (
	"reflect"
	"testing"
)

const sampleText = "11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124"

func Test_parseranges(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want []idrange
	}{
		{name: "sampleText", args: args{sampleText},
			want: []idrange{{11, 22}, {95, 115}, {998, 1012}, {1188511880, 1188511890}, {222220, 222224},
				{1698522, 1698528}, {446443, 446449}, {38593856, 38593862}, {565653, 565659},
				{824824821, 824824827}, {2121212118, 2121212124}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseranges(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseranges() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_testIdIsValid(t *testing.T) {
	type args struct {
		id int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "single digit", args: args{5}, want: true},
		{name: "2 digit valid", args: args{12}, want: true},
		{name: "2 digit invalid", args: args{11}, want: false},
		{name: "2 digit invalid", args: args{22}, want: false},
		{name: "3 digit valid", args: args{123}, want: true},
		{name: "4 digit valid", args: args{1234}, want: true},
		{name: "4 digit valid", args: args{1212}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := testIdIsValidSingleRepeat(tt.args.id); got != tt.want {
				t.Errorf("testIdIsValidSingleRepeat() = %v, want %v", got, tt.want)
			}
		})
	}
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
		{name: "sampleText", args: args{sampleText}, want: 1227775554},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run1(tt.args.input); got != tt.want {
				t.Errorf("run1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_testIdIsValidMultipleRepeats(t *testing.T) {
	type args struct {
		id int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "single digit", args: args{5}, want: true},
		{name: "123x3", args: args{123123123}, want: false},
		{name: "12x6", args: args{121212121212}, want: false},
		{name: "1x7", args: args{1111111}, want: false},
		{name: "1x7 +2", args: args{11111112}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := testIdIsValidMultipleRepeats(tt.args.id); got != tt.want {
				t.Errorf("testIdIsValidMultipleRepeats() = %v, want %v", got, tt.want)
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
		{name: "sampleText", args: args{sampleText}, want: 4174379265},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run2(tt.args.input); got != tt.want {
				t.Errorf("run2() = %v, want %v", got, tt.want)
			}
		})
	}
}
