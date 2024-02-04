package main

import (
	"fmt"
)

func countTestedDevices(batteryPercentages []int) int {
	ans := 0
	tests := 0

	for i := range batteryPercentages {
		if batteryPercentages[i]-tests > 0 {
			ans++
			tests++
		}
	}

	return ans
}

func main() {
	testCases := []struct {
		batteryPercentages []int
		want               int
	}{
		{
			batteryPercentages: []int{1, 1, 2, 1, 3},
			want:               3,
		},
		{
			batteryPercentages: []int{0, 1, 2},
			want:               2,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := countTestedDevices(tc.batteryPercentages)
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
