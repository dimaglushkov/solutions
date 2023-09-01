package main

import (
	"fmt"
)

func isAcronym(words []string, s string) bool {
	if len(words) != len(s) {
		return false
	}

	for i, w := range words {
		if s[i] != w[0] {
			return false
		}
	}

	return true
}

func main() {
	testCases := []struct {
		words []string
		s     string
		want  bool
	}{
		{
			words: []string{"alice", "bob", "charlie"},
			s:     "abc",
			want:  true,
		},
		{
			words: []string{"an", "apple"},
			s:     "a",
			want:  false,
		},
		{
			words: []string{"never", "gonna", "give", "up", "on", "you"},
			s:     "ngguoy",
			want:  true,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := isAcronym(tc.words, tc.s)
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
