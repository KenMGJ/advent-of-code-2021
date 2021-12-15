package runner

import (
	"fmt"

	"github.com/KenMGJ/advent-of-code-2021/internal/util"
)

func (r *Runner) Day15Part1(lines []string) {
	nodesByKey, h, w := parseCaveNodeGraph(lines)

	start := util.WeightedNodeKey(0, 0)
	goal := util.WeightedNodeKey(h-1, w-1)
	uniformCostSearch(nodesByKey, start, goal)
}

func (r *Runner) Day15Part2(lines []string) {
	nodesByKey, h, w := parseCaveNodeGraphPart2(lines)

	start := util.WeightedNodeKey(0, 0)
	goal := util.WeightedNodeKey(h-1, w-1)
	uniformCostSearch(nodesByKey, start, goal)
}

func parseCaveNodeGraph(lines []string) (map[string]*util.WeightedNode, int, int) {

	height := len(lines)
	width := len(lines[0])

	nodes := []*util.WeightedNode{}

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			weight := util.MustConvertDecimalStringToInt(string(lines[i][j]))
			nodes = append(nodes, util.NewWeightedNode(i, j, weight))
		}
	}

	nodesByKey := map[string]*util.WeightedNode{}
	for _, node := range nodes {
		nodesByKey[node.Key()] = node
	}

	return nodesByKey, height, width
}

func parseCaveNodeGraphPart2(lines []string) (map[string]*util.WeightedNode, int, int) {

	vals := util.StringArrayToIntArrayArray(lines)

	oldHeight := len(vals)
	height := oldHeight * 5
	oldWidth := len(vals[0])
	width := oldWidth * 5

	new := [][]int{}
	for i := 0; i < height; i++ {
		newLine := []int{}
		for j := 0; j < width; j++ {

			idxA := i % oldHeight
			idxB := j % oldWidth

			val := vals[idxA][idxB] + (i / oldHeight) + (j / oldWidth)
			if val > 9 {
				val = val - 9
			}

			newLine = append(newLine, val)
		}
		new = append(new, newLine)
	}

	newLines := []string{}
	for _, i := range new {
		line := ""
		for _, j := range i {
			line = fmt.Sprintf("%s%d", line, j)
		}
		newLines = append(newLines, line)
	}

	return parseCaveNodeGraph(newLines)
}

func uniformCostSearch(nodesByKey map[string]*util.WeightedNode, start, goal string) {

	startNode := nodesByKey[start]
	goalNode := nodesByKey[goal]

	frontier := util.NewWeightedNodePriorityQueue()

	frontier.Insert(startNode, 0)

	cameFrom := map[string]*util.WeightedNode{}
	costSoFar := map[string]int{}

	cameFrom[start] = nil
	costSoFar[start] = 0

	for !frontier.IsEmpty() {
		current := frontier.Get()

		if current == goalNode {
			break
		}

		for _, next := range neighbors(nodesByKey, current) {
			newCost := costSoFar[current.Key()] + next.Weight
			nextKey := next.Key()
			nextCost, ok := costSoFar[nextKey]

			if !ok || newCost < nextCost {
				costSoFar[nextKey] = newCost
				priority := newCost
				frontier.Insert(next, priority)
				cameFrom[nextKey] = current
			}
		}
	}

	fmt.Println(costSoFar[goal])
}

func neighbors(nodesByKey map[string]*util.WeightedNode, node *util.WeightedNode) []*util.WeightedNode {

	neighbors := []*util.WeightedNode{}

	up := util.WeightedNodeKey(node.X-1, node.Y)
	upNode, ok := nodesByKey[up]
	if ok {
		neighbors = append(neighbors, upNode)
	}

	right := util.WeightedNodeKey(node.X, node.Y+1)
	rightNode, ok := nodesByKey[right]
	if ok {
		neighbors = append(neighbors, rightNode)
	}

	down := util.WeightedNodeKey(node.X+1, node.Y)
	downNode, ok := nodesByKey[down]
	if ok {
		neighbors = append(neighbors, downNode)
	}

	left := util.WeightedNodeKey(node.X, node.Y-1)
	leftNode, ok := nodesByKey[left]
	if ok {
		neighbors = append(neighbors, leftNode)
	}

	return neighbors
}
