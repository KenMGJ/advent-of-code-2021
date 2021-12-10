package runner

import (
	"fmt"
	"sort"
	"strings"

	"github.com/KenMGJ/advent-of-code-2021/internal/stacks"
)

var filteredDay10 []string

func (r *Runner) Day10Part1(lines []string) {

	counts := make(map[string]int)

	filteredDay10 = []string{}
	for _, l := range lines {
		corrupt, char := isCorruptDay10Line(l)
		if corrupt {
			counts[char]++
		} else {
			filteredDay10 = append(filteredDay10, l)
		}
	}

	score := 0
	for k, v := range counts {
		switch k {
		case ")":
			score += (3 * v)
		case "]":
			score += (57 * v)
		case "}":
			score += (1197 * v)
		case ">":
			score += (25137 * v)
		}
	}
	fmt.Println(score)
}

func (r *Runner) Day10Part2(lines []string) {

	scores := []int{}

	for _, l := range filteredDay10 {
		chars := completeDay10Line(l)

		total := 0

		for _, c := range chars {
			total *= 5

			val := 0
			switch c {
			case ")":
				val = 1
			case "]":
				val = 2
			case "}":
				val = 3
			case ">":
				val = 4
			}

			total += val
		}

		scores = append(scores, total)
	}

	sort.Ints(scores)
	fmt.Println(scores[len(scores)/2])
}

func completeDay10Line(line string) []string {

	stack := stacks.NewStringStack()
	out := []string{}

	for _, l := range strings.Split(line, "") {
		if isOpenCharDay10(l) {
			stack.Push(l)
		} else {
			stack.Pop()
		}
	}

	for !stack.IsEmpty() {
		out = append(out, matchingCloseCharDay10(stack.Pop()))
	}

	return out
}

func isCorruptDay10Line(line string) (bool, string) {

	stack := stacks.NewStringStack()
	for _, l := range strings.Split(line, "") {
		if isOpenCharDay10(l) {
			stack.Push(l)
		} else if stack.Peek() == matchingOpenCharDay10(l) {
			stack.Pop()
		} else {
			return true, l
		}
	}

	return false, ""
}

func matchingCloseCharDay10(char string) string {
	switch char {
	case "(":
		return ")"
	case "[":
		return "]"
	case "{":
		return "}"
	case "<":
		return ">"
	}
	return ""
}

func matchingOpenCharDay10(char string) string {
	switch char {
	case ")":
		return "("
	case "]":
		return "["
	case "}":
		return "{"
	case ">":
		return "<"
	}
	return ""
}

func isOpenCharDay10(char string) bool {
	return char == "(" || char == "[" || char == "{" || char == "<"
}
