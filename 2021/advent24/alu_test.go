package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_alu_div(t *testing.T) {
	a := &alu{7, 17, 11, 13}
	s := step{instruction: "div", rTo: X, rFrom: NOREG, value: 6}
	a.div(s)
	assert.Equal(t, 2, a.get(X))
}
func Test_alu_div2(t *testing.T) {
	a := &alu{7, 18, 11, 13}
	s := step{instruction: "div", rTo: X, rFrom: NOREG, value: 6}
	a.div(s)
	assert.Equal(t, 3, a.get(X))
}
func Test_alu_div_neg(t *testing.T) {
	a := &alu{7, -17, 11, 13}
	s := step{instruction: "div", rTo: X, rFrom: NOREG, value: 6}
	a.div(s)
	assert.Equal(t, -2, a.get(X))
}

func Test_alu_mod(t *testing.T) {
	a := &alu{7, 17, 11, 13}
	s := step{instruction: "mod", rTo: X, rFrom: NOREG, value: 6}
	a.mod(s)
	assert.Equal(t, 5, a.get(X))
}

func Test_alu_mod2(t *testing.T) {
	a := &alu{7, 18, 11, 13}
	s := step{instruction: "mod", rTo: X, rFrom: NOREG, value: 6}
	a.mod(s)
	assert.Equal(t, 0, a.get(X))
}

func Test_alu_mod_neg2(t *testing.T) {
	a := &alu{7, 18, 11, 13}
	s := step{instruction: "mod", rTo: X, rFrom: NOREG, value: -6}
	a.mod(s)
	assert.Equal(t, 0, a.get(X))
}

func Test_alu_eql(t *testing.T) {
	a := &alu{7, 17, 11, 17}
	s := step{instruction: "eql", rTo: X, rFrom: NOREG, value: 17}
	a.eql(s)
	assert.Equal(t, 1, a.get(X))

	a = &alu{7, 17, 11, 17}
	s = step{instruction: "eql", rTo: X, rFrom: NOREG, value: 8}
	a.eql(s)
	assert.Equal(t, 0, a.get(X))

	a = &alu{7, 17, 11, 17}
	s = step{instruction: "eql", rTo: X, rFrom: Z}
	a.eql(s)
	assert.Equal(t, 1, a.get(X))

	a = &alu{7, 17, 11, 17}
	s = step{instruction: "eql", rTo: X, rFrom: Y}
	a.eql(s)
	assert.Equal(t, 0, a.get(X))
}

func Test_alu_inp(t *testing.T) {
	a := &alu{0, 0, 0, 0}
	s := step{instruction: "inp", rTo: X}
	inputs := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 1, 2, 3, 4, 5}
	assert.Len(t, inputs, 14)
	a.inp(s, &inputs)
	assert.Equal(t, 1, a.get(X))
	assert.Len(t, inputs, 13)
	assert.Equal(t, []int{2, 3, 4, 5, 6, 7, 8, 9, 1, 2, 3, 4, 5}, inputs)
}
