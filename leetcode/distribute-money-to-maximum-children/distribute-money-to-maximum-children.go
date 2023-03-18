package main

import (
	"fmt"
)

// source: https://leetcode.com/problems/distribute-money-to-maximum-children/

func distMoney(money int, children int) int {
	if money < children {
		return -1
	}
	if money < 8 {
		return 0
	}
	d, m, l := money/8, money%8, children-money/8
	if d == children && m == 0 {
		return children
	}
	if d >= children {
		return children - 1
	}
	if m == 4 && l == 1 {
		return children - 2
	}
	for m < l {
		l++
		m += 8
		d--
	}
	return d
}

func main() {
	testCases := []struct {
		money    int
		children int
		want     int
	}{
		{
			money:    20,
			children: 3,
			want:     1,
		},
		{
			money:    16,
			children: 2,
			want:     2,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := distMoney(tc.money, tc.children)
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
