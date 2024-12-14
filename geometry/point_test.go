package geometry

import (
	"reflect"
	"testing"
)

func TestCalculateOffsets(t *testing.T) {
	type args struct {
		a Point2
		b Point2
	}
	tests := []struct {
		name  string
		args  args
		wantX int
		wantY int
	}{
		{name: "simple", args: args{a: Point2{2, 2}, b: Point2{4, 5}}, wantX: 2, wantY: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotX, gotY := CalculateOffsets(tt.args.a, tt.args.b)
			if gotX != tt.wantX {
				t.Errorf("CalculateOffsets() gotX = %v, want %v", gotX, tt.wantX)
			}
			if gotY != tt.wantY {
				t.Errorf("CalculateOffsets() gotY = %v, want %v", gotY, tt.wantY)
			}
		})
	}
}

func TestNewPoint2(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want Point2
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewPoint2(tt.args.x, tt.args.y); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPoint2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewPoint3(t *testing.T) {
	type args struct {
		x int
		y int
		z int
	}
	tests := []struct {
		name string
		args args
		want Point3
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewPoint3(tt.args.x, tt.args.y, tt.args.z); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPoint3() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPoint2_Dist(t *testing.T) {
	type fields struct {
		X int
		Y int
	}
	type args struct {
		y Point2
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := Point2{
				X: tt.fields.X,
				Y: tt.fields.Y,
			}
			if got := x.Dist(tt.args.y); got != tt.want {
				t.Errorf("Dist() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPoint2_DistSqr(t *testing.T) {
	type fields struct {
		X int
		Y int
	}
	type args struct {
		y Point2
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := Point2{
				X: tt.fields.X,
				Y: tt.fields.Y,
			}
			if got := x.DistSqr(tt.args.y); got != tt.want {
				t.Errorf("DistSqr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPoint2_Equal(t *testing.T) {
	type fields struct {
		X int
		Y int
	}
	type args struct {
		y Point2
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := Point2{
				X: tt.fields.X,
				Y: tt.fields.Y,
			}
			if got := x.Equal(tt.args.y); got != tt.want {
				t.Errorf("Equal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPoint2_MovePoint2(t *testing.T) {
	type fields struct {
		X int
		Y int
	}
	type args struct {
		dx int
		dy int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Point2
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Point2{
				X: tt.fields.X,
				Y: tt.fields.Y,
			}
			if got := p.MovePoint2(tt.args.dx, tt.args.dy); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MovePoint2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPoint3_Dist(t *testing.T) {
	type fields struct {
		X int
		Y int
		Z int
	}
	type args struct {
		y Point3
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := Point3{
				X: tt.fields.X,
				Y: tt.fields.Y,
				Z: tt.fields.Z,
			}
			if got := x.Dist(tt.args.y); got != tt.want {
				t.Errorf("Dist() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPoint3_DistSqr(t *testing.T) {
	type fields struct {
		X int
		Y int
		Z int
	}
	type args struct {
		y Point3
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := Point3{
				X: tt.fields.X,
				Y: tt.fields.Y,
				Z: tt.fields.Z,
			}
			if got := x.DistSqr(tt.args.y); got != tt.want {
				t.Errorf("DistSqr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPoint3_Equal(t *testing.T) {
	type fields struct {
		X int
		Y int
		Z int
	}
	type args struct {
		y Point3
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := Point3{
				X: tt.fields.X,
				Y: tt.fields.Y,
				Z: tt.fields.Z,
			}
			if got := x.Equal(tt.args.y); got != tt.want {
				t.Errorf("Equal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPoint3_MovePoint3(t *testing.T) {
	type fields struct {
		X int
		Y int
		Z int
	}
	type args struct {
		dx int
		dy int
		dz int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Point3
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Point3{
				X: tt.fields.X,
				Y: tt.fields.Y,
				Z: tt.fields.Z,
			}
			if got := p.MovePoint3(tt.args.dx, tt.args.dy, tt.args.dz); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MovePoint3() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPoint2_MovePoint2WithWrap(t *testing.T) {
	type fields struct {
		X int
		Y int
	}
	type args struct {
		dx   int
		dy   int
		maxX int
		maxY int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Point2
	}{
		{name: "simple0", fields: fields{1, 1}, args: args{dx: 5, dy: 6, maxX: 30, maxY: 30}, want: Point2{X: 6, Y: 7}},
		{name: "simple1", fields: fields{1, 1}, args: args{dx: 5, dy: 6, maxX: 3, maxY: 3}, want: Point2{X: 2, Y: 3}},
		{name: "simple2", fields: fields{1, 1}, args: args{dx: 6, dy: 7, maxX: 3, maxY: 3}, want: Point2{X: 3, Y: 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Point2{
				X: tt.fields.X,
				Y: tt.fields.Y,
			}
			if got := p.MovePoint2WithWrap(tt.args.dx, tt.args.dy, tt.args.maxX, tt.args.maxY); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MovePoint2WithWrap() = %v, want %v", got, tt.want)
			}
		})
	}
}
