package stacks

type StringStack struct {
	stack []string
}

func NewStringStack() *StringStack {
	s := []string{}
	return &StringStack{
		stack: s,
	}
}

func (s *StringStack) IsEmpty() bool {
	return len(s.stack) == 0
}

func (s *StringStack) Peek() string {
	top := ""
	l := len(s.stack)
	if len(s.stack) > 0 {
		top = s.stack[l-1]
	}
	return top
}

func (s *StringStack) Pop() string {
	top := ""
	l := len(s.stack)
	if len(s.stack) > 0 {
		top = s.stack[l-1]
		s.stack = s.stack[:l-1]
	}
	return top
}

func (s *StringStack) Push(str string) {
	s.stack = append(s.stack, str)
}
