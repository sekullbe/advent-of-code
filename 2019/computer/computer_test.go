package computer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testInput = []int64{1, 9, 10, 3, 2, 3, 11, 0, 99, 30, 40, 50}

func Test_ops(t *testing.T) {

	c := NewComputer(testInput, []int64{})
	c.Operate()
	assert.Equal(t, int64(4), c.ptr)
	assert.Equal(t, int64(70), c.Get(3))
	c.Operate()
	assert.Equal(t, int64(8), c.ptr)
	assert.Equal(t, int64(3500), c.Get(0))
	assert.Equal(t, []int64{3500, 9, 10, 70, 2, 3, 11, 0, 99, 30, 40, 50}, c.GetMemory())
}

/*
func Test_FinalStates(t *testing.T) {
	assert.Equal(t, []int64{2, 0, 0, 0, 99}, SetupAndRun(1, 0, 0, 0, 99))
	assert.Equal(t, []int64{2, 3, 0, 6, 99}, SetupAndRun(2, 3, 0, 3, 99))
	assert.Equal(t, []int64{2, 4, 4, 5, 99, 9801}, SetupAndRun(2, 4, 4, 5, 99, 0))
	assert.Equal(t, []int64{30, 1, 1, 4, 2, 5, 6, 0, 99}, SetupAndRun(1, 1, 1, 4, 99, 5, 6, 0, 99))
}
*/

func Test_echo(t *testing.T) {
	c := NewComputer([]int64{3, 0, 4, 0, 99}, []int64{42})
	c.Run()
	assert.Equal(t, []string{"42"}, c.outputs)
}

