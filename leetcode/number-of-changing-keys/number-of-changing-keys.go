package main

import (
	"fmt"
	"strings"
)

func countKeyChanges(s string) int {
	s = strings.ToLower(s)
	ans := 0
	p := s[0]

	for i := 1; i < len(s); i++ {
		if s[i] != p {
			p = s[i]
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
			s:    "aAbBcC",
			want: 2,
		},
		{
			s:    "AaAaAaaA",
			want: 0,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := countKeyChanges(tc.s)
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
