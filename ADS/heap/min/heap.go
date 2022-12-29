package max

import "container/heap"

type IntHeap []int

func NewIntHeap(values ...int) *IntHeap {
	h := IntHeap(values)
	heap.Init(&h)
	return &h
}

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
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
