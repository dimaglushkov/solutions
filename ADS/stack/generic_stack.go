package stack

type GenStack[T any] struct {
	values []T
}

func NewGenStack[T any]() *GenStack[T] {
	return &GenStack[T]{}
}

func (s *GenStack[T]) Push(val T) {
	s.values = append(s.values, val)
}

func (s *GenStack[T]) Pop() (val T, ok bool) {
	if len(s.values) == 0 {
		return val, false
	}
	val = s.values[len(s.values)-1]
	s.values = s.values[:len(s.values)-1]
	return val, true
}

func (s *GenStack[T]) Len() int {
	return len(s.values)
}

func (s *GenStack[T]) Top() T {
	return s.values[len(s.values)-1]
}
