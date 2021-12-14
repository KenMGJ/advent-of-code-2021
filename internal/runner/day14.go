package runner

import (
	"fmt"
	"regexp"
	"strings"
)

func (r *Runner) Day14Part1(lines []string) {
	template, rules := parseDay14Input(lines)
	val := runDay14(template, rules, 10)
	fmt.Println(val)
}

func (r *Runner) Day14Part2(lines []string) {
	template, rules := parseDay14Input(lines)
	val := runDay14(template, rules, 40)
	fmt.Println(val)
}

func runDay14(template string, rules map[string]string, iterations int) int {

	charCount := map[string]int{}
	pairCounts := map[string]int{}

	for _, c := range strings.Split(template, "") {
		charCount[c]++
	}

	for i := 1; i < len(template); i += 1 {
		s := template[(i - 1):(i + 1)]
		pairCounts[s]++
	}

	for i := 0; i < iterations; i++ {
		newPairCounts := map[string]int{}
		for k, v := range pairCounts {
			mid := rules[k]
			charCount[mid] += v
			newPairCounts[string(k[0])+mid] += v
			newPairCounts[mid+string(k[1])] += v
		}
		pairCounts = newPairCounts
	}

	vals := []int{}
	for _, v := range charCount {
		vals = append(vals, v)
	}

	min, max := MinAndMaxOfIntList(vals)
	return max - min
}

func parseDay14Input(lines []string) (string, map[string]string) {
	var template string
	rules := map[string]string{}

	templateMatcher := regexp.MustCompile(`^(.+)$`)
	rulesMatcher := regexp.MustCompile(`^(.+) -> (.+)$`)

	for _, l := range lines {
		matches := rulesMatcher.FindStringSubmatch(l)
		if len(matches) == 3 {
			rules[matches[1]] = matches[2]
			continue
		}

		matches = templateMatcher.FindStringSubmatch(l)
		if len(matches) == 2 {
			template = matches[1]
			continue
		}
	}

	return template, rules
}
