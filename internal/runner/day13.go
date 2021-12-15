package runner

import (
	"fmt"
	"regexp"

	"github.com/KenMGJ/advent-of-code-2021/internal/pairs"
	"github.com/KenMGJ/advent-of-code-2021/internal/util"
)

func (r *Runner) Day13Part1(lines []string) {
	paper, folds := parseDay13Input(lines)

	for i := 0; i < 1; i++ {
		fold := folds[i]
		if fold.Along == FOLD_OVER_Y {
			paper = foldUp(paper, fold.Value)
		} else if fold.Along == FOLD_OVER_X {
			paper = foldLeft(paper, fold.Value)
		}
	}

	count := 0
	for _, l := range paper {
		for _, c := range l {
			if c {
				count++
			}
		}
	}

	fmt.Println(count)
}

func (r *Runner) Day13Part2(lines []string) {
	paper, folds := parseDay13Input(lines)

	for _, fold := range folds {
		if fold.Along == FOLD_OVER_Y {
			paper = foldUp(paper, fold.Value)
		} else if fold.Along == FOLD_OVER_X {
			paper = foldLeft(paper, fold.Value)
		}
	}

	PrintSliceOfBoolSlice(paper)
}

func foldUp(v [][]bool, y int) [][]bool {

	height := len(v)
	temp := [][]bool{}

	i := y - 1
	oi := y + 1

	for oi < height || i >= 0 {
		line := []bool{}
		for j := 0; j < len(v[0]); j++ {
			a := false
			if oi < height {
				a = v[oi][j]
			}
			b := false
			if i >= 0 {
				b = v[i][j]
			}

			if a || b {
				line = append(line, true)
			} else {
				line = append(line, false)
			}
		}
		temp = append(temp, line)
		i--
		oi++
	}

	new := [][]bool{}
	for i := len(temp) - 1; i >= 0; i-- {
		new = append(new, temp[i])
	}

	return new
}

func foldLeft(v [][]bool, x int) [][]bool {

	new := [][]bool{}
	for i := 0; i < len(v); i++ {
		width := len(v[i])
		line := []bool{}
		j := x + 1
		oj := x - 1
		for j < width || oj >= 0 {
			a := false
			if j < width && v[i][j] {
				a = true
			}
			b := false
			if oj >= 0 && v[i][oj] {
				a = true
			}

			if a || b {
				line = append(line, true)
			} else {
				line = append(line, false)
			}
			j++
			oj--
		}

		newLine := []bool{}
		for i := len(line) - 1; i >= 0; i-- {
			newLine = append(newLine, line[i])
		}

		new = append(new, newLine)
	}

	return new
}

func PrintSliceOfBoolSlice(v [][]bool) {
	for _, line := range v {
		for _, char := range line {
			if char {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

type day13FoldOver string

const FOLD_OVER_Y = day13FoldOver("y")
const FOLD_OVER_X = day13FoldOver("x")

type day13Fold struct {
	Along day13FoldOver
	Value int
}

func parseDay13Input(lines []string) ([][]bool, []day13Fold) {

	coordinateMatcher := regexp.MustCompile(`^(\d+),(\d+)`)
	foldAlongYMatcher := regexp.MustCompile(`^fold along y=(\d+)$`)
	foldAlongXMatcher := regexp.MustCompile(`^fold along x=(\d+)$`)

	coordinates := []pairs.IntPair{}
	folds := []day13Fold{}

	width := 0
	height := 0

	for _, line := range lines {
		matches := coordinateMatcher.FindStringSubmatch(line)
		if len(matches) == 3 {
			a := util.MustConvertDecimalStringToInt(matches[1])
			if a > width {
				width = a
			}

			b := util.MustConvertDecimalStringToInt(matches[2])
			if b > height {
				height = b
			}

			coordinates = append(coordinates, pairs.IntPair{A: a, B: b})
			continue
		}

		matches = foldAlongYMatcher.FindStringSubmatch(line)
		if len(matches) == 2 {
			val := util.MustConvertDecimalStringToInt(matches[1])
			folds = append(folds, day13Fold{Along: FOLD_OVER_Y, Value: val})
			continue
		}

		matches = foldAlongXMatcher.FindStringSubmatch(line)
		if len(matches) == 2 {
			val := util.MustConvertDecimalStringToInt(matches[1])
			folds = append(folds, day13Fold{Along: FOLD_OVER_X, Value: val})
			continue
		}
	}

	width++
	height++

	paper := [][]bool{}

	for i := 0; i < height; i++ {
		line := []bool{}
		for j := 0; j < width; j++ {
			line = append(line, false)
		}
		paper = append(paper, line)
	}

	for _, p := range coordinates {
		paper[p.B][p.A] = true
	}

	return paper, folds
}
