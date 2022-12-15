package main

import (
	"fmt"
)
// source: https://leetcode.com/problems/power-of-two/

func isPowerOfTwo(n int) bool {
	if n < 0 {
		return false
	}
	if n == 0 {
		return true
	}
    for i := 1; i < 1 << 32; i = i << 1 {
		if n & i == n {
			return true
		}
	}
	return false
}

func main() {
    testCases := []struct {
		n int
		want bool
    }{
		{
			n:  1,
			want:  true,
		},
		{
			n:  16,
			want:  true,
		},
		{
			n:  3,
			want:  false,
		},
	}


    successes := 0
    for _, tc := range testCases {
        x := isPowerOfTwo(tc.n)
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
        fmt.Printf("===\nFAIL: %d tests failed\n", l - successes)
    }

}
