package main

import (
	"fmt"
)

func candy(ratings []int) int {
	if len(ratings) == 0 {
		return 0
	}

	ans := len(ratings)
	up, down, peak := 0, 0, 0

	for i := 0; i < len(ratings)-1; i++ {
		prev, cur := ratings[i], ratings[i+1]

		if prev < cur {
			up, down, peak = up+1, 0, up
			ans += up
		} else if prev == cur {
			up, down, peak = 0, 0, 0
		} else {
			up, down = 0, down+1
			ans += down
			if peak >= down {
				ans--
			}
		}
	}

	return ans
}

func main() {
	testCases := []struct {
		ratings []int
		want    int
	}{
		{
			ratings: []int{1, 0, 2},
			want:    5,
		},
		{
			ratings: []int{1, 2, 2},
			want:    4,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := candy(tc.ratings)
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
