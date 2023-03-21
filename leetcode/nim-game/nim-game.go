package main

import (
	"fmt"
)

// source: https://leetcode.com/problems/nim-game/

func canWinNim(n int) bool {
	return n%4 != 0
}

func main() {
	testCases := []struct {
		n    int
		want bool
	}{
		{
			n:    4,
			want: false,
		},
		{
			n:    1,
			want: true,
		},
		{
			n:    2,
			want: true,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := canWinNim(tc.n)
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
