package sets

import (
	"sort"
	"strings"
)

var exists = struct{}{}

type StringSet struct {
	set map[string]struct{}
}

func NewStringSet() *StringSet {
	return &StringSet{
		set: make(map[string]struct{}),
	}
}

func (s *StringSet) Add(element string) {
	s.set[element] = exists
}

func (s *StringSet) AddAllFromSet(other *StringSet) {
	for _, v := range other.Vals() {
		s.set[v] = exists
	}
}

func (s *StringSet) Contains(element string) bool {
	_, ok := s.set[element]
	return ok
}

func (s *StringSet) ContainsAll(elements []string) bool {
	for _, e := range elements {
		_, ok := s.set[e]
		if !ok {
			return false
		}
	}
	return true
}

func (s *StringSet) Remove(element string) {
	delete(s.set, element)
}

func (s *StringSet) Size() int {
	return len(s.set)
}

func (s *StringSet) Vals() []string {
	keys := []string{}
	for k := range s.set {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	return keys
}

func (s *StringSet) String() string {
	keys := s.Vals()
	return "Set{ " + strings.Join(keys, ", ") + " }"
}
