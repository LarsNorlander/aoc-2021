package stack

import "errors"

type Stack[T any] struct {
	contents []T
}

func Empty[T any]() *Stack[T] {
	return &Stack[T]{}
}

func (s *Stack[T]) Push(item T) {
	s.contents = append(s.contents, item)
}

func (s *Stack[T]) Pop() (T, error) {
	if len(s.contents) == 0 {
		return *new(T), errors.New("the stack is empty")
	}
	n := len(s.contents) - 1
	result := s.contents[n]
	s.contents = s.contents[:n]
	return result, nil
}

func (s *Stack[T]) Peek() T {
	if len(s.contents) == 0 {
		return *new(T)
	}
	n := len(s.contents) - 1
	return s.contents[n]
}
