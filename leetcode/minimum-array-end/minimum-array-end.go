package main

// source: https://leetcode.com/problems/minimum-array-end/
// [n]int, for each element nums[i] < nums[i+1] and ANDing all elements equals x
// return min possible value for nums[n-1] last element

func getBits(x int) []int {
	bits := make([]int, 0, 32)
	for x > 0 {
		bits = append(bits, x&1)
		x >>= 1
	}

	return bits
}

func minEnd(n int, x int) int64 {
	nb := getBits(n - 1)
	xb := getBits(x)

	ni := 0
	// add n to x with saving x's bits unchanged
	for xi := 0; xi < len(xb) && ni < len(nb); xi++ {
		if xb[xi] == 0 {
			xb[xi] = nb[ni]
			ni++
		}
	}

	// add remaining
	if ni < len(nb) {
		xb = append(xb, nb[ni:]...)
	}

	var ans int64
	for i := len(xb) - 1; i >= 0; i-- {
		ans <<= 1
		ans += int64(xb[i])
	}

	return ans
}

func minEnd_TLE(n int, x int) int64 {
	var ans int64 = 0
	x64 := int64(x)
	for c := int64(x); n > 0; c++ {
		if c&x64 == x64 {
			n--
			ans = c
		}
	}

	return ans
}

func main() {
	minEnd(3, 4)
	minEnd(2, 7)
}
