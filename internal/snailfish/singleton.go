package snailfish

import "fmt"

type Singleton struct {
	parent Element
	Val    int
}

func NewSingleton(parent Element, value int) *Singleton {
	return &Singleton{
		parent: parent,
		Val:    value,
	}
}

func (s *Singleton) SetParent(parent Element) {
	s.parent = parent
}

func (s *Singleton) Parent() Element {
	return s.parent
}

func (s *Singleton) SetValue(value int) {
	s.Val = value
}

func (s *Singleton) Value() int {
	return s.Val
}

func (s *Singleton) String() string {
	return fmt.Sprintf("%d", s.Val)
}
