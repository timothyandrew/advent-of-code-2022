package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	f, _ := os.Open("input.txt")
	data, _ := ioutil.ReadAll(f)
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")

	max := 0
	totals := []int{}
	running := 0

	for _, line := range lines {
		if line == "" {
			if running > max {
				max = running
			}

			totals = append(totals, running)
			running = 0

			continue
		}

		n, _ := strconv.ParseInt(line, 10, 64)
		running += int(n)
	}

	fmt.Println("PART 1:", max)

	sort.Slice(totals, func(i, j int) bool { return totals[j] < totals[i] })
	sum := 0
	for _, n := range totals[:3] {
		sum += n
	}

	fmt.Println("PART 2:", sum)
}
