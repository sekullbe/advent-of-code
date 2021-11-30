package advent11

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_seat_flip(t *testing.T) {

	var s seat
	s.occupant = '.'
	sp := &s
	
	b := sp.flip()
	assert.Equal(t, false, b)
	assert.True(t, '.' == sp.occupant)

	s.occupant = 'L'
	b = sp.flip()
	assert.Equal(t, true, b)
	assert.True(t, '#' == sp.occupant)

	s.occupant = '#'
	b = sp.flip()
	assert.Equal(t, true, b)
	assert.True(t, 'L' == sp.occupant)

}

func Test_countSurroundingOccupiedSeats(t *testing.T) {

	type args struct {
		seats seatsMap
		x     int
		y     int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "corner 00",
			args: args{ seats: parseSeats("###\n###\n###"),
				x: 0, y: 0},
			want: 3,
		},
		{
			name: "corner 02",
			args: args{ seats: parseSeats("###\n###\n###"),
				x: 0, y: 0},
			want: 3,
		},
		{
			name: "corner 20",
			args: args{ seats: parseSeats("###\n###\n###"),
				x: 0, y: 0},
			want: 3,
		},
		{
			name: "corner 22",
			args: args{ seats: parseSeats("###\n###\n###"),
				x: 0, y: 0},
			want: 3,
		},
		{
			name: "full",
			args: args{ seats: parseSeats("###\n###\n###"),
				x: 1, y: 1},
			want: 8,
		},
		{
			name: "empty",
			args: args{ seats: parseSeats("...\n...\n..."),
				x: 0, y: 0},
			want: 0,
		},
	}
		for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countSurroundingOccupiedSeats(tt.args.seats, tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("countSurroundingOccupiedSeats() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseSeats(t *testing.T) {

	inputText := "L.LL.LL.LL\nLLLLLLL.LL\nL.L.L..L..\nLLLL.LL.LL\nL.LL.LL.LL\nL.LLLLL.LL\n..L.L.....\nLLLLLLLLLL\nL.LLLLLL.L"
	s := parseSeats(inputText)
	assert.Equal(t, 8, s.maxRowNum)
	assert.Equal(t, 9, s.maxColNum)
	// [row][col]
	assert.Equal(t, 'L', s.seats[0][0].occupant)
	assert.Equal(t, '.', s.seats[0][1].occupant)
	assert.Equal(t, 'L', s.seats[0][2].occupant)
	assert.Equal(t, 'L', s.seats[0][3].occupant)
	assert.Equal(t, '.', s.seats[0][4].occupant)
	assert.Equal(t, 'L', s.seats[0][5].occupant)
	assert.Equal(t, 'L', s.seats[0][6].occupant)
	assert.Equal(t, '.', s.seats[0][7].occupant)
	assert.Equal(t, 'L', s.seats[0][8].occupant)
	assert.Equal(t, 'L', s.seats[0][9].occupant)

	assert.Equal(t, 'L', s.seats[1][0].occupant)
	assert.Equal(t, 'L', s.seats[1][1].occupant)
	assert.Equal(t, 'L', s.seats[1][2].occupant)
	assert.Equal(t, 'L', s.seats[1][3].occupant)
	assert.Equal(t, 'L', s.seats[1][4].occupant)
	assert.Equal(t, 'L', s.seats[1][5].occupant)
	assert.Equal(t, 'L', s.seats[1][6].occupant)
	assert.Equal(t, '.', s.seats[1][7].occupant)
	assert.Equal(t, 'L', s.seats[1][8].occupant)
	assert.Equal(t, 'L', s.seats[1][9].occupant)

	assert.Equal(t, 'L', s.seats[8][0].occupant)
	assert.Equal(t, '.', s.seats[8][1].occupant)
	assert.Equal(t, 'L', s.seats[8][2].occupant)
	assert.Equal(t, 'L', s.seats[8][3].occupant)
	assert.Equal(t, 'L', s.seats[8][4].occupant)
	assert.Equal(t, 'L', s.seats[8][5].occupant)
	assert.Equal(t, 'L', s.seats[8][6].occupant)
	assert.Equal(t, 'L', s.seats[8][7].occupant)
	assert.Equal(t, '.', s.seats[8][8].occupant)
	assert.Equal(t, 'L', s.seats[8][9].occupant)



}

func Test_run1(t *testing.T) {
	type args struct {
		inputText string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "empty",
			args: args {inputText: "...\n...\n..."},
			want: 0,
		},
		{
			name: "example",
			args: args {inputText: "L.LL.LL.LL\nLLLLLLL.LL\nL.L.L..L..\nLLLL.LL.LL\nL.LL.LL.LL\nL.LLLLL.LL\n..L.L.....\nLLLLLLLLLL\nL.LLLLLL.L\nL.LLLLL.LL"},
			want: 37,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run1(tt.args.inputText); got != tt.want {
				t.Errorf("run1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_run2(t *testing.T) {
	type args struct {
		inputText string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "empty",
			args: args {inputText: "...\n...\n..."},
			want: 0,
		},
		{
			name: "example",
			args: args {inputText: "L.LL.LL.LL\nLLLLLLL.LL\nL.L.L..L..\nLLLL.LL.LL\nL.LL.LL.LL\nL.LLLLL.LL\n..L.L.....\nLLLLLLLLLL\nL.LLLLLL.L\nL.LLLLL.LL"},
			want: 26,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run2(tt.args.inputText); got != tt.want {
				t.Errorf("run1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_computeNextCoordinateInRay(t *testing.T) {
	type args struct {
		row       int
		col       int
		direction int
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 int
	}{
		{
			name: "N",
			args: args{row: 5, col: 5, direction: N},
			want: 4,
			want1: 5,
		},
		{
			name: "E",
			args: args{row: 5, col: 5, direction: E},
			want: 5,
			want1: 6,
		},
		{
			name: "SW",
			args: args{row: 5, col: 5, direction: SW},
			want: 6,
			want1:4,
		},
	}
		for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := computeNextCoordinateInRay(tt.args.row, tt.args.col, tt.args.direction)
			if got != tt.want {
				t.Errorf("computeNextCoordinateInRay() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("computeNextCoordinateInRay() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_lookForSeat(t *testing.T) {
	type args struct {
		seats     seatsMap
		row       int
		col       int
		direction int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1",
			args: args {
				seats: parseSeats("LLL\nLLL\nLLL"),
				row: 1, col: 1, direction: N,
			},
			want: false,
		},
		{
			name: "2",
			args: args {
				seats: parseSeats("L#L\nLLL\nLLL"),
				row: 1, col: 1, direction: N,
			},
			want: true,
		},
		{
			name: "3",
			args: args {
				seats: parseSeats("....................\n...................#\n....................\n"),
				row: 1, col: 0, direction: E,
			},
			want: true,
		},
		{
			name: "4",
			args: args {
				seats: parseSeats("....................\n...................L\n....................\n"),
				row: 1, col: 0, direction: E,
			},
			want: false,
		},
		{
			name: "4",
			args: args {
				seats: parseSeats("....................\n..........L.........#\n....................\n"),
				row: 1, col: 0, direction: E,
			},
			want: false,
		},
	}
		for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lookForSeat(tt.args.seats, tt.args.row, tt.args.col, tt.args.direction); got != tt.want {
				t.Errorf("lookForSeat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_countVisibleOccupiedSeats(t *testing.T) {
	type args struct {
		seats seatsMap
		x     int
		y     int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "all around",
			args: args {
				seats: parseSeats("###\n###\n###"), x:1, y:1,
			},
			want: 8,
		},
		{
			name: "none",
			args: args {
				seats: parseSeats("LLL\nLLL\nLLL"), x:1, y:1,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countVisibleOccupiedSeats(tt.args.seats, tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("countVisibleOccupiedSeats() = %v, want %v", got, tt.want)
			}
		})
	}
}
