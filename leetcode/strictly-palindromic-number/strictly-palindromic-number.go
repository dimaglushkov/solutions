package main

import (
	"fmt"
	"strconv"
)

func isStrictlyPalindromic(n int) bool {
	for i := 2; i < n-1; i++ {
		s := strconv.FormatInt(int64(n), i)
		l := len(s)
		for j := 0; j < l/2; j++ {
			if s[j] != s[l-j-1] {
				return false
			}
		}
	}

	return true
}

func main() {
	testCases := []struct {
		n    int
		want bool
	}{
		{
			n:    9,
			want: false,
		},
		{
			n:    4,
			want: false,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := isStrictlyPalindromic(tc.n)
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
