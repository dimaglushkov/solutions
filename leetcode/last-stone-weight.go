package main

import (
	"container/heap"
	"fmt"
)

// source: https://leetcode.com/problems/last-stone-weight/
// first version uses container/heap

type IntHeap []int

func NewIntHeap(values ...int) *IntHeap {
	h := make(IntHeap, 0, len(values))
	h = append(h, values...)
	return &h
}

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func lastStoneWeight(stones []int) int {
	pq := NewIntHeap(stones...)
	heap.Init(pq)
	for pq.Len() > 1 {
		x, y := heap.Pop(pq).(int), heap.Pop(pq).(int)
		if x != y {
			heap.Push(pq, x-y)
		}
	}
	if pq.Len() == 0 {
		return 0
	}
	return pq.Pop().(int)
}

/*
// second version is not working properly, unfortunately no time to fix
func heapify(v *[]int, i int) {
	var largest = i
	var il = 2*i + 1
	var ir = il + 1

	if il < len(*v) && (*v)[il] < (*v)[largest] {
		largest = il
	}

	if ir < len(*v) && (*v)[ir] < (*v)[largest] {
		largest = ir
	}

	if largest != i {
		(*v)[i], (*v)[largest] = (*v)[largest], (*v)[i]
		heapify(v, largest)
	}
}

func push(v *[]int, x int) {
	*v = append(*v, x)
	for i := len(*v)/2 - 1; i >= 0; i-- {
		heapify(v, i)
	}
}

func pop(v *[]int) int {
	old := *v
	n := len(old)
	x := old[n-1]
	*v = old[0 : n-1]

	for i := len(*v)/2 - 1; i >= 0; i-- {
		heapify(v, i)
	}
	return x
}

func lastStoneWeight(stones []int) int {
	stonesPtr := &stones
	// heapify source slice
	for i := len(stones)/2 - 1; i >= 0; i-- {
		heapify(stonesPtr, i)
	}
	for len(*stonesPtr) > 1 {
		x, y := pop(stonesPtr), pop(stonesPtr)
		if x != y {
			push(stonesPtr, x-y)
		}
	}
	if len(*stonesPtr) == 0 {
		return 0
	}
	return pop(stonesPtr)
}*/

func main() {
	// Example 1
	var stones1 = []int{2, 7, 4, 1, 8, 1}
	fmt.Println("Expected: 1	Output: ", lastStoneWeight(stones1))

	// Example 2
	var stones2 = []int{1}
	fmt.Println("Expected: 1	Output: ", lastStoneWeight(stones2))

	// Example 3
	var stones3 = []int{1, 1}
	fmt.Println("Expected: 0	Output: ", lastStoneWeight(stones3))

}
