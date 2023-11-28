package computer

import "fmt"

type Computer struct {
	ptr    int
	memory map[int]int
}

func NewComputer(initialState []int) Computer {
	c := Computer{
		ptr:    0,
		memory: make(map[int]int),
	}
	for i, v := range initialState {
		c.memory[i] = v
	}
	return c
}

func (c *Computer) OpcodeAt(ptr int) int {
	return c.memory[ptr]
}
func (c *Computer) OperandOneAt(ptr int) int {
	return c.memory[ptr+1]
}
func (c *Computer) OperandTwoAt(ptr int) int {
	return c.memory[ptr+2]
}
func (c *Computer) DestinationAt(ptr int) int {
	return c.memory[ptr+3]
}
func (c *Computer) Next() {
	c.ptr += 4
}
func (c *Computer) Get(ptr int) int {
	return c.memory[ptr]
}
func (c *Computer) Set(ptr int, val int) int {
	c.memory[ptr] = val
	return val
}
func (c *Computer) Add(ptr int) int {
	return c.Set(c.DestinationAt(ptr), c.Get(c.OperandOneAt(ptr))+c.Get(c.OperandTwoAt(ptr)))
}
func (c *Computer) Mult(ptr int) int {
	return c.Set(c.DestinationAt(ptr), c.Get(c.OperandOneAt(ptr))*c.Get(c.OperandTwoAt(ptr)))
}
func (c *Computer) GetMemory() []int {
	memslice := make([]int, len(c.memory))
	for p := range c.memory {
		memslice[p] = c.Get(p)
	}
	return memslice
}

func (c *Computer) Operate() bool {
	opcode := c.OpcodeAt(c.ptr)
	switch opcode {
	case 99:
		return false
	case 1:
		c.Add(c.ptr)
	case 2:
		c.Mult(c.ptr)
	default:
		panic(fmt.Sprintf("Unhandled opcode %d at %d", opcode, c.ptr))
	}
	c.Next()
	return true
}

func (c *Computer) Run() []int {
	for c.Operate() {
	}
	return c.GetMemory()
}

func SetupAndRun(initialState ...int) []int {
	c := NewComputer(initialState)
	return c.Run()
}
