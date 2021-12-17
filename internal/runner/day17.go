package runner

import (
	"fmt"
	"regexp"

	"github.com/KenMGJ/advent-of-code-2021/internal/util"
)

func (r *Runner) Day17Part1(lines []string) {
	target := parseDay17(lines)

	max := 0
	for i := 0; i < 10000; i++ {
		for j := target.MinY; j < 10000; j++ {
			m, ok := withTargetArea(target, i, j)
			if ok {
				if m > max {
					max = m
				}
			}
		}
	}

	fmt.Println(max)
}

func (r *Runner) Day17Part2(lines []string) {
	target := parseDay17(lines)

	count := 0
	for i := 0; i < 10000; i++ {
		for j := target.MinY; j < 10000; j++ {
			_, ok := withTargetArea(target, i, j)
			if ok {
				count++
			}
		}
	}

	fmt.Println(count)
}

func withTargetArea(target day17Target, vX, vY int) (int, bool) {

	x, y, maxY := 0, 0, 0

	for {
		x += vX
		y += vY

		if y > maxY {
			maxY = y
		}

		if x >= target.MinX && x <= target.MaxX && y >= target.MinY && y <= target.MaxY {
			return maxY, true
		}

		if x > target.MaxX || y < target.MinY {
			return 0, false
		}
		if vX > 0 {
			vX--
		} else if vX < 0 {
			vX++
		}
		vY--
	}
}

type day17Target struct {
	MinX int
	MaxX int
	MinY int
	MaxY int
}

func parseDay17(lines []string) day17Target {

	line := lines[0]

	matcher := regexp.MustCompile(`^target area: x=(-?\d+)..(-?\d+), y=(-?\d+)..(-?\d+)$`)
	matches := matcher.FindStringSubmatch(line)

	t := day17Target{
		MinX: util.MustConvertDecimalStringToInt(matches[1]),
		MaxX: util.MustConvertDecimalStringToInt(matches[2]),
		MinY: util.MustConvertDecimalStringToInt(matches[3]),
		MaxY: util.MustConvertDecimalStringToInt(matches[4]),
	}

	return t
}
