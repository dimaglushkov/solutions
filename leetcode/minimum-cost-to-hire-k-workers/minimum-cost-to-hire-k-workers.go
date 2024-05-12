package main

import (
	"container/heap"
	"fmt"
	"math"
	"sort"
)

type worker struct {
	quality, wage int
}

// Max heap for quality
type maxHeap []int

func (h maxHeap) Len() int {
	return len(h)
}
func (h maxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h maxHeap) Less(i, j int) bool {
	return h[i] > h[j]
}
func (h *maxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *maxHeap) Pop() interface{} {
	x := (*h)[h.Len()-1]
	*h = (*h)[0 : h.Len()-1]
	return x
}

func mincostToHireWorkers(quality []int, wage []int, k int) float64 {
	workers := make([]worker, len(quality))
	for i := range workers {
		workers[i] = worker{quality[i], wage[i]}
	}

	// Sort workers by ratio of wage/quality
	sort.Slice(workers, func(i, j int) bool {
		return workers[i].wage*workers[j].quality < workers[j].wage*workers[i].quality
	})

	h, qSum, minCost := make(maxHeap, 0), 0, math.MaxFloat64
	for _, w := range workers {
		heap.Push(&h, w.quality)
		qSum += w.quality
		if h.Len() > k {
			qSum -= heap.Pop(&h).(int)
		}
		cost := float64(w.wage*qSum) / float64(w.quality)
		if h.Len() == k && cost < minCost {
			minCost = cost
		}
	}

	return minCost
}

func main() {
	testCases := []struct {
		quality []int
		wage    []int
		k       int
		want    float64
	}{
		{
			quality: []int{10, 20, 5},
			wage:    []int{70, 50, 30},
			k:       2,
			want:    105.00000,
		},
		{
			quality: []int{3, 1, 10, 10, 1},
			wage:    []int{4, 8, 2, 2, 7},
			k:       3,
			want:    30.66667,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := mincostToHireWorkers(tc.quality, tc.wage, tc.k)
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
