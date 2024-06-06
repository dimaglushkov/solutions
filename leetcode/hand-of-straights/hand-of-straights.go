package main

import (
	"fmt"
	"sort"
)

func isNStraightHand(hand []int, groupSize int) bool {
	n := len(hand)

	sort.Ints(hand)
	cnt := make(map[int]int)
	for i := range hand {
		cnt[hand[i]]++
	}

	if n%groupSize != 0 {
		return false
	}

	for i := range hand {
		if cnt[hand[i]] > 0 {
			cnt[hand[i]]--

			for j := 1; j < groupSize; j++ {
				if cnt[hand[i]+j] == 0 {
					return false
				}

				cnt[hand[i]+j]--
			}
		}
	}

	return true
}

func main() {
	testCases := []struct {
		hand      []int
		groupSize int
		want      bool
	}{
		{
			hand:      []int{8, 10, 12},
			groupSize: 3,
			want:      false,
		},
		{
			hand:      []int{1, 2, 3, 6, 2, 3, 4, 7, 8},
			groupSize: 3,
			want:      true,
		},
		{
			hand:      []int{1, 2, 3, 4, 5},
			groupSize: 4,
			want:      false,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := isNStraightHand(tc.hand, tc.groupSize)
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
