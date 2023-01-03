package main

import (
	"fmt"
)

// source: https://leetcode.com/problems/delete-columns-to-make-sorted/

func minDeletionSize(strs []string) int {
	res := 0
	n, m := len(strs), len(strs[0])

	for j := 0; j < m; j++ {
		for i := 0; i < n-1; i++ {
			if strs[i][j] > strs[i+1][j] {
				res++
				break
			}
		}
	}
	return res
}

func main() {
	testCases := []struct {
		strs []string
		want int
	}{
		{
			strs: []string{"cba", "daf", "ghi"},
			want: 1,
		},
		{
			strs: []string{"a", "b"},
			want: 0,
		},
		{
			strs: []string{"zyx", "wvu", "tsr"},
			want: 3,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := minDeletionSize(tc.strs)
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
