package stack

type Stack[T any] struct {
	values []T
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{}
}

func (s *Stack[T]) Push(val T) {
	s.values = append(s.values, val)
}

func (s *Stack[T]) Pop() (val T, ok bool) {
	if len(s.values) == 0 {
		return val, false
	}
	val = s.values[len(s.values)-1]
	s.values = s.values[:len(s.values)-1]
	return val, true
}

func (s *Stack[T]) Len() int {
	return len(s.values)
}

func (s *Stack[T]) Top() T {
	return s.values[len(s.values)-1]
}
