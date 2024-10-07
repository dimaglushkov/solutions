package main

import (
	"fmt"
)

func minLength(s string) int {
	var stack []byte

	for i := 0; i < len(s); i++ {
		if len(stack) > 0 && ((stack[len(stack)-1] == 'A' && s[i] == 'B') || (stack[len(stack)-1] == 'C' && s[i] == 'D')) {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}

	return len(stack)
}

func main() {
	testCases := []struct {
		s    string
		want int
	}{
		{
			s:    "ABFCACDB",
			want: 2,
		},
		{
			s:    "ACBBD",
			want: 5,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := minLength(tc.s)
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
