package runner

import (
	"fmt"
	"strconv"
	"strings"
)

func AbsInt(a int) int {
	if a < 0 {
		return -1 * a
	}
	return a
}

func BinaryStringsToInts(lines []string) []int64 {

	ints := []int64{}

	for _, line := range lines {
		i, err := strconv.ParseInt(line, 2, 64)
		if err != nil {
			panic(err)
		}
		ints = append(ints, i)
	}

	return ints
}

func MaxInt(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func MaxOfIntList(ints []int) int {
	if len(ints) < 1 {
		panic("empty list")
	}

	max := ints[0]
	for i := 1; i < len(ints); i++ {
		if ints[i] > max {
			max = ints[i]
		}
	}

	return max
}

func MinInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func MinOfIntList(ints []int) int {
	if len(ints) < 1 {
		panic("empty list")
	}

	min := ints[0]
	for i := 1; i < len(ints); i++ {
		if ints[i] < min {
			min = ints[i]
		}
	}

	return min
}

func MustConvertDecimalStringToInt(s string) int {
	v, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return v
}

func PrintCharArrayArray(charsArr [][]rune) {
	for _, chars := range charsArr {
		for _, c := range chars {
			if c == '0' {
				fmt.Print(0)
			} else if c == '1' {
				fmt.Print(1)
			}
		}
		fmt.Print(",")
	}
	fmt.Println()
}

func StringArrayToCharArrayArray(lines []string) [][]rune {
	arr := [][]rune{}
	for _, l := range lines {
		arr = append(arr, []rune(l))
	}
	return arr
}

func OneLineCommaSeparatedToIntSlice(lines []string) []int {

	ints := []int{}

	firstLine := strings.Split(lines[0], ",")
	for _, l := range firstLine {
		i, err := strconv.Atoi(l)
		if err != nil {
			panic(err)
		}
		ints = append(ints, i)
	}

	return ints
}
