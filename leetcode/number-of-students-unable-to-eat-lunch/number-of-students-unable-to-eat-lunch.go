package main

import (
	"fmt"
)

type queue []int

func (q *queue) push(x int) {
	*q = append(*q, x)
}
func (q *queue) pop() int {
	x := (*q)[0]
	*q = (*q)[1:]
	return x
}

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

func reverse(x []int) {
	n := len(x)
	for i := 0; i < n/2; i++ {
		x[i], x[n-1-i] = x[n-1-i], x[i]
	}
}

func countStudents(students []int, sandwiches []int) int {
	q := queue(students)

	reverse(sandwiches)

	s := stack(sandwiches)

	var ans int
	for {
		ans = len(q)
		for i := 0; i < len(q); i++ {
			x := q.pop()
			if x == s.top() {
				s.pop()
			} else {
				q.push(x)
			}
		}

		if ans == len(q) {
			return ans
		}
	}
}

func main() {
	testCases := []struct {
		students   []int
		sandwiches []int
		want       int
	}{
		{
			students:   []int{1, 1, 0, 0},
			sandwiches: []int{0, 1, 0, 1},
			want:       0,
		},
		{
			students:   []int{1, 1, 1, 0, 0, 1},
			sandwiches: []int{1, 0, 0, 0, 1, 1},
			want:       3,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := countStudents(tc.students, tc.sandwiches)
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
