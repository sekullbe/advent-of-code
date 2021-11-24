package advent3

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var inputs string

// landscape is a map[int] of mapint] of boolean
// so you can call landscape[row][column] to get a tree or not
type landscape map[int]map[int]bool

type speed struct {
	dx int
	dy int
}

func Run() {
	vel := speed{dx:3, dy:1}
	fmt.Printf("Trees for speed (%d, %d): %d\n",vel.dx, vel.dy, countTreesInRun(vel))

	var runs []speed
	runs = append(runs, speed{dx:1, dy:1})
	runs = append(runs, speed{dx:3, dy:1})
	runs = append(runs, speed{dx:5, dy:1})
	runs = append(runs, speed{dx:7, dy:1})
	runs = append(runs, speed{dx:1, dy:2})

	product := 1
	for _, run := range runs {
		trees :=  countTreesInRun(run)
		fmt.Printf("Trees for speed (%d, %d): %d\n",vel.dx, vel.dy, trees)
		product *= trees
	}
	fmt.Printf("product: %d\n", product)
}

func countTreesInRun(vel speed) int{


	// parse the input into some structure- I think a 2D map
	thehill, colCount, rowCount := parseLandscape(inputs)

	//starting locations
	x := 0
	y := 0

	trees := 0

	for i := 0; i < rowCount; i++ {
		x,y = move(x, y, colCount - 1, vel.dx, vel.dy)
		if thehill[y][x] {
			trees++
		}
	}

	return trees
}

// MaxX is the index of the end column, so the row is maxX + 1 columns wide (0..maxX)
func move(x, y, maxX, dx, dy int) (newX,newY int) {
	newY = y + dy
	newX = (x + dx) % (maxX+1)
	return newX, newY
}


func parseLandscape(inputText string) (landscape, int, int) {

	rows := strings.Fields(inputText)

	out := make(landscape)
	var rowCount, colCount int
	for i, row := range rows {
		rowMap := make(map[int]bool)
		if colCount == 0 {
			colCount = len(row)
		}
		for i2, s := range strings.Split(row, "") {
			rowMap[i2] = s == "#"
		}
		out[i] = rowMap
		rowCount++
	}
	return out, colCount, rowCount
}
