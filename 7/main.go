package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"strings"
)

type File struct {
	parent *Dir
	size   int
}

type Dir struct {
	parent *Dir
	files  map[string]*File
	dirs   map[string]*Dir
}

func NewDir(parent *Dir) *Dir {
	return &Dir{
		parent: parent,
		files:  make(map[string]*File),
		dirs:   make(map[string]*Dir),
	}

}

func (d *Dir) Sizes() (ret []int) {
	total := 0

	for _, file := range d.files {
		total += file.size
	}

	for _, dir := range d.dirs {
		ret = append(ret, dir.Sizes()...)
		total += ret[len(ret)-1]
	}

	ret = append(ret, total)

	return
}

func (d *Dir) Inspect(level int) {
	for k, v := range d.dirs {
		for i := 0; i < level; i++ {
			fmt.Print("  ")
		}
		fmt.Println(k, "(dir)")

		v.Inspect(level + 1)
	}

	for k, v := range d.files {
		for i := 0; i < level; i++ {
			fmt.Print("  ")
		}
		fmt.Println(k, v.size)
	}
}

type Navigation struct {
	cwd  *Dir
	root *Dir
}

func (n *Navigation) ChangeDir(target string) {
	switch target {
	case "..":
		n.cwd = n.cwd.parent
	case "/":
		n.cwd = n.root
	default:
		n.cwd = n.cwd.dirs[target]
	}
}

func (n *Navigation) ParseListing(listing string) {
	if strings.HasPrefix(listing, "dir") {
		dirname := listing[4:]
		if _, ok := n.cwd.dirs[dirname]; !ok {
			n.cwd.dirs[dirname] = NewDir(n.cwd)
		}
	} else {
		split := strings.Split(listing, " ")

		size, _ := strconv.ParseInt(split[0], 10, 64)
		filename := split[1]

		if _, ok := n.cwd.files[filename]; !ok {
			n.cwd.files[filename] = &File{
				size:   int(size),
				parent: n.cwd,
			}
		}
	}
}

func (n *Navigation) Inspect() {
	fmt.Println("/")
	n.root.Inspect(1)
}

func NewNavigation() Navigation {
	root := NewDir(nil)

	return Navigation{
		cwd:  root,
		root: root,
	}
}

func partOne(lines []string) {
	navigation := NewNavigation()

	for i := 0; i < len(lines); i++ {
		line := lines[i]

		if strings.HasPrefix(line, "$ cd") {
			target := line[5:]
			navigation.ChangeDir(target)
		}

		if strings.HasPrefix(line, "$ ls") {
			i++

			for i < len(lines) && !strings.HasPrefix(lines[i], "$") {
				navigation.ParseListing(lines[i])
				i++
			}
			i--
		}
	}

	// navigation.Inspect()
	sizes := navigation.root.Sizes()
	answer := 0

	for _, size := range sizes {
		if size <= 100000 {
			answer += size
		}
	}

	fmt.Println("PART 1", answer)
}

func partTwo(lines []string) {
	navigation := NewNavigation()

	for i := 0; i < len(lines); i++ {
		line := lines[i]

		if strings.HasPrefix(line, "$ cd") {
			target := line[5:]
			navigation.ChangeDir(target)
		}

		if strings.HasPrefix(line, "$ ls") {
			i++

			for i < len(lines) && !strings.HasPrefix(lines[i], "$") {
				navigation.ParseListing(lines[i])
				i++
			}
			i--
		}
	}

	sizes := navigation.root.Sizes()

	total := 70_000_000
	used := sizes[len(sizes)-1]
	unused := total - used
	required := 30_000_000
	toFree := required - unused

	min := math.MaxInt

	for _, size := range sizes {
		if size < min && size >= toFree {
			min = size
		}
	}

	fmt.Println("PART 2", min)
}

func main() {
	f, _ := os.Open("input.txt")
	data, _ := ioutil.ReadAll(f)
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")

	partOne(lines)
	partTwo(lines)
}
