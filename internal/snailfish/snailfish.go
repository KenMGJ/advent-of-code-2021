package snailfish

type Element interface {
	Parent() Element
	SetParent(Element)
	SetValue(int)
	Value() int
}
