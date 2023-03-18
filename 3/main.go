package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func calcPriority(c rune) int {
	if c >= 'a' && c <= 'z' {
		return int(c) - 96
	}

	if c >= 'A' && c <= 'Z' {
		return int(c) - 38
	}

	return 0
}

func partOne(lines []string) {
	total := 0

OUTER:
	for _, line := range lines {
		mid := len(line) / 2
		left := line[:mid]
		right := line[mid:]

		seen := make(map[rune]bool)

		for _, c := range left {
			seen[c] = true
		}

		for _, c := range right {
			if _, ok := seen[c]; ok {
				total += calcPriority(c)
				continue OUTER
			}
		}
	}

	fmt.Println("PART 1:", total)
}

func partTwo(lines []string) {
	total := 0

OUTER:
	for i := 0; i < len(lines); i++ {
		seen := make(map[rune]int8)

		for ; i < i+3; i++ {
			line := lines[i]
			for _, c := range line {
				// Each of the lowest three bits of the int represents
				// one of the three rucksacks in each group. If all three
				// bits are raised (0x07), all three rucksacks contain
				// that rune.
				seen[c] |= 1 << (i % 3)

				if seen[c] == 0x07 {
					total += calcPriority(c)
					continue OUTER
				}
			}
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
