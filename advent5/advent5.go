package advent5

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"

	"github.com/deckarep/golang-set"
)

//go:embed input.txt
var inputs string


func Run() {

	seenSeatIds := mapset.NewSet()
	i := computeSeatId(0,1)
	for i < 835 {
		seenSeatIds.Add(i)
		i += 1
	}

	for _, seatCode := range strings.Fields(inputs) {
		seenSeatIds.Remove(computeSeatId(parseInputRow(seatCode)))
	}

	ids := seenSeatIds.ToSlice()
	for _, id := range ids {
		intid := id.(int64)
		plusone := intid + 1
		minusone := intid - 1
		if !seenSeatIds.Contains(plusone) && !seenSeatIds.Contains(minusone) {
			fmt.Printf("Your seat ID: %d\n", intid)
		}

	}


}

func Run1() {

	var maxSeatId int64
	for _, seatCode := range strings.Fields(inputs) {
		seatId := computeSeatId(parseInputRow(seatCode))
		if seatId > maxSeatId {
			maxSeatId = seatId
		}
	}
	fmt.Printf("Max Seat ID: %d\n", maxSeatId)
}


func parseInputRow(seatCode string) (row, col int64) {
	// split to 7/3
	// F=0 B=1 L=0 R=1
	seatCode = strings.ReplaceAll(seatCode,"F", "0")
	seatCode = strings.ReplaceAll(seatCode,"B", "1")
	seatCode = strings.ReplaceAll(seatCode,"L", "0")
	seatCode = strings.ReplaceAll(seatCode,"R", "1")
	rowCode := seatCode[:7]
	colCode := seatCode[7:]
	row, _ = strconv.ParseInt(rowCode, 2, 0)
	col, _ = strconv.ParseInt(colCode, 2, 0)
	return row,col
}

func computeSeatId(row,col int64) int64 {
	return row*8+col
}
