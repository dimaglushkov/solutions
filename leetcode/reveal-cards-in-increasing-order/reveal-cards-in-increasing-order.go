package main

import (
	"fmt"
	"sort"
)

func deckRevealedIncreasing(deck []int) []int {
	sort.Ints(deck)

	ans := make([]int, len(deck))
	ans[0] = deck[0]
	ii := 0

	for _, x := range deck[1:] {
		for j := 0; j < 2; j++ {
			ii++
			if ii == len(deck) {
				ii = 0
			}
			for ans[ii] != 0 {
				ii++
				if ii == len(deck) {
					ii = 0
				}
			}
		}
		ans[ii] = x
	}

	return ans
}

func main() {
	testCases := []struct {
		deck []int
		want []int
	}{
		{
			deck: []int{17, 13, 11, 2, 3, 5, 7},
			want: []int{2, 13, 3, 11, 5, 17, 7},
		},
		{
			deck: []int{1, 1000},
			want: []int{1, 1000},
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := deckRevealedIncreasing(tc.deck)
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
