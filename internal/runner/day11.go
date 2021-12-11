package runner

import (
	"fmt"
	"strings"

	"github.com/KenMGJ/advent-of-code-2021/internal/pairs"
)

func (r *Runner) Day11Part1(lines []string) {
	grid := parseOctopuses(lines)

	total := 0
	for times := 0; times < 100; times++ {
		total += day11RunOneIteration(grid)
	}

	fmt.Println(total)
}

func (r *Runner) Day11Part2(lines []string) {
	grid := parseOctopuses(lines)

	size := len(grid)
	size *= size

	times := 0
	for {
		count := day11RunOneIteration(grid)

		times++
		if count == size {
			break
		}
	}

	fmt.Println(times)
}

func day11RunOneIteration(grid [][]Octopus) int {

	// step 0 : reset flash
	for i, r := range grid {
		for j := range r {
			grid[i][j].Flashed = false
		}
	}

	// step 1 : increase energy level
	for i, r := range grid {
		for j := range r {
			grid[i][j].Level++
		}
	}

	// step 2 : flash
	for i, r := range grid {
		for j := range r {
			day11FlashAndIncAdjacent(grid, i, j)
		}
	}

	// step 3 : reset level of flashed
	for i, r := range grid {
		for j := range r {
			if grid[i][j].Flashed {
				grid[i][j].Level = 0
			}
		}
	}

	// step 4 : count flashes
	count := 0
	for _, r := range grid {
		for _, o := range r {
			if o.Flashed {
				count++
			}
		}
	}

	return count
}

func day11FlashAndIncAdjacent(grid [][]Octopus, row, col int) {
	octopus := grid[row][col]
	if octopus.Flashed {
		return
	}

	if octopus.Level > 9 {
		grid[row][col].Flashed = true
		h := len(grid)
		w := len(grid[row])
		for _, o := range pairs.AdjacentPairs(row, col, h, w) {
			grid[o.A][o.B].Level++
			day11FlashAndIncAdjacent(grid, o.A, o.B)
		}
	}
}

type Octopus struct {
	Level   int
	Flashed bool
}

func parseOctopuses(lines []string) [][]Octopus {
	octopusGrid := [][]Octopus{}

	for _, l := range lines {
		octopusRow := []Octopus{}
		for _, s := range strings.Split(l, "") {
			v := MustConvertDecimalStringToInt(s)
			octopusRow = append(octopusRow, Octopus{Level: v})
		}
		octopusGrid = append(octopusGrid, octopusRow)
	}

	return octopusGrid
}
