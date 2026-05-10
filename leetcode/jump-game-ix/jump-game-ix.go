package main

// source: https://leetcode.com/problems/jump-game-ix/
func maxValue(nums []int) []int {
	n := len(nums)

	pre := make([]int, n)
	suf := make([]int, n)
	ans := make([]int, n)

	pre[0] = nums[0]
	for i := 1; i < n; i++ {
		pre[i] = max(pre[i-1], nums[i])
	}

	suf[n-1] = nums[n-1]
	for i := n - 2; i >= 0; i-- {
		suf[i] = min(suf[i+1], nums[i])
	}

	ans[n-1] = pre[n-1]
	for i := n - 2; i >= 0; i-- {
		if pre[i] > suf[i+1] {
			ans[i] = ans[i+1]
		} else {
			ans[i] = pre[i]
		}
	}

	return ans
}

func main() {
	maxValue([]int{2, 1, 3})
	maxValue([]int{2, 3, 1})
}
