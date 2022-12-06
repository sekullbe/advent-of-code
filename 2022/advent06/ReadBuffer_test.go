package main

import (
	"reflect"
	"testing"
)

func TestNewReadBuffer(t *testing.T) {
	type args struct {
		size int
	}
	tests := []struct {
		name string
		args args
		want *Readbuffer
	}{
		{name: "Example1", args: args{size: 4}, want: &Readbuffer{size: 4, buf: []rune{0, 0, 0, 0}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewReadBuffer(tt.args.size); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewReadBuffer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReadbuffer_add(t *testing.T) {
	type fields struct {
		size int
		buf  []rune
	}
	type args struct {
		r rune
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{name: "Example1", fields: fields{size: 4, buf: []rune{'b', 'b', 'b', 'b'}}, args: args{r: 'a'}, want: true},
		{name: "Example1", fields: fields{size: 4, buf: []rune{0, 0, 0, 0}}, args: args{r: 'b'}, want: true},
		{name: "Example1", fields: fields{size: 4, buf: []rune{0, 0, 0, 0}}, args: args{r: 'c'}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Readbuffer{
				size: tt.fields.size,
				buf:  tt.fields.buf,
			}
			if got := b.add(tt.args.r); got != tt.want {
				t.Errorf("add() = %v, want %v", got, tt.want)
			}
			if b.buf[3] != tt.args.r {
				t.Errorf("didn't append %c", tt.args.r)
			}
		})
	}
}

func TestReadbuffer_contentsAllUnique(t *testing.T) {
	type fields struct {
		size int
		buf  []rune
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{name: "Example1", fields: fields{size: 4, buf: []rune{'b', 'b', 'b', 'b'}}, want: false},
		{name: "Example2", fields: fields{size: 4, buf: []rune{'b', 'b', 'b', 'c'}}, want: false},
		{name: "Example3", fields: fields{size: 4, buf: []rune{'a', 'b', 'c', 'd'}}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Readbuffer{
				size: tt.fields.size,
				buf:  tt.fields.buf,
			}
			if got := b.contentsAllUnique(); got != tt.want {
				t.Errorf("contentsAllUnique() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReadbuffer_countRunes(t *testing.T) {
	type fields struct {
		size int
		buf  []rune
	}
	type args struct {
		r rune
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		wantCount int
	}{
		{name: "Example1", fields: fields{size: 4, buf: []rune{'b', 'b', 'b', 'b'}}, args: args{r: 'b'}, wantCount: 4},
		{name: "Example2", fields: fields{size: 4, buf: []rune{'a', 'b', 'c', 'd'}}, args: args{r: 'b'}, wantCount: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Readbuffer{
				size: tt.fields.size,
				buf:  tt.fields.buf,
			}
			if gotCount := b.countRunes(tt.args.r); gotCount != tt.wantCount {
				t.Errorf("countRunes() = %v, want %v", gotCount, tt.wantCount)
			}
		})
	}
}
