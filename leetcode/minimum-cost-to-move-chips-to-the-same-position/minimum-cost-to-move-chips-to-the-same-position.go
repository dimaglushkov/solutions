package main

import (
	"fmt"
)

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func minCostToMoveChips(position []int) int {
	odd := 0

	for _, p := range position {
		if p%2 == 1 {
			odd++
		}
	}

	return min(odd, len(position)-odd)
}

func main() {
	testCases := []struct {
		position []int
		want     int
	}{
		{
			position: []int{1, 2, 3},
			want:     1,
		},
		{
			position: []int{2, 2, 2, 3, 3},
			want:     2,
		},
		{
			position: []int{1, 1000000000},
			want:     1,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := minCostToMoveChips(tc.position)
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
