package main

import (
	"github.com/sekullbe/advent/parsers"
	"strings"
)

const (
	LOW  = false
	HIGH = true
)
const (
	OFF = false
	ON  = true
)

type queuedPulse struct {
	sourceName string
	level      bool
	target     string
}

type pulser interface {
	// input the queuedPulse - the pulser needs to send who it is for conjunctions to work
	pulse(sender string, pulse bool) []queuedPulse
	setOutgoing([]pulser) // []pulser?
	setOutgoingNames([]string)
	getOutgoingNames() []string
	addOutgoing(pulser) // pulser?
	addIncoming(string) // pulser?
	getName() string
	getLows() int
	getHighs() int
	incLow()
	incHigh()
	inc(level bool)
}

type module struct {
	name          string
	incoming      []string
	outgoing      []pulser
	outgoingNames []string
	lows, highs   int
}

func (m *module) getLows() int {
	return m.lows
}

func (m *module) getHighs() int {
	return m.highs
}

func (m *module) incLow() {
	m.lows++
}

func (m *module) incHigh() {
	m.highs++
}

func (m *module) getName() string {
	return m.name
}
func (m *module) inc(pl bool) {
	if pl == HIGH {
		m.highs++
	} else {
		m.lows++
	}
}

func (m *module) setOutgoing(ps []pulser) {
	m.outgoing = ps
}
func (m *module) addOutgoing(p pulser) {
	m.outgoing = append(m.outgoing, p)
}
func (m *module) setOutgoingNames(ns []string) {
	m.outgoingNames = ns
}
func (m *module) getOutgoingNames() []string {
	return m.outgoingNames
}

// these only really apply to Conjunctions

func (m *module) addIncoming(s string) {
}
func (c *conjunction) addIncoming(s string) {
	c.state[s] = LOW
}

type flipflop struct {
	module
	state bool
}

func parseModules(lines []string) map[string]pulser {
	pulsers := make(map[string]pulser)
	for _, line := range lines {
		p := parseModule(line)
		pulsers[p.getName()] = p
	}
	// loop again, this time attaching the actual pulsers
	for pn, p := range pulsers {
		for _, on := range p.getOutgoingNames() {
			op, ok := pulsers[on]
			if !ok {
				// target of this pulse does not exist, so generate a no-op and add it to the map
				op = &noop{module: module{name: on}}
				pulsers[on] = op
			}
			p.addOutgoing(op)
			op.addIncoming(pn) // conjunctions also need to know the names of their incomings
			pulsers[on] = op
		}
	}

	return pulsers
}

// not sure what to return here...
func parseModule(line string) pulser {
	// Going to need two passes; first one creates all the modules and the second wires them up
	// i think pulser will need methods to set the outgoing and incoming
	lineParts := strings.Split(line, " -> ")
	name := lineParts[0]
	dests := parsers.SplitByCommasAndTrim(lineParts[1])
	_ = dests
	var p pulser
	switch name[0] {
	case 'b':
		p = &broadcast{
			module: module{
				name: "broadcaster",
			},
		}
	case '%':
		p = &flipflop{
			module: module{
				name: name[1:],
			},
			state: LOW,
		}
	case '&':
		p = &conjunction{
			module: module{
				name: name[1:],
			},
			state: make(map[string]bool),
		}
	default:
		p = &noop{module: module{name: name}}
	}
	p.setOutgoingNames(dests)
	return p
}

func (f *flipflop) pulse(sender string, pulse bool) (nextPulses []queuedPulse) {

	f.inc(pulse)
	if pulse == LOW {
		f.state = !f.state
		//fmt.Printf("%s received LOW, state toggled to %t\n", f.name, f.state)
		for _, p := range f.outgoing {
			//fmt.Printf("%s -%s -> %s\n", f.name, printPulse(f.state), p.getName())
			//p.Pulse(f.name, f.state)
			nextPulses = append(nextPulses, queuedPulse{sourceName: f.name, level: f.state, target: p.getName()})
		}
	} else {
		// if queuedPulse = HIGH we do nothing
		//fmt.Printf("%s sending NO pulses; input is HIGH\n", f.getName())
	}
	return nextPulses
}

type conjunction struct {
	module
	state map[string]bool
}

func (c *conjunction) pulse(sender string, pulse bool) (nextPulses []queuedPulse) {
	c.inc(pulse)
	c.state[sender] = pulse
	mypulse := LOW
	// check if any state is low; we send LOW if all are HIGH, else HIGH
	for _, s := range c.state {
		if s == LOW {
			mypulse = HIGH
			break
		}
	}
	for _, p := range c.outgoing {
		//fmt.Printf("%s -%s -> %s\n", c.getName(), printPulse(mypulse), p.getName())
		nextPulses = append(nextPulses, queuedPulse{sourceName: c.name, level: mypulse, target: p.getName()})
		//p.pulse(c.name, mypulse)
	}
	return nextPulses
}

// there's only one of these, and it has no inputs
type broadcast struct {
	module
}

func (b *broadcast) pulse(sender string, pulse bool) (nextPulses []queuedPulse) {
	b.inc(pulse)
	for _, p := range b.outgoing {
		//fmt.Printf("%s -%s -> %s\n", b.getName(), printPulse(pulse), p.getName())
		nextPulses = append(nextPulses, queuedPulse{sourceName: b.name, level: pulse, target: p.getName()})
		//p.pulse(b.name, pulse)
	}
	return nextPulses
}

type noop struct {
	module
}

func (n *noop) pulse(sender string, pulse bool) (nextPulses []queuedPulse) {
	n.inc(pulse)
	return []queuedPulse{}
}

func printPulse(p bool) string {
	if p {
		return "high"
	}
	return "low"
}
