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

func canReach(arr []int, start int) bool {
	var q queue
	visited := make([]bool, len(arr))

	q.push(start)

	for len(q) > 0 {
		x := q.pop()
		if visited[x] {
			continue
		}
		visited[x] = true
		if arr[x] == 0 {
			return true
		}
		if x-arr[x] >= 0 {
			q.push(x - arr[x])
		}
		if x+arr[x] < len(arr) {
			q.push(x + arr[x])
		}
	}

	return false
}

func main() {
	testCases := []struct {
		arr   []int
		start int
		want  bool
	}{
		{
			arr:   []int{4, 2, 3, 0, 3, 1, 2},
			start: 5,
			want:  true,
		},
		{
			arr:   []int{4, 2, 3, 0, 3, 1, 2},
			start: 0,
			want:  true,
		},
		{
			arr:   []int{3, 0, 2, 1, 2},
			start: 2,
			want:  false,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := canReach(tc.arr, tc.start)
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
