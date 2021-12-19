package runner

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/KenMGJ/advent-of-code-2021/internal/snailfish"
	"github.com/KenMGJ/advent-of-code-2021/internal/stacks"
	"github.com/KenMGJ/advent-of-code-2021/internal/util"
)

func (r *Runner) Day18Part1(lines []string) {
	nums := parseDay18(lines)

	a := nums[0]
	for i := 1; i < len(nums); i++ {
		b := nums[i]

		n := &snailfish.Number{}
		a.SetParent(n)
		n.Left = a
		b.SetParent(n)
		n.Right = b

		reduce(n)

		a = n
	}

	magnitude := calcMagnitude(a)
	fmt.Println(magnitude)
}

func (r *Runner) Day18Part2(lines []string) {
	nums := parseDay18(lines)

	max := 0
	for i := 0; i < len(nums)-1; i++ {
		for j := 1; j < len(nums); j++ {
			nums = parseDay18(lines)
			a := nums[i]
			b := nums[j]

			n := &snailfish.Number{}
			a.SetParent(n)
			n.Left = a
			b.SetParent(n)
			n.Right = b

			reduce(n)

			magnitude := calcMagnitude(n)
			if magnitude > max {
				max = magnitude
			}

			nums = parseDay18(lines)
			a = nums[j]
			b = nums[i]

			n = &snailfish.Number{}
			a.SetParent(n)
			n.Left = a
			b.SetParent(n)
			n.Right = b

			reduce(n)

			magnitude = calcMagnitude(n)
			if magnitude > max {
				max = magnitude
			}
		}
	}

	fmt.Println(max)
}

func calcMagnitude(e snailfish.Element) int {
	switch v := e.(type) {
	case *snailfish.Singleton:
		return v.Value()
	case *snailfish.Number:
		return 3*calcMagnitude(v.Left) + 2*calcMagnitude(v.Right)
	}

	return 0
}

func reduce(n *snailfish.Number) {

	for {
		vals := traverseSnailfishNumber(n, 0)
		exploded := explode(vals)
		if exploded {
			continue
		}

		didSplit := split(vals)
		if didSplit {
			continue
		} else {
			break
		}
	}
}

func split(vals []ValueLevelPair) bool {
	for _, v := range vals {
		if v.IsSingleton && v.Value >= 10 {

			newValLeft := v.Value / 2
			newValRight := (v.Value / 2) + (v.Value % 2)

			newNumber := &snailfish.Number{}
			newNumber.Left = snailfish.NewSingleton(newNumber, newValLeft)
			newNumber.Right = snailfish.NewSingleton(newNumber, newValRight)

			parent := v.Self.Parent()
			newNumber.SetParent(parent)

			p, ok := parent.(*snailfish.Number)
			if !ok {
				panic("wrong type")
			}

			if p.Left == v.Self {
				p.Left = newNumber
			} else if p.Right == v.Self {
				p.Right = newNumber
			}

			return true
		}
	}
	return false
}

func explode(vals []ValueLevelPair) bool {
	for i, v := range vals {
		if v.Level == 4 {
			right := vals[i+1]
			if right.Level != 4 {
				panic("level mismatch")
			}

			for j := i - 1; j >= 0; j-- {
				next := vals[j]
				if next.IsSingleton {
					next.Self.SetValue(next.Self.Value() + v.Value)
					break
				}
			}

			for j := i + 2; j < len(vals); j++ {
				next := vals[j]
				if next.IsSingleton {
					next.Self.SetValue(next.Self.Value() + right.Value)
					break
				}
			}

			parent := v.Self.Parent()
			grandParent := parent.Parent()

			gp, ok := grandParent.(*snailfish.Number)
			if !ok {
				panic("wrong type")
			}

			if gp.Left == parent {
				gp.Left = snailfish.NewSingleton(gp, 0)
			} else if gp.Right == parent {
				gp.Right = snailfish.NewSingleton(gp, 0)
			}
			return true
		}
	}
	return false
}

type ValueLevelPair struct {
	Value       int
	Level       int
	IsSingleton bool
	Self        snailfish.Element
}

func traverseSnailfishNumber(n *snailfish.Number, level int) []ValueLevelPair {
	vals := []ValueLevelPair{
		{
			IsSingleton: false,
			Self:        n,
		},
	}

	switch left := n.Left.(type) {
	case *snailfish.Singleton:
		vals = append(vals, ValueLevelPair{Value: left.Val, Level: level, IsSingleton: true, Self: left})
	case *snailfish.Number:
		vals = append(vals, traverseSnailfishNumber(left, level+1)...)
	}

	switch right := n.Right.(type) {
	case *snailfish.Singleton:
		vals = append(vals, ValueLevelPair{Value: right.Val, Level: level, IsSingleton: true, Self: right})
	case *snailfish.Number:
		vals = append(vals, traverseSnailfishNumber(right, level+1)...)
	}

	return vals
}

func parseDay18(lines []string) []*snailfish.Number {
	numbers := []*snailfish.Number{}
	for _, l := range lines {
		numbers = append(numbers, parseSnailfishNumber(l))
	}
	return numbers
}

var snailfishMatcher = regexp.MustCompile(`(\d*),(\d*)`)

func parseSnailfishNumber(line string) *snailfish.Number {

	charStack := stacks.NewStringStack()
	numberStack := snailfish.NewNumberStack()

	for _, c := range strings.Split(line, "") {

		if c == "[" {
			numberStack.PushNew()
		} else if c == "]" {

			s := ""

			ch := charStack.Pop()
			for ch != "[" {
				s = ch + s
				ch = charStack.Pop()
			}

			matches := snailfishMatcher.FindStringSubmatch(s)

			var left snailfish.Element
			var right snailfish.Element

			popLeft := len(matches[1]) == 0
			popRight := len(matches[2]) == 0

			if popRight {
				right = numberStack.Pop()
			} else {
				right = &snailfish.Singleton{
					Val: util.MustConvertDecimalStringToInt(matches[2]),
				}
			}

			if popLeft {
				left = numberStack.Pop()
			} else {
				left = &snailfish.Singleton{
					Val: util.MustConvertDecimalStringToInt(matches[1]),
				}
			}

			current := numberStack.Pop()
			left.SetParent(current)
			right.SetParent(current)
			current.Left = left
			current.Right = right
			numberStack.Push(current)

			continue
		}

		charStack.Push(c)
	}

	return numberStack.Pop()
}
