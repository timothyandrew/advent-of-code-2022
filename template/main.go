package main

import (
	"io/ioutil"
	"os"
	"strings"
)

func partOne(lines []string) {
}

func partTwo(lines []string) {
}

func main() {
	f, _ := os.Open("input.txt")
	data, _ := ioutil.ReadAll(f)
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")

	partOne(lines)
	partTwo(lines)
}
