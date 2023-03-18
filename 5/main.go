package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Stack struct {
	crates []rune
}

func (s *Stack) Push(c rune) {
	s.crates = append(s.crates, c)
}

func (s *Stack) Pop() rune {
	c := s.crates[len(s.crates)-1]
	s.crates = s.crates[:len(s.crates)-1]
	return c
}

func (s *Stack) Inspect() {
	for _, c := range s.crates {
		fmt.Print(string(c), " ")
	}

	fmt.Println("")
}

type Universe struct {
	stacks []Stack
}

func (u *Universe) Inspect() {
	fmt.Println(" ")
	for _, stack := range u.stacks {
		stack.Inspect()
	}
}

func (u *Universe) ApplyInstruction(instruction string) {
	regex := regexp.MustCompile(`move (\d+) from (\d+) to (\d+)`)
	matches := regex.FindStringSubmatch(instruction)

	count, _ := strconv.ParseInt(matches[1], 10, 64)
	from, _ := strconv.ParseInt(matches[2], 10, 64)
	to, _ := strconv.ParseInt(matches[3], 10, 64)

	for i := 0; i < int(count); i++ {
		c := u.stacks[from-1].Pop()
		u.stacks[to-1].Push(c)
	}
}

func (u *Universe) ApplyInstructionNoReorder(instruction string) {
	regex := regexp.MustCompile(`move (\d+) from (\d+) to (\d+)`)
	matches := regex.FindStringSubmatch(instruction)

	count, _ := strconv.ParseInt(matches[1], 10, 64)
	from, _ := strconv.ParseInt(matches[2], 10, 64)
	to, _ := strconv.ParseInt(matches[3], 10, 64)

	buffer := []rune{}

	for i := 0; i < int(count); i++ {
		buffer = append(buffer, u.stacks[from-1].Pop())
	}

	for i := 0; i < len(buffer); i++ {
		u.stacks[to-1].Push(buffer[len(buffer)-i-1])
	}
}

func NewUniverse(lines []string) Universe {
	u := Universe{
		stacks: make([]Stack, 9),
	}

	for j := 0; j < len(lines); j++ {
		line := lines[len(lines)-j-1]

		for i := 0; i < 9; i++ {
			cell := line[(i*4)+1]
			if cell != ' ' {
				u.stacks[i].crates = append(u.stacks[i].crates, rune(cell))
			}
		}
	}

	return u
}

func partOne(lines []string) {
	universe := NewUniverse(lines[:8])

	for _, instruction := range lines[10:] {
		universe.ApplyInstruction(instruction)
	}
	universe.Inspect()
}

func partTwo(lines []string) {
	universe := NewUniverse(lines[:8])

	for _, instruction := range lines[10:] {
		universe.ApplyInstructionNoReorder(instruction)
	}
	universe.Inspect()
}

func main() {
	f, _ := os.Open("input.txt")
	data, _ := ioutil.ReadAll(f)
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")

	partOne(lines)
	partTwo(lines)
}
