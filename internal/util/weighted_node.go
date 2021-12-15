package util

import "fmt"

type WeightedNode struct {
	X      int
	Y      int
	Weight int
}

func NewWeightedNode(x, y, weight int) *WeightedNode {
	return &WeightedNode{X: x, Y: y, Weight: weight}
}

func (w *WeightedNode) String() string {
	return fmt.Sprintf("{ x: %d, y: %d, w: %d }", w.X, w.Y, w.Weight)
}

const weightedNodeKeyFormat = "%d,%d"

func (w *WeightedNode) Key() string {
	return fmt.Sprintf(weightedNodeKeyFormat, w.X, w.Y)
}

func WeightedNodeKey(x, y int) string {
	return fmt.Sprintf(weightedNodeKeyFormat, x, y)
}

type WeightedNodePriorityQueue struct {
	nodes map[int][]*WeightedNode
}

func NewWeightedNodePriorityQueue() *WeightedNodePriorityQueue {
	nodes := map[int][]*WeightedNode{}
	return &WeightedNodePriorityQueue{
		nodes: nodes,
	}
}

func (w *WeightedNodePriorityQueue) IsEmpty() bool {
	return len(w.nodes) == 0
}

func (w *WeightedNodePriorityQueue) Insert(node *WeightedNode, priority int) {
	priorityNodes, ok := w.nodes[priority]
	if !ok {
		priorityNodes = []*WeightedNode{}
	}

	priorityNodes = append(priorityNodes, node)
	w.nodes[priority] = priorityNodes
}

func (w *WeightedNodePriorityQueue) Get() *WeightedNode {
	p := []int{}
	for k, _ := range w.nodes {
		p = append(p, k)
	}
	min := MinOfIntList(p)

	node := w.nodes[min][0]
	w.nodes[min] = w.nodes[min][1:]
	if len(w.nodes[min]) == 0 {
		delete(w.nodes, min)
	}
	return node
}
