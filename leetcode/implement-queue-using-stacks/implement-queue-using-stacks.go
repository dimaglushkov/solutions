package main

// source: https://leetcode.com/problems/implement-queue-using-stacks/

type Stack struct {
	values []int
}

func NewStack() Stack {
	return Stack{values: make([]int, 0)}
}

func (s *Stack) Push(val int) {
	s.values = append(s.values, val)
}

func (s *Stack) Pop() int {

	if len(s.values) == 0 {
		return -1
	}
	val := s.values[len(s.values)-1]
	s.values = s.values[:len(s.values)-1]
	return val
}

func (s *Stack) Len() int {
	return len(s.values)
}

func (s *Stack) Peek() int {
	if len(s.values) == 0 {
		return -1
	}
	return s.values[len(s.values)-1]
}

//======

type MyQueue struct {
	s1 Stack
	s2 Stack
}

func Constructor() MyQueue {
	return MyQueue{
		s1: NewStack(),
		s2: NewStack(),
	}
}

func (q *MyQueue) Push(x int) {
	for q.s2.Len() != 0 {
		q.s1.Push(q.s2.Pop())
	}
	q.s1.Push(x)
}

func (q *MyQueue) Pop() int {
	for q.s1.Len() != 0 {
		q.s2.Push(q.s1.Pop())
	}
	return q.s2.Pop()
}

func (q *MyQueue) Peek() int {
	for q.s1.Len() != 0 {
		q.s2.Push(q.s1.Pop())
	}
	return q.s2.Peek()
}

func (q *MyQueue) Empty() bool {
	return q.s1.Len() == 0 && q.s2.Len() == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
