package main

import (
	"fmt"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

type queue []int

func (q *queue) push(x int) {
	*q = append(*q, x)
}
func (q *queue) pop() int {
	x := (*q)[0]
	*q = (*q)[1:]
	return x
}
func (q *queue) top() int {
	return (*q)[0]
}

func shortestToChar(s string, c byte) []int {
	n := len(s)
	ans := make([]int, n)
	q := queue{}

	for i, l := range s {
		if byte(l) == c {
			q.push(i)
		}
	}
	q.push(2 * n)

	cur := q.pop()
	for i := range s {
		if abs(cur-i) > abs(q.top()-i) {
			cur = q.pop()
		}
		ans[i] = abs(cur - i)
	}

	return ans
}

func main() {
	testCases := []struct {
		s    string
		c    byte
		want []int
	}{
		{
			s:    "loveleetcode",
			c:    'e',
			want: []int{3, 2, 1, 0, 1, 0, 0, 1, 2, 2, 1, 0},
		},
		{
			s:    "aaab",
			c:    'b',
			want: []int{3, 2, 1, 0},
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := shortestToChar(tc.s, tc.c)
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
