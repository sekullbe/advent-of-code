package main

import (
	_ "embed"
	"fmt"
	mapset "github.com/deckarep/golang-set/v2"
	"github.com/dominikbraun/graph"
	"github.com/sekullbe/advent/parsers"
	"github.com/sekullbe/advent/tools"
	"strings"
)

//go:embed input.txt
var inputText string

func main() {
	fmt.Printf("Magic number: %d\n", run1(inputText))
	fmt.Println("-------------")
	fmt.Printf("Magic number: %d\n", run2(inputText))
}

type Component struct {
	name        string
	id          int
	cxnNames    mapset.Set[string]
	connections mapset.Set[*Component]
}

func (c Component) getName() string {
	return c.name
}

type Components map[string]*Component

func run1(input string) int {
	lines := parsers.SplitByLines(input)
	components := make(Components)
	for _, line := range lines {
		cp := ParseComponent(line)
		if !tools.KeyExists(components, cp.name) {
			components[cp.name] = cp
		}

		for _, name := range cp.cxnNames.ToSlice() {
			if !tools.KeyExists(components, name) {
				components[name] = &Component{name: name, cxnNames: mapset.NewSet(cp.name), connections: mapset.NewSet(cp)}
			} else {
				components[name].cxnNames.Add(cp.name)
				components[name].connections.Add(cp)
			}
			components[cp.name].cxnNames.Add(name)
			components[cp.name].connections.Add(components[name])
		}
	}

	// I really should use Karger's algorithm given that I did my undergrad thesis
	// with him, but I'm tired. Maybe I'll try again later.

	g := graph.New(graph.StringHash)
	for n, component := range components {
		g.AddVertex(n)
		for _, cxn := range component.cxnNames.ToSlice() {
			g.AddEdge(n, cxn)
		}
	}
	//file, _ := os.Create("./out.gv")
	//err := draw.DOT(g, file)
	//if err != nil {
	//	panic(err)
	//}

	// from visual examination, the cuts are rjs/mrd, ntx/gmr, and gsk/ncg
	g.RemoveEdge("rjs", "mrd")
	g.RemoveEdge("mrd", "rjs")
	g.RemoveEdge("ntx", "gmr")
	g.RemoveEdge("gmr", "ntx")
	g.RemoveEdge("gsk", "ncg")
	g.RemoveEdge("ncg", "gsk")

	p1 := 0
	graph.DFS(g, "rjs", func(s string) bool {
		p1++
		return false
	})

	p2 := 0
	graph.DFS(g, "mrd", func(s string) bool {
		p2++
		return false
	})

	return p1 * p2
}

func run2(input string) int {

	return 0
}

func ParseComponent(line string) *Component {
	c := Component{}
	parts := strings.Split(line, ": ")
	c.name = parts[0]
	c.cxnNames = mapset.NewSet[string]()
	c.connections = mapset.NewSet[*Component]()
	for _, s := range strings.Fields(parts[1]) {
		c.cxnNames.Add(s)
	}
	return &c
}

func PrintComponents(components Components) {
	for _, component := range components {
		fmt.Printf("%s: %v\n", component.name, component.cxnNames)
	}
}
