package main

import (
	"fmt"
)

// source: https://leetcode.com/problems/determine-if-string-halves-are-alike/

func halvesAreAlike(s string) bool {
	vowels := map[byte]bool{
		'a': true,
		'e': true,
		'i': true,
		'o': true,
		'u': true,
		'A': true,
		'E': true,
		'I': true,
		'O': true,
		'U': true,
	}
	res := 0
	n := len(s)
	for i := 0; i < n; i++ {
		if vowels[s[i]] {
			if i < n/2 {
				res++
			} else {
				res--
			}
		}
	}
	return res == 0
}
func main() {
	testCases := []struct {
		s    string
		want bool
	}{
		{
			s:    "book",
			want: true,
		},
		{
			s:    "textbook",
			want: false,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := halvesAreAlike(tc.s)
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
