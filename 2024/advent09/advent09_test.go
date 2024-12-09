package main

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"slices"
	"testing"
)

const sampleText = `2333133121414131402`

func Test_run1(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "simple", args: args{input: "12345"}, want: 60},
		{name: "sampletext", args: args{input: sampleText}, want: 1928},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run1(tt.args.input); got != tt.want {
				t.Errorf("run1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_run2(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "simple", args: args{input: "12345"}, want: 132}, // this one doesn't actually move at all
		{name: "simple", args: args{input: "12355"}, want: 92},  // this one moves the 2s to the left
		{name: "sampletext", args: args{input: sampleText}, want: 2858},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run2(tt.args.input); got != tt.want {
				t.Errorf("run2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_decodeFileSystem(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want filesystem
	}{
		{name: "simple", args: args{input: "12345"}, want: filesystem{
			disk:            []int{0, -1, -1, 1, 1, 1, -1, -1, -1, -1, 2, 2, 2, 2, 2},
			firstBlankIndex: 1, lastFileBlock: 14},
		},
		{name: "simple with end padding", args: args{input: "123452"}, want: filesystem{
			disk:            []int{0, -1, -1, 1, 1, 1, -1, -1, -1, -1, 2, 2, 2, 2, 2, -1, -1},
			firstBlankIndex: 1, lastFileBlock: 14},
		},
		{name: "sample", args: args{input: sampleText}, want: filesystem{
			disk:            []int{0, 0, -1, -1, -1, 1, 1, 1, -1, -1, -1, 2, -1, -1, -1, 3, 3, 3, -1, 4, 4, -1, 5, 5, 5, 5, -1, 6, 6, 6, 6, -1, 7, 7, 7, -1, 8, 8, 8, 8, 9, 9},
			firstBlankIndex: 2, lastFileBlock: 41},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := decodeFileSystem(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("decodeFileSystem() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_filesystem_moveBlockFromEnd(t *testing.T) {
	fs := filesystem{disk: []int{0, -1, -1, 1, 1, 1, -1, -1, -1, -1, 2, 2, 2, 2, 2}, firstBlankIndex: 1, lastFileBlock: 14}
	fs.moveBlockFromEnd()
	if !slices.Equal(fs.disk, []int{0, 2, -1, 1, 1, 1, -1, -1, -1, -1, 2, 2, 2, 2, -1}) {
		t.Errorf("moveBlockFromEnd() = %v, want %v", fs.disk, []int{0, 2, -1, 1, 1, 1, -1, -1, -1, -1, 2, 2, 2, 2, -1})
	}
	if fs.firstBlankIndex != 2 {
		t.Errorf("firstBlankIndex = %v, want 2", fs.firstBlankIndex)
	}
	if fs.lastFileBlock != 13 {
		t.Errorf("lastFileBlock = %v, want 13", fs.lastFileBlock)
	}

	fs.moveBlockFromEnd()
	if !slices.Equal(fs.disk, []int{0, 2, 2, 1, 1, 1, -1, -1, -1, -1, 2, 2, 2, -1, -1}) {
		t.Errorf("moveBlockFromEnd() = %v, want %v", fs.disk, []int{0, 2, 2, 1, 1, 1, -1, -1, -1, -1, 2, 2, 2, -1, -1})
	}
	if fs.firstBlankIndex != 6 {
		t.Errorf("firstBlankIndex = %v, want 2", fs.firstBlankIndex)
	}
	if fs.lastFileBlock != 12 {
		t.Errorf("lastFileBlock = %v, want 13", fs.lastFileBlock)
	}

	fs.moveBlockFromEnd()
	if !slices.Equal(fs.disk, []int{0, 2, 2, 1, 1, 1, 2, -1, -1, -1, 2, 2, -1, -1, -1}) {
		t.Errorf("moveBlockFromEnd() = %v, want %v", fs.disk, []int{0, 2, 2, 1, 1, 1, 2, -1, -1, -1, 2, 2, -1, -1, -1})
	}
	if fs.firstBlankIndex != 7 {
		t.Errorf("firstBlankIndex = %v, want 2", fs.firstBlankIndex)
	}
	if fs.lastFileBlock != 11 {
		t.Errorf("lastFileBlock = %v, want 13", fs.lastFileBlock)
	}

}

func Test_filesystem_computeChecksum(t *testing.T) {
	fs := filesystem{disk: []int{0, 2, 2, 1, 1, 1, 2, 2, 2, -1, -1, -1, -1, -1, -1}, firstBlankIndex: 9, lastFileBlock: 8}
	cs := fs.computeChecksum()
	if cs != 0+2+4+3+4+5+12+14+16 {
		t.Errorf("computeChecksum() = %v, want %v", cs, 0+2+4+3+4+5+12+14+16)
	}

	fs = filesystem{disk: []int{0, 0, 9, 9, 8, 1, 1, 1, 8, 8, 8, 2, 7, 7, 7, 3, 3, 3, 6, 4, 4, 6, 5, 5, 5, 5, 6, 6, -1, -1, -1, -1, -1, -1}, firstBlankIndex: 28, lastFileBlock: 27}
	cs = fs.computeChecksum()
	if cs != 1928 {
		t.Errorf("computeChecksum() = %v, want %v", cs, 1928)
	}

}

func Test_filesystem_infoForBlockAt(t *testing.T) {
	type fields struct {
		disk            []int
		firstBlankIndex int
		lastFileBlock   int
	}
	type args struct {
		idx int
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantId     int
		wantLength int
		wantErr    bool
	}{
		{name: "sample0",
			fields: fields{
				disk:            []int{0, -1, -1, 1, 1, 1, -1, -1, -1, -1, 2, 2, 2, 2, 2},
				firstBlankIndex: 1,
				lastFileBlock:   14},
			args:   args{0},
			wantId: 0, wantLength: 1, wantErr: false,
		},
		{name: "sample1",
			fields: fields{
				disk:            []int{0, -1, -1, 1, 1, 1, -1, -1, -1, -1, 2, 2, 2, 2, 2},
				firstBlankIndex: 1,
				lastFileBlock:   14},
			args:   args{1},
			wantId: -1, wantLength: 2, wantErr: false,
		},
		{name: "sample4",
			fields: fields{
				disk:            []int{0, -1, -1, 1, 1, 1, -1, -1, -1, -1, 2, 2, 2, 2, 2},
				firstBlankIndex: 1,
				lastFileBlock:   14},
			args:   args{4},
			wantId: 1, wantLength: 3, wantErr: false,
		},
		{name: "sample14",
			fields: fields{
				disk:            []int{0, -1, -1, 1, 1, 1, -1, -1, -1, -1, 2, 2, 2, 2, 2},
				firstBlankIndex: 1,
				lastFileBlock:   14},
			args:   args{14},
			wantId: 2, wantLength: 5, wantErr: false,
		},
		{name: "sample15 off the end",
			fields: fields{
				disk:            []int{0, -1, -1, 1, 1, 1, -1, -1, -1, -1, 2, 2, 2, 2, 2},
				firstBlankIndex: 1,
				lastFileBlock:   14},
			args:   args{15},
			wantId: -1, wantLength: -1, wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fs := &filesystem{
				disk:            tt.fields.disk,
				firstBlankIndex: tt.fields.firstBlankIndex,
				lastFileBlock:   tt.fields.lastFileBlock,
			}
			gotId, gotLength, err := fs.infoForBlockAt(tt.args.idx)
			if (err != nil) != tt.wantErr {
				t.Errorf("infoForBlockAt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotId != tt.wantId {
				t.Errorf("infoForBlockAt() gotId = %v, want %v", gotId, tt.wantId)
			}
			if gotLength != tt.wantLength {
				t.Errorf("infoForBlockAt() gotLength = %v, want %v", gotLength, tt.wantLength)
			}
		})
	}
}

func Test_filesystem_firstFreeSpace(t *testing.T) {
	type fields struct {
		disk            []int
		firstBlankIndex int
		lastFileBlock   int
	}
	type args struct {
		length int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{name: "sample2",
			fields: fields{
				disk:            []int{0, -1, -1, 1, 1, 1, -1, -1, -1, -1, 2, 2, 2, 2, 2},
				firstBlankIndex: 1,
				lastFileBlock:   14},
			args: args{2},
			want: 1,
		},
		{name: "sample4",
			fields: fields{
				disk:            []int{0, -1, -1, 1, 1, 1, -1, -1, -1, -1, 2, 2, 2, 2, 2},
				firstBlankIndex: 1,
				lastFileBlock:   14},
			args: args{4},
			want: 6,
		},
		{name: "sample5 too long",
			fields: fields{
				disk:            []int{0, -1, -1, 1, 1, 1, -1, -1, -1, -1, 2, 2, 2, 2, 2},
				firstBlankIndex: 1,
				lastFileBlock:   14},
			args: args{5},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fs := &filesystem{
				disk:            tt.fields.disk,
				firstBlankIndex: tt.fields.firstBlankIndex,
				lastFileBlock:   tt.fields.lastFileBlock,
			}
			if got := fs.firstFreeSpace(tt.args.length); got != tt.want {
				t.Errorf("firstFreeSpace() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_filesystem_firstFreeSpaceAllowingOverwrite(t *testing.T) {
	type fields struct {
		disk            []int
		firstBlankIndex int
		lastFileBlock   int
	}
	type args struct {
		length int
		mover  int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{name: "sample2",
			fields: fields{
				disk:            []int{0, -1, -1, 1, 1, 1, -1, -1, -1, -1, 2, 2, 2, 2, 2},
				firstBlankIndex: 1,
				lastFileBlock:   14},
			args: args{2, 1},
			want: 1,
		},
		{name: "sample41",
			fields: fields{
				disk:            []int{0, -1, -1, 1, 1, 1, -1, -1, -1, -1, 2, 2, 2, 2, 2},
				firstBlankIndex: 1,
				lastFileBlock:   14},
			args: args{4, 1},
			want: 1,
		},
		{name: "sample42",
			fields: fields{
				disk:            []int{0, -1, -1, 1, 1, 1, -1, -1, -1, -1, 2, 2, 2, 2, 2},
				firstBlankIndex: 1,
				lastFileBlock:   14},
			args: args{4, 2},
			want: 6,
		},
		{name: "sample5 too long",
			fields: fields{
				disk:            []int{0, -1, -1, 1, 1, 1, -1, -1, -1, -1, 2, 2, 2, 2, 2},
				firstBlankIndex: 1,
				lastFileBlock:   14},
			args: args{5, 3},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fs := &filesystem{
				disk:            tt.fields.disk,
				firstBlankIndex: tt.fields.firstBlankIndex,
				lastFileBlock:   tt.fields.lastFileBlock,
			}
			if got := fs.firstFreeSpaceAllowingOverwrite(tt.args.length, tt.args.mover); got != tt.want {
				t.Errorf("firstFreeSpace() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_filesystem_moveFileFromEnd(t *testing.T) {

	// note this fs has an extra -1 before the 2s
	fs := filesystem{
		disk:            []int{0, -1, -1, 1, 1, 1, -1, -1, -1, -1, -1, 2, 2, 2, 2, 2},
		firstBlankIndex: 1,
		lastFileBlock:   15,
	}

	moved := fs.moveFileFromEnd(15, 2, 5)
	assert.True(t, moved)
	assert.Equal(t, []int{0, -1, -1, 1, 1, 1, 2, 2, 2, 2, 2, -1, -1, -1, -1, -1}, fs.disk)
	assert.Equal(t, fs.firstBlankIndex, 1)

	moved = fs.moveFileFromEnd(5, 1, 3)
	assert.False(t, moved)
	assert.Equal(t, []int{0, -1, -1, 1, 1, 1, 2, 2, 2, 2, 2, -1, -1, -1, -1, -1}, fs.disk)
	assert.Equal(t, fs.firstBlankIndex, 1)
}

// 22222....
