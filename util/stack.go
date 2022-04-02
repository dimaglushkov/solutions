package util

type Stack[T any] struct {
	values []T
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{}
}

func (s *Stack[T]) Push(val T) {
	s.values = append[T](s.values, val)
}

func (s *Stack[T]) Pop() (val T, ok bool) {
	if len[T](s.values) == 0 {
		return val, false
	}
	val = s.values[len[T](s.values)-1]
	s.values = s.values[:len[T](s.values)-1]
	return val, true
}

func (s *Stack[T]) Size() int {
	return len[T](s.values)
}

func (s *Stack[T]) Top() T {
	return s.values[len[T](s.values)-1]
}
