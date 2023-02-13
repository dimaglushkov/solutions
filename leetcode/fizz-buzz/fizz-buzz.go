package main

import (
	"fmt"
	"strconv"
)

// source: https://leetcode.com/problems/fizz-buzz/

func itoa(x int) string {
	return strconv.FormatInt(int64(x), 10)
}

func fizzBuzz(n int) []string {
	ans := make([]string, n+1)
	for i := 1; i <= n; i++ {
		if d3, d5 := i%3 == 0, i%5 == 0; d3 && d5 {
			ans[i] = "FizzBuzz"
		} else if d5 {
			ans[i] = "Buzz"
		} else if d3 {
			ans[i] = "Fizz"
		} else {
			ans[i] = itoa(i)
		}
	}
	return ans[1:]
}

func main() {
	testCases := []struct {
		n    int
		want []string
	}{
		{
			n:    3,
			want: []string{"1", "2", "Fizz"},
		},
		{
			n:    5,
			want: []string{"1", "2", "Fizz", "4", "Buzz"},
		},
		{
			n:    15,
			want: []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"},
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := fizzBuzz(tc.n)
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
