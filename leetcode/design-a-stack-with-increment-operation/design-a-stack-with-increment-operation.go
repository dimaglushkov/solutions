package main

// source: https://leetcode.com/problems/design-a-stack-with-increment-operation/

type CustomStack struct {
	n int
	v []int
}

func Constructor(maxSize int) CustomStack {
	return CustomStack{
		n: maxSize,
		v: make([]int, 0, maxSize),
	}
}

func (s *CustomStack) Push(x int) {
	if len(s.v) < s.n {
		s.v = append(s.v, x)
	}
}

func (s *CustomStack) Pop() int {
	if len(s.v) == 0 {
		return -1
	}

	x := s.v[len(s.v)-1]
	s.v = s.v[:len(s.v)-1]
	return x
}

func (s *CustomStack) Increment(k int, val int) {
	for i := 0; i < min(k, len(s.v)); i++ {
		s.v[i] += val
	}
}

func main() {
	stk := Constructor(3)
	stk.Push(1)
	stk.Push(2)
	stk.Pop()
	stk.Push(2)
	stk.Push(3)
	stk.Push(4)
	stk.Increment(5, 100)
	stk.Increment(2, 100)
	stk.Pop()
	stk.Pop()
	stk.Pop()
	stk.Pop()

}

/**
 * Your CustomStack object will be instantiated and called as such:
 * obj := Constructor(maxSize);
 * obj.Push(x);
 * param_2 := obj.Pop();
 * obj.Increment(k,val);
 */
