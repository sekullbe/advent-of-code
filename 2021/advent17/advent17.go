package main

import (
	_ "embed"
	"fmt"
	"github.com/sekullbe/advent/tools"
	"log"
	"regexp"
	"strconv"
)

//go:embed input.txt
var inputText string

type target struct {
	minX int
	maxX int
	minY int
	maxY int
}

type shot struct {
	x, y, dx, dy int
}

func main() {
	fmt.Printf("Magic number: %d\n", run1(inputText))
	fmt.Println("-------------")
	fmt.Printf("Magic number: %d\n", run2(inputText))
}

func run1(inputText string) int {

	target := parseTarget(inputText)
	possibleDx := target.findPossibleDxRange(0)
	maxYSoFar := 0
	for dy := 0; dy < 1000; dy++ {
		for _, dx := range possibleDx {
			s := shot{x: 0, y: 0, dx: dx, dy: dy}
			hit, maxY := launch(s, target)
			if hit {
				//log.Printf("hit: (%d,%d) max %d", dx, dy, maxY)
				if maxY > maxYSoFar {
					maxYSoFar = maxY
				}
			}
		}
	}
	return maxYSoFar
}

func run2(inputText string) int {
	target := parseTarget(inputText)
	possibleDx := target.findPossibleDxRange(0)
	count := 0
	for dy := -1000; dy < 10000; dy++ {
		for _, dx := range possibleDx {
			s := shot{x: 0, y: 0, dx: dx, dy: dy}
			hit, _ := launch(s, target)
			if hit {
				count++
				//log.Printf("hit: (%d,%d) count %d", dx, dy, count)
			}
		}
	}
	return count
}

func launch(s shot, t target) (bool, int) {
	maxY := 0
	for ; !t.pastTarget(s); s.tick() {
		if s.y > maxY {
			maxY = s.y
		}
		//s.log()
		if t.inTarget(s) {
			return true, maxY
		}
	}
	return false, 0
}

func (t target) findPossibleDxRange(initX int) (velocities []int) {
	for i := 1; i <= t.maxX-initX; i++ {
		x := i
		for j := i - 1; j >= 0; j-- {
			if x >= t.minX && x <= t.maxX && !tools.ContainsInt(velocities, i) {
				velocities = append(velocities, i)
			}
			x += j
		}
	}
	return
}

func (t target) findPossibleDyRange(initY int) (velocities []int) {
	for dy := 0; dy < t.minY; dy-- {
		velocities = append(velocities, dy)
	}
	return
}

func (s *shot) tick() {
	s.x = s.x + s.dx
	s.y = s.y + s.dy
	s.dx = reduceOrZero(s.dx, 1)
	s.dy = s.dy - 1
}

func (s shot) log() {
	log.Printf("shot: (%d,%d) dx=%d dy=%d", s.x, s.y, s.dx, s.dy)
}

func reduceOrZero(x int, dx int) (newX int) {
	newX = x - dx
	if newX < 0 {
		newX = 0
	}
	return newX
}

func (t target) inTarget(s shot) bool {
	return s.x >= t.minX && s.x <= t.maxX && s.y >= t.minY && s.y <= t.maxY
}

func (t target) pastTarget(s shot) bool {
	return s.x > t.maxX || s.y < t.minY
}

func parseTarget(inputText string) target {
	re := regexp.MustCompile(`x=([-0-9]+)\.\.([-0-9]+), y=([-0-9]+)\.\.([-0-9]+)`)
	matches := re.FindStringSubmatch(inputText)
	x0, err := strconv.Atoi(matches[1])
	if err != nil {
		panic("bad x0")
	}
	x1, err := strconv.Atoi(matches[2])
	if err != nil {
		panic("bad x1")
	}
	y0, err := strconv.Atoi(matches[3])
	if err != nil {
		panic("bad y0")
	}
	y1, err := strconv.Atoi(matches[4])
	if err != nil {
		panic("bad y1")
	}
	// just compute min,max and have done with it?
	return target{x0, x1, y0, y1}
}
