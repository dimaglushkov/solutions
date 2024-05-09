package main

import (
	"fmt"
	"sort"
)

func maximumHappinessSum(happiness []int, k int) int64 {
	sort.Sort(sort.Reverse(sort.IntSlice(happiness)))

	var ans int64
	for i := 0; i < k; i++ {
		if happiness[i] > i {
			ans += int64(happiness[i] - i)
		}
	}

	return ans
}

func main() {
	testCases := []struct {
		happiness []int
		k         int
		want      int64
	}{
		{
			happiness: []int{1, 2, 3},
			k:         2,
			want:      4,
		},
		{
			happiness: []int{1, 1, 1, 1},
			k:         2,
			want:      1,
		},
		{
			happiness: []int{2, 3, 4, 5},
			k:         1,
			want:      5,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := maximumHappinessSum(tc.happiness, tc.k)
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
