package main

import (
	_ "embed"
	"fmt"
	mapset "github.com/deckarep/golang-set/v2"
	combos "github.com/mxschmitt/golang-combinations"
	"github.com/sekullbe/advent/parsers"
	"github.com/sekullbe/advent/tools"
	"sort"
	"strings"
	"time"
)

//go:embed input.txt
var inputText string

func main() {
	fmt.Printf("Magic number: %d\n", run1(inputText))
	fmt.Println("-------------")
	fmt.Printf("Magic number: %d\n", run2(inputText))
}

/*
one way
for each computer beginning with t
look at all pairs of its links
call them A and B
if A links to B you have a three cycle
*/
func run1(input string) int {
	defer tools.Track(time.Now(), "Part 1 Time")
	network := createNetwork(parsers.SplitByLines(input))

	triangles := make(map[string]bool)
	for _, c0 := range network {
		if c0.startswith != 't' {
			continue
		}
		links := c0.links
		// get all pairs of links (c1 c2) in that Set
		// and see if c1 links to c2
		var pairs [][]*computer = combos.Combinations(links.ToSlice(), 2)
		for _, pair := range pairs {
			c1 := pair[0]
			c2 := pair[1]
			if c1.links.Contains(c2) {
				triangles[nameTriangle(c0, c1, c2)] = true
			}
		}
	}
	//fmt.Println(triangles)
	return len(triangles)
}

func run2(input string) int {
	defer tools.Track(time.Now(), "Part 2 Time")

	return 0
}

// input is a set of pairs of computer names that are connected
type computer struct {
	name       string
	startswith rune
	links      mapset.Set[*computer]
}

func nameTriangle(c0, c1, c2 *computer) string {
	names := []string{c0.name, c1.name, c2.name}
	sort.Strings(names)
	return strings.Join(names, ",")
}

type network map[string]*computer

func (n network) get(name string) *computer {
	return n[name]
}
func (n network) linksof(name string) mapset.Set[*computer] {
	return n[name].links
}

func parseComputer(s string) (computer, computer) {
	components := strings.Split(s, "-")
	c1 := computer{name: components[0], startswith: rune(s[0]), links: mapset.NewSet[*computer]()}
	c2 := computer{name: components[1], startswith: rune(components[1][0]), links: mapset.NewSet[*computer]()}
	return c1, c2
}

func createNetwork(lines []string) map[string]*computer {
	network := make(map[string]*computer)
	// build all the computers
	for _, s := range lines {
		c1, c2 := parseComputer(s)
		// there will be duplicate inserts; it does not matter
		network[c1.name] = &c1
		network[c2.name] = &c2
	}
	//now wire them up
	for _, s := range lines {
		cs := strings.Split(s, "-")
		c1 := network[cs[0]]
		c2 := network[cs[1]]

		c1.links.Add(network[c2.name])
		c2.links.Add(network[c1.name])
	}

	return network
}
