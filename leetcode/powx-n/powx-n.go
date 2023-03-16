package main

import (
	"fmt"
	"math"
)

// source: https://leetcode.com/problems/powx-n/

func myPow(x float64, n int) float64 {
	return math.Pow(x, float64(n))
}

func main() {
	testCases := []struct {
		x    float64
		n    int
		want float64
	}{
		{
			x:    2.00000,
			n:    10,
			want: 1024.00000,
		},
		{
			x:    2.10000,
			n:    3,
			want: 9.26100,
		},
		{
			x:    2.00000,
			n:    -2,
			want: 0.25000,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := myPow(tc.x, tc.n)
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
