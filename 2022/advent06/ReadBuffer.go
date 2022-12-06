package main

type Readbuffer struct {
	size int
	buf  []rune
}

func NewReadBuffer(size int) *Readbuffer {
	return &Readbuffer{size, make([]rune, size, size)}
}

// returns true if the buffer filled
func (b *Readbuffer) add(r rune) bool {
	b.buf = append(b.buf, r)
	if len(b.buf) > b.size {
		b.buf = b.buf[1:]
		return true
	}
	return false
}

// returns how many r there are in the buffer
func (b *Readbuffer) countRunes(r rune) (count int) {
	for _, r2 := range b.buf {
		if r == r2 {
			count++
		}
	}
	return
}

func (b *Readbuffer) contentsAllUnique() bool {
	for _, r := range b.buf {
		if b.countRunes(r) > 1 {
			return false
		}
	}
	return true
}
