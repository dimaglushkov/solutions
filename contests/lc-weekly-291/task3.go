package main

import (
	"sort"
	"strconv"
)

func countDistinct(nums []int, k int, p int) int {
	var subarrays = make(map[string]bool)
	var div, nondiv = make([]int, 0, 201), make([]int, 0, 201)

	toString := func(nums []int) string {
		var s string
		for i := range nums {
			s += strconv.FormatInt(int64(i), 10)
		}
		return s
	}

	for i := range nums {
		if nums[i]%p == 0 {
			div = append(div, nums[i])
		} else {
			nondiv = append(nondiv, nums[i])
		}
	}

	sort.Ints(div)

	for i := 0; i <= k; i++ {
		for
	}

	return 0
}

func permute(nums []int) [][]int {
	res := make([][]int, 0)
	var permuteUtil func(nums, cur []int)
	permuteUtil = func(nums, cur []int) {
		if len(nums) == 0 {
			res = append(res, cur)
			return
		}
		for i := 0; i < len(nums); i++ {
			var newNums = make([]int, len(nums))
			copy(newNums, nums)
			permuteUtil(append(newNums[:i], newNums[i+1:]...), append(cur, nums[i]))
		}
	}

	permuteUtil(nums, []int{})
	return res
}

func main() {
	countDistinct([]int{2, 3, 3, 2, 2}, 2, 2)
}
