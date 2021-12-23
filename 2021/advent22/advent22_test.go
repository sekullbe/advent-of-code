package main

import (
	_ "embed"
	"github.com/stretchr/testify/assert"
	"testing"
)

//go:embed example1.txt
var example1Text string

//go:embed example2.txt
var example2Text string

func TestRun1_examples(t *testing.T) {

	input := "on x=10..12,y=10..12,z=10..12\non x=11..13,y=11..13,z=11..13\noff x=9..11,y=9..11,z=9..11\non x=10..10,y=10..10,z=10..10"
	on := run1(input)
	assert.Equal(t, 39, on)

	on = run1(example1Text)
	assert.Equal(t, 590784, on)

	input = example2Text
	on = run1(input)
	assert.Equal(t, 474140, on)

}

func TestRun2(t *testing.T) {
	on := run2(example2Text)
	assert.Equal(t, 2758514936282235, on)

}
