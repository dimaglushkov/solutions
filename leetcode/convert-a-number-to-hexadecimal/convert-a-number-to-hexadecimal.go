package main

import (
	"fmt"
)

// source: https://leetcode.com/problems/convert-a-number-to-hexadecimal/

func toHex(num int) string {
	if num == 0 {
		return "0"
	}
	if num < 0 {
		num = 1<<32 + num
	}
	hex := "0123456789abcdef"
	ans := make([]byte, 0)

	for num > 0 {
		ans = append([]byte{hex[num%16]}, ans...)
		num /= 16
	}
	return string(ans)
}

func main() {
	testCases := []struct {
		num  int
		want string
	}{
		{
			num:  26,
			want: "1a",
		},
		{
			num:  -1,
			want: "ffffffff",
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := toHex(tc.num)
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
