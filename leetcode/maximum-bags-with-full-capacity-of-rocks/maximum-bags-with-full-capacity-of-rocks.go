package main

import (
	"fmt"
	"sort"
)

// source: https://leetcode.com/problems/maximum-bags-with-full-capacity-of-rocks/

func maximumBags(c []int, r []int, ar int) int {
	n := len(c)
	av := make([]int, n)
	for i := range c {
		av[i] = c[i] - r[i]
	}
	sort.Ints(av)

	res := 0
	for i := 0; i < n && ar > 0; i++ {
		if ar >= av[i] {
			ar -= av[i]
			res++
		}
	}
	return res
}

func main() {
	testCases := []struct {
		capacity        []int
		rocks           []int
		additionalRocks int
		want            int
	}{
		{
			capacity:        []int{2, 3, 4, 5},
			rocks:           []int{1, 2, 4, 4},
			additionalRocks: 2,
			want:            3,
		},
		{
			capacity:        []int{10, 2, 2},
			rocks:           []int{2, 2, 0},
			additionalRocks: 100,
			want:            3,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := maximumBags(tc.capacity, tc.rocks, tc.additionalRocks)
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