func Test_breakdownOp(t *testing.T) {

	type args struct {
		op int64
	}
	tests := []struct {
		name       string
		args       args
		wantOpcode int64
		wantP1mode int64
		wantP2mode int64
		wantP3mode int64
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
		{
			name:       "22205",
			args:       args{op: 22205},
			wantOpcode: 5,
			wantP1mode: 2,
			wantP2mode: 2,
			wantP3mode: 2,
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

func Test_jumps_positionMode(t *testing.T) {
	c := SetupAndRunWithInputs([]int64{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, -1, 0, 1, 9}, []int64{0})
	assert.Equal(t, []string{"0"}, c.outputs)
	c = SetupAndRunWithInputs([]int64{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, -1, 0, 1, 9}, []int64{1})
	assert.Equal(t, []string{"1"}, c.outputs)
}

func Test_jumps_immediateMode(t *testing.T) {
	c := SetupAndRunWithInputs([]int64{3, 3, 1105, -1, 9, 1101, 0, 0, 12, 4, 12, 99, 1}, []int64{0})
	assert.Equal(t, []string{"0"}, c.outputs)
	c = SetupAndRunWithInputs([]int64{3, 3, 1105, -1, 9, 1101, 0, 0, 12, 4, 12, 99, 1}, []int64{1})
	assert.Equal(t, []string{"1"}, c.outputs)

}

func Test_ComparesEqual_Position(t *testing.T) {
	c := SetupAndRunWithInputs([]int64{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8}, []int64{8})
	assert.Equal(t, []string{"1"}, c.outputs)
	c = SetupAndRunWithInputs([]int64{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8}, []int64{42})
	assert.Equal(t, []string{"0"}, c.outputs)
}

func Test_ComparesLessThan_Position(t *testing.T) {
	c := SetupAndRunWithInputs([]int64{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8}, []int64{7})
	assert.Equal(t, []string{"1"}, c.outputs)
	c = SetupAndRunWithInputs([]int64{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8}, []int64{8})
	assert.Equal(t, []string{"0"}, c.outputs)
	c = SetupAndRunWithInputs([]int64{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8}, []int64{9})
	assert.Equal(t, []string{"0"}, c.outputs)
}

func Test_ComparesEqual_Immediate(t *testing.T) {
	c := SetupAndRunWithInputs([]int64{3, 3, 1108, -1, 8, 3, 4, 3, 99}, []int64{8})
	assert.Equal(t, []string{"1"}, c.outputs)
	c = SetupAndRunWithInputs([]int64{3, 3, 1108, -1, 8, 3, 4, 3, 99}, []int64{42})
	assert.Equal(t, []string{"0"}, c.outputs)
}

func Test_ComparesLessThan_Immediate(t *testing.T) {
	c := SetupAndRunWithInputs([]int64{3, 3, 1107, -1, 8, 3, 4, 3, 99}, []int64{7})
	assert.Equal(t, []string{"1"}, c.outputs)
	c = SetupAndRunWithInputs([]int64{3, 3, 1107, -1, 8, 3, 4, 3, 99}, []int64{8})
	assert.Equal(t, []string{"0"}, c.outputs)
	c = SetupAndRunWithInputs([]int64{3, 3, 1107, -1, 8, 3, 4, 3, 99}, []int64{9})
	assert.Equal(t, []string{"0"}, c.outputs)
}

func Test_JumpAndCompare(t *testing.T) {
	bigprog := []int64{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31,
		1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104,
		999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99}
	c := SetupAndRunWithInputs(bigprog, []int64{7})
	assert.Equal(t, []string{"999"}, c.outputs)
	c = SetupAndRunWithInputs(bigprog, []int64{8})
	assert.Equal(t, []string{"1000"}, c.outputs)
	c = SetupAndRunWithInputs(bigprog, []int64{9})
	assert.Equal(t, []string{"1001"}, c.outputs)
}

func Test_RelativeMode(t *testing.T) {
	c := SetupAndRunWithInputs([]int64{109, 1, 204, -1, 1001, 100, 1, 100, 1008, 100, 16, 101, 1006, 101, 0, 99}, []int64{})
	assert.Equal(t, []string{"109", "1", "204", "-1", "1001", "100", "1", "100", "1008", "100", "16", "101", "1006", "101", "0", "99"}, c.outputs)

	c = SetupAndRunWithInputs([]int64{1102, 34915192, 34915192, 7, 4, 7, 99, 0}, []int64{})
	assert.Equal(t, "1219070632396864", c.outputs[0])
	assert.Equal(t, 16, len(c.outputs[0]))

	c = SetupAndRunWithInputs([]int64{104, 1125899906842624, 99}, []int64{})
	assert.Equal(t, "1125899906842624", c.outputs[0])

	c = SetupAndRunWithInputs([]int64{109, -1, 4, 1, 99}, []int64{})
	assert.Equal(t, "-1", c.outputs[0])

	c = SetupAndRunWithInputs([]int64{109, -1, 104, 1, 99}, []int64{})
	assert.Equal(t, "1", c.outputs[0])
	c = SetupAndRunWithInputs([]int64{109, -1, 204, 1, 99}, []int64{})
	assert.Equal(t, "109", c.outputs[0])
	c = SetupAndRunWithInputs([]int64{109, 1, 9, 2, 204, -6, 99}, []int64{})
	assert.Equal(t, "204", c.outputs[0])
	c = SetupAndRunWithInputs([]int64{109, 1, 109, 9, 204, -6, 99}, []int64{})
	assert.Equal(t, "204", c.outputs[0])
	c = SetupAndRunWithInputs([]int64{109, 1, 209, -1, 204, -106, 99}, []int64{})
	assert.Equal(t, "204", c.outputs[0])
	c = SetupAndRunWithInputs([]int64{109, 1, 3, 3, 204, 2, 99}, []int64{42})
	assert.Equal(t, "42", c.outputs[0])
}

func Test_thisOneFails(t *testing.T) {
	c := SetupAndRunWithInputs([]int64{109, 1, 203, 2, 204, 2, 99}, []int64{42})
	assert.Equal(t, "42", c.outputs[0])
}
