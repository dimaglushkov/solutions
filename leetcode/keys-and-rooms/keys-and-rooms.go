package main

import (
	"fmt"
)

// source: https://leetcode.com/problems/keys-and-rooms/

func canVisitAllRooms(rooms [][]int) bool {
	n, cnt := len(rooms), 0
	visited := make([]bool, n)
	queue := []int{0}

	for len(queue) > 0 {
		v := queue[0]
		queue = queue[1:]

		if visited[v] {
			continue
		}
		visited[v] = true
		cnt++

		for _, k := range rooms[v] {
			if visited[k] {
				continue
			}
			queue = append(queue, k)
		}
	}

	return cnt == n
}

func main() {
	testCases := []struct {
		rooms [][]int
		want  bool
	}{
		{
			rooms: [][]int{{1}, {2}, {3}, {}},
			want:  true,
		},
		{
			rooms: [][]int{{1, 3}, {3, 0, 1}, {2}, {0}},
			want:  false,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := canVisitAllRooms(tc.rooms)
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
