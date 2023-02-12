package main

import "strconv"

func atoi(x string) int64 {
	y, _ := strconv.ParseInt(x, 10, 32)
	return y
}

func itoa(x int) string {
	return strconv.FormatInt(int64(x), 10)
}

func findTheArrayConcVal(nums []int) int64 {
	var ans int64
	for len(nums) > 0 {
		if len(nums) == 1 {
			ans += int64(nums[0])
			return ans
		}
		x, y := itoa(nums[0]), itoa(nums[len(nums)-1])
		nums = nums[1 : len(nums)-1]
		ans += atoi(x + y)
	}
	return ans
}
