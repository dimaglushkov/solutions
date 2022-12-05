package main

import (
	"fmt"
)

// source: https://leetcode.com/problems/pascals-triangle-ii/

func getRow(rowIndex int) []int {
	res := make([]int, rowIndex+1)
	res[0] = 1
	for i := 0; i <= rowIndex; i++ {
		for j := i; j >= 1; j-- {
			res[j] += res[j-1]
		}
	}
	return res
}

func main() {
	testCases := []struct {
		rowIndex int
		want     []int
	}{
		{
			rowIndex: 3,
			want:     []int{1, 3, 3, 1},
		},
		{
			rowIndex: 0,
			want:     []int{1},
		},
		{
			rowIndex: 1,
			want:     []int{1, 1},
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := getRow(tc.rowIndex)
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
