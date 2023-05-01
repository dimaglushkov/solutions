package main

import (
	"fmt"
	"sort"
)

func average(salary []int) float64 {
	sort.Ints(salary)
	var sum int64
	for i := 1; i < len(salary)-1; i++ {
		sum += int64(salary[i])
	}
	return float64(sum) / float64(len(salary)-2)
}

func main() {
	testCases := []struct {
		salary []int
		want   float64
	}{
		{
			salary: []int{4000, 3000, 1000, 2000},
			want:   2500.00000,
		},
		{
			salary: []int{1000, 2000, 3000},
			want:   2000.00000,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := average(tc.salary)
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
