package main

import (
	"fmt"
)

func triangleType(nums []int) string {
	if nums[0]+nums[1] <= nums[2] || nums[1]+nums[2] <= nums[0] || nums[0]+nums[2] <= nums[1] {
		return "none"
	}

	if nums[0] == nums[1] && nums[1] == nums[2] && nums[0] == nums[2] {
		return "equilateral"
	} else if nums[0] == nums[1] || nums[1] == nums[2] || nums[0] == nums[2] {
		return "isosceles"
	} else {
		return "scalene"
	}
}

func main() {
	testCases := []struct {
		nums []int
		want string
	}{
		{
			nums: []int{3, 3, 3},
			want: "equilateral",
		},
		{
			nums: []int{3, 4, 5},
			want: "scalene",
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := triangleType(tc.nums)
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
