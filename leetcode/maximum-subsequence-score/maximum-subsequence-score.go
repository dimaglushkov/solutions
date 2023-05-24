package main

import (
	"container/heap"
	"fmt"
	"sort"
)

type MinHeap [][2]int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i][0] < h[j][0] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.([2]int))
}

func (h *MinHeap) Pop() interface{} {
	x := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return x
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func maxScore(nums1 []int, nums2 []int, k int) int64 {
	arr := make([][2]int, len(nums1))
	for i := range nums1 {
		arr[i] = [2]int{nums1[i], nums2[i]}
	}

	sort.Slice(arr, func(i, j int) bool {
		return arr[i][1] > arr[j][1]
	})

	minHeap := &MinHeap{}
	sum := 0
	for i := 0; i < k; i++ {
		heap.Push(minHeap, arr[i])
		sum += arr[i][0]
	}

	ans := sum * arr[k-1][1]
	for i := k; i < len(arr); i++ {
		x := heap.Pop(minHeap).([2]int)
		sum += arr[i][0] - x[0]

		heap.Push(minHeap, arr[i])
		ans = max(ans, sum*arr[i][1])
	}

	return int64(ans)
}

func main() {
	testCases := []struct {
		nums1 []int
		nums2 []int
		k     int
		want  int64
	}{
		{
			nums1: []int{1, 3, 3, 2},
			nums2: []int{2, 1, 3, 4},
			k:     3,
			want:  12,
		},
		{
			nums1: []int{4, 2, 3, 1, 1},
			nums2: []int{7, 5, 10, 9, 6},
			k:     1,
			want:  30,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := maxScore(tc.nums1, tc.nums2, tc.k)
		status := "ERROR"
		if fmt.Sprint(x) == fmt.Sprint(tc.want) {
			status = "OK"
			successes++
		}
		fmt.Println(status, "	Expected: ", tc.want, " Actual: ", x)
	}
	if l := len(testCases); successes == len(testCases) {
		fmt.Printf("===\nSUCCESS: %d of %d tests ended successfully\n", successes, l)
	} else {
		fmt.Printf("===\nFAIL: %d tests failed\n", l-successes)
	}

}
