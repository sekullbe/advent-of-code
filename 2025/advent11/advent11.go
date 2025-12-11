package main

import (
	_ "embed"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/dominikbraun/graph"
	"github.com/sekullbe/advent/parsers"
	"github.com/sekullbe/advent/tools"
)

//go:embed input.txt
var inputText string

func main() {
	fmt.Printf("Magic number: %d\n", run1(inputText))
	fmt.Println("-------------")
	fmt.Printf("Magic number: %d\n", run2(inputText))
}

// Did this originally then backported my part 2 solution
func run1_slow(input string) int {
	defer tools.Track(time.Now(), "Part 1 Time")

	g := buildGraph(input)
	paths := computeAllPaths(g, "you", "out")

	return len(paths)
}

func run1(input string) int {
	defer tools.Track(time.Now(), "Part 2 Time")
	g := make(map[string][]string)
	for _, s := range parsers.SplitByLines(input) {
		device, outputs := parseLine(s)
		g[device] = outputs
	}
	memo := make(map[string]int)
	return dfs(g, "you", "out", memo)
}

func run2(input string) int {
	defer tools.Track(time.Now(), "Part 2 Time")
	g := make(map[string][]string)
	for _, s := range parsers.SplitByLines(input) {
		device, outputs := parseLine(s)
		g[device] = outputs
	}
	memo := make(map[string]int)

	// From examination, I know that only fft->dac exists
	// So break the path into the necessary chunks and multiply the parts
	a := dfs(g, "svr", "fft", memo)
	clear(memo)
	b := dfs(g, "fft", "dac", memo)
	clear(memo)
	c := dfs(g, "dac", "out", memo)
	return a * b * c
}

func dfs(g map[string][]string, curr, target string, memo map[string]int) int {
	if curr == target {
		return 1
	}
	if tools.KeyExists(memo, curr) {
		return memo[curr]
	}
	sum := 0
	for _, e := range g[curr] {
		sum += dfs(g, e, target, memo)
	}
	memo[curr] = sum
	return memo[curr]

}

// This passes the sample but is way too slow for the real data,
// even with the small optimization of breaking up the paths.
func run2_way_too_slow(input string) int {
	defer tools.Track(time.Now(), "Part 2 Time")

	g := buildGraph(input)

	pathsSF := computeAllPaths(g, "svr", "fft")
	log.Printf("pathsSF=%d", len(pathsSF))
	pathsFD := computeAllPaths(g, "fft", "dac")
	log.Printf("pathsFD=%d", len(pathsFD))
	pathsDE := computeAllPaths(g, "dac", "out")
	log.Printf("pathsDE=%d", len(pathsDE))
	count := len(pathsSF) * len(pathsFD) * len(pathsDE)

	return count
}

func parseLine(s string) (string, []string) {
	parts := strings.Split(s, ":")
	device := parts[0]
	outputs := strings.Fields(parts[1])
	return device, outputs
}

func buildGraph(input string) graph.Graph[string, string] {
	g := graph.New(graph.StringHash, graph.Directed(), graph.Acyclic(), graph.PreventCycles())
	for _, s := range parsers.SplitByLines(input) {
		device, outputs := parseLine(s)

		err := g.AddVertex(device)
		if err != nil { // this amount of error checking is silly
			if errors.Is(err, graph.ErrVertexAlreadyExists) {
				//log.Printf("already added %s, that's ok", device)
			} else {
				log.Panicf("that ain't right adding %s", device)
			}
		}
		for _, output := range outputs {
			err = g.AddVertex(output)
			if err != nil {
				if errors.Is(err, graph.ErrVertexAlreadyExists) {
					//log.Printf("already added %s, that's ok", device)
				} else {
					log.Panicf("that ain't right adding %s", device)
				}
			}
		}
		for _, output := range outputs {
			err = g.AddEdge(device, output)
			if err != nil {
				if errors.Is(err, graph.ErrVertexAlreadyExists) {
					log.Printf("already added %s, that's ok", device)
				} else {
					log.Panicf("that ain't right adding %s", device)
				}
			}
		}
	}
	return g
}

func computeAllPaths(g graph.Graph[string, string], start string, end string) [][]string {

	paths, err := graph.AllPathsBetween(g, start, end)
	if err != nil {
		panic(err)
	}
	return paths
}
