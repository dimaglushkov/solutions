package main

const maxLen = 1001

func intersection(nums [][]int) []int {
	var cnt [maxLen]int
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums[i]); j++ {
			cnt[nums[i][j]]++
		}
	}

	res := make([]int, 0)
	for i := 0; i < maxLen; i++ {
		if cnt[i] == len(nums) {
			res = append(res, i)
		}
	}

	return res
}

func main() {
	println(intersection([][]int{{3, 1, 2, 4, 5}, {1, 2, 3, 4}, {3, 4, 5, 6}}))
	println(intersection([][]int{{1, 2, 3}, {4, 5, 6}}))
}
