package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

const (
	ROCK     = 'X'
	PAPER    = 'Y'
	SCISSORS = 'Z'

	LOSE = 'X'
	DRAW = 'Y'
	WIN  = 'Z'

	OPP_ROCK     = 'A'
	OPP_PAPER    = 'B'
	OPP_SCISSORS = 'C'
)

func partOne(lines []string) {
	total := 0
	for _, line := range lines {
		l := line[0]
		r := line[2]

		isWinner := false
		isDraw := false
		score := 0

		if r == ROCK && l == OPP_SCISSORS {
			isWinner = true
		}
		if r == PAPER && l == OPP_ROCK {
			isWinner = true
		}
		if r == SCISSORS && l == OPP_PAPER {
			isWinner = true
		}

		if r == ROCK && l == OPP_ROCK {
			isDraw = true
		}
		if r == PAPER && l == OPP_PAPER {
			isDraw = true
		}
		if r == SCISSORS && l == OPP_SCISSORS {
			isDraw = true
		}

		if isDraw {
			score += 3
		}
		if isWinner {
			score += 6
		}

		switch r {
		case ROCK:
			score += 1
		case PAPER:
			score += 2
		case SCISSORS:
			score += 3
		}

		total += score
	}

	fmt.Println("PART 1:", total)
}

type Choice struct {
	lose rune
	draw rune
	win  rune
}

func partTwo(lines []string) {
	choices := map[rune]Choice{
		OPP_ROCK:     {lose: SCISSORS, win: PAPER, draw: ROCK},
		OPP_PAPER:    {lose: ROCK, win: SCISSORS, draw: PAPER},
		OPP_SCISSORS: {lose: PAPER, win: ROCK, draw: SCISSORS},
	}

	totalScore := 0

	for _, line := range lines {
		l := rune(line[0])
		r := rune(line[2])

		choice := choices[l]
		move := choice.lose
		score := 0

		switch r {
		case 'Y':
			move = choice.draw
			score += 3
		case 'Z':
			move = choice.win
			score += 6
		}

		switch move {
		case ROCK:
			score += 1
		case PAPER:
			score += 2
		case SCISSORS:
			score += 3
		}

		totalScore += score
	}

	fmt.Println("PART 2:", totalScore)
}

func main() {
	f, _ := os.Open("input.txt")
	data, _ := ioutil.ReadAll(f)
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")

	partOne(lines)
	partTwo(lines)
}
