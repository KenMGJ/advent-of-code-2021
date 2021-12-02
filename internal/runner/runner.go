package runner

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strconv"
)

type Runner struct{}

func New() *Runner {
	return &Runner{}
}

func (r *Runner) Run(day int, test bool) error {

	var dirname string
	if test {
		dirname = "test"
	} else {
		dirname = "input"
	}

	filename := fmt.Sprintf("%s/%02d.txt", dirname, day)

	lines, err := readFile(filename)
	if err != nil {
		return err
	}

	funcNamePart1 := fmt.Sprintf("Day%02dPart1", day)
	funcNamePart2 := fmt.Sprintf("Day%02dPart2", day)

	input := make([]reflect.Value, 1)
	input[0] = reflect.ValueOf(lines)
	reflect.ValueOf(r).MethodByName(funcNamePart1).Call(input)
	reflect.ValueOf(r).MethodByName(funcNamePart2).Call(input)

	return nil
}

func readFile(filename string) ([]string, error) {

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	lines := []string{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

func stringToNumber(lines []string) []int {
	numberLines := []int{}
	for _, l := range lines {
		v, err := strconv.Atoi(l)
		if err != nil {
			break
		}

		numberLines = append(numberLines, v)
	}
	return numberLines
}
