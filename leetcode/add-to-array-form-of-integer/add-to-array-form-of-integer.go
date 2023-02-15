package main

import (
	"fmt"
)

// source: https://leetcode.com/problems/add-to-array-form-of-integer/

func reverse(x []int) {
	n := len(x)
	for i := 0; i < n/2; i++ {
		x[i], x[n-1-i] = x[n-1-i], x[i]
	}
}

func addToArrayForm(num1 []int, k int) []int {
	num2 := make([]int, 0)
	for k > 0 {
		num2 = append(num2, k%10)
		k /= 10
	}
	reverse(num2)
	max, min := num1, num2
	if len(num2) > len(num1) {
		max, min = min, max
	}
	nMin, nMax := len(min), len(max)

	iAns := nMax
	ans := make([]int, iAns+1)
	overflow := false
	for i := 0; i < len(max); i++ {
		c := max[nMax-1-i]
		if nMin-1-i >= 0 {
			c += min[nMin-1-i]
		}
		if overflow {
			c++
		}
		if c > 9 {
			overflow, c = true, c-10
		} else {
			overflow = false
		}
		ans[iAns] = c
		iAns--
	}
	if overflow {
		ans[iAns] = 1
		iAns--
	}
	return ans[iAns+1:]
}

func main() {
	testCases := []struct {
		num  []int
		k    int
		want []int
	}{
		{
			num:  []int{1, 2, 0, 0},
			k:    34,
			want: []int{1, 2, 3, 4},
		},
		{
			num:  []int{2, 7, 4},
			k:    181,
			want: []int{4, 5, 5},
		},
		{
			num:  []int{2, 1, 5},
			k:    806,
			want: []int{1, 0, 2, 1},
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := addToArrayForm(tc.num, tc.k)
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
