package main

import (
	"fmt"
)

func averageWaitingTime(customers [][]int) float64 {
	totalWaitTime := float64(0)
	prevEndTime := float64(0)

	for i := range customers {
		s, t := float64(customers[i][0]), float64(customers[i][1])
		delay := float64(0)
		if prevEndTime > s {
			delay = prevEndTime - s
		}
		totalWaitTime += delay + t
		prevEndTime = s + delay + t
	}

	return totalWaitTime / float64(len(customers))
}

func main() {
	testCases := []struct {
		customers [][]int
		want      float64
	}{
		{
			customers: [][]int{{1, 2}, {2, 5}, {4, 3}},
			want:      5.00000,
		},
		{
			customers: [][]int{{5, 2}, {5, 4}, {10, 3}, {20, 1}},
			want:      3.25000,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := averageWaitingTime(tc.customers)
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
