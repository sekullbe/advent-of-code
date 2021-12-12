package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_exploreNeighbor(t *testing.T) {

	_, startCave := parseCaveSystem("start-A\nstart-b\nA-c\nA-b\nb-d\nA-end\nb-end")
	assert.Equal(t, 10, exploreNeighbors(*startCave, map[string]bool{"start": true}, []string{"start"}))

	_, startCave = parseCaveSystem("fs-end\nhe-DX\nfs-he\nstart-DX\npj-DX\nend-zg\nzg-sl\nzg-pj\npj-he\nRW-he\nfs-DX\npj-RW\nzg-RW\nstart-pj\nhe-WI\nzg-he\npj-fs\nstart-RW")
	assert.Equal(t, 226, exploreNeighbors(*startCave, map[string]bool{"start": true}, []string{"start"}))
}

func Test_exploreNeighbor_trivial(t *testing.T) {
	_, startCave := parseCaveSystem("start-a\na-b\na-c\nb-c\nc-end\nb-end")
	assert.Equal(t, 4, exploreNeighbors(*startCave, map[string]bool{"start": true}, []string{"start"}))

}

func Test_exploreNeighbor2_trivial(t *testing.T) {
	_, startCave := parseCaveSystem("start-a\na-b\na-c\nb-c\nc-end\nb-end")
	assert.Equal(t, 8, exploreNeighbors2(*startCave, map[string]int{"start": 1}, []string{"start"}, false))

}

func Test_exploreNeighbor2(t *testing.T) {
	_, startCave := parseCaveSystem("start-A\nstart-b\nA-c\nA-b\nb-d\nA-end\nb-end")
	assert.Equal(t, 36, exploreNeighbors2(*startCave, map[string]int{"start": 1}, []string{"start"}, false))

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
		// TODO: Add test cases.
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
