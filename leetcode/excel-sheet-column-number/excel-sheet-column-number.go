package main

import (
	"fmt"
	"math"
)

// source: https://leetcode.com/problems/excel-sheet-column-number/

func titleToNumber(columnTitle string) int {
	res := 0
	n := len(columnTitle)
	p := int(math.Pow(float64(26), float64(n)))
	for i := 0; i < n; i++ {
		p /= 26

		x := int(columnTitle[i] - 'A' + 1)
		res += x * p
	}
	return res
}

func main() {
	testCases := []struct {
		columnTitle string
		want        int
	}{
		{
			columnTitle: "ZY",
			want:        701,
		},
		{
			columnTitle: "BUYC",
			want:        50001,
		},

		{
			columnTitle: "A",
			want:        1,
		},
		{
			columnTitle: "AB",
			want:        28,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := titleToNumber(tc.columnTitle)
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
