package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var sampleData = `$ cd /
$ ls
dir a
14848514 b.txt
8504156 c.dat
dir d
$ cd a
$ ls
dir e
29116 f
2557 g
62596 h.lst
$ cd e
$ ls
584 i
$ cd ..
$ cd ..
$ cd d
$ ls
4060174 j
8033020 d.log
5626152 d.ext
7214296 k`

func Test_filesystem_parseCommand(t *testing.T) {

	fs := initializeEmptyFilesystem()
	assert.Equal(t, "/", fs.cwd)
	fs.parseCommand("cd foo")
	assert.Equal(t, "/foo/", fs.cwd)
	fs.parseCommand("cd bar")
	assert.Equal(t, "/foo/bar/", fs.cwd)
	fs.parseCommand("ls")
	assert.Equal(t, "/foo/bar/", fs.cwd)
	fs.parseCommand("cd ..")
	assert.Equal(t, "/foo/", fs.cwd)
	fs.parseCommand("cd baz")
	assert.Equal(t, "/foo/baz/", fs.cwd)
	fs.parseCommand("cd /")
	assert.Equal(t, "/", fs.cwd)
	fs.parseCommand("cd ..")
	assert.Equal(t, "/", fs.cwd)
	fs.parseCommand("cd quux")
	assert.Equal(t, "/quux/", fs.cwd)

}

func Test_filesystem_parseLsResponse(t *testing.T) {
	fs := initializeEmptyFilesystem()
	fs.parseLsResponse("123 fileone")
	fs.parseCommand("cd foo")
	fs.parseLsResponse("2000 filetwo")
	fs.parseLsResponse("3000 filethree")
	fs.parseLsResponse("4000 filefour")

	assert.Contains(t, fs.files, "/fileone")
	assert.Equal(t, 123, fs.files["/fileone"])
	assert.Contains(t, fs.files, "/foo/filetwo")
	assert.Equal(t, 2000, fs.files["/foo/filetwo"])
	assert.Contains(t, fs.files, "/foo/filethree")
	assert.Equal(t, 3000, fs.files["/foo/filethree"])
	assert.Contains(t, fs.files, "/foo/filefour")
	assert.Equal(t, 4000, fs.files["/foo/filefour"])
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
			name: "sample",
			args: args{inputText: sampleData},
			want: 95437,
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
			name: "sample",
			args: args{inputText: sampleData},
			want: 24933642,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run2(tt.args.inputText); got != tt.want {
				t.Errorf("run2() = %v, want %v", got, tt.want)
			}
		})
	}
}
