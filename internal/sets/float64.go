package sets

import (
	"fmt"
	"sort"
	"strings"
)

type Float64Set struct {
	set map[float64]struct{}
}

func NewFloat64Set() *Float64Set {
	return &Float64Set{
		set: make(map[float64]struct{}),
	}
}

func (f *Float64Set) Add(element float64) {
	f.set[element] = exists
}

func (f *Float64Set) AddAllFromSet(other *Float64Set) {
	for _, v := range other.Vals() {
		f.set[v] = exists
	}
}

func (f *Float64Set) Contains(element float64) bool {
	_, ok := f.set[element]
	return ok
}

func (f *Float64Set) Intersect(other *Float64Set) *Float64Set {

	intersect := NewFloat64Set()

	for _, val := range f.Vals() {
		if other.Contains(val) {
			intersect.Add(val)
		}
	}

	return intersect
}

func (f *Float64Set) ContainsAll(elements []float64) bool {
	for _, e := range elements {
		_, ok := f.set[e]
		if !ok {
			return false
		}
	}
	return true
}

func (f *Float64Set) Remove(element float64) {
	delete(f.set, element)
}

func (f *Float64Set) Size() int {
	return len(f.set)
}

func (f *Float64Set) Vals() []float64 {
	keys := []float64{}
	for k := range f.set {
		keys = append(keys, k)
	}
	sort.Float64s(keys)

	return keys
}

func (f *Float64Set) String() string {
	keys := f.Vals()
	keyStr := []string{}

	for _, k := range keys {
		keyStr = append(keyStr, fmt.Sprintf("%v", k))
	}

	return "Set{ " + strings.Join(keyStr, ", ") + " }"
}
