package main

import (
	"fmt"
)

func minSteps(n int) int {
	d := 2
	ans := 0

	for n > 1 {
		for n%d == 0 {
			ans += d
			n /= d
		}
		d++
	}

	return ans
}

func main() {
	testCases := []struct {
		n    int
		want int
	}{
		{
			n:    8,
			want: 6,
		},
		{
			n:    4,
			want: 4,
		},
		{
			n:    3,
			want: 3,
		},
		{
			n:    1,
			want: 0,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := minSteps(tc.n)
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
