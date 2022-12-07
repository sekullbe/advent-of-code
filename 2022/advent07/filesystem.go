package main

import "strings"

type dirsize map[string]int
type filesize map[string]int
type filesystem struct {
	dirs  dirsize
	files filesize // This ended up not being necessary
	cwd   string   // Always maintains the final /
}

func initializeEmptyFilesystem() filesystem {
	var fs filesystem
	fs.dirs = make(dirsize)
	fs.files = make(filesize)
	fs.dirs["/"] = 0
	fs.cwd = "/"
	return fs
}

func (fs *filesystem) pushdir(dirname string) {
	fs.cwd = pushdirname(fs.cwd, dirname)
}

func (fs *filesystem) popdir() {
	fs.cwd = popdirname(fs.cwd)
}

func (fs *filesystem) goroot() {
	fs.cwd = "/"
}

// Given a path like /foo/bar/, chops off the last element
func popdirname(path string) string {
	idx := strings.LastIndex(path[0:len(path)-1], "/")
	if idx <= 0 {
		return "/" // handle the case where you're already at the root
	}
	return path[0 : idx+1]
}

func pushdirname(path, newdir string) string {
	return path + newdir + "/"
}
