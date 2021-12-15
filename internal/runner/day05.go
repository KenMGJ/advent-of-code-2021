package runner

import (
	"fmt"
	"regexp"

	"github.com/KenMGJ/advent-of-code-2021/internal/util"
)

func (r *Runner) Day05Part1(lines []string) {
	vents := parseDay05Input(lines)

	floor := make(map[Point]int)

	for _, v := range vents {
		if v.A.X == v.B.X {
			minY := util.MinInt(v.A.Y, v.B.Y)
			maxY := util.MaxInt(v.A.Y, v.B.Y)

			for i := minY; i <= maxY; i++ {
				point := Point{X: v.A.X, Y: i}

				val := floor[point]
				floor[point] = val + 1
			}
		} else if v.A.Y == v.B.Y {
			minX := util.MinInt(v.A.X, v.B.X)
			maxX := util.MaxInt(v.A.X, v.B.X)

			for i := minX; i <= maxX; i++ {
				point := Point{X: i, Y: v.A.Y}

				val := floor[point]
				floor[point] = val + 1
			}
		}
	}

	count := 0
	for _, v := range floor {
		if v >= 2 {
			count++
		}
	}

	fmt.Println(count)
}

func (r *Runner) Day05Part2(lines []string) {
	vents := parseDay05Input(lines)

	floor := make(map[Point]int)

	for _, v := range vents {
		if v.A.X == v.B.X {
			minY := util.MinInt(v.A.Y, v.B.Y)
			maxY := util.MaxInt(v.A.Y, v.B.Y)

			for i := minY; i <= maxY; i++ {
				point := Point{X: v.A.X, Y: i}

				val := floor[point]
				floor[point] = val + 1
			}
		} else if v.A.Y == v.B.Y {
			minX := util.MinInt(v.A.X, v.B.X)
			maxX := util.MaxInt(v.A.X, v.B.X)

			for i := minX; i <= maxX; i++ {
				point := Point{X: i, Y: v.A.Y}

				val := floor[point]
				floor[point] = val + 1
			}
		} else {

			diffX := util.AbsInt(v.A.X - v.B.X)
			diffY := util.AbsInt(v.A.Y - v.B.Y)

			if diffX != diffY {
				continue
			}

			xInc := 1
			if v.B.X < v.A.X {
				xInc = -1
			}

			yInc := 1
			if v.B.Y < v.A.Y {
				yInc = -1
			}

			for i := 0; i <= diffX; i++ {
				newX := v.A.X + (i * xInc)
				newY := v.A.Y + (i * yInc)
				point := Point{X: newX, Y: newY}

				val := floor[point]
				floor[point] = val + 1
			}
		}
	}

	count := 0
	for _, v := range floor {
		if v >= 2 {
			count++
		}
	}

	fmt.Println(count)
}

type Point struct {
	X int
	Y int
}

type LineSegment struct {
	A Point
	B Point
}

func parseDay05Input(lines []string) []LineSegment {

	matcher := regexp.MustCompile(`^(\d+),(\d+) -> (\d+),(\d+)$`)

	vents := []LineSegment{}

	for _, l := range lines {
		matches := matcher.FindStringSubmatch(l)

		x1 := util.MustConvertDecimalStringToInt(matches[1])
		y1 := util.MustConvertDecimalStringToInt(matches[2])
		x2 := util.MustConvertDecimalStringToInt(matches[3])
		y2 := util.MustConvertDecimalStringToInt(matches[4])

		vents = append(vents, LineSegment{
			A: Point{X: x1, Y: y1},
			B: Point{X: x2, Y: y2},
		})
	}

	return vents
}
