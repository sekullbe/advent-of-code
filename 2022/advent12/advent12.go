package main

import (
	_ "embed"
	"fmt"
	"github.com/beefsack/go-astar"
	"github.com/sekullbe/advent/parsers"
	"log"
)

//go:embed input.txt
var inputText string

// Cheating somewhat by using the astar library, but I really dislike pathing

func main() {
	fmt.Printf("Magic number: %d\n", run1(inputText))
	fmt.Println("-------------")
	fmt.Printf("Magic number: %f\n", run2(inputText))
}

func run1(inputText string) int {

	world, from, to := parseWorld(parsers.SplitByLines(inputText))
	_ = world
	p, dist, found := astar.Path(from, to)
	_ = p
	if !found {
		log.Panicf("could not find a path, that's bad")
	}

	return int(dist)
}

func run2(inputText string) float64 {

	world, from, to := parseWorld(parsers.SplitByLines(inputText))

	minHike := 528.0

	for p, tile := range world {
		if tile.height == 0 {
			from = world[p]
			_, dist, found := astar.Path(from, to)
			if found && dist < minHike {
				minHike = dist
			}
		}
	}
	return minHike
}
