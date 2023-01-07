package main

import (
	"fmt"
	"sort"
)

// source: https://leetcode.com/problems/maximum-ice-cream-bars/

func maxIceCream(costs []int, coins int) int {
	sort.Ints(costs)
	res := 0
	for _, i := range costs {
		if nc := coins - i; nc >= 0 {
			coins = nc
			res++
		}
	}
	return res
}

func main() {
	testCases := []struct {
		costs []int
		coins int
		want  int
	}{
		{
			costs: []int{1, 3, 2, 4, 1},
			coins: 7,
			want:  4,
		},
		{
			costs: []int{10, 6, 8, 7, 7, 8},
			coins: 5,
			want:  0,
		},
		{
			costs: []int{1, 6, 3, 1, 2, 5},
			coins: 20,
			want:  6,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := maxIceCream(tc.costs, tc.coins)
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
