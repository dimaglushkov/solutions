package main

import (
	"fmt"
)

// source: https://leetcode.com/problems/add-strings/

func addStrings(num1 string, num2 string) string {
	max, min := num1, num2
	if len(num2) > len(num1) {
		max, min = min, max
	}
	nMin, nMax := len(min), len(max)

	ans := ""
	var overflow byte = '0'
	for i := 0; i < len(max); i++ {
		c := max[nMax-1-i]
		if nMin-1-i >= 0 {
			c += min[nMin-1-i] - '0'
		}
		if overflow > '0' {
			c += overflow - '0'
		}
		if c > '9' {
			overflow, c = '1', c-10
		} else {
			overflow = '0'
		}
		ans = string(c) + ans
	}
	if overflow != '0' {
		ans = string(overflow) + ans
	}

	return ans
}

func main() {
	testCases := []struct {
		num1 string
		num2 string
		want string
	}{
		{
			num1: "9",
			num2: "9",
			want: "18",
		},
		{
			num1: "456",
			num2: "77",
			want: "533",
		},
		{
			num1: "11",
			num2: "123",
			want: "134",
		},

		{
			num1: "0",
			num2: "0",
			want: "0",
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := addStrings(tc.num1, tc.num2)
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
