package computer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testInput = []int{1, 9, 10, 3, 2, 3, 11, 0, 99, 30, 40, 50}

func Test_ops(t *testing.T) {

	c := NewComputer(testInput)
	c.Operate()
	assert.Equal(t, 4, c.ptr)
	assert.Equal(t, 70, c.Get(3))
	c.Operate()
	assert.Equal(t, 8, c.ptr)
	assert.Equal(t, 3500, c.Get(0))
	assert.Equal(t, []int{3500, 9, 10, 70, 2, 3, 11, 0, 99, 30, 40, 50}, c.GetMemory())
}

func Test_FinalStates(t *testing.T) {
	assert.Equal(t, []int{2, 0, 0, 0, 99}, SetupAndRun(1, 0, 0, 0, 99))
	assert.Equal(t, []int{2, 3, 0, 6, 99}, SetupAndRun(2, 3, 0, 3, 99))
	assert.Equal(t, []int{2, 4, 4, 5, 99, 9801}, SetupAndRun(2, 4, 4, 5, 99, 0))
	assert.Equal(t, []int{30, 1, 1, 4, 2, 5, 6, 0, 99}, SetupAndRun(1, 1, 1, 4, 99, 5, 6, 0, 99))
}
