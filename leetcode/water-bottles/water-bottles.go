package main

import (
	"fmt"
)

func numWaterBottles(numBottles int, numExchange int) int {
	totalBottles := numBottles

	for numBottles >= numExchange {
		totalBottles += numBottles / numExchange
		numBottles = (numBottles / numExchange) + (numBottles % numExchange)
	}

	return totalBottles
}

func main() {
	testCases := []struct {
		numBottles  int
		numExchange int
		want        int
	}{
		{
			numBottles:  9,
			numExchange: 3,
			want:        13,
		},
		{
			numBottles:  15,
			numExchange: 4,
			want:        19,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := numWaterBottles(tc.numBottles, tc.numExchange)
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
