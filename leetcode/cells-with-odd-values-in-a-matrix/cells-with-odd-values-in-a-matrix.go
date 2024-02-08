package main

import (
	"fmt"
)

func oddCells(m int, n int, indices [][]int) int {
	rows := make([]int, m)
	cols := make([]int, n)

	for i := range indices {
		rows[indices[i][0]]++
		cols[indices[i][1]]++
	}

	ans := 0

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if (rows[i]+cols[j])%2 == 1 {
				ans++
			}
		}
	}

	return ans
}

func main() {
	testCases := []struct {
		m       int
		n       int
		indices [][]int
		want    int
	}{
		{
			m:       2,
			n:       3,
			indices: [][]int{{0, 1}, {1, 1}},
			want:    6,
		},
		{
			m:       2,
			n:       2,
			indices: [][]int{{1, 1}, {0, 0}},
			want:    0,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := oddCells(tc.m, tc.n, tc.indices)
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
