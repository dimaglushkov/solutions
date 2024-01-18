package main

import (
	"fmt"
	"strings"
)

func toLowerCase(s string) string {
	return strings.ToLower(s)
}

func main() {
	testCases := []struct {
		s    string
		want string
	}{
		{
			s:    "Hello",
			want: "hello",
		},
		{
			s:    "here",
			want: "here",
		},
		{
			s:    "LOVELY",
			want: "lovely",
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := toLowerCase(tc.s)
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
