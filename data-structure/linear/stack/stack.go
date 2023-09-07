package stack

import "errors"

type Stack[T any] []T

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{}
}

func (s *Stack[T]) Push(value T) {
	*s = append(*s, value)
}

func (s *Stack[T]) Pop() (T, error) {
	if len(*s) == 0 {
		return *new(T), errors.New("stack is empty")
	}

	value := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]

	return value, nil
}
