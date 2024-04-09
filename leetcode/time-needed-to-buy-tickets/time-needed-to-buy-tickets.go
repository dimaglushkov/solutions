package main

import (
	"fmt"
)

func timeRequiredToBuy(tickets []int, k int) int {
	result := 0
	c := tickets[k]
	for i, n := range tickets {
		if i == k {
			result += c
		} else if i < k {
			result += min(c, n)
		} else {
			result += min(c-1, n)
		}
	}
	return result
}

func main() {
	testCases := []struct {
		tickets []int
		k       int
		want    int
	}{
		{
			tickets: []int{2, 3, 2},
			k:       2,
			want:    6,
		},
		{
			tickets: []int{5, 1, 1, 1},
			k:       0,
			want:    8,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := timeRequiredToBuy(tc.tickets, tc.k)
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
