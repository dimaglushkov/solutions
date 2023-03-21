package main

// source: https://leetcode.com/problems/range-sum-query-immutable/

// ["NumArray", "sumRange", "sumRange", "sumRange"]
// [[], [0, 2], [2, 5], [0, 5]]
// Output
// [null, 1, -1, -3]
// [-2,  0, 3, -5,  2, -1]
// [-2, -2, 1, -4, -2, -3]
//   0   1  2   3   4   5
// 1: 1 - 2
//

type NumArray struct {
	ps []int
}

func Constructor(nums []int) NumArray {
	n := len(nums)
	ps := make([]int, n)
	ps[0] = nums[0]
	for i := 1; i < n; i++ {
		ps[i] = nums[i] + ps[i-1]
	}
	return NumArray{ps: ps}
}

func (na *NumArray) SumRange(left int, right int) int {
	if left == 0 {
		return na.ps[right]
	}
	return na.ps[right] - na.ps[left-1]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */
