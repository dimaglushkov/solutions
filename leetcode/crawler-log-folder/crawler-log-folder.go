package main

import (
	"fmt"
)

func minOperations(logs []string) int {
	ans := 0

	for _, l := range logs {
		switch l {
		case "../":
			if ans > 0 {
				ans--
			}
		case "./":

		default:
			ans++
		}
	}

	return ans
}

func main() {
	testCases := []struct {
		logs []string
		want int
	}{
		{
			logs: []string{"./", "wz4/", "../", "mj2/", "../", "../", "ik0/", "il7/"},
			want: 2,
		},
		{
			logs: []string{"d1/", "d2/", "../", "d21/", "./"},
			want: 2,
		},
		{
			logs: []string{"d1/", "d2/", "./", "d3/", "../", "d31/"},
			want: 3,
		},
		{
			logs: []string{"d1/", "../", "../", "../"},
			want: 0,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := minOperations(tc.logs)
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
