package main

import (
	"container/heap"
	"fmt"
)

// source: https://leetcode.com/problems/minimize-deviation-in-array/

type IntHeap []int

func NewIntHeap(values ...int) *IntHeap {
	h := IntHeap(values)
	heap.Init(&h)
	return &h
}

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func (h IntHeap) Peek() int {
	return h[0]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func minimumDeviation(nums []int) int {
	minVal := 1<<31 - 1

	for i := range nums {
		if nums[i]%2 == 1 {
			nums[i] *= 2
		}
		minVal = min(minVal, nums[i])
	}

	pq := NewIntHeap(nums...)
	ans := pq.Peek() - minVal

	for pq.Peek()%2 == 0 {
		x := heap.Pop(pq).(int)
		x /= 2
		minVal = min(minVal, x)
		heap.Push(pq, x)
		ans = min(ans, pq.Peek()-minVal)
	}

	return ans
}

func main() {
	testCases := []struct {
		nums []int
		want int
	}{
		{
			nums: []int{3, 5},
			want: 1,
		},
		{
			nums: []int{10, 4, 3},
			want: 2,
		},
		{
			nums: []int{4, 1, 5, 20, 3},
			want: 3,
		},

		{
			nums: []int{1, 11},
			want: 9,
		},
		{
			nums: []int{2, 10, 8},
			want: 3,
		},
		{
			nums: []int{1, 2, 3, 4},
			want: 1,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := minimumDeviation(tc.nums)
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
