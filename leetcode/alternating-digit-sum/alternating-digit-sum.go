package main

import (
	"fmt"
	"strconv"
)

// source: https://leetcode.com/problems/alternating-digit-sum/

func itoa(x int) string {
	return strconv.FormatInt(int64(x), 10)
}

func alternateDigitSum(n int) int {
	var ans int
	sign := 1

	for _, c := range itoa(n) {
		ans += int(c-'0') * sign
		sign *= -1
	}

	return ans
}

func main() {
	testCases := []struct {
		n    int
		want int
	}{
		{
			n:    521,
			want: 4,
		},
		{
			n:    111,
			want: 1,
		},
		{
			n:    886996,
			want: 0,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := alternateDigitSum(tc.n)
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
