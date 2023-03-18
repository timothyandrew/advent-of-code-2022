package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type Range struct {
	min, max int64
}

func (r1 Range) Within(r2 Range) bool {
	return r1.min >= r2.min && r1.max <= r2.max
}

func (r1 Range) Overlap(r2 Range) bool {
	if r1.max < r2.min || r2.max < r1.min {
		return false
	}

	return true
}

type Assignments struct {
	left, right Range
}

func parse(line string) Assignments {
	split := strings.Split(line, ",")

	left := strings.Split(split[0], "-")
	right := strings.Split(split[1], "-")

	leftMin, _ := strconv.ParseInt(left[0], 10, 64)
	leftMax, _ := strconv.ParseInt(left[1], 10, 64)
	rightMin, _ := strconv.ParseInt(right[0], 10, 64)
	rightMax, _ := strconv.ParseInt(right[1], 10, 64)

	return Assignments{
		left:  Range{leftMin, leftMax},
		right: Range{rightMin, rightMax},
	}
}

func partOne(lines []string) {
	total := 0
	for _, line := range lines {
		assn := parse(line)

		if assn.left.Within(assn.right) || assn.right.Within(assn.left) {
			total++
		}
	}

	fmt.Println("PART 1:", total)
}

func partTwo(lines []string) {
	total := 0
	for _, line := range lines {
		assn := parse(line)
		if assn.left.Overlap(assn.right) {
			total++
		}
	}

	fmt.Println("PART 2:", total)
}

func main() {
	f, _ := os.Open("input.txt")
	data, _ := ioutil.ReadAll(f)
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")

	partOne(lines)
	partTwo(lines)
}
