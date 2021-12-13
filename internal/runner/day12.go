package runner

import (
	"fmt"
	"regexp"
	"strings"
)

var day12Matcher = regexp.MustCompile(`^[a-z]+$`)
var day12Nodes map[string]*Node

func (r *Runner) Day12Part1(lines []string) {
	day12Nodes = parseDay12Input(lines)
	paths := traverseDay12("start", map[string]int{}, false)

	fmt.Println(len(paths))
}

func (r *Runner) Day12Part2(lines []string) {
	day12Nodes = parseDay12Input(lines)
	paths := traverseDay12("start", map[string]int{}, true)

	fmt.Println(len(paths))
}

func traverseDay12(start string, seen map[string]int, atMost bool) []string {

	node := day12Nodes[start]

	if node.Value == "end" {
		return []string{"end"}
	}

	paths := []string{}

	for _, n := range node.Children {
		value := n.Value
		if value == "start" {
			continue
		}

		_, ok := seen[value]
		if ok && !atMost {
			continue
		}

		s := make(map[string]int)
		for k, v := range seen {
			s[k] = v
		}
		if day12Matcher.MatchString(node.Value) {
			newValue := 1
			_, ok := s[node.Value]
			if ok {
				newValue++
			}
			s[node.Value] = newValue
		}

		if atMost {
			found := false
			for _, v := range s {
				if v > 1 {
					found = true
					break
				}
			}
			if ok && found {
				continue
			}
		}

		for _, t := range traverseDay12(value, s, atMost) {
			path := node.Value + "," + t
			paths = append(paths, path)
		}
	}

	return paths
}

type Node struct {
	Value    string
	Children []*Node
}

func parseDay12Input(lines []string) map[string]*Node {

	nodes := make(map[string]*Node)

	for _, l := range lines {
		v := strings.Split(l, "-")
		p := v[0]
		c := v[1]

		parent, ok := nodes[p]
		if !ok {
			parent = &Node{
				Value:    p,
				Children: []*Node{},
			}
		}

		child, ok := nodes[c]
		if !ok {
			child = &Node{
				Value:    c,
				Children: []*Node{},
			}
		}

		parent.Children = append(parent.Children, child)
		child.Children = append(child.Children, parent)
		nodes[p] = parent
		nodes[c] = child
	}

	return nodes
}
