package main

import (
	"fmt"
	"math"
)

func judgeSquareSum(c int) bool {
	l, r := 0, int(math.Sqrt(float64(c)))
	for l <= r {
		if d := int64(l*l + r*r); d == int64(c) {
			return true
		} else if d < int64(c) {
			l++
		} else {
			r--
		}
	}

	return false
}

func main() {
	testCases := []struct {
		c    int
		want bool
	}{
		{
			c:    1,
			want: true,
		},
		{
			c:    0,
			want: true,
		},
		{
			c:    5,
			want: true,
		},
		{
			c:    3,
			want: false,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := judgeSquareSum(tc.c)
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
