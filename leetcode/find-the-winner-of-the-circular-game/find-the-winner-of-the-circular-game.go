package main

import (
	"fmt"
)

func findTheWinner(n int, k int) int {
	var util func(n, k int) int
	util = func(n, k int) int {
		if n == 1 {
			return 0
		}
		return (util(n-1, k) + k) % n
	}

	return util(n, k) + 1
}

func main() {
	testCases := []struct {
		n    int
		k    int
		want int
	}{
		{
			n:    5,
			k:    2,
			want: 3,
		},
		{
			n:    6,
			k:    5,
			want: 1,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := findTheWinner(tc.n, tc.k)
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
