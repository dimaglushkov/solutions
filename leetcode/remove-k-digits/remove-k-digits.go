package main

import (
	"fmt"
)

func removeKdigits(num string, k int) string {
	res := make([]rune, 0)

	for _, c := range num {
		for len(res) > 0 && res[len(res)-1] > c && k > 0 {
			res = res[:len(res)-1]
			k--
		}

		if len(res) > 0 || c != '0' {
			res = append(res, c)
		}
	}

	for len(res) > 0 && k > 0 {
		res = res[:len(res)-1]
		k--
	}

	if len(res) == 0 {
		return "0"
	}

	return string(res)
}

func main() {
	testCases := []struct {
		num  string
		k    int
		want string
	}{
		{
			num:  "10001",
			k:    4,
			want: "0",
		},
		{
			num:  "10001",
			k:    1,
			want: "1",
		},
		{
			num:  "112",
			k:    1,
			want: "11",
		},
		{
			num:  "10",
			k:    1,
			want: "0",
		},
		{
			num:  "1432219",
			k:    3,
			want: "1219",
		},
		{
			num:  "9211249327",
			k:    4,
			want: "112327",
		},
		{
			num:  "10200",
			k:    1,
			want: "200",
		},
		{
			num:  "10",
			k:    2,
			want: "0",
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := removeKdigits(tc.num, tc.k)
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
