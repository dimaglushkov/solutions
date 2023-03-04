package main

import (
	"fmt"
	"sort"
	"strconv"
)

// source: https://leetcode.com/problems/split-with-minimum-sum/

func atoi(x string) int {
	y, _ := strconv.ParseInt(x, 10, 32)
	return int(y)
}

func itoa(x int) string {
	return strconv.FormatInt(int64(x), 10)
}

func sortStr(x string) string {
	r := []rune(x)
	sort.Slice(r, func(i, j int) bool {
		return r[i] < r[j]
	})
	return string(r)
}

func splitNum(num int) int {
	s := sortStr(itoa(num))
	var s1, s2 string

	for i := 0; i < len(s); i++ {
		s1 += string(s[i])
		if i+1 < len(s) {
			i++
			s2 += string(s[i])
		}
	}
	return atoi(s1) + atoi(s2)
}

func main() {
	testCases := []struct {
		num  int
		want int
	}{
		{
			num:  4325,
			want: 59,
		},
		{
			num:  687,
			want: 75,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := splitNum(tc.num)
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
