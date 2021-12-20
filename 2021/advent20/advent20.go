package main

import (
	_ "embed"
	"fmt"
	"github.com/sekullbe/advent/parsers"
)

//go:embed input.txt
var inputText string

type pixel struct {
	row int
	col int
}

type pixels map[pixel]bool

type image struct {
	pixels   pixels
	min, max int // min or max coordinate in any direction
}

type algorithm map[int]bool

func main() {
	fmt.Printf("Magic number: %d\n", run1(inputText))
	fmt.Println("-------------")
	fmt.Printf("Magic number: %d\n", run2(inputText))
}

func run1(inputText string) int {
	return run(inputText, 2)
}

func run2(inputText string) int {
	return run(inputText, 50)
}

func run(inputText string, iterations int) int {
	im, al := parseInput(inputText)
	background := false
	//printImage(im)
	for i := 0; i < iterations; i++ {
		im = enhance(im, al, background)
		// toggle the rest of the world, anything we haven't seen yet
		if background {
			background = al[511]
		} else {
			background = al[0]
		}
		//printImage(im)
	}

	count := 0
	for _, lit := range im.pixels {
		if lit {
			count++
		}
	}
	return count
}

func parseInput(inputText string) (image, algorithm) {
	// algorithm, newline, grid
	lines := parsers.SplitByLines(inputText)
	algorithm := parseAlgorithm(lines[0])
	image := parseToMap(lines[2:])

	return image, algorithm
}

func parseAlgorithm(line string) algorithm {
	a := make(algorithm)
	for i, p := range line {
		if p == '#' {
			a[i] = true
		}
	}
	return a
}

func parseToMap(lines []string) (im image) {
	im = image{}
	im.pixels = make(pixels)
	im.max = len(lines)
	im.min = 0
	for row, line := range lines {
		for col, pr := range line {
			lit := pr == '#'
			p := pixel{row: row, col: col}
			im.pixels[p] = lit
		}
	}

	return im
}

func (p pixel) neighbors() (neighbors []pixel) {
	n := []int{-1, 0, 1}
	// do not worry about writing off the 'edge'
	// we want the row first then the column
	for _, dx := range n {
		for _, dy := range n {
			neighbors = append(neighbors, pixel{row: p.row + dx, col: p.col + dy})
		}
	}
	return neighbors
}

func enhancePixel(im image, p pixel, a algorithm, bg bool) (lit bool) {
	var enhanceKey int
	n := p.neighbors()
	for _, neighbor := range n {
		enhanceKey <<= 1
		neighborLit, known := im.pixels[neighbor]
		if !known { // i.e. outside the original image
			neighborLit = bg
		}
		if neighborLit {
			enhanceKey |= 1
		}
	}
	return a[enhanceKey]
}

func enhance(im image, a algorithm, bg bool) (newImage image) {
	fillInPixels(&im, bg)
	newImage.pixels = make(pixels)
	newImage.min = im.min
	newImage.max = im.max
	for p, _ := range im.pixels {
		newImage.pixels[p] = enhancePixel(im, p, a, bg)
	}
	return
}

func fillInPixels(im *image, bg bool) {
	im.min -= 1
	im.max += 1

	for r := im.min; r < im.max; r++ {
		if r == im.min || r == im.max-1 {
			for c := im.min; c < im.max; c++ {
				_, known := im.pixels[pixel{r, c}]
				//log.Printf("max=%d min=%d row=%d col=%d", im.max, im.min, r, c)
				if !known {
					p := pixel{r, c}
					im.pixels[p] = bg
				}
			}
		} else {
			c := im.min
			_, known := im.pixels[pixel{r, c}]
			//log.Printf("max=%d min=%d row=%d col=%d", im.max, im.min, r, c)
			if !known {
				p := pixel{r, c}
				im.pixels[p] = bg
			}
			c = im.max - 1
			_, known = im.pixels[pixel{r, c}]
			//log.Printf("max=%d min=%d row=%d col=%d", im.max, im.min, r, c)
			if !known {
				p := pixel{r, c}
				im.pixels[p] = bg
			}
		}

	}
	//log.Println("-----")
}

func printImage(im image) {
	for r := im.min - 1; r < im.max+1; r++ {
		for c := im.min - 1; c < im.max+1; c++ {
			lit := im.pixels[pixel{r, c}]
			if lit {
				fmt.Printf("#")
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Println()
	}
	fmt.Println("----------")
}
