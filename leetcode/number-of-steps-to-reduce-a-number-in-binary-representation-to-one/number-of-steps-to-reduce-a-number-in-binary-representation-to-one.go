package main

import (
	"fmt"
)

func numSteps(s string) int {
	ans := 0

	buf := []byte(s)

	for len(buf) != 1 && buf[0] != 1 {
		ans++
		n := len(buf)

		if buf[n-1] == '0' {
			buf = buf[:n-1]
		} else {
			overflow := true

			for i := n - 1; i >= 0; i-- {
				if buf[i] == '1' {
					buf[i] = '0'
				} else {
					buf[i] = '1'
					overflow = false
					break
				}
			}

			if overflow {
				buf = append([]byte{'1'}, buf...)
			}
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
			s:    "1101",
			want: 6,
		},
		{
			s:    "10",
			want: 1,
		},
		{
			s:    "1",
			want: 0,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := numSteps(tc.s)
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
