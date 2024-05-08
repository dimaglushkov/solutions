package main

import (
	"fmt"
	"sort"
	"strconv"
)

func findRelativeRanks(score []int) []string {
	var temp []int
	var ans []string
	mp := map[int]int{}

	temp = append(temp, score...)

	sort.Sort(sort.Reverse(sort.IntSlice(temp)))

	for i := range temp {
		mp[temp[i]] = i + 1
	}

	for i := range score {
		switch mp[score[i]] {
		case 1:
			ans = append(ans, "Gold Medal")
		case 2:
			ans = append(ans, "Silver Medal")
		case 3:
			ans = append(ans, "Bronze Medal")
		default:
			ans = append(ans, strconv.Itoa(mp[score[i]]))
		}
	}

	return ans
}

func main() {
	testCases := []struct {
		score []int
		want  []string
	}{
		{
			score: []int{5, 4, 3, 2, 1},
			want:  []string{"Gold Medal", "Silver Medal", "Bronze Medal", "4", "5"},
		},
		{
			score: []int{10, 3, 8, 9, 4},
			want:  []string{"Gold Medal", "5", "Bronze Medal", "Silver Medal", "4"},
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := findRelativeRanks(tc.score)
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
