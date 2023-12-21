package main

import (
	"fmt"
	"sort"
)

func maxWidthOfVerticalArea(p [][]int) int {
	sort.Slice(p, func(i, j int) bool {
		return p[i][0] < p[j][0]
	})

	ans := 0

	for i := 0; i < len(p)-1; i++ {
		if d := p[i+1][0] - p[i][0]; d > ans {
			ans = d
		}
	}

	return ans
}

func main() {
	testCases := []struct {
		points [][]int
		want   int
	}{
		{
			points: [][]int{{8, 7}, {9, 9}, {7, 4}, {9, 7}},
			want:   1,
		},
		{
			points: [][]int{{3, 1}, {9, 0}, {1, 0}, {1, 4}, {5, 3}, {8, 8}},
			want:   3,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := maxWidthOfVerticalArea(tc.points)
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
