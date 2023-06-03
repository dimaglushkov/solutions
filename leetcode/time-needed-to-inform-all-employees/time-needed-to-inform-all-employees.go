package main

import (
	"fmt"
)

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func calc(curSum, root int, edges map[int][]int, informTime []int) int {
	curSum += informTime[root]
	maxSum := curSum

	for _, ch := range edges[root] {
		maxSum = max(maxSum, calc(curSum, ch, edges, informTime))
	}

	return maxSum
}

func numOfMinutes(n int, headID int, manager []int, informTime []int) int {
	edges := make(map[int][]int)
	for i, h := range manager {
		if h == -1 {
			continue
		}
		edges[h] = append(edges[h], i)
	}
	return calc(0, headID, edges, informTime)
}

func main() {
	testCases := []struct {
		n          int
		headID     int
		manager    []int
		informTime []int
		want       int
	}{
		{
			n:          6,
			headID:     2,
			manager:    []int{2, 2, -1, 2, 2, 2},
			informTime: []int{0, 0, 1, 0, 0, 0},
			want:       1,
		},
		{
			n:          1,
			headID:     0,
			manager:    []int{-1},
			informTime: []int{0},
			want:       0,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := numOfMinutes(tc.n, tc.headID, tc.manager, tc.informTime)
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
