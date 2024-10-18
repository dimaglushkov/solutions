package main

import (
	"fmt"
	"strconv"
)

//You are given an integer num. You can swap two digits at most once to get the maximum valued number.
//Return the maximum valued number you can get.

func maximumSwap(num int) int {
	numStr := []rune(strconv.Itoa(num))
	n := len(numStr)

	last := make([]int, 10)
	for i := 0; i < n; i++ {
		last[numStr[i]-'0'] = i
	}

	for i := 0; i < n; i++ {
		for d := 9; d > int(numStr[i]-'0'); d-- {
			if last[d] > i {
				numStr[i], numStr[last[d]] = numStr[last[d]], numStr[i]
				ans, _ := strconv.Atoi(string(numStr))
				return ans
			}
		}
	}

	return num
}

func main() {
	testCases := []struct {
		num  int
		want int
	}{
		{
			num:  1331,
			want: 3311,
		},
		{
			num:  3131,
			want: 3311,
		},
		{
			num:  1133,
			want: 3131,
		},
		{
			num:  1313,
			want: 3311,
		},
		{
			num:  2736,
			want: 7236,
		},
		{
			num:  9973,
			want: 9973,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := maximumSwap(tc.num)
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
