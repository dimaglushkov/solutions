package main

import (
	"fmt"
)

// source: https://leetcode.com/problems/greatest-common-divisor-of-strings/

func gcd(a, b int) int {
	if a == 0 {
		return b
	}
	return gcd(b%a, a)
}

func gcdOfStrings(str1 string, str2 string) string {
	if str1+str2 != str2+str1 {
		return ""
	}
	return str1[0:gcd(len(str1), len(str2))]
}

func main() {
	testCases := []struct {
		str1 string
		str2 string
		want string
	}{
		{
			str1: "ABCABC",
			str2: "ABC",
			want: "ABC",
		},
		{
			str1: "ABABAB",
			str2: "ABAB",
			want: "AB",
		},
		{
			str1: "LEET",
			str2: "CODE",
			want: "",
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := gcdOfStrings(tc.str1, tc.str2)
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
