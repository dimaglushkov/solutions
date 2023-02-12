package main

import "sort"

func countFairPairs(nums []int, lower int, upper int) int64 {
	n := len(nums)
	var ans int64
	sortedNums := make([]int, len(nums))
	copy(sortedNums, nums)
	sort.Ints(sortedNums)

	for _, x := range nums {
		li, ui := sort.SearchInts(sortedNums, lower-x), sort.SearchInts(sortedNums, upper-x)
		if x >= lower-x && x <= upper-x {
			ans--
		}
		for ui < n && sortedNums[ui] == upper-x {
			ui++
		}
		ans += int64(ui - li)
	}

	return ans / 2
}
