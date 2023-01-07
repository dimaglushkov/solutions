package main

import (
	"fmt"
)

// source: https://leetcode.com/problems/gas-station/

func canCompleteCircuit(gas []int, cost []int) int {
	res, s1, s2 := 0, 0, 0
	for i := range gas {
		s1 += gas[i] - cost[i]
		s2 += gas[i] - cost[i]
		if s2 < 0 {
			res = i + 1
			s2 = 0
		}
	}

	if s1 < 0 {
		return -1
	}
	return res
}

func main() {
	testCases := []struct {
		gas  []int
		cost []int
		want int
	}{
		{
			gas:  []int{1, 2, 3, 4, 5},
			cost: []int{3, 4, 5, 1, 2},
			want: 3,
		},
		{
			gas:  []int{2, 3, 4},
			cost: []int{3, 4, 3},
			want: -1,
		},
		{
			gas:  []int{5, 8, 2, 8},
			cost: []int{6, 5, 6, 6},
			want: 3,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := canCompleteCircuit(tc.gas, tc.cost)
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
