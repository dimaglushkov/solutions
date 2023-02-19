package main

import (
	"fmt"
	"strconv"
	"strings"
)

// source: https://leetcode.com/problems/minimum-operations-to-reduce-an-integer-to-0/

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func minOperations(x int) int {
	ans := 1 << 31
	ones := strings.Count(strconv.FormatInt(int64(x), 2), "1")
	var calculate func(val int, additions int)
	calculate = func(val int, additions int) {
		s := strconv.FormatInt(int64(val), 2)
		n := len(s)
		for i, c := range s {
			if c == '0' && additions+1 < ones {
				calculate(val+1<<(n-i-1), additions+1)
			}
		}

		tmp := 0
		for ind := strings.Index(s, "111"); ind != -1; ind = strings.Index(s, "111") {
			li := ind
			for li < n && s[li] == '1' {
				li++
			}
			tmp++
			val += 1 << (n - li)
			s = strconv.FormatInt(int64(val), 2)

		}
		ans = min(ans, tmp+additions+strings.Count(s, "1"))
		n = len(s)
	}
	calculate(x, 0)

	return ans
}

func main() {
	testCases := []struct {
		n    int
		want int
	}{
		{
			n:    39,
			want: 3,
		},
		{
			n:    54,
			want: 3,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := minOperations(tc.n)
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
