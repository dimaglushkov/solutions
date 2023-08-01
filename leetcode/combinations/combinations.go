package main

import (
	"fmt"
)

func combine(n int, k int) [][]int {
	var res [][]int
	var util func(cur []int, start, left int)
	util = func(cur []int, start, left int) {
		if left == 0 {
			res = append(res, cur)
		}

		for i := start; i <= n; i++ {
			newCur := make([]int, len(cur), len(cur)+1)
			copy(newCur, cur)
			newCur = append(newCur, i)
			util(newCur, start+1, left-1)
			start++
		}
	}

	util(make([]int, 0), 1, k)

	return res
}

func main() {
	testCases := []struct {
		n    int
		k    int
		want [][]int
	}{
		{
			n:    5,
			k:    4,
			want: [][]int{{1, 2, 3, 4}, {1, 2, 3, 5}, {1, 2, 4, 5}, {1, 3, 4, 5}, {2, 3, 4, 5}},
		},
		{
			n:    4,
			k:    2,
			want: [][]int{{1, 2}, {1, 3}, {1, 4}, {2, 3}, {2, 4}, {3, 4}},
		},
		{
			n:    1,
			k:    1,
			want: [][]int{{1}},
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := combine(tc.n, tc.k)
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
