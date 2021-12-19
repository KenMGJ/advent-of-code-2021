package snailfish

import "fmt"

type Number struct {
	parent Element
	Left   Element
	Right  Element
}

func (n *Number) SetParent(parent Element) {
	n.parent = parent
}

func (n *Number) Parent() Element {
	return n.parent
}

func (n *Number) SetValue(value int) {
}

func (n *Number) Value() int {
	return 0
}

func (s *Number) String() string {
	return fmt.Sprintf("[%v,%v]", s.Left, s.Right)
}

type NumberStack struct {
	stack []*Number
}

func NewNumberStack() *NumberStack {
	n := []*Number{}
	return &NumberStack{
		stack: n,
	}
}

func (n *NumberStack) IsEmpty() bool {
	return len(n.stack) == 0
}

func (n *NumberStack) Peek() *Number {
	var top *Number
	l := len(n.stack)
	if len(n.stack) > 0 {
		top = n.stack[l-1]
	}
	return top
}

func (n *NumberStack) Pop() *Number {
	var top *Number
	l := len(n.stack)
	if len(n.stack) > 0 {
		top = n.stack[l-1]
		n.stack = n.stack[:l-1]
	}
	return top
}

func (n *NumberStack) Push(num *Number) {
	n.stack = append(n.stack, num)
}

func (n *NumberStack) PushNew() {
	n.stack = append(n.stack, &Number{})
}
