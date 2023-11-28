package computer

import "fmt"

type Computer struct {
	ptr      int
	memory   map[int]int
	inputs   []int
	inputPtr int
	outputs  []string
}

func NewComputer(initialState []int, inputs []int) Computer {
	c := Computer{
		ptr:      0,
		memory:   make(map[int]int),
		inputs:   inputs,
		inputPtr: 0,
		outputs:  []string{},
	}
	for i, v := range initialState {
		c.memory[i] = v
	}
	return c
}

func (c *Computer) AdvancePointer(steps int) {
	c.ptr += steps
}
func (c *Computer) Get(ptr int) int {
	return c.memory[ptr]
}
func (c *Computer) GetOffset(offset int) int {
	return c.memory[c.ptr+offset]
}
func (c *Computer) Set(ptr int, val int) int {
	c.memory[ptr] = val
	return val
}

func (c *Computer) GetMemory() []int {
	memslice := make([]int, len(c.memory))
	for p := range c.memory {
		memslice[p] = c.Get(p)
	}
	return memslice
}

func (c *Computer) getCurrentInput() int {
	c.inputPtr = c.inputPtr + 1
	return c.inputs[c.inputPtr-1]
}

func (c *Computer) GetOutputs() []string {
	return c.outputs
}

func breakdownOp(op int) (opcode, p1mode, p2mode, p3mode int) {
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
		/// maybe stop using the methods and do it all here? the methods make assumptions about the values
		v1 := c.ResolveOperand(p1mode, c.GetOffset(1))
		v2 := c.ResolveOperand(p2mode, c.GetOffset(2))
		target := c.GetOffset(3)
		c.Set(target, v1+v2)
		c.AdvancePointer(4)
	case 2:
		v1 := c.ResolveOperand(p1mode, c.GetOffset(1))
		v2 := c.ResolveOperand(p2mode, c.GetOffset(2))
		target := c.GetOffset(3)
		c.Set(target, v1*v2)
		c.AdvancePointer(4)
	case 3:
		target := c.GetOffset(1)
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
		target := c.GetOffset(3)
		if v1 < v2 {
			c.Set(target, 1)
		} else {
			c.Set(target, 0)
		}
		c.AdvancePointer(4)
	case 8: //  1 == 2 -> 3
		v1 := c.ResolveOperand(p1mode, c.GetOffset(1))
		v2 := c.ResolveOperand(p2mode, c.GetOffset(2))
		target := c.GetOffset(3)
		if v1 == v2 {
			c.Set(target, 1)
		} else {
			c.Set(target, 0)
		}
		c.AdvancePointer(4)
	default:
		panic(fmt.Sprintf("Unhandled opcode %d at %d", opcode, c.ptr))
	}

	return true
}

func (c *Computer) Run() []int {
	for c.Operate() {
	}
	return c.GetMemory()
}

func SetupAndRun(initialState ...int) []int {
	c := NewComputer(initialState, []int{})
	return c.Run()
}

func SetupAndRunWithInputs(initialState []int, inputs []int) *Computer {
	c := NewComputer(initialState, inputs)
	c.Run()
	return &c
}

func (c *Computer) ResolveOperand(opMode int, opValue int) int {
	switch opMode {
	case 0:
		return c.Get(opValue)
	case 1:
		return opValue
	default:
		panic(fmt.Sprintf("bad opmode %d", opMode))
	}
}
