package main

import (
	"fmt"
)

// source: https://leetcode.com/problems/excel-sheet-column-title/

func convertToTitle(n int) string {
	if n == 0 {
		return ""
	}
	n--
	return convertToTitle(n/26) + string(rune('A'+(n%26)))
}

func main() {
	testCases := []struct {
		columnNumber int
		want         string
	}{
		{
			columnNumber: 50001,
			want:         "BUYC",
		},
		{
			columnNumber: 701,
			want:         "ZY",
		},
		{
			columnNumber: 1,
			want:         "A",
		},
		{
			columnNumber: 28,
			want:         "AB",
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := convertToTitle(tc.columnNumber)
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
