package set

import (
	"fmt"
	"strings"
)

type Set[T comparable] struct {
	contents map[T]bool
}

func (set Set[T]) String() string {
	var s string
	s += "("
	for elem, _ := range set.contents {
		s += fmt.Sprintf("%v, ", elem)
	}
	s = strings.TrimSuffix(s, ", ")
	s += ")"
	return s
}

func Of[T comparable](s ...T) *Set[T] {
	contents := make(map[T]bool)
	for _, val := range s {
		contents[val] = true
	}
	return &Set[T]{contents: contents}
}

func Union[T comparable](sets ...*Set[T]) *Set[T] {
	union := make(map[T]bool)
	for _, set := range sets {
		for elem, _ := range set.contents {
			union[elem] = true
		}
	}
	return &Set[T]{contents: union}
}

func (set *Set[T]) Add(elem ...T) {
	for _, val := range elem {
		set.contents[val] = true
	}
}

func (set *Set[T]) Contains(elem T) bool {
	return set.contents[elem]
}
