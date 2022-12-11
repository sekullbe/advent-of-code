package main

import (
	_ "embed"
	"fmt"
	"github.com/sekullbe/advent/parsers"
	"github.com/sekullbe/advent/tools"
	"log"
	"sort"
	"strings"
)

//go:embed input.txt
var inputText string

func main() {
	fmt.Printf("Magic number: %d\n", run1(inputText))
	fmt.Println("-------------")
	fmt.Printf("Magic number: %d\n", run2(inputText))
}

type monkey struct {
	id           int
	items        []int
	op           rune
	opArg        string
	testNum      int // always 'divisible by'
	targetTrue   int
	targetFalse  int
	inspectCount int
}

// type barrel map[int]monkey
type barrel struct {
	lcd     int
	monkeys map[int]monkey
}

func newBarrel() barrel {
	return barrel{
		lcd:     1,
		monkeys: make(map[int]monkey),
	}
}

func eMonkey() monkey {
	return monkey{
		items:   make([]int, 0),
		testNum: 1,
	}
}

func parseMonkeys(lines []string) barrel {
	b := newBarrel()
	var monkeyLineGroup []string
	for _, line := range lines {
		if line != "" {
			monkeyLineGroup = append(monkeyLineGroup, line)
		} else {
			m := parseOneMonkey(monkeyLineGroup)
			b.monkeys[m.id] = m
			monkeyLineGroup = monkeyLineGroup[:0]
		}
	}
	// get the last one
	if len(monkeyLineGroup) == 6 {
		m := parseOneMonkey(monkeyLineGroup)
		b.monkeys[m.id] = m
	}
	return b
}

func parseOneMonkey(lines []string) monkey {
	m := eMonkey()
	fmt.Sscanf(lines[0], "Monkey %d:", &m.id)
	sistrings := strings.Split(lines[1], ": ")
	m.items = parsers.StringsWithCommasToIntSlice(sistrings[1])
	n, err := fmt.Sscanf(lines[2], "  Operation: new = old %c %s", &m.op, &m.opArg)
	if n != 2 || err != nil {
		log.Panicf("parse failed: %s", lines[2])
	}
	fmt.Sscanf(lines[3], "  Test: divisible by %d", &m.testNum)
	fmt.Sscanf(lines[4], "    If true: throw to monkey %d", &m.targetTrue)
	fmt.Sscanf(lines[5], "    If false: throw to monkey %d", &m.targetFalse)

	return m
}

func (b *barrel) monkeyTurn(monkeyNum int, worryReductionFactor int) {
	// inspect each item
	m := (*b).monkeys[monkeyNum]
	items := m.items
	for _, item := range items {
		m.inspectCount++
		//   modify the item by the monkey's op
		opArg := m.opArg
		var opArgNum int = 1
		if opArg == "old" {
			opArgNum = item
		} else {
			opArgNum = tools.Atoi(opArg)
		}
		switch m.op {
		case '*':
			item = item * opArgNum
		case '+':
			item = item + opArgNum
		}

		// reduce worry
		item = item / worryReductionFactor // this rounds towards zero; if worry doesn't go negative we're ok
		// throw item to *end* of target's list (append not shift)
		pass := item%m.testNum == 0
		if b.lcd > 1 {
			item %= b.lcd
		}
		if pass {
			b.throwToMonkey(m.targetTrue, item)
		} else {
			b.throwToMonkey(m.targetFalse, item)
		}
	}

	m.items = m.items[:0]
	(*b).monkeys[m.id] = m
}

func (b *barrel) throwToMonkey(monkeyId int, item int) {
	target := (*b).monkeys[monkeyId]
	target.items = append(target.items, item)
	(*b).monkeys[monkeyId] = target
}

func run1(inputText string) int {

	b := parseMonkeys(parsers.SplitByLinesNoTrim(inputText))
	rounds := 20
	for i := 0; i < rounds; i++ {
		for mId := 0; mId < len(b.monkeys); mId++ {
			b.monkeyTurn(mId, 3)
		}
	}
	var inspectCounts []int
	for _, m := range b.monkeys {
		log.Printf("Monkey %d inspected items %d times\n", m.id, m.inspectCount)
		inspectCounts = append(inspectCounts, m.inspectCount)
	}
	// get top 2 and multiply them
	sort.Ints(inspectCounts)
	return inspectCounts[len(inspectCounts)-1] * inspectCounts[len(inspectCounts)-2]
}

func run2(inputText string) int {
	b := parseMonkeys(parsers.SplitByLinesNoTrim(inputText))
	for _, m := range b.monkeys {
		b.lcd *= m.testNum
	}
	rounds := 10000
	for round := 1; round <= rounds; round++ {
		for mId := 0; mId < len(b.monkeys); mId++ {
			b.monkeyTurn(mId, 1)
		}

		//if round == 20 || round%1000 == 0 {
		//	log.Printf("== After round %d ==", round)
		//	b.reportInspections()
		//}
	}
	var inspectCounts []int
	for _, m := range b.monkeys {
		inspectCounts = append(inspectCounts, m.inspectCount)
	}
	// get top 2 and multiply them
	sort.Ints(inspectCounts)
	return inspectCounts[len(inspectCounts)-1] * inspectCounts[len(inspectCounts)-2]
}

func (b barrel) reportInspections() {
	for _, m := range b.monkeys {
		log.Printf("Monkey %d inspected items %d times\n", m.id, m.inspectCount)
	}
}
