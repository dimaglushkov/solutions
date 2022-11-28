package stack

type Stack []int

func NewStack(values ...int) Stack {
	s := make(Stack, 0, len(values))
	s = append(s, values...)
	return s
}

func (s *Stack) Push(x int) {
	*s = append(*s, x)
}

func (s *Stack) Pop() int {
	n := len(*s)
	x := (*s)[n-1]
	*s = (*s)[:n-1]
	return x
}

func (s *Stack) Len() int {
	return len(*s)
}

func (s *Stack) Top() int {
	if len(*s) > 0 {
		return (*s)[len(*s)-1]
	}
	return -1
}
