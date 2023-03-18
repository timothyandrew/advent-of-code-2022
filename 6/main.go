package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func partOne(line string) {
	for i := 0; i < len(line); i++ {
		seen := make(map[byte]bool)

		if i+3 < len(line) {
			seen[line[i]] = true
			seen[line[i+1]] = true
			seen[line[i+2]] = true
			seen[line[i+3]] = true

			if len(seen) == 4 {
				fmt.Println("PART 1", i+4)
				return
			}
		}
	}
}

func partTwo(line string) {
	for i := 0; i < len(line); i++ {
		seen := make(map[byte]bool)

		if i+13 < len(line) {
			for j := 0; j < 14; j++ {
				seen[line[i+j]] = true
			}

			if len(seen) == 14 {
				fmt.Println("PART 2", i+14)
				return
			}
		}
	}
}

func main() {
	f, _ := os.Open("input.txt")
	data, _ := ioutil.ReadAll(f)
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")

	partOne(lines[0])
	partTwo(lines[0])
}
