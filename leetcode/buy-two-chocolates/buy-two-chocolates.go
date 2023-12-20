package main

import (
	"fmt"
)

func buyChoco(prices []int, money int) int {
	p1, p2 := 1<<32-1, 1<<32-1
	for _, p := range prices {
		if p < p1 {
			p1, p2 = p, p1
		} else if p < p2 {
			p2 = p
		}
	}

	if p1+p2 > money {
		return money
	}

	return money - p1 - p2
}

func main() {
	testCases := []struct {
		prices []int
		money  int
		want   int
	}{
		{
			prices: []int{1, 2, 2},
			money:  3,
			want:   0,
		},
		{
			prices: []int{3, 2, 3},
			money:  3,
			want:   3,
		},
		{
			prices: []int{69, 91, 78, 19, 40, 13},
			money:  94,
			want:   62,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := buyChoco(tc.prices, tc.money)
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
