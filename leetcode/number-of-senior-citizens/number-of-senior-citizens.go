package main

import (
	"fmt"
)

func countSeniors(details []string) int {
	ans := 0

	for _, d := range details {
		if d[11] > '6' || (d[11] == '6' && d[12] > '0') {
			ans++
		}
	}

	return ans
}

func main() {
	testCases := []struct {
		details []string
		want    int
	}{
		{
			details: []string{"7868190130M7522", "5303914400F9211", "9273338290F4010"},
			want:    2,
		},
		{
			details: []string{"1313579440F2036", "2921522980M5644"},
			want:    0,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := countSeniors(tc.details)
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
