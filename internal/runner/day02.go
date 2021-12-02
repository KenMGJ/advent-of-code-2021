package runner

import (
	"fmt"
	"strconv"
	"strings"
)

func (r *Runner) Day02Part1(lines []string) {

	h, d := 0, 0

	for _, l := range lines {

		c, v := lineToCommandValuePair(l)

		switch c {
		case "forward":
			h += v
		case "down":
			d += v
		case "up":
			d -= v
		}
	}

	fmt.Println(h * d)
}

func (r *Runner) Day02Part2(lines []string) {

	h, d, aim := 0, 0, 0

	for _, l := range lines {

		c, v := lineToCommandValuePair(l)

		switch c {
		case "forward":
			h += v
			d += aim * v
		case "down":
			aim += v
		case "up":
			aim -= v
		}
	}

	fmt.Println(h * d)
}

func lineToCommandValuePair(line string) (string, int) {

	sl := strings.Split(line, " ")
	c := sl[0]
	v, err := strconv.Atoi(sl[1])
	if err != nil {
		panic(err)
	}

	return c, v
}
