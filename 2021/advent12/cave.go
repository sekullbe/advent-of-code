package main

import "strings"

type cave struct {
	id        string
	large     bool
	neighbors []*cave
}

type caveSystem map[string]*cave

func newCave(id string) *cave {
	cave := cave{id: id, neighbors: []*cave{}}
	if strings.ToUpper(id) == id {
		cave.large = true
	}
	return &cave
}
func (c cave) isStart() bool {
	return strings.ToLower(c.id) == "start"
}
func (c cave) isEnd() bool {
	return strings.ToLower(c.id) == "end"
}

// just to make syntax cleaner and not mix "c.isStart()" and "c.large"
func (c cave) isSmall() bool {
	return !c.large
}
