package advent5

import "testing"

func Test_parseInputRow(t *testing.T) {
	type args struct {
		seatCode string
	}
	tests := []struct {
		name    string
		args    args
		wantRow int64
		wantCol int64
	}{
		{
			name: "example1",
			args: args{seatCode: "FBFBBFFRLR"},
			wantRow: 44,
			wantCol: 5,
		},
		{
			name: "example2",
			args: args{seatCode: "BFFFBBFRRR"},
			wantRow: 70,
			wantCol: 7,
		},
		{
			name: "example3",
			args: args{seatCode: "FFFBBBFRRR"},
			wantRow: 14,
			wantCol: 7,
		},
		{
			name: "example4",
			args: args{seatCode: "BBFFBBFRLL"},
			wantRow: 102,
			wantCol: 4,
		},
		{
			name: "zero",
			args: args{seatCode: "FFFFFFFLLL"},
			wantRow: 0,
			wantCol: 0,
		},
		{
			name: "one",
			args: args{seatCode: "BBBBBBBRRR"},
			wantRow: 127,
			wantCol: 7,
		},



	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRow, gotCol := parseInputRow(tt.args.seatCode)
			if gotRow != tt.wantRow {
				t.Errorf("parseInputRow() gotRow = %v, want %v", gotRow, tt.wantRow)
			}
			if gotCol != tt.wantCol {
				t.Errorf("parseInputRow() gotCol = %v, want %v", gotCol, tt.wantCol)
			}
		})
	}
}
