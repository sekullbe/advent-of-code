package main

import (
	_ "embed"
	"errors"
	"fmt"
	"log"
	"slices"
)

//go:embed input.txt
var inputText string

// input is a 20k line of digits

func main() {
	fmt.Printf("Magic number: %d\n", run1(inputText))
	fmt.Println("-------------")
	fmt.Printf("Magic number: %d\n", run2(inputText))
}

type filesystem struct {
	disk            []int // the id of the file at this point
	firstBlankIndex int
	lastFileBlock   int
}

func run1(input string) int {
	fs := decodeFileSystem(input)
	// starting from the right, move blocks one at a time to the first blank space
	for fs.moveBlockFromEnd() {
	}

	return fs.computeChecksum()
}

func run2(input string) int {
	fs := decodeFileSystem(input)

	// for each file starting at the end, try to move it
	lastfid, _, err := fs.infoForBlockAt(fs.lastFileBlock)
	if err != nil {
		panic(err) // that really shouldn't happen looking at the last block
	}

	// process each file once
	idx := len(fs.disk) - 1
	for fid := lastfid; fid > 0; fid-- { // > 0 because you can't move 0
		// find the last block of the file
		for i := idx; i > 0; i-- {
			if fs.disk[i] == fid {
				idx = i
				break
			}
		}
		_, flen, err := fs.infoForBlockAt(idx)
		if err != nil {
			log.Println(err)
			continue
		}
		moved := fs.moveFileFromEnd(idx, fid, flen)
		_ = moved
		//if !moved {
		//	log.Printf("couldn't move %d", fid)
		//}
	}

	return fs.computeChecksum()
}

func decodeFileSystem(input string) filesystem {
	fs := filesystem{[]int{}, -1, -1}
	fileId := 0
	for i, r := range input {
		flen := int(r - '0')
		if flen < 0 || flen > 9 {
			//log.Fatalf("flen %c out of range", flen)
			continue // this is probably the \n at the end
		}
		var f []int
		if i%2 == 0 { // every even index is a file, every odd free space
			f = slices.Repeat([]int{fileId}, flen)
			fileId++
			// initialize the very first firstBlankIndex
			if fs.firstBlankIndex < 0 {
				fs.firstBlankIndex = flen
			}
			fs.disk = append(fs.disk, f...)
			fs.lastFileBlock = len(fs.disk) - 1
		} else {
			f = slices.Repeat([]int{-1}, flen)
			fs.disk = append(fs.disk, f...)
		}
	}
	return fs
}

// return true if a block could be moved, else false
func (fs *filesystem) moveBlockFromEnd() bool {

	if fs.firstBlankIndex >= fs.lastFileBlock {
		return false
	}
	// starting from the end, iterate backwards until id >= 0 or just use the lfb
	fileId := fs.disk[fs.lastFileBlock]
	fs.disk[fs.firstBlankIndex] = fileId
	fs.disk[fs.lastFileBlock] = -1
	// now move the pointers
	for i, n := range fs.disk[fs.firstBlankIndex:] {
		if n < 0 {
			fs.firstBlankIndex += i
			break
		}
	}
	for i := fs.lastFileBlock - 1; i >= 0; i-- {
		n := fs.disk[i]
		if n > 0 {
			fs.lastFileBlock = i
			break
		}
	}
	return true
}

// move the file at fidx as close to the beginning as possible
// return true if it could be moved, else false
// fidx is the *end* of the file block!
func (fs *filesystem) moveFileFromEnd(fidx, fid, flen int) bool {
	ffs := fs.firstFreeSpace(flen)
	if ffs < 0 || fidx < ffs {
		return false
	}
	for i := 0; i < flen; i++ {
		fs.disk[i+ffs] = fid
		if fidx-i > ffs+i { // don't mark the moved block's spaces free
			fs.disk[fidx-i] = -1
		}
	}
	if ffs == fs.firstBlankIndex {
		for i := ffs + flen; i < len(fs.disk); i++ {
			if fs.disk[i] < 0 {
				fs.firstBlankIndex = i
				break
			}
		}
	}
	//log.Printf("moved %dx%d from %d to %d", fid, flen, fidx, ffs)
	return true
}

func (fs *filesystem) infoForBlockAt(idx int) (id, length int, err error) {

	if idx < 0 || idx >= len(fs.disk) {
		return -1, -1, errors.New("out of range")
	}

	id = fs.disk[idx]
	length = 1
	for i := idx + 1; i < len(fs.disk) && fs.disk[i] == id; i++ {
		length++
	}
	for i := idx - 1; i > 0 && fs.disk[i] == id; i-- {
		length++
	}

	return id, length, nil
}

// This needs to NOT take into account that the file being moved counts as free space
// ie    ...22222 is NOT moveable to 22222...
func (fs *filesystem) firstFreeSpace(length int) int {

	if length <= 1 {
		return fs.firstBlankIndex
	}
	for i := fs.firstBlankIndex; i < len(fs.disk); i++ {
		if i+length >= len(fs.disk) { // can't go off the end
			return -1
		}
		// look ahead length blocks and see if it consists only of -1 or mover
		if slices.Equal(fs.disk[i:i+length], slices.Repeat([]int{-1}, length)) {
			return i
		}
	}
	return -1
}

// you can't actually  move ...1111 to 1111... - the free space needs to completely exist
func (fs *filesystem) firstFreeSpaceAllowingOverwrite(length int, mover int) int {

	if length <= 1 {
		return fs.firstBlankIndex
	}
	for i := fs.firstBlankIndex; i < len(fs.disk); i++ {
		if i+length > len(fs.disk) {
			return -1
		}
		// look ahead length blocks and see if it consists only of -1 or mover
		if blockContainsOnly(fs.disk[i:i+length], mover) {
			return i
		}
	}
	return -1
}

func blockContainsOnly(block []int, ok int) bool {
	for _, id := range block {
		if id > 0 && id != ok {
			return false
		}
	}
	return true
}

func (fs *filesystem) computeChecksum() int {
	cs := 0
	for i := 0; i <= fs.lastFileBlock; i++ {
		id := fs.disk[i]
		if id >= 0 {
			cs += i * id
		}
	}
	return cs
}
