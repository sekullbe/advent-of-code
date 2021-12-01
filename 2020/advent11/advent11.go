package advent11

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var inputText string

const (
	N int = iota
	NE
	E
	SE
	S
	SW
	W
	NW
)




type seat struct {
	occupant rune
	flipped  bool
}

func (s *seat) flip() bool {
	switch (*s).occupant {
	case '.':
		return false
	case 'L':
		(*s).occupant = '#'
		return true
	case '#':
		(*s).occupant = 'L'
		return true
	}
	return false
}

// contains . L #
type seatsMap struct {
	seats map[int]map[int]*seat
	maxRowNum int
	maxColNum int
}

func Run() {
	c := run1(inputText)
	fmt.Printf("Run 1 seat count: %d\n", c)

	c = run2(inputText)
	fmt.Printf("Run 2 seat count: %d\n", c)
	// 2050 is too high
}

func run1(inputText string) int{

	seats := parseSeats(inputText)
	var anyFlips bool
	var iterations int
	var countOccupied int

	for {
		anyFlips = false
		iterations++
		// outputSeats(seats)
		countOccupied = 0
		for i := 0; i <= seats.maxRowNum; i++ {
			for j := 0; j <= seats.maxColNum; j++ {
				c := countSurroundingOccupiedSeats(seats, i,j)
				s := seats.seats[i][j]
				if s.occupant == 'L' && c == 0 {
					s.flipped = true
					anyFlips = true
				} else if s.occupant == '#' && c >= 4 {
					s.flipped = true
					anyFlips = true
					continue
				}
				if s.occupant == '#' {
					countOccupied++
				}
			}
		}
		// execute the flips. yeh this kinda sucks.
		if anyFlips {
			for i := 0; i <= seats.maxRowNum; i++ {
				for j := 0; j <= seats.maxColNum; j++ {
					s := seats.seats[i][j]
					if s.flipped {
						s.flip()
						s.flipped = false
					}
				}
			}
		} else {
			break
		}
	}
	return countOccupied
}

func run2(inputText string) int {
	// could consolidate a lot of code with run2 if i cared which I don't
	seats := parseSeats(inputText)
	var anyFlips bool
	var iterations int
	var countOccupied int

	for {
		anyFlips = false
		iterations++
		//outputSeats(seats)
		fmt.Print(".")
		countOccupied = 0
		for i := 0; i <= seats.maxRowNum; i++ {
			for j := 0; j <= seats.maxColNum; j++ {
				c := countVisibleOccupiedSeats(seats, i,j)
				s := seats.seats[i][j]
				if s.occupant == 'L' && c == 0 {
					s.flipped = true
					anyFlips = true
				} else if s.occupant == '#' && c >= 5 {
					s.flipped = true
					anyFlips = true
					continue
				}
				if s.occupant == '#' {
					countOccupied++
				}
			}
		}
		// execute the flips. yeh this kinda sucks.
		if anyFlips {
			for i := 0; i <= seats.maxRowNum; i++ {
				for j := 0; j <= seats.maxColNum; j++ {
					s := seats.seats[i][j]
					if s.flipped {
						s.flip()
						s.flipped = false
					}
				}
			}
		} else {
			break
		}
	}
	fmt.Println()
	return countOccupied
}

func countSurroundingOccupiedSeats(seats seatsMap, x int, y int) int {
	count := 0
	for i := -1; i <= 1 ; i++ {
		if  x + i < 0 || x + i > seats.maxRowNum {
			continue
		}
		for j := -1; j <= 1; j++ {
			if y + j < 0 || y + j > seats.maxColNum {
				continue
			}
			if i == 0 && j == 0 {
				continue
			}
			if seats.seats[x+i][y+j].occupant =='#' {
				count++
			}
		}
	}
	return count
}

func countVisibleOccupiedSeats(seats seatsMap, x int, y int) int {
	count := 0
	// look for seats in all 8 directions 0-7 or N-NW
	for dir := N; dir <= NW; dir++ {
		seen := lookForSeat(seats, x, y, dir)
		if seen {
			count++
		}
	}
	return count
}


func parseSeats(inputText string) seatsMap {

	rows := strings.Fields(inputText)

	out := seatsMap{ seats:make(map[int]map[int]*seat), maxRowNum: 0 , maxColNum: 0}
	var rowCount, colCount int
	for i, row := range rows {
		rowMap := make(map[int]*seat)
		if colCount == 0 {
			colCount = len(row)
		}
		for i2, s := range row {
			rowMap[i2] = &seat{occupant: s}
		}
		out.seats[i] = rowMap
		rowCount++
	}
	out.maxRowNum = rowCount - 1
	out.maxColNum = colCount - 1

	return out
}

func outputSeats(seats seatsMap) {
	for i := 0; i <= seats.maxRowNum; i++ {
		for j := 0; j <= seats.maxColNum; j++ {
			fmt.Printf("%c", seats.seats[i][j].occupant)
		}
		fmt.Println()
	}
}

func lookForSeat(seats seatsMap, row int, col int, direction int) bool{

	r2,c2 := computeNextCoordinateInRay(row, col, direction)
	for {
		//fmt.Printf("Looking at row %d col %d\n", r2, c2)
		if r2 < 0 || c2 < 0 || r2 > seats.maxRowNum || c2 > seats.maxColNum {
			return false
		}
		if seats.seats[r2][c2].occupant == 'L' {
			return false
		}
		if seats.seats[r2][c2].occupant == '#' {
			return true
		}
		r2,c2 = computeNextCoordinateInRay(r2, c2, direction)
	}
}

func computeNextCoordinateInRay(row int, col int, direction int) (int, int) {
	switch direction {
	case N:
		return row - 1, col
	case NE:
	 	return row - 1, col + 1
	case E:
		return row, col + 1
	case SE:
		return row + 1, col +1
	case S:
		return row + 1, col
	case SW:
		return row + 1, col -1
	case W:
		return row, col -1
	case NW:
		return row - 1, col -1
	}
	// can't get there from here
	panic("bad direction")
}
