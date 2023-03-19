package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type Map struct {
	heights    [][]int
	visibility [][]int
	distance   [][]int
}

func (m *Map) SetVisibility(i, j int, value int) {
	if m.visibility[i][j] == 1 {
		return
	}

	m.visibility[i][j] = value
}

func (m *Map) AnnotateDistance() {
	for i := 0; i < len(m.heights); i++ {
		for j := 0; j < len(m.heights[0]); j++ {
			cell := m.heights[i][j]

			left := 0
			right := 0
			top := 0
			bottom := 0

			for current := j - 1; current >= 0; current-- {
				left++

				if m.heights[i][current] >= cell {
					break
				}
			}

			for current := j + 1; current < len(m.heights[0]); current++ {
				right++

				if m.heights[i][current] >= cell {
					break
				}
			}

			for current := i - 1; current >= 0; current-- {
				top++

				if m.heights[current][j] >= cell {
					break
				}
			}

			for current := i + 1; current < len(m.heights); current++ {
				bottom++

				if m.heights[current][j] >= cell {
					break
				}
			}

			distance := top * bottom * left * right
			m.distance[i][j] = distance
		}
	}
}

func (m *Map) AnnotateVisibility() {
	width := len(m.heights[0])
	height := len(m.heights)

	for i := 1; i < len(m.heights); i++ {
		maxLeft := m.heights[i][0]
		maxRight := m.heights[i][width-1]
		maxTop := m.heights[0][i]
		maxBottom := m.heights[height-1][i]

		for j := 1; j < len(m.heights[i]); j++ {
			if m.heights[i][j] > maxLeft {
				m.SetVisibility(i, j, 1)
				maxLeft = m.heights[i][j]
			} else if m.heights[i][j] == maxLeft {
				m.SetVisibility(i, j, 0)
			}

			if m.heights[i][width-j-1] > maxRight {
				m.SetVisibility(i, width-j-1, 1)
				maxRight = m.heights[i][width-j-1]
			} else if m.heights[i][width-j-1] == maxRight {
				m.SetVisibility(i, width-j-1, 0)
			}

			if m.heights[j][i] > maxTop {
				m.SetVisibility(j, i, 1)
				maxTop = m.heights[j][i]
			} else if m.heights[i][j] == maxTop {
				m.SetVisibility(j, i, 0)
			}

			if m.heights[height-j-1][i] > maxBottom {
				m.SetVisibility(height-j-1, i, 1)
				maxBottom = m.heights[height-j-1][i]
			} else if m.heights[i][height-j-1] == maxBottom {
				m.SetVisibility(height-j-1, i, 0)
			}
		}
	}

}

func (m *Map) Inspect() {
	fmt.Println("HEIGHTS")
	for _, row := range m.heights {
		for _, cell := range row {
			fmt.Print(cell, " ")
		}

		fmt.Println()
	}

	fmt.Println("\nVISIBILITY")
	for _, row := range m.visibility {
		for _, cell := range row {
			fmt.Print(cell, " ")
		}

		fmt.Println()
	}

	fmt.Println("\nDISTANCE")
	for _, row := range m.distance {
		for _, cell := range row {
			fmt.Print(cell, " ")
		}

		fmt.Println()
	}
}

func NewMap(lines []string) Map {
	heights := [][]int{}
	visibility := [][]int{}
	distance := [][]int{}

	for i, line := range lines {
		h := []int{}
		v := []int{}
		d := []int{}

		for j, c := range line {
			n, _ := strconv.ParseInt(string(c), 10, 64)
			h = append(h, int(n))
			d = append(d, 0)

			if i == 0 || j == 0 {
				v = append(v, 1)
			} else if i == len(lines)-1 || j == len(line)-1 {
				v = append(v, 1)
			} else {
				v = append(v, 0)
			}
		}

		heights = append(heights, h)
		visibility = append(visibility, v)
		distance = append(distance, d)
	}

	return Map{
		heights:    heights,
		visibility: visibility,
		distance:   distance,
	}
}

func partOne(lines []string) {
	m := NewMap(lines)
	m.AnnotateVisibility()

	total := 0
	for _, row := range m.visibility {
		for _, cell := range row {
			if cell > 0 {
				total++
			}
		}
	}

	fmt.Println("PART 1", total)
}

func partTwo(lines []string) {
	m := NewMap(lines)
	m.AnnotateVisibility()
	m.AnnotateDistance()
	// m.Inspect()

	max := 0
	for _, row := range m.distance {
		for _, cell := range row {
			if cell > max {
				max = cell
			}
		}
	}

	fmt.Println("PART 2", max)
}

func main() {
	f, _ := os.Open("input.txt")
	data, _ := ioutil.ReadAll(f)
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")

	partOne(lines)
	partTwo(lines)
}
