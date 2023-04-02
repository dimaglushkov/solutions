package main

func main() {
	println(minNumber([]int{1}, []int{1}))

}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func minNumber(nums1 []int, nums2 []int) int {
	var cnt1, cnt2 [10]int
	for _, i := range nums1 {
		cnt1[i]++
	}
	for _, i := range nums2 {
		cnt2[i]++
	}

	v1, v2 := 10, 10
	for i := range cnt1 {
		if cnt1[i] == 1 && cnt2[i] == 1 {
			return i
		}
		if cnt1[i] != 0 {
			v1 = min(i, v1)
		}
		if cnt2[i] != 0 {
			v2 = min(i, v2)
		}
	}

	return min(v1, v2)*10 + max(v1, v2)
}
