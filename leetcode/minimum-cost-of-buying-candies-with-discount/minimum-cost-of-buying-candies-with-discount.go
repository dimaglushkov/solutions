package main

import (
	"fmt"
	"sort"
)

func minimumCost(cost []int) int {
	sort.Slice(cost, func(i, j int) bool { return cost[i] > cost[j] })

	ans := 0

	for i := range cost {
		if (i+1)%3 != 0 {
			ans += cost[i]
		}
	}

	return ans
}

func main() {
	testCases := []struct {
		cost []int
		want int
	}{
		{
			cost: []int{1, 2, 3},
			want: 5,
		},
		{
			cost: []int{6, 5, 7, 9, 2, 2},
			want: 23,
		},
		{
			cost: []int{5, 5},
			want: 10,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := minimumCost(tc.cost)
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
