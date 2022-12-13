package main

import (
	"encoding/json"
	"log"
	"reflect"
	"testing"
)

var testinput = `[1,1,3,1,1]
[1,1,5,1,1]

[[1],[2,3,4]]
[[1],4]

[9]
[[8,7,6]]

[[4,4],4,4]
[[4,4],4,4,4]

[7,7,7,7]
[7,7,7]

[]
[3]

[[[]]]
[[]]

[1,[2,[3,[4,[5,6,7]]]],8,9]
[1,[2,[3,[4,[5,6,0]]]],8,9]`

func Test_unmarshaling(t *testing.T) {

	packets := [][]any{}

	var packet []any
	err := json.Unmarshal([]byte("[[1],[2,3,4]]"), &packet)
	if err != nil {
		log.Panicf("unmarshal err: %v", err)
	}
	packets = append(packets, packet)

}

// ignore for now
func test_parsePackets(t *testing.T) {
	type args struct {
		packetLines []string
	}
	tests := []struct {
		name string
		args args
		want [][]any
	}{
		{name: "example", args: args{[]string{"[[1],[2,3,4]]"}}, want: [][]any{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parsePackets(tt.args.packetLines); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parsePackets() = %v, want %v", got, tt.want)
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
		{name: "example", args: args{testinput}, want: 13},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run1(tt.args.inputText); got != tt.want {
				t.Errorf("run1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_run2(t *testing.T) {
	type args struct {
		inputText string
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
			if got := run2(tt.args.inputText); got != tt.want {
				t.Errorf("run2() = %v, want %v", got, tt.want)
			}
		})
	}
}
