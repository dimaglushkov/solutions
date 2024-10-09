package main

import (
	"fmt"
)

func minAddToMakeValid(s string) int {
	ops := 0
	ans := 0
	for _, c := range s {
		if c == '(' {
			ops++
		} else {
			if ops > 0 {
				ops--
			} else {
				ans++
			}
		}
	}

	return ans + ops
}

func main() {
	testCases := []struct {
		s    string
		want int
	}{
		{
			s:    "())",
			want: 1,
		},
		{
			s:    "(((",
			want: 3,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := minAddToMakeValid(tc.s)
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
