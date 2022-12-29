package main

import (
	"container/heap"
	"fmt"
)

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

// PushVal pushes value into heap and fixed the heap
func (h *IntHeap) PushVal(x int) {
	heap.Push(h, x)
}

// PopVal removes a top value from a heap and returns it
func (h *IntHeap) PopVal() int {
	return heap.Pop(h).(int)
}

// Fix fixes heap after changes at ith index
func (h *IntHeap) Fix(i int) {
	heap.Fix(h, i)
}

func sum(x ...int) int {
	res := 0
	for _, i := range x {
		res += i
	}
	return res
}

func minStoneSum(piles []int, k int) int {
	h := *NewIntHeap(piles...)

	for i := 0; i < k; i++ {
		h[0] -= h[0] / 2
		h.Fix(0)
	}

	return sum(piles...)
}

func main() {
	testCases := []struct {
		piles []int
		k     int
		want  int
	}{
		{
			piles: []int{10000},
			k:     10000,
			want:  1,
		},
		{
			piles: []int{5, 4, 9},
			k:     2,
			want:  12,
		},
		{
			piles: []int{4, 3, 6, 7},
			k:     3,
			want:  12,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := minStoneSum(tc.piles, tc.k)
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
