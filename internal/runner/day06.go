package runner

import (
	"fmt"

	"github.com/KenMGJ/advent-of-code-2021/internal/util"
)

const day06Part1Iterations = 80
const day06Part2Iterations = 256

func (r *Runner) Day06Part1(lines []string) {
	runDay06(lines, day06Part1Iterations)
}

func (r *Runner) Day06Part2(lines []string) {
	runDay06(lines, day06Part2Iterations)
}

func runDay06(lines []string, iterations int) {
	fish := util.OneLineCommaSeparatedToIntSlice(lines)

	var itCount [9]int
	for _, f := range fish {
		itCount[f] += 1
	}

	for day := 0; day < iterations; day++ {
		zero := itCount[0]
		for i := 0; i < len(itCount)-1; i++ {
			itCount[i] = itCount[i+1]
		}
		itCount[8] = zero
		itCount[6] += zero
	}

	sum := 0
	for _, i := range itCount {
		sum += i
	}
	fmt.Println(sum)
}
