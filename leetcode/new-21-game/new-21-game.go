package main

import (
	"fmt"
)

func main() {
	testCases := []struct {
		n      int
		k      int
		maxPts int
		want   float64
	}{
		{
			n:      10,
			k:      1,
			maxPts: 10,
			want:   1.00000,
		},
		{
			n:      6,
			k:      1,
			maxPts: 10,
			want:   0.60000,
		},
		{
			n:      21,
			k:      17,
			maxPts: 10,
			want:   0.73278,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := new21Game(tc.n, tc.k, tc.maxPts)
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
