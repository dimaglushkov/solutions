package main

import (
	"fmt"
)

// source: https://leetcode.com/problems/jump-game-iv/

type queue [][2]int

func (q *queue) push(x, y int) {
	*q = append(*q, [2]int{x, y})
}
func (q *queue) pop() (int, int) {
	x := (*q)[0]
	*q = (*q)[1:]
	return x[0], x[1]
}

func minJumps(arr []int) int {
	n := len(arr)
	dist := make([]int, n)
	for i := range dist {
		dist[i] = -1
	}
	copies := make(map[int]map[int]bool)

	for i, v := range arr {
		if copies[v] == nil {
			copies[v] = make(map[int]bool)
		}
		copies[v][i] = true
	}

	var q queue
	q.push(0, 0)
	for len(q) > 0 {
		id, d := q.pop()
		if dist[id] != -1 {
			continue
		}
		dist[id] = d
		for i, ok := range copies[arr[id]] {
			if ok {
				q.push(i, d+1)
			}
		}
		copies[arr[id]] = nil

		if id < n-1 {
			q.push(id+1, d+1)
		}
		if id > 0 {
			q.push(id-1, d+1)
		}
	}
	return dist[n-1]
}

func main() {
	testCases := []struct {
		arr  []int
		want int
	}{
		{
			arr:  []int{100, -23, -23, 404, 100, 23, 23, 23, 3, 404},
			want: 3,
		},
		{
			arr:  []int{7},
			want: 0,
		},
		{
			arr:  []int{7, 6, 9, 6, 9, 6, 9, 7},
			want: 1,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := minJumps(tc.arr)
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
