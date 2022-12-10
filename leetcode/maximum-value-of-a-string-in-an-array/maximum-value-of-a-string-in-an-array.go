package main

import (
	"fmt"
	"strconv"
)

// source: https://leetcode.com/problems/maximum-value-of-a-string-in-an-array/

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func maximumValue(strs []string) int {
	res := -1

	for _, s := range strs {
		x, err := strconv.ParseInt(s, 10, 32)
		v := int(x)
		if err != nil {
			v = len(s)
		}
		res = max(res, v)
	}

	return res
}

func main() {
	testCases := []struct {
		strs []string
		want int
	}{
		{
			strs: []string{"alic3", "bob", "3", "4", "00000"},
			want: 5,
		},
		{
			strs: []string{"1", "01", "001", "0001"},
			want: 1,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := maximumValue(tc.strs)
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
