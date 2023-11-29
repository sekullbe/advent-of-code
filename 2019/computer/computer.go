package computer

import "fmt"

type Computer struct {
	ptr      int64
	memory   map[int64]int64
	inputs   []int64
	inputPtr int64
	outputs  []string
	relBase  int64
}

func NewComputer(initialState []int64, inputs []int64) Computer {
	c := Computer{
		ptr:      0,
		memory:   make(map[int64]int64),
		inputs:   inputs,
		inputPtr: 0,
		outputs:  []string{},
		relBase:  0,
	}
	for i, v := range initialState {
		c.memory[int64(i)] = v
	}
	return c
}

func (c *Computer) AdvancePointer(steps int) {
	c.ptr += int64(steps)
}
func (c *Computer) Get(ptr int64) int64 {
	return c.memory[ptr]
}
func (c *Computer) GetOffset(offset int) int64 {
	return c.memory[c.ptr+int64(offset)]
}
func (c *Computer) Set(ptr int64, val int64) int64 {
	c.memory[ptr] = val
	return val
}
func (c *Computer) AdjustRelBase(offset int64) {
	c.relBase = c.relBase + offset
}

func (c *Computer) GetMemory() []int64 {
	memslice := make([]int64, len(c.memory))
	for p := range c.memory {
		memslice[p] = c.Get(p)
	}
	return memslice
}

func (c *Computer) getCurrentInput() int64 {
	c.inputPtr = c.inputPtr + 1
	return c.inputs[c.inputPtr-1]
}

func (c *Computer) GetOutputs() []string {
	return c.outputs
}

func breakdownOp(op int64) (opcode, p1mode, p2mode, p3mode int64) {
	opcode = op % 100
	p1mode = op / 100 % 10
	p2mode = op / 1000 % 10
	p3mode = op / 10000 % 10
	return
}

func (c *Computer) Operate() bool {
	opcode, p1mode, p2mode, p3mode := breakdownOp(c.Get(c.ptr))
	_ = p3mode
	// TODO think about moving out the resolves

	switch opcode {
	case 99:
		return false
	case 1:
		v1 := c.ResolveOperand(p1mode, c.GetOffset(1))
		v2 := c.ResolveOperand(p2mode, c.GetOffset(2))
		target := c.ResolveTarget(p3mode, c.GetOffset(3))
		c.Set(target, v1+v2)
		c.AdvancePointer(4)
	case 2:
		v1 := c.ResolveOperand(p1mode, c.GetOffset(1))
		v2 := c.ResolveOperand(p2mode, c.GetOffset(2))
		target := c.ResolveTarget(p3mode, c.GetOffset(3))

		c.Set(target, v1*v2)
		c.AdvancePointer(4)
	case 3:
		target := c.ResolveTarget(p1mode, c.GetOffset(1))
		c.Set(target, c.getCurrentInput())
		c.AdvancePointer(2)
	case 4:
		c.outputs = append(c.outputs, fmt.Sprintf("%d", c.ResolveOperand(p1mode, c.GetOffset(1))))
		c.AdvancePointer(2)
	case 5: // jump if true
		v1 := c.ResolveOperand(p1mode, c.GetOffset(1))
		v2 := c.ResolveOperand(p2mode, c.GetOffset(2))
		if v1 != 0 {
			c.ptr = v2
		} else {
			c.AdvancePointer(3)
		}
	case 6: // jump if false
		v1 := c.ResolveOperand(p1mode, c.GetOffset(1))
		v2 := c.ResolveOperand(p2mode, c.GetOffset(2))
		if v1 == 0 {
			c.ptr = v2
		} else {
			c.AdvancePointer(3)
		}
	case 7: //  1 < 2 -> 3
		v1 := c.ResolveOperand(p1mode, c.GetOffset(1))
		v2 := c.ResolveOperand(p2mode, c.GetOffset(2))
		target := c.ResolveTarget(p3mode, c.GetOffset(3))
		//target := c.GetOffset(3)
		//if p3mode == 2 {
		//	target += c.relBase
		//}
		if v1 < v2 {
			c.Set(target, 1)
		} else {
			c.Set(target, 0)
		}
		c.AdvancePointer(4)
	case 8: //  1 == 2 -> 3
		v1 := c.ResolveOperand(p1mode, c.GetOffset(1))
		v2 := c.ResolveOperand(p2mode, c.GetOffset(2))
		target := c.ResolveTarget(p3mode, c.GetOffset(3))
		if v1 == v2 {
			c.Set(target, 1)
		} else {
			c.Set(target, 0)
		}
		c.AdvancePointer(4)
	case 9:
		v1 := c.ResolveOperand(p1mode, c.GetOffset(1))
		c.AdjustRelBase(v1)
		c.AdvancePointer(2)
	default:
		panic(fmt.Sprintf("Unhandled opcode %d at %d", opcode, c.ptr))
	}

	return true
}

func (c *Computer) RunGetMem() []int64 {
	for c.Operate() {
	}
	return c.GetMemory()
}

func (c *Computer) Run() {
	for c.Operate() {
	}
}

func SetupAndRun(initialState ...int64) {
	c := NewComputer(initialState, []int64{})
	c.Run()
}

func SetupAndRunWithInputs(initialState []int64, inputs []int64) *Computer {
	c := NewComputer(initialState, inputs)
	c.Run()
	return &c
}

func (c *Computer) ResolveOperand(opMode int64, opValue int64) int64 {
	switch opMode {
	case 0:
		return c.Get(opValue)
	case 1:
		return opValue
	case 2:
		return c.Get(opValue + c.relBase)
	default:
		panic(fmt.Sprintf("bad opmode %d", opMode))
	}
}

func (c *Computer) ResolveTarget(opMode int64, opValue int64) int64 {
	switch opMode {
	case 0:
		return opValue
	case 1:
		panic("can't happen")
	case 2:
		return opValue + c.relBase
	default:
		panic(fmt.Sprintf("bad opmode %d", opMode))
	}
}
