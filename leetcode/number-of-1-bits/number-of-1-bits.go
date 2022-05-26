package main

// source: https://leetcode.com/problems/number-of-1-bits/

func hammingWeight(num uint32) int {
	var res uint32

	for num > 0 {
		res += num & 1
		num >>= 1
	}

	return int(res)
}
