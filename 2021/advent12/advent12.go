package main

import (
	_ "embed"
	"fmt"
	"github.com/sekullbe/advent/parsers"
	"log"
	"strings"
)

//go:embed input.txt
var inputText string

func main() {
	fmt.Printf("Magic number: %d\n", run1(inputText))
	fmt.Println("-------------")
	fmt.Printf("Magic number: %d\n", run2(inputText))
}

func run1(inputText string) int {

	_, startCave := parseCaveSystem(inputText)
	return exploreNeighbors(*startCave, map[string]bool{"start": true}, []string{"start"})
}

func run2(inputText string) int {
	_, startCave := parseCaveSystem(inputText)
	return exploreNeighbors2(*startCave, map[string]int{"start": 1}, []string{"start"}, false)
}

// inspired heavily by https://github.com/alexchao26/advent-of-code-go/blob/main/2021/day12/main.go
// and rewritten until it fit my data model and I understood it mostly
func exploreNeighbors(node cave, visited map[string]bool, trail []string) int {
	if node.isEnd() {
		//log.Printf("Finished: %s\n", strings.Join(trail, "-"))
		return 1
	}

	pathsToEnd := 0
	visited[node.id] = true

	for _, neighbor := range node.neighbors {
		if visited[neighbor.id] && neighbor.isSmall() {
			continue
		}

		//log.Printf("Seen %s, exploring %s-%s\n", strings.Join(trail, "-"), node.id, neighbor.id)
		trail = append(trail, neighbor.id)
		pathsToEnd += exploreNeighbors(*neighbor, visited, trail)

		// this took me a bit to grok...
		// back up one and DFS again from that point
		// i.e. if you have s-a-b-end and s-a-c-end
		// after it sees b-end, it'll forget it's seen end and try c-end next
		//log.Printf("Backtracking: forget about %s\n", neighbor.id)
		visited[neighbor.id] = false
		trail = trail[:len(trail)-1]
	}

	return pathsToEnd
}

// for part 2 - you can revisit a small cave twice not once but start/end only once
// secondtime means we're descending from a small cave we've already been through once
func exploreNeighbors2(node cave, visits map[string]int, trail []string, secondTime bool) int {
	if node.isEnd() {
		log.Printf("Finished: %s\n", strings.Join(trail, "-"))
		return 1
	}

	pathsToEnd := 0
	visits[node.id]++

	for _, neighbor := range node.neighbors {

		// explicitly can't go back to start; we already catch hitting end and don't recurse
		if neighbor.isStart() {
			continue
		}

		log.Printf("Seen %s, exploring %s-%s\n", strings.Join(trail, "-"), node.id, neighbor.id)

		// if the neighbor cave is small AND we've already seen it once
		// and the node we are recursing from cannot be visited again
		// then don't visit it, else set the flag
		if neighbor.isSmall() && visits[neighbor.id] > 0 {
			if secondTime {
				// if secondtime is true, don't visit any small caves we've seen befor
				log.Printf("secondtime true , not visiting %s", neighbor.id)
				continue
			} else {
				secondTime = true
				log.Printf("secondtime set TRUE because looking at %s", neighbor.id)
			}
		}

		trail = append(trail, neighbor.id)
		pathsToEnd += exploreNeighbors2(*neighbor, visits, trail, secondTime)

		log.Printf("Backtracking: forget about %s\n", neighbor.id)
		visits[neighbor.id]--
		// If backtracking from small cave we've been to twice before, undo secondTime as well
		// so we can recurse from the original node that set it some more
		if neighbor.isSmall() && visits[neighbor.id] == 1 {
			log.Printf("secondtime set FALSE backtracking from %s", neighbor.id)
			secondTime = false

		}
		trail = trail[:len(trail)-1]
	}

	return pathsToEnd
}

func exploreNeighbors2a(node cave, visits map[string]int, trail []string, secondTime bool) int {
	if node.isEnd() {
		//log.Printf("Finished: %s\n", strings.Join(trail, "-"))
		return 1
	}

	pathsToEnd := 0
	visits[node.id]++

	for _, neighbor := range node.neighbors {

		// explicitly can't go back to start; we already catch hitting end and don't recurse
		if neighbor.isStart() {
			continue
		}

		//log.Printf("Seen %s, exploring %s-%s\n", strings.Join(trail, "-"), node.id, neighbor.id)
		trail = append(trail, neighbor.id)
		if visits[neighbor.id] > 2 && neighbor.isSmall() {
			pathsToEnd += exploreNeighbors2(*neighbor, visits, trail, secondTime)
		}

		//log.Printf("Backtracking: forget about %s\n", neighbor.id)
		visits[neighbor.id]--
		if neighbor.isSmall() && visits[neighbor.id] == 1 {
			secondTime = false

		}
		trail = trail[:len(trail)-1]
	}

	return pathsToEnd
}

// for part1 anyway, the actual caveSystem wasn't necessary because
// all the relevant information was in the neighbors
func parseCaveSystem(inputText string) (caveSystem, *cave) {
	caveSystem := make(caveSystem)
	var startCave *cave
	for _, line := range parsers.SplitByLines(inputText) {
		ids := strings.Split(line, "-")
		from := ids[0]
		to := ids[1]
		var fromCave, toCave *cave
		if fc, exists := caveSystem[from]; !exists {
			fromCave = newCave(from)
			if fromCave.isStart() {
				startCave = fromCave
			}
		} else {
			fromCave = fc
		}
		if tc, exists := caveSystem[to]; !exists {
			toCave = newCave(to)
		} else {
			toCave = tc
		}
		fromCave.neighbors = append(fromCave.neighbors, toCave)
		toCave.neighbors = append(toCave.neighbors, fromCave)
		caveSystem[from] = fromCave
		caveSystem[to] = toCave
	}
	return caveSystem, startCave
}
