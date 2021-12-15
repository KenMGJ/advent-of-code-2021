package util

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
	_, max := MinAndMaxOfIntList(ints)
	return max
}

func MinInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func MinAndMaxOfIntList(ints []int) (int, int) {
	if len(ints) < 1 {
		panic("empty list")
	}

	min, max := ints[0], ints[0]
	for i := 1; i < len(ints); i++ {
		if ints[i] < min {
			min = ints[i]
		} else if ints[i] > max {
			max = ints[i]
		}
	}

	return min, max
}

func MinOfIntList(ints []int) int {
	min, _ := MinAndMaxOfIntList(ints)
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

func StringArrayToIntArrayArray(lines []string) [][]int {
	arr := [][]int{}
	for _, l := range lines {
		intLine := []int{}
		for _, c := range strings.Split(l, "") {
			intLine = append(intLine, MustConvertDecimalStringToInt(c))
		}
		arr = append(arr, intLine)
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
