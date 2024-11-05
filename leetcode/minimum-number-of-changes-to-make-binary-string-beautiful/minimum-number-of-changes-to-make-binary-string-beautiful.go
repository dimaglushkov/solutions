package main

import (
	"fmt"
)

func minChanges(s string) int {
	ans := 0
	for i := 1; i < len(s); i += 2 {
		if s[i] != s[i-1] {
			ans++
		}
	}
	return ans
}

func main() {
	testCases := []struct {
		s    string
		want int
	}{
		{
			s:    "1001",
			want: 2,
		},
		{
			s:    "10",
			want: 1,
		},
		{
			s:    "0000",
			want: 0,
		},
		{
			s:    "11000111",
			want: 1,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := minChanges(tc.s)
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
