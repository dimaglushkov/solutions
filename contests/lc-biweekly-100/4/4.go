package main

import (
	"container/heap"
	"fmt"
	"sort"
)

// source: https://leetcode.com/problems/minimum-time-to-repair-cars/

type pair struct {
	v int64
	i int
}

type IntHeap []pair

func NewIntHeap(values ...pair) *IntHeap {
	h := IntHeap(values)
	heap.Init(&h)
	return &h
}

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i].v < h[j].v }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(pair))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func repairCars(ranks []int, cars int) int64 {
	n := len(ranks)
	cnt := make([]int, n)
	sort.Ints(ranks)
	pq := NewIntHeap()

	for i := range cnt {
		pq.Push(pair{
			v: int64(ranks[i]),
			i: i,
		})
	}

	for cars > 0 {
		next := heap.Pop(pq).(pair)
		cnt[next.i]++
		cars--
		heap.Push(pq, pair{
			v: int64((cnt[next.i] + 1) * (cnt[next.i] + 1) * ranks[next.i]),
			i: next.i,
		})
	}

	var ans int64
	for i := range cnt {
		if x := int64(cnt[i] * cnt[i] * ranks[i]); x > ans {
			ans = x
		}
	}
	return ans
}

func main() {
	testCases := []struct {
		ranks []int
		cars  int
		want  int64
	}{
		{
			ranks: []int{4, 2, 3, 1},
			cars:  10,
			want:  16,
		},
		{
			ranks: []int{5, 1, 8},
			cars:  6,
			want:  16,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := repairCars(tc.ranks, tc.cars)
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
