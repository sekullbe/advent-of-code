package main

import (
	_ "embed"
	"fmt"
	"github.com/sekullbe/advent/tools"
	"regexp"
	"strings"
)

//go:embed input.txt
var inputText string

func main() {
	fmt.Printf("Magic number: %d\n", run1(inputText))
	fmt.Println("-------------")
	fmt.Printf("Magic number: %d\n", run2(inputText))
}

func run1(input string) int {

	sum := 0
	// without the trim it picked up a newline at the end of the string
	steps := strings.Split(strings.TrimSpace(input), ",")
	for _, step := range steps {
		sum += hash(step)
	}

	return sum
}

type box struct {
	id     int
	lenses map[string]lens
}
type lens struct {
	focal    int
	label    string
	position int
}

func initBoxes() []*box {
	var boxes []*box = make([]*box, 256)
	for i := 0; i <= 255; i++ {
		boxes[i] = &box{id: i, lenses: make(map[string]lens)}
	}
	return boxes
}

func run2(input string) int {

	var boxes []*box = initBoxes()

	steps := strings.Split(strings.TrimSpace(input), ",")
	for _, steps := range steps {
		label, boxnum, op, focal := decodeStep(steps)
		switch op {
		case "-":
			boxes[boxnum].removeLens(label)
		case "=":
			boxes[boxnum].insertLens(label, focal)
		default:
			panic("unknown operator")
		}
	}
	pow := 0
	for _, b := range boxes {
		pow += b.calculatePower()
	}

	return pow
}

func (b *box) insertLens(label string, focal int) {
	l := lens{focal: focal, label: label} // leave position blank
	oldLens, exists := b.lenses[label]
	if exists {
		// just update the focal, the rest is the same
		oldLens.focal = focal
		b.lenses[label] = oldLens
	} else {
		l.position = b.getMaxLensPosition() + 1
		b.lenses[label] = l
	}
}

func (b *box) getMaxLensPosition() int {
	maxPos := 0
	for _, l := range b.lenses {
		maxPos = max(maxPos, l.position)
	}
	return maxPos
}

func (b *box) removeLens(label string) {
	l, exists := b.lenses[label]
	if !exists {
		return // removing a lens that doesn't exist
	}
	p := l.position
	delete(b.lenses, label)
	for _, l2 := range b.lenses {
		if l2.position > p {
			l2.position = l2.position - 1
			b.lenses[l2.label] = l2
		}
	}
}

func (b *box) calculatePower() int {
	boxPower := 0
	for _, l := range b.lenses {
		boxPower += (b.id + 1) * l.position * l.focal
	}
	return boxPower
}

func decodeStep(step string) (label string, boxnum int, op string, focal int) {
	re := regexp.MustCompile(`(.*?)([-=])(\d*)`)
	matches := re.FindStringSubmatch(step)
	label = matches[1]
	boxnum = hash(matches[1])
	op = matches[2]
	if op == "=" {
		focal = tools.Atoi(matches[3])
	}
	return label, boxnum, op, focal
}

func hash(step string) int {
	val := 0
	stepInts := stringToAsciiInt(step)
	for _, u := range stepInts {
		val += u
		val *= 17
		val %= 256
	}
	return val
}

func stringToAsciiInt(str string) []int {
	b := []byte(str)
	i := []int{}
	for _, b2 := range b {
		i = append(i, int(b2))
	}
	return i
}
