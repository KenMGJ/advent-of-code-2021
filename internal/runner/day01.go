package runner

import "fmt"

func (r *Runner) Day01Part1(lines []string) {

	numberLines := stringToNumber(lines)

	prev := -1
	inc := 0
	dec := 0
	for _, l := range numberLines {

		if prev != -1 {
			if l > prev {
				inc += 1
			} else {
				dec += 1
			}
		}

		prev = l
	}

	fmt.Printf("%d\n", inc)
}

func (r *Runner) Day01Part2(lines []string) {

	numberLines := stringToNumber(lines)

	prev := -1
	inc := 0
	dec := 0

	for i := 0; i < len(numberLines)-2; i++ {

		sum := numberLines[i] + numberLines[i+1] + numberLines[i+2]

		if prev != -1 {
			if sum > prev {
				inc += 1
			} else if sum < prev {
				dec += 1
			}
		}

		prev = sum
	}

	fmt.Printf("%d\n", inc)
}
