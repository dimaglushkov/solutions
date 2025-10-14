package main

// source: https://leetcode.com/problems/adjacent-increasing-subarrays-detection-i/
func hasIncreasingSubarrays(nums []int, k int) bool {
	if k == 1 {
		return true
	}

	sa := make(map[int]bool)

	c := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			c++
			if c >= k && i >= k-1 {
				sa[i-k+1] = true
			}
		} else {
			c = 1
		}
	}

	for x := range sa {
		if sa[x+k] {
			return true
		}
	}

	return false
}

func main() {
	hasIncreasingSubarrays([]int{-15, 19}, 1)
	// 0, 1, 2, 5
	hasIncreasingSubarrays([]int{2, 5, 7, 8, 9, 2, 3, 4, 3, 1}, 3)
}
