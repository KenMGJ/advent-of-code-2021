package runner

import (
	"fmt"
	"strconv"
)

func (r *Runner) Day03Part1(lines []string) {
	chars := stringArrayToCharArrayArray(lines)
	diags := bitCounts(chars)

	gamma, epsilon := 0, 0
	for _, d := range diags {
		gamma = gamma << 1
		epsilon = epsilon << 1
		if d.OneCount < d.ZeroCount {
			gamma += 1
		} else {
			epsilon += 1
		}
	}

	fmt.Println(gamma * epsilon)
}

func (r *Runner) Day03Part2(lines []string) {
	chars := stringArrayToCharArrayArray(lines)

	oRating, err := filterUntilFound(chars, true)
	if err != nil {
		panic(err)
	}

	cRating, err := filterUntilFound(chars, false)
	if err != nil {
		panic(err)
	}

	fmt.Println(oRating * cRating)
}

func filterUntilFound(chars [][]rune, keepMostCommon bool) (int64, error) {

	filteredChars := [][]rune{}
	for _, c := range chars {
		item := []rune{}
		item = append(item, c...)
		filteredChars = append(filteredChars, item)
	}

	diags := bitCounts(chars)

	for i := 0; i < len(diags); i++ {
		if len(filteredChars) < 2 {
			break
		}

		f := [][]rune{}
		for _, c := range filteredChars {

			if keepMostCommon {
				if diags[i].OneCount >= diags[i].ZeroCount {
					if c[i] == '1' {
						f = append(f, c)
					}
				} else {
					if c[i] == '0' {
						f = append(f, c)
					}
				}

			} else {

				if !keepMostCommon {
					if diags[i].OneCount >= diags[i].ZeroCount {
						if c[i] == '0' {
							f = append(f, c)
						}
					} else {
						if c[i] == '1' {
							f = append(f, c)
						}
					}
				}
			}

		}

		filteredChars = f
		diags = bitCounts(filteredChars)
	}

	return strconv.ParseInt(string(filteredChars[0]), 2, 64)
}

type binaryDiagnostic struct {
	ZeroCount int
	OneCount  int
}

func bitCounts(lines [][]rune) []binaryDiagnostic {

	diags := []binaryDiagnostic{}

	for i, line := range lines {
		for j, c := range line {
			if i == 0 {
				diags = append(diags, binaryDiagnostic{})
			}

			if c == '0' {
				diags[j].ZeroCount += 1
			} else if c == '1' {
				diags[j].OneCount += 1
			}
		}
	}

	return diags
}

func stringArrayToCharArrayArray(lines []string) [][]rune {
	arr := [][]rune{}
	for _, l := range lines {
		arr = append(arr, []rune(l))
	}
	return arr
}

/*
func printCharArrayArray(charsArr [][]rune) {
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

func binaryStringsToInts(lines []string) []int64 {

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
*/
