package main

import (
	"fmt"
)

func checkValidString(s string) bool {
	pairsMin, pairsMax := 0, 0
	for _, c := range s {
		switch c {
		case '(':
			pairsMin++
			pairsMax++
		case ')':
			pairsMin--
			pairsMax--
		default:
			pairsMin--
			pairsMax++
		}

		if pairsMax < 0 {
			return false
		}

		if pairsMin < 0 {
			pairsMin = 0
		}
	}

	return pairsMin == 0
}

func main() {
	testCases := []struct {
		s    string
		want bool
	}{
		{
			s:    "()",
			want: true,
		},
		{
			s:    "(*)",
			want: true,
		},
		{
			s:    "(*))",
			want: true,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := checkValidString(tc.s)
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
