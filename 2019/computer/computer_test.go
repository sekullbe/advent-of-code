package computer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testInput = []int{1, 9, 10, 3, 2, 3, 11, 0, 99, 30, 40, 50}

func Test_ops(t *testing.T) {

	c := NewComputer(testInput, []int{})
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

func Test_echo(t *testing.T) {
	c := NewComputer([]int{3, 0, 4, 0, 99}, []int{42})
	c.Run()
	assert.Equal(t, []string{"42"}, c.outputs)
}

func Test_breakdownOp(t *testing.T) {

	type args struct {
		op int
	}
	tests := []struct {
		name       string
		args       args
		wantOpcode int
		wantP1mode int
		wantP2mode int
		wantP3mode int
	}{
		{
			name:       "1002",
			args:       args{op: 1002},
			wantOpcode: 2,
			wantP1mode: 0,
			wantP2mode: 1,
			wantP3mode: 0,
		},
		{
			name:       "11005",
			args:       args{op: 11005},
			wantOpcode: 5,
			wantP1mode: 0,
			wantP2mode: 1,
			wantP3mode: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotOpcode, gotP1mode, gotP2mode, gotP3mode := breakdownOp(tt.args.op)
			if gotOpcode != tt.wantOpcode {
				t.Errorf("breakdownOp() gotOpcode = %v, want %v", gotOpcode, tt.wantOpcode)
			}
			if gotP1mode != tt.wantP1mode {
				t.Errorf("breakdownOp() gotP1mode = %v, want %v", gotP1mode, tt.wantP1mode)
			}
			if gotP2mode != tt.wantP2mode {
				t.Errorf("breakdownOp() gotP2mode = %v, want %v", gotP2mode, tt.wantP2mode)
			}
			if gotP3mode != tt.wantP3mode {
				t.Errorf("breakdownOp() gotP3mode = %v, want %v", gotP3mode, tt.wantP3mode)
			}
		})
	}
}

func Test_jumps(t *testing.T) {
	c := SetupAndRunWithInputs([]int{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, -1, 0, 1, 9}, []int{0})
	assert.Equal(t, []string{"0"}, c.outputs)
	c = SetupAndRunWithInputs([]int{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, -1, 0, 1, 9}, []int{1})
	assert.Equal(t, []string{"1"}, c.outputs)

	c = SetupAndRunWithInputs([]int{3, 3, 1105, -1, 9, 1101, 0, 0, 12, 4, 12, 99, 1}, []int{0})
	assert.Equal(t, []string{"0"}, c.outputs)
	c = SetupAndRunWithInputs([]int{3, 3, 1105, -1, 9, 1101, 0, 0, 12, 4, 12, 99, 1}, []int{1})
	assert.Equal(t, []string{"1"}, c.outputs)

}

func Test_ComparesEqual_Position(t *testing.T) {
	c := SetupAndRunWithInputs([]int{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8}, []int{8})
	assert.Equal(t, []string{"1"}, c.outputs)
	c = SetupAndRunWithInputs([]int{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8}, []int{42})
	assert.Equal(t, []string{"0"}, c.outputs)
}

func Test_ComparesLessThan_Position(t *testing.T) {
	c := SetupAndRunWithInputs([]int{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8}, []int{7})
	assert.Equal(t, []string{"1"}, c.outputs)
	c = SetupAndRunWithInputs([]int{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8}, []int{8})
	assert.Equal(t, []string{"0"}, c.outputs)
	c = SetupAndRunWithInputs([]int{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8}, []int{9})
	assert.Equal(t, []string{"0"}, c.outputs)
}

func Test_ComparesEqual_Immediate(t *testing.T) {
	c := SetupAndRunWithInputs([]int{3, 3, 1108, -1, 8, 3, 4, 3, 99}, []int{8})
	assert.Equal(t, []string{"1"}, c.outputs)
	c = SetupAndRunWithInputs([]int{3, 3, 1108, -1, 8, 3, 4, 3, 99}, []int{42})
	assert.Equal(t, []string{"0"}, c.outputs)
}

func Test_ComparesLessThan_Immediate(t *testing.T) {
	c := SetupAndRunWithInputs([]int{3, 3, 1107, -1, 8, 3, 4, 3, 99}, []int{7})
	assert.Equal(t, []string{"1"}, c.outputs)
	c = SetupAndRunWithInputs([]int{3, 3, 1107, -1, 8, 3, 4, 3, 99}, []int{8})
	assert.Equal(t, []string{"0"}, c.outputs)
	c = SetupAndRunWithInputs([]int{3, 3, 1107, -1, 8, 3, 4, 3, 99}, []int{9})
	assert.Equal(t, []string{"0"}, c.outputs)
}

func Test_JumpAndCompare(t *testing.T) {
	bigprog := []int{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31,
		1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104,
		999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99}
	c := SetupAndRunWithInputs(bigprog, []int{7})
	assert.Equal(t, []string{"999"}, c.outputs)
	c = SetupAndRunWithInputs(bigprog, []int{8})
	assert.Equal(t, []string{"1000"}, c.outputs)
	c = SetupAndRunWithInputs(bigprog, []int{9})
	assert.Equal(t, []string{"1001"}, c.outputs)
}
