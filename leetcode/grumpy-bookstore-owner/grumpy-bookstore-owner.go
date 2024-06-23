package main

import (
	"fmt"
)

func maxSatisfied(customers []int, grumpy []int, minutes int) int {
	n := len(customers)

	maxUnsatisfiedCount := 0
	for i := 0; i < minutes; i++ {
		if grumpy[i] == 1 {
			maxUnsatisfiedCount += customers[i]
		}
	}
	ans := maxUnsatisfiedCount

	for i := 0; i < n-minutes; i++ {
		if grumpy[i] == 1 {
			maxUnsatisfiedCount -= customers[i]
		}
		if grumpy[i+minutes] == 1 {
			maxUnsatisfiedCount += customers[i+minutes]
		}

		if maxUnsatisfiedCount > ans {
			ans = maxUnsatisfiedCount
		}
	}

	for i := range customers {
		if grumpy[i] == 0 {
			ans += customers[i]
		}
	}

	return ans
}

func main() {
	testCases := []struct {
		customers []int
		grumpy    []int
		minutes   int
		want      int
	}{
		{
			customers: []int{1, 0, 1, 2, 1, 1, 7, 5},
			grumpy:    []int{0, 1, 0, 1, 0, 1, 0, 1},
			minutes:   3,
			want:      16,
		},
		{
			customers: []int{1},
			grumpy:    []int{0},
			minutes:   1,
			want:      1,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := maxSatisfied(tc.customers, tc.grumpy, tc.minutes)
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
