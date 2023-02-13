package main

import (
	"fmt"
)

// source: https://leetcode.com/problems/find-the-difference/

func findTheDifference(s string, t string) byte {
	var l [26]int
	for _, c := range t {
		l[c-'a']++
	}
	for _, c := range s {
		l[c-'a']--
	}
	for i := range l {
		if l[i] == 1 {
			return byte(i + 'a')
		}
	}
	return 0
}

func main() {
	testCases := []struct {
		s    string
		t    string
		want byte
	}{
		{
			s:    "abcd",
			t:    "abcde",
			want: 'e',
		},
		{
			s:    "",
			t:    "y",
			want: 'y',
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := findTheDifference(tc.s, tc.t)
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
