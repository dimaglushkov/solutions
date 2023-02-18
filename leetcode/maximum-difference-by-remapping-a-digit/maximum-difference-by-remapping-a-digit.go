package main

import (
	"fmt"
	"strconv"
	"strings"
)

// source: https://leetcode.com/problems/maximum-difference-by-remapping-a-digit/

func itoa(x int) string {
	return strconv.FormatInt(int64(x), 10)
}

func atoi(x string) int {
	y, _ := strconv.ParseInt(x, 10, 32)
	return int(y)
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func minMaxDifference(num int) int {
	numStr := itoa(num)
	minV, maxV := num, num
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if i == j {
				continue
			}
			newNumStr := strings.ReplaceAll(numStr, string(rune(i+'0')), string(rune(j+'0')))
			newNum := atoi(newNumStr)
			minV = min(minV, newNum)
			maxV = max(maxV, newNum)
		}
	}
	return maxV - minV
}

func main() {
	testCases := []struct {
		num  int
		want int
	}{
		{
			num:  11891,
			want: 99009,
		},
		{
			num:  90,
			want: 99,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := minMaxDifference(tc.num)
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
