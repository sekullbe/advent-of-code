package main

import (
	_ "embed"
	"fmt"
	"github.com/sekullbe/advent/parsers"
	"log"
	"strings"
)

//go:embed input.txt
var inputText string

func main() {
	fmt.Printf("Magic number: %d\n", run1(inputText))
	fmt.Println("-------------")
	fmt.Printf("Magic number: %d\n", run2(inputText))
}

func run1(inputText string) int {
	lines := parsers.SplitByLines(inputText)
	fs := initializeEmptyFilesystem()
	fs.evaluateCommands(lines)
	totalSize := 0
	for _, size := range fs.dirs {
		if size <= 100000 {
			totalSize += size
		}
	}
	return totalSize
}

func run2(inputText string) int {
	lines := parsers.SplitByLines(inputText)
	fs := initializeEmptyFilesystem()
	fs.evaluateCommands(lines)
	var updateSize = 30000000
	var diskSize = 70000000
	var freeSpaceAtStart = diskSize - fs.dirs["/"]
	var smallestThatsBigEnough = diskSize
	var doomed string
	for dir, size := range fs.dirs {
		if freeSpaceAtStart+size >= updateSize && size < smallestThatsBigEnough {
			smallestThatsBigEnough = size
			doomed = dir
		}
	}
	fmt.Printf("deleting: %s\n", doomed)
	return smallestThatsBigEnough
}

func (fs *filesystem) evaluateCommands(commands []string) {
	for _, line := range commands {
		// if a line begins with $ it's a command
		// else it's a response
		if line[0] == '$' {
			fs.parseCommand(line[2:]) // strip off "$ "
		} else {
			fs.parseLsResponse(line)
		}
	}
}

func (fs *filesystem) parseCommand(command string) {
	// ignore ls, we'll recognize output because it has no $
	if strings.HasPrefix("ls", command) {
		return
	}
	var cmd, arg string
	n, err := fmt.Sscan(command, &cmd, &arg)
	if n != 2 || err != nil || cmd != "cd" {
		log.Fatalf("something went badly wrong parsing command '%s': %s", command, err)
	}
	switch arg {
	case "/":
		fs.goroot()
	case "..":
		fs.popdir()
	default:
		fs.pushdir(arg)
	}
}

func (fs *filesystem) parseLsResponse(line string) {
	// we don't really care that a directory exists; that'll only matter if we move into it
	if strings.HasPrefix(line, "dir") {
		return
	}
	var filename string
	var filesize int
	n, err := fmt.Sscanf(line, "%d %s", &filesize, &filename)
	if n != 2 || err != nil {
		log.Fatalf("something went badly wrong parsing file size '%s': %s", line, err)
	}
	fs.files[fs.cwd+filename] = filesize // this isn't really necessary
	// do this for cwd and every subdir of cwd
	for d := fs.cwd; d != "/"; d = popdirname(d) {
		fs.dirs[d] += filesize
	}
	fs.dirs["/"] += filesize // because that loop stops at "/", do that manually

}
