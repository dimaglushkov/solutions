package main

import (
	"fmt"
)

type stack []int

func (s *stack) push(x int) {
	*s = append(*s, x)
}
func (s *stack) pop() int {
	n := len(*s)
	x := (*s)[n-1]
	*s = (*s)[:n-1]
	return x
}
func (s *stack) top() int {
	return (*s)[len(*s)-1]
}

func asteroidCollision(asteroids []int) []int {
	var s stack
	for _, ast := range asteroids {
		for len(s) > 0 && (s.top() > 0 && ast < 0) {
			d := s.top() + ast
			switch {
			case d < 0:
				s.pop()
			case d == 0:
				ast = 0
				s.pop()
			case d > 0:
				ast = 0
			}
		}

		if ast != 0 {
			s.push(ast)
		}
	}
	return s
}

func main() {
	testCases := []struct {
		asteroids []int
		want      []int
	}{
		{
			asteroids: []int{5, 10, -5},
			want:      []int{5, 10},
		},
		{
			asteroids: []int{8, -8},
			want:      []int{},
		},
		{
			asteroids: []int{10, 2, -5},
			want:      []int{10},
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := asteroidCollision(tc.asteroids)
		status := "ERROR"
		if fmt.Sprint(x) == fmt.Sprint(tc.want) {
			status = "OK"
			successes++
		}
		fmt.Println(status, "	Expected: ", tc.want, " Actual: ", x)
	}
	if l := len(testCases); successes == len(testCases) {
		fmt.Printf("===\nSUCCESS: %d of %d tests ended successfully\n", successes, l)
	} else {
		fmt.Printf("===\nFAIL: %d tests failed\n", l-successes)
	}

}
