package main

import (
	"fmt"
)

func minFlips(a int, b int, c int) int {
	ans := 0
	for a > 0 || b > 0 || c > 0 {
		ai, bi, ci := a%2, b%2, c%2
		a, b, c = a/2, b/2, c/2
		if ai|bi != ci {
			ans++
			if ci == 0 && ai == 1 && bi == 1 {
				ans++
			}
		}

	}
	return ans
}

func main() {
	testCases := []struct {
		a    int
		b    int
		c    int
		want int
	}{
		{
			a:    5, //  101
			b:    2, //  010
			c:    8, // 1000
			want: 4,
		},
		{
			a:    8, // 1000
			b:    3, // 0011
			c:    5, // 0101
			want: 3,
		},
		{
			a:    2,
			b:    6,
			c:    5,
			want: 3,
		},
		{
			a:    4,
			b:    2,
			c:    7,
			want: 1,
		},
		{
			a:    1,
			b:    2,
			c:    3,
			want: 0,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := minFlips(tc.a, tc.b, tc.c)
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
