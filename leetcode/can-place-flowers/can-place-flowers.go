package main

import (
	"fmt"
)

// source: https://leetcode.com/problems/can-place-flowers/

func canPlaceFlowers(flowerbed []int, n int) bool {
	l := len(flowerbed)
	for i := range flowerbed {
		if flowerbed[i] == 0 && (i == 0 || flowerbed[i-1] == 0) && (i == l-1 || flowerbed[i+1] == 0) {
			flowerbed[i] = 1
			n--
		}
	}

	return n <= 0
}

func main() {
	testCases := []struct {
		flowerbed []int
		n         int
		want      bool
	}{
		{
			flowerbed: []int{1, 0, 0, 0, 1},
			n:         1,
			want:      true,
		},
		{
			flowerbed: []int{1, 0, 0, 0, 1},
			n:         2,
			want:      false,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := canPlaceFlowers(tc.flowerbed, tc.n)
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
