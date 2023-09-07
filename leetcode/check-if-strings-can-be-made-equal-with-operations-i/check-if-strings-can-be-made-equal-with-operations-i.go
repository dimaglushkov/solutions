package main

import (
	"fmt"
)

func swap(s string, i, j int) string {
	b := []byte(s)
	b[i], b[j] = b[j], b[i]

	return string(b)
}

func canBeEqual(s1 string, s2 string) bool {
	return s1 == s2 ||
		swap(s1, 0, 2) == s2 ||
		swap(s1, 1, 3) == s2 ||
		swap(swap(s1, 0, 2), 1, 3) == s2
}

func main() {
	testCases := []struct {
		s1   string
		s2   string
		want bool
	}{
		{
			s1:   "abcd",
			s2:   "cdab",
			want: true,
		},
		{
			s1:   "abcd",
			s2:   "dacb",
			want: false,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := canBeEqual(tc.s1, tc.s2)
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
