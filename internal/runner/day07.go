package runner

import (
	"fmt"

	"github.com/KenMGJ/advent-of-code-2021/internal/util"
)

func (r *Runner) Day07Part1(lines []string) {
	fuel := day07Runner(lines, calcFuel)
	fmt.Println(fuel)
}

func (r *Runner) Day07Part2(lines []string) {
	fuel := day07Runner(lines, calcFuelPart2)
	fmt.Println(fuel)
}

type calFuelFunc func(counts []int, pos int) int

func day07Runner(lines []string, fn calFuelFunc) int {
	ints := util.OneLineCommaSeparatedToIntSlice(lines)

	min := util.MinOfIntList(ints)
	if min != 0 {
		panic("assert min = 0")
	}
	max := util.MaxOfIntList(ints)

	counts := make([]int, max+1)
	for _, i := range ints {
		counts[i] += 1
	}

	fuel := -1
	for i := 0; i < max+1; i++ {
		f := fn(counts, i)
		if fuel != -1 && f >= fuel {
			break
		}
		fuel = f
	}

	return fuel
}

func calcFuel(counts []int, pos int) int {
	fuel := 0

	for i, c := range counts {
		fuel += c * util.AbsInt(i-pos)
	}

	return fuel
}

func calcFuelPart2(counts []int, pos int) int {
	fuel := 0

	for i, c := range counts {
		d := util.AbsInt(i - pos)
		cost := 0
		for i := 0; i <= d; i++ {
			cost += i
		}
		fuel += c * cost
	}

	return fuel
}
