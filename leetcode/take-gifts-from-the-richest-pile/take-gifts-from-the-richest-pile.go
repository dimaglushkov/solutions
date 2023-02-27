package main

import (
	"container/heap"
	"fmt"
	"math"
)

// source: https://leetcode.com/problems/take-gifts-from-the-richest-pile/

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
	*h = old[0 : n-1]
	return x
}

func pickGifts(gifts []int, k int) int64 {
	pq := NewIntHeap(gifts...)
	for k > 0 {
		x := heap.Pop(pq).(int)
		nx := int(math.Sqrt(float64(x)))
		heap.Push(pq, nx)
		k--
	}

	var ans int64
	for _, i := range *pq {
		ans += int64(i)
	}
	return ans
}

func main() {
	testCases := []struct {
		gifts []int
		k     int
		want  int64
	}{
		{
			gifts: []int{25, 64, 9, 4, 100},
			k:     4,
			want:  29,
		},
		{
			gifts: []int{1, 1, 1, 1},
			k:     4,
			want:  4,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := pickGifts(tc.gifts, tc.k)
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
