package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type file struct {
	name string
	size int
}

type directory struct {
	name   string
	parent *directory
	files  []*file
	dirs   []*directory
}

func (d *directory) addDir(name string) {
	d.dirs = append(d.dirs, &directory{name: name, parent: d})
}

func (d *directory) addFile(name string, size int) {
	d.files = append(d.files, &file{name: name, size: size})
}

func (d *directory) size() int {
	size := 0
	for _, file := range d.files {
		size += file.size
	}
	for _, dir := range d.dirs {
		size += dir.size()
	}
	return size
}

func parseFilesystem(s string) *directory {
	lines := strings.Split(s, "\n")
	root := &directory{name: "/"}
	wd := root
	for _, line := range lines {
		fields := strings.Split(line, " ")
		switch fields[0] {
		case "$":
			switch fields[1] {
			case "cd":
				switch fields[2] {
				case "/":
					wd = root
				case "..":
					wd = wd.parent
				default:
					for _, dir := range wd.dirs {
						if dir.name == fields[2] {
							wd = dir
							break
						}
					}
				}
			case "ls":
				continue
			}
		case "dir":
			wd.addDir(fields[1])
		default:
			size, err := strconv.Atoi(fields[0])
			if err != nil {
				panic("no way")
			}
			wd.addFile(fields[1], size)
		}
	}
	return root
}

func (d *directory) collectSizes(list *[]int) {
	*list = append(*list, d.size())
	for _, dir := range d.dirs {
		dir.collectSizes(list)
	}
}

// part 1: 1243729
// part 2: 4443914
func main() {
	fs := parseFilesystem(input)

	sizes := []int{}
	fs.collectSizes(&sizes)
	sort.Ints(sizes)

	sizeSum := 0
	for _, size := range sizes {
		if size <= 100_000 {
			sizeSum += size
		} else {
			break
		}
	}
	fmt.Printf("part 1: %d\n", sizeSum)

	tds := 70_000_000
	need := 30_000_000
	used := fs.size()
	mustFree := need - (tds - used)

	dirSize := 0
	for _, size := range sizes {
		if size >= mustFree {
			dirSize = size
			break
		}
	}
	fmt.Printf("part 2: %d\n", dirSize)

}
