package sets

import (
	"sort"
	"strconv"
	"strings"
)

type IntSet struct {
	set map[int]struct{}
}

func NewIntSet() *IntSet {
	return &IntSet{
		set: make(map[int]struct{}),
	}
}

func (i *IntSet) Add(element int) {
	i.set[element] = exists
}

func (i *IntSet) AddAllFromSet(other *IntSet) {
	for _, v := range other.Vals() {
		i.set[v] = exists
	}
}

func (i *IntSet) Contains(element int) bool {
	_, ok := i.set[element]
	return ok
}

func (i *IntSet) ContainsAll(elements []int) bool {
	for _, e := range elements {
		_, ok := i.set[e]
		if !ok {
			return false
		}
	}
	return true
}

func (i *IntSet) Intersect(other *IntSet) *IntSet {

	intersect := NewIntSet()

	for _, val := range i.Vals() {
		if other.Contains(val) {
			intersect.Add(val)
		}
	}

	return intersect
}

func (i *IntSet) Remove(element int) {
	delete(i.set, element)
}

func (i *IntSet) Size() int {
	return len(i.set)
}

func (i *IntSet) Vals() []int {
	keys := []int{}
	for k := range i.set {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	return keys
}

func (i *IntSet) String() string {
	keys := i.Vals()
	ks := []string{}
	for _, k := range keys {
		ks = append(ks, strconv.Itoa(k))
	}
	return "IntSet{ " + strings.Join(ks, ", ") + " }"
}
