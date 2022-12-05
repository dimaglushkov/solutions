package main

import (
	"fmt"
	"strings"
)

// source: https://leetcode.com/problems/length-of-last-word/

func lengthOfLastWord(s string) int {
	s = strings.TrimRight(s, " ")
	w := strings.Split(s, " ")
	return len(w[len(w)-1])
}

func main() {
	testCases := []struct {
		s    string
		want int
	}{
		{
			s:    "Hello World",
			want: 5,
		},
		{
			s:    "   fly me   to   the moon  ",
			want: 4,
		},
		{
			s:    "luffy is still joyboy",
			want: 6,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := lengthOfLastWord(tc.s)
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
